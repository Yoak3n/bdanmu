package blivedanmu

import (
	"math/rand/v2"
	"time"
)

var Queue = make(chan int64, 1000)

func SendMsg(target int64) {
	Queue <- target
}

func Start() {
	for {
		select {
		case target := <-Queue:
			go getMedalTargetUserInfo(target)
		}
		r := rand.IntN(3) + 2
		time.Sleep(time.Second * time.Duration(r))
		// 处理接收到的消息

	}
}
