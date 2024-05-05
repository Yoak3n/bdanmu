package blivedanmu

import (
	"bdanmu/package/model"
	"math/rand/v2"
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
		Medal: make(chan int64),
		User:  make(chan int64),
		Reply: make(chan map[int64]*model.User),
	}
}

func Start() {
	initQueue()
	go handler()
	select {}

}
func handler() {
	for {
		select {
		case target := <-Queue.Medal:
			go getMedalTargetUserInfo(target)
		case user := <-Queue.User:
			go func() {
				u := getUserInfo(user)
				if u != nil {
					reply := make(map[int64]*model.User)
					reply[user] = u
					Queue.Reply <- reply
				}
			}()
		}
		r := rand.IntN(3) + 2
		time.Sleep(time.Second * time.Duration(r))
	}
}
