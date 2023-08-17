package datastorefactory

import (
	"user-service/app/config"
	"user-service/app/container"
)

type DataStoreInterface interface {}

type dsFbInterface interface {
	Build(container.ContainerInterface, *config.DataStoreConfig) (DataStoreInterface, error)
}

var dsFbMap = map[string]dsFbInterface {
	config.SQLDB: &sqlFactory{},
}

func GetDataStoreFb(key string) dsFbInterface {
	return dsFbMap[key]
}