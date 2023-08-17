package app

import (
	"user-service/app/config"
	"user-service/app/container"
	"user-service/app/container/servicecontainer"
	"user-service/app/logger"

	logConfig "github.com/jfeng45/glogger/config"
	logFactory "github.com/jfeng45/glogger/factory"
	"github.com/pkg/errors"
)

func InitApp(filename ...string) (container.ContainerInterface, error) {
	config, err := config.BuildConfig(filename...)
	if err != nil {
		return nil, errors.Wrap(err, "BuildConfig")
	}
	err = initLogger(&config.LogConfig)
	if err != nil {
		return nil, err
	}
	return initContainer(config)
}

func initLogger(lc *logConfig.Logging) error {
	log, err := logFactory.Build(lc)
	if err != nil {
		return errors.Wrap(err, "load logger")
	}
	logger.SetLogger(log)
	return nil
}

func initContainer(config *config.AppConfig) (container.ContainerInterface, error) {
	factoryMap := make(map[string]interface{})
	c := servicecontainer.ServiceContainer{FactoryMap: factoryMap, AppConfig: config}
	return &c, nil
}