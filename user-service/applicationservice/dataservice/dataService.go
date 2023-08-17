package dataservice

import "user-service/domain/model"

type UserDataInterface interface {
	Insert(user *model.User) (resultUser *model.User, err error)
	FindByEmail(email string) (user *model.User, err error)
}