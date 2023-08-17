package usecasefactory

import (
	"user-service/app/config"
	"user-service/app/container"
	"user-service/app/container/dataservicefactory"
	"user-service/applicationservice/dataservice"

	"github.com/pkg/errors"
)

func buildUserData(c container.ContainerInterface, dc *config.DataConfig) (dataservice.UserDataInterface, error) {
	dsi, err := dataservicefactory.GetDataServiceFb(dc.Code).Build(c, dc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	udi := dsi.(dataservice.UserDataInterface)
	return udi, nil
}