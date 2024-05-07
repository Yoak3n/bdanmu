package blivedanmu

import (
	"bdanmu/api/router/ws"
	"bdanmu/package/model"
	"sync/atomic"
	"time"
)

type QueueChan struct {
	Medal chan int64
	User  chan int64
	Reply chan map[int64]*model.User
}

var Queue *QueueChan

func SendMedalMsg(target int64) {
	Queue.Medal <- target
}

func SendUserMsg(user int64) {
	Queue.User <- user
}

func initQueue() {
	Queue = &QueueChan{
		Medal: make(chan int64, 1000),
		User:  make(chan int64, 1000),
		Reply: make(chan map[int64]*model.User),
	}
}

func Start() {
	initQueue()
	go handler()
	go collectUserId()
	select {}

}
func handler() {
	for {
		select {
		case target := <-Queue.Medal:
			go getMedalTargetUserInfo(target)
		case reply := <-Queue.Reply:
			for _, u := range reply {
				ws.UpdateUser(u)
			}
		}
	}
}

func collectUserId() {
	sendIds := make([]int64, 0)
	timer := time.NewTimer(time.Second * 2)
	var flag int32
	go func() {
		for {
			<-timer.C
			atomic.StoreInt32(&flag, 1)
			timer.Reset(time.Second * 2)
		}
	}()
	for {
		select {
		case user := <-Queue.User:
			sendIds = append(sendIds, user)
			if atomic.LoadInt32(&flag) > 0 || len(sendIds) > 10 {
				atomic.StoreInt32(&flag, 0)
				go func(ids []int64) {
					users := getUserInfoMultiply(ids)
					if len(users) > 0 {
						for _, user := range users {
							reply := make(map[int64]*model.User)
							reply[user.UID] = user
							Queue.Reply <- reply
						}
					}
				}(sendIds)
				sendIds = sendIds[0:0]
			}

		}
	}
}
