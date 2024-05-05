package blivedanmu

import (
	"bdanmu/config"
	"bdanmu/package/logger"
	"bdanmu/package/model"
	"bdanmu/package/request"
	"bdanmu/package/util"
	"bdanmu/service"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func getUserInfoMultiply(uids []int64) (users []*model.User) {
	users = make([]*model.User, 0)
	uidsStr := make([]string, 0)
	for _, uid := range uids {
		if user := service.ReadUserRecord(uid); user != nil {
			logger.Logger.Println("use local user info:", user.UID)
			users = append(users, user)
		} else {
			uidsStr = append(uidsStr, strconv.FormatInt(uid, 10))
		}
	}
	if len(uidsStr) == 0 {
		return users
	}
	count := 0
	targets := strings.Join(uidsStr, ",")
	logger.Logger.Println(targets)
	for {
		res, err := request.Get("https://api.vc.bilibili.com/account/v1/user/cards", fmt.Sprintf("uids=%s", targets))
		if err != nil {
			continue
		}
		result := gjson.ParseBytes(res)
		logger.Logger.Println(string(res))
		if code := result.Get("code"); code.Exists() && code.Int() != 0 {
			logger.Logger.Debugln("getUserInfoMultiply:", result.Get("message").String())
			count += 1
			if count > 5 {
				return nil
			}
			time.Sleep(time.Second)
			continue
		}
		data := result.Get("data").Array()
		for _, v := range data {
			u := &model.User{
				UID:    v.Get("mid").Int(),
				Avatar: v.Get("face").String(),
				Name:   v.Get("uname").String(),
				Sex:    util.TransSex(v.Get("sex").String()),
			}
			logger.Logger.Println("create user record:", u.UID)
			go service.CreateUserRecord(u)
			users = append(users, u)
		}
		return users
	}
}

func getUserInfo(uid int64) *model.User {
	// local database to avoid anti-crawler
	if user := service.ReadUserRecord(uid); user != nil {
		return user
	}
	count := 0
	for {
		res, err := request.Get("https://api.bilibili.com/x/web-interface/card", fmt.Sprintf("mid=%d", uid))
		if err != nil {
			continue
		}
		result := gjson.ParseBytes(res)
		if code := result.Get("code"); code.Exists() && code.Int() != 0 {
			logger.Logger.Debugln("getUserInfo:", result.Get("message").String())
			count += 1
			if count > 5 {
				return nil
			}
			time.Sleep(time.Second)
			continue
		}
		data := result.Get("data")
		u := &model.User{
			UID:           uid,
			Avatar:        data.Get("card.face").String(),
			Name:          data.Get("card.name").String(),
			Sex:           util.TransSex(data.Get("card.sex").String()),
			FollowerCount: data.Get("follower").Int(),
		}
		go service.UpdateUserRecord(u)
		return u
	}
}

func getUserInfoWithWBI(uid int64) *model.Medal {
	mux.Lock()
	defer mux.Unlock()
	if user := service.ReadUserRecord(uid); user != nil {
		log.Debugln("use local user info")
		return user.Medal
	}

	count := 0
	for {
		res, err := request.GetUserDetail(uid)
		if err != nil {
			log.Errorln(err)
			time.Sleep(time.Second * 3)
		}
		log.Debugln(string(res))
		result := gjson.ParseBytes(res)
		if code := result.Get("code"); code.Exists() && code.Int() != 0 {
			count += 1
			if count > 5 {
				log.Errorf("get user %d medal failed", uid)
				return nil
			}
			r := rand.IntN(3) + 2
			err = config.RefreshCookie()
			if err != nil {
				log.Errorln(err)
			}
			time.Sleep(time.Second * time.Duration(r))
			continue
		}
		if M := result.Get("data.fans_medal.medal"); M.Get("medal_name").Exists() {
			medal := &model.Medal{
				Name:     M.Get("medal_name").String(),
				OwnerID:  uid,
				TargetID: M.Get("target_id").Int(),
				Level:    int(M.Get("level").Int()),
			}
			return medal
		}
		return nil
	}
}

// getMedalTargetUserInfo 通过勋章递归获取用户信息
func getMedalTargetUserInfo(targetId int64) {
	if roomInfo.User.UID != 0 && targetId == roomInfo.User.UID {
		return
	}
	u := getUserInfo(targetId)
	if u == nil {
		return
	}
	// 需要减少请求量，wbi接口非常容易风控检验失败
	medal := getUserInfoWithWBI(targetId)
	if medal != nil {
		u.Medal = medal
		// 互相带对方牌子则会无限查询
		if targetId == u.Medal.OwnerID {
			return
		}
		SendMedalMsg(medal.TargetID)
	}
	go service.CreateUserRecord(u)
}

func NewUserInformation(s string) *model.User {
	result := gjson.Get(s, "data")
	user := &model.User{
		UID:    result.Get("uid").Int(),
		Avatar: result.Get("uinfo.base.face").String(),
		Name:   result.Get("uname").String(),
	}
	log.Debugln("user:", user.UID)
	if user.UID == 0 {
		err := config.RefreshCookie()
		if err != nil {
			panic(err)
		}
	}
	if m := result.Get("fans_medal"); m.Get("medal_name").String() != "" {
		medal := &model.Medal{
			Name:     m.Get("medal_name").String(),
			Level:    int(m.Get("medal_level").Int()),
			OwnerID:  user.UID,
			TargetID: m.Get("target_id").Int(),
		}
		if config.Conf.Extension {
			SendMedalMsg(medal.TargetID)
		}
		user.Medal = medal
	}
	if config.Conf.Extension {
		u := getUserInfo(user.UID)
		if u != nil {
			user.Sex = u.Sex
			user.FollowerCount = u.FollowerCount
		} else {
			user.Sex = -1
			user.FollowerCount = 0
		}
	}
	if config.Conf.Database != nil && user.UID != 0 {
		go service.CreateUserRecord(user)
	}

	m := &model.Message{
		Type: 1,
		Data: user,
	}
	data, err := json.Marshal(m)
	if err != nil {
		logger.Logger.Errorln(err)
	}
	logger.Logger.Debugln(string(data))
	return user
}
