package datastorefactory

import (
	"database/sql"
	"fmt"
	"user-service/app/config"
	"user-service/app/container"
	"user-service/app/logger"

	databaseConfig "github.com/jfeng45/gtransaction/config"
	"github.com/jfeng45/gtransaction/factory"
	"github.com/jfeng45/gtransaction/gdbc"
)

type sqlFactory struct {}

func (sf *sqlFactory) Build(c container.ContainerInterface, dsc *config.DataStoreConfig) (DataStoreInterface, error) {
	logger.Log.Debug("sqlFactory")
	key := dsc.Code

	// only non-transaction connection is cached
	if !dsc.Tx {
		if value, found := c.Get(key); found {
			logger.Log.Debug("found db in container for key: " + key)
			return value, nil
		}
	}
	tdbc := databaseConfig.DatabaseConfig{DriverName: dsc.DriverName, DataSourceName: dsc.UrlAddress, Tx: dsc.Tx}
	fmt.Println(tdbc)
	db, err := factory.BuildSqlDB(&tdbc)
	if err != nil {
		return nil, err
	}
	mygdbc, err := buildGdbc(db, dsc.Tx)
	if err != nil {
		return nil, err
	}
	// only non-transaction connection is cached
	if !dsc.Tx {
		c.Put(key, mygdbc)
	}
	return mygdbc, nil
}

func buildGdbc(sdb *sql.DB, tx bool) (gdbc.SqlGdbc, error) {
	var sdt gdbc.SqlGdbc
	if tx {
		tx, err := sdb.Begin()
		if err != nil {
			return nil, err
		}
		sdt = &gdbc.SqlConnTx{DB: tx}
		logger.Log.Debug("buildGdbc(), create TX:")
	} else {
		sdt = &gdbc.SqlDBTx{DB: sdb}
		logger.Log.Debug("buildGdbc(), create DB:")
	}
	return sdt, nil
}