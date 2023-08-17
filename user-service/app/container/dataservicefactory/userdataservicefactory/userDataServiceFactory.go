package userdataservicefactory

import (
	"user-service/app/config"
	"user-service/app/container"
	"user-service/applicationservice/dataservice"
)

type userDataServiceFbInterface interface {
	Build(container.ContainerInterface, *config.DataConfig) (dataservice.UserDataInterface, error)
}

var udsFbMap = map[string]userDataServiceFbInterface {
	config.SQLDB: &sqlUserDataServiceFactory{},
}

func GetUserDataServiceFb(key string) userDataServiceFbInterface {
	return udsFbMap[key]
}