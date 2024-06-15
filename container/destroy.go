package container

import "github.com/sirupsen/logrus"

type EndFunc func() error
var EndFn EndFunc

func SetEndFn(endFunc EndFunc) {
	EndFn = func() error {
		logrus.Infof("@container.endfunc start...")

		err := endFunc()
		if err!=nil{
			logrus.Errorf("@container.endfunc error: %s",err.Error())
			return err
		}

		logrus.Infof("@container.endfunc end...")
		return nil
	}
}