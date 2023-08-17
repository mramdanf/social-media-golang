package usecasefactory

import (
	"user-service/app/config"
	"user-service/app/container"
	"user-service/domain/usecase/registration"

	"github.com/pkg/errors"
)

type RegistrationFactory struct {}

func (rf *RegistrationFactory) Build(c container.ContainerInterface, appConfig *config.AppConfig) (UseCaseInterface, error) {
	uc := appConfig.UseCaseConfig.Registration
	udi, err := buildUserData(c, &uc.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ruc := registration.RegistrationUseCase{UserDataServiceInterface: udi}
	return &ruc, nil
}