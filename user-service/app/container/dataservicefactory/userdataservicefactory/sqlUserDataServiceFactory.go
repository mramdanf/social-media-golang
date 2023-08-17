package userdataservicefactory

import (
	"user-service/app/config"
	"user-service/app/container"
	"user-service/app/container/datastorefactory"
	"user-service/app/logger"
	"user-service/applicationservice/dataservice"
	"user-service/applicationservice/dataservice/userdata/sqldb"

	"github.com/jfeng45/gtransaction/gdbc"
	"github.com/pkg/errors"
)

type sqlUserDataServiceFactory struct {}

func (sudsf *sqlUserDataServiceFactory) Build(c container.ContainerInterface, dataConfig *config.DataConfig) (dataservice.UserDataInterface, error) {
	logger.Log.Debug("sqlUserDataServiceFactory")
	dsc := dataConfig.DataStoreConfig
	dsi, err := datastorefactory.GetDataStoreFb(dsc.Code).Build(c, &dsc)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	ds := dsi.(gdbc.SqlGdbc)
	uds := sqldb.UserDataSql{DB: ds}
	logger.Log.Debug("uds: ", uds.DB)
	return &uds, nil
}