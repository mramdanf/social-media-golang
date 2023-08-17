package usecasefactory

import (
	"user-service/app/config"
	"user-service/app/container"
)

type UseCaseInterface interface {}

type UseCaseFbInterface interface {
	Build(c container.ContainerInterface, appConfig *config.AppConfig) (UseCaseInterface, error)
}

var useCaseFactoryBuilderMap = map[string]UseCaseFbInterface {
	config.REGISTRATION: &RegistrationFactory{},
}

func GetUseCaseFb(key string) UseCaseFbInterface {
	return useCaseFactoryBuilderMap[key]
}