package config

import (
	"github.com/pkg/errors"
)

// database code. Need to map to the database code (DataStoreConfig) in the configuration yaml file.
const (
	SQLDB      string = "sqldb"
)

// constant for logger code, it needs to match log code (logConfig)in configuration
const (
	ZAP    string = "zap"
)

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
const (
	REGISTRATION string = "registration"
)

// data service code. Need to map to the data service code (DataConfig) in the configuration yaml file.
const (
	USER_DATA string = "userData"
)

func validateConfig(appConfig AppConfig) error {
	err := validateDataStore(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = validateLogger(appConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	useCase := appConfig.UseCaseConfig
	err = validateUseCase(useCase)
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

func validateLogger(appConfig AppConfig) error {
	zc := appConfig.ZapConfig
	key := zc.Code
	zcMsg := " in validateLogger doesn't match key = "
	if ZAP != key {
		errMsg := ZAP + zcMsg + key
		return errors.New(errMsg)
	}
	return nil
}

func validateUseCase(useCaseConfig UseCaseConfig) error {
	err := validateRegistration(useCaseConfig)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}

func validateRegistration(useCaseConfig UseCaseConfig) error {
	rc := useCaseConfig.Registration
	key := rc.Code
	rcMsg := " in validateRegistration doens't match key = "
	if REGISTRATION != key {
		errMsg := REGISTRATION + rcMsg + key
		return errors.New(errMsg)
	}
	key = rc.UserDataConfig.Code
	if USER_DATA != key {
		errMsg := USER_DATA + rcMsg + key
		return errors.New(errMsg)
	}
	return nil
}