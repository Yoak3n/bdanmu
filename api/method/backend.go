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

func ChangeBackend() error {
	cl := blivedanmu.GetClient()
	if cl != nil {
		cl.Stop()
	}
	err := blivedanmu.InitHub()
	if err != nil {
		logger.Logger.Errorln(err)
		return err
	}
	err = blivedanmu.GetClient().Start()
	if err != nil {
		logger.Logger.Errorln(err)
		return err
	}
	return nil
}
