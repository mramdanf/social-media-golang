package config

import (
	"fmt"
	"os"

	logConfig "github.com/jfeng45/glogger/config"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// AppConfig represents the application config
type AppConfig struct {
	SQLConfig DataStoreConfig `yaml:"sqlConfig"`
	LogConfig logConfig.Logging `yaml:"logConfig"`
	ZapConfig logConfig.Logging `yaml:"zapConfig"`
	UseCaseConfig UseCaseConfig `yaml:"useCaseConfig"`
}

// DataConfig represents data service
type DataConfig struct {
	Code string `yaml:"code"`
	DataStoreConfig DataStoreConfig `yaml:"dataStoreConfig"`
}

// DataStoreConfig represents handlers for data store. It can be a database or a gRPC connection
type DataStoreConfig struct {
	Code string `yaml:"code"`
	DriverName string `yaml:"driverName"`
	UrlAddress string `yaml:"urlAddress"`
	DbName string `yaml:"dbName"`
	Tx bool `yaml:"tx"`
}

// UseCaseConfig represents different use cases
type UseCaseConfig struct {
	Registration RegistrationConfig `yaml:"registration"`
}

// RegistrationConfig represents registration use case
type RegistrationConfig struct {
	Code string `yaml:"code"`
	UserDataConfig DataConfig `yaml:"userDataConfig"`
}

func BuildConfig(filename ...string) (*AppConfig, error) {
	return buildConfigFromFile(filename[0])
}

func buildConfigFromFile(filename string) (*AppConfig, error) {
	var ac AppConfig
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "read error")
	}
	err = yaml.Unmarshal(file, &ac)

	if err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}
	err = validateConfig(ac)
	if err != nil {
		return nil, errors.Wrap(err, "validate config")
	}

	mysqlUserName := os.Getenv("USER_SERVICE_MYSQL_USERNAME")
	mysqlPassword := os.Getenv("USER_SERVICE_MYSQL_PASSWORD")
	mysqlDbName := os.Getenv("USER_SERVICE_MYSQL_DB_NAME")
	mysqlPort := os.Getenv("USER_SERVICE_MYSQL_PORT")
	mysqlHost := os.Getenv("USER_SERVICE_MYSQL_HOST")
	ac.UseCaseConfig.Registration.UserDataConfig.DataStoreConfig.UrlAddress = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mysqlUserName, mysqlPassword, mysqlHost, mysqlPort, mysqlDbName)

	fmt.Println("appConfig:", ac)
	return &ac, nil
}