package app

import (
	"fmt"
	"user-service/app/config"
)

func InitApp(filename ...string) {
	_, err := config.BuildConfig(filename...)
	if err != nil {
		fmt.Println(err)
	}
}