package blivedanmu

import (
	"bdanmu/package/model"
	"math/rand/v2"
	"sync/atomic"
	"time"
)

type QueueChan struct {
	Medal chan int64
	User  chan int64
	Users chan []int64
	ids   []int64
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
		Medal: make(chan int64),
		User:  make(chan int64),
		ids:   make([]int64, 0),
		Users: make(chan []int64), // 用户IDs
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
	go func() {
		for {
			r := rand.IntN(3)
			time.Sleep(time.Second * time.Duration(r))
			Queue.Users <- Queue.ids
			Queue.ids = Queue.ids[0:0]
		}
	}()
	for {
		select {
		case target := <-Queue.Medal:
			go getMedalTargetUserInfo(target)
		case user := <-Queue.User:
			Queue.ids = append(Queue.ids, user)
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
		case users := <-Queue.Users:
			sendIds = append(sendIds, users...)
			if atomic.LoadInt32(&flag) > 0 || len(sendIds) > 20 {
				atomic.StoreInt32(&flag, 0)
				go func() {
					users := getUserInfoMultiply(sendIds)
					sendIds = sendIds[0:0]
					if len(users) > 0 {
						for _, user := range users {
							reply := make(map[int64]*model.User)
							reply[user.UID] = user
							Queue.Reply <- reply
						}
					}

				}()

			}

		}
	}
}
