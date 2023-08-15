package app

import (
	"fmt"
	"user-service/app/config"
	"user-service/app/logger"

	logConfig "github.com/jfeng45/glogger/config"
	logFactory "github.com/jfeng45/glogger/factory"
	"github.com/pkg/errors"
)

func InitApp(filename ...string) {
	config, err := config.BuildConfig(filename...)
	if err != nil {
		fmt.Println(err)
	}
	err = initLogger(&config.LogConfig)
}

func initLogger(lc *logConfig.Logging) error {
	log, err := logFactory.Build(lc)
	if err != nil {
		return errors.Wrap(err, "load logger")
	}
	logger.SetLogger(log)
	return nil
}