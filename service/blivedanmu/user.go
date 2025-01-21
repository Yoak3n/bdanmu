package blivedanmu

import (
	"bdanmu/config"
	"bdanmu/internal/model"
	"bdanmu/package/logger"
	"bdanmu/package/request"
	"bdanmu/package/util"
	"bdanmu/service"
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/tidwall/gjson"
)

func getUserInfoMultiply(uids []int64) (users []*model.User) {
	users = make([]*model.User, 0)
	uidsStr := make([]string, 0)
	// 简单去重，因为 uid 重复的概率极低，所以这里不考虑性能
	newArr := make([]int64, 0)
	for i := 0; i < len(uids); i++ {
		repeat := false
		for j := i + 1; j < len(uids); j++ {
			if uids[i] == uids[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, uids[i])
		}
	}

	for _, uid := range newArr {
		uidsStr = append(uidsStr, strconv.FormatInt(uid, 10))
	}
	if len(uidsStr) == 0 {
		return users
	}
	count := 0
	targets := strings.Join(uidsStr, ",")
	for {
		res, err := request.Get("https://api.vc.bilibili.com/account/v1/user/cards", fmt.Sprintf("uids=%s", targets))
		if err != nil {
			continue
		}
		result := gjson.ParseBytes(res)
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
				Name:   v.Get("name").String(),
				Sex:    util.TransSex(v.Get("sex").String()),
			}
			users = append(users, u)
		}
		// 一点冗余更新
		go service.CreateUserAndUpdateStack(users)
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
		go service.CreateUserRecord(u)
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
			r := rand.Intn(3) + 2
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
	if RoomInfo.User.UID != 0 && targetId == RoomInfo.User.UID {
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
