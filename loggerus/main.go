package main

import (
	"test/logger"

	loggrus "github.com/sirupsen/logrus"
)

func main() {
	defer fileClose()

	logger.InitLogrus()
	logger.StdLogrus()

	user := logger.NewUser() // 객체에 logrus를 할당하고 사용하는 방법
	user.Errorln("this is error")
}

func fileClose() {
	if err := logger.LoggerFile.Close(); err != nil {
		loggrus.Errorln("file close error")
	}
}
/workspaces/components/logger-us