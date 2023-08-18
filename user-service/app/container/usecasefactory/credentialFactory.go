package usecasefactory

import (
	"user-service/app/config"
	"user-service/app/container"
	"user-service/domain/usecase/credential"

	"github.com/pkg/errors"
)

type CredentialFactory struct {}

func (cf *CredentialFactory) Build(c container.ContainerInterface, appConfig *config.AppConfig) (UseCaseInterface, error) {
	cc := appConfig.UseCaseConfig.Credential
	udi, err := buildUserData(c, &cc.UserDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	cuc := credential.CredentialUseCase{UserDataInterface: udi}
	return &cuc, nil
}