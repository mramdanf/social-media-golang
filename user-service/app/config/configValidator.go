package config

import (
	"github.com/pkg/errors"
)

const (
	SQLDB      string = "sqldb"
)

func validateConfig(appConfig AppConfig) error {
	err := validateDataStore(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateDataStore(appConfig AppConfig) error {
	sc := appConfig.SQLConfig
	key := sc.Code
	scMsg := " in validateDataStore doesn't match key = "
	if SQLDB != key {
		errMsg := SQLDB + scMsg + key
		return errors.New(errMsg)
	}
	return nil
}