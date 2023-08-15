package config

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	SQLConfig DataStoreConfig `yaml:"sqlConfig"`
}

type DataStoreConfig struct {
	Code string `yaml:"code"`
	DriverName string `yaml:"driverName"`
	UrlAddress string `yaml:"urlAddress"`
	DbName string `yaml:"dbName"`
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
	fmt.Println("appConfig:", ac)
	return &ac, nil
}