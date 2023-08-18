package usecase

import (
	"user-service/domain/model"
)

type RegistrationUseCaseInterface interface {
	SignUp(user *model.User) (resultUser *model.User, err error)
}

type CredentialUseCaseInterface interface {
	Login(email, password string) (resultUser *model.User, err error)
}