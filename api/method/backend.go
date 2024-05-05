package method

import (
	"bdanmu/package/logger"
	"bdanmu/service/blivedanmu"
)

func InitBackend() {
	go blivedanmu.Start()
	blivedanmu.InitHub()
	c := blivedanmu.GetClient()
	err := c.Start()
	if err != nil {
		logger.Logger.Errorln(err)
	}
}
