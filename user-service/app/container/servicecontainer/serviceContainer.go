package servicecontainer

import (
	"user-service/app/config"
	"user-service/app/container/usecasefactory"
)

type ServiceContainer struct {
	FactoryMap map[string]interface{}
	AppConfig *config.AppConfig
}

func (sc *ServiceContainer) BuildUseCase(code string) (interface{}, error) {
	return usecasefactory.GetUseCaseFb(code).Build(sc, sc.AppConfig)
}

func (sc *ServiceContainer) Get(code string) (interface{}, bool) {
	value, found := sc.FactoryMap[code]
	return value, found
}

func (sc *ServiceContainer) Put(code string, value interface{}) {
	sc.FactoryMap[code] = value
} 