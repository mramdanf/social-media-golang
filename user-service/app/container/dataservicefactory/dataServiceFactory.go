package dataservicefactory

import (
	"user-service/app/config"
	"user-service/app/container"
)

type DataServiceInterface interface {}

type dataServiceFbInterface interface {
	Build(container.ContainerInterface, *config.DataConfig) (DataServiceInterface, error)
}

var dsFbMap = map[string]dataServiceFbInterface{
	config.USER_DATA: &userDataServiceFactoryWrapper{},
}

func GetDataServiceFb(key string) dataServiceFbInterface {
	return dsFbMap[key]
}