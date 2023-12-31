package credential

import (
	"user-service/app/utils"
	"user-service/applicationservice/dataservice"
	"user-service/domain/model"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/pkg/errors"
)

type CredentialUseCase struct {
	UserDataInterface dataservice.UserDataInterface
}

func (cuc *CredentialUseCase) Login(email, password string) (*model.User, error) {
	credential := struct {
		email string
		password string
	} {
		email: email,
		password: password,
	}
	err := validation.ValidateStruct(&credential,
		validation.Field(&credential.email, validation.Required, is.Email),
		validation.Field(&credential.password, validation.Required))
	if err != nil {
		return nil, errors.Wrap(err, "credential validation failed")
	}
	userByEmail, err := cuc.UserDataInterface.FindByEmail(email)
	if err != nil {
		return nil, errors.New("email not found")
	}
	match := utils.CheckPasswordHash(credential.password, userByEmail.Password)
	if !match {
		return nil, errors.New("invalid credential")
	}
	return userByEmail, nil
}