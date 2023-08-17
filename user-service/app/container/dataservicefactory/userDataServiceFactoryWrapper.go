package dataservicefactory

import (
	"user-service/app/config"
	"user-service/app/container"
	"user-service/app/container/dataservicefactory/userdataservicefactory"
	"user-service/app/logger"

	"github.com/pkg/errors"
)

type userDataServiceFactoryWrapper struct {}

func (udsfw *userDataServiceFactoryWrapper) Build(c container.ContainerInterface, dataConfig *config.DataConfig) (DataServiceInterface, error) {
	logger.Log.Debug("UserDataServiceFactory")
	key := dataConfig.DataStoreConfig.Code
	udsi, err := userdataservicefactory.GetUserDataServiceFb(key).Build(c, dataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return udsi, nil
}