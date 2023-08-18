package containerhelper

import (
	"user-service/app/config"
	"user-service/app/container"
	"user-service/domain/usecase"

	"github.com/pkg/errors"
)

func GetRegistrationUseCase(c container.ContainerInterface) (usecase.RegistrationUseCaseInterface, error) {
	key := config.REGISTRATION
	value, err := c.BuildUseCase(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.RegistrationUseCaseInterface), nil
}

func GetCredentialUseCase(c container.ContainerInterface) (usecase.CredentialUseCaseInterface, error) {
	key := config.CREDENTIAL
	value, err := c.BuildUseCase(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return value.(usecase.CredentialUseCaseInterface), nil
}