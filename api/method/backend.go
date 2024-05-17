package method

import (
	"bdanmu/package/logger"
	"bdanmu/service/blivedanmu"
)

func InitBackend() {
	blivedanmu.InitHub()
	c := blivedanmu.GetClient()
	err := c.Start()
	if err != nil {
		logger.Logger.Errorln(err)
	}
}

func ChangeBackend() {
	cl := blivedanmu.GetClient()
	if cl != nil {
		cl.Stop()
	}
	blivedanmu.InitHub()
	err := blivedanmu.GetClient().Start()
	if err != nil {
		logger.Logger.Errorln(err)
	}
}
