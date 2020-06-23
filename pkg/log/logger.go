package log

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"gitlab.haochang.tv/gopkg/logger"
)

var (
	Al *logrus.Logger
	Hl *logrus.Logger
)

func init() {
	logger.MaxStackTrace = 20
}

func initAppLogger() error {
	var err error
	Al, err = logger.NewLogger(logger.APPLogsV1)
	Al.SetLevel(logrus.DebugLevel)
	Al.SetOutput(os.Stdout)
	return err
}

func initHttpLogger() error {
	var err error
	Hl, err = logger.NewLogger(logger.HTTPRequestV1)
	return err
}

//InitLogger 初始化日志
func InitLogger() error {
	var err error
	if err = initAppLogger(); err != nil {
		return fmt.Errorf("初始化日志失败 %w", err)
	}
	return initHttpLogger()
}
