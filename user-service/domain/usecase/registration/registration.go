package registration

import (
	"user-service/app/logger"
	"user-service/applicationservice/dataservice"
	"user-service/domain/model"

	"github.com/pkg/errors"
)

type RegistrationUseCase struct {
	UserDataServiceInterface dataservice.UserDataInterface
}

func (ruc *RegistrationUseCase) SignUp(user *model.User) (*model.User, error) {
	logger.Log.Debug("Signing up user...")
	err := user.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "user validation failed")
	}
	isDup, err := ruc.isDuplicate(user.Email)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if isDup {
		return nil, errors.New("duplicate user for " + user.Email)
	}

	resultUser, err := ruc.UserDataServiceInterface.Insert(user)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return resultUser, nil
}

func (ruc *RegistrationUseCase) isDuplicate(email string) (bool, error) {
	user, err := ruc.UserDataServiceInterface.FindByEmail(email)
	if err != nil {
		return false, errors.Wrap(err, "")
	}
	if user != nil {
		return true, nil
	}
	return false, nil
}