package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"user-service/app"
	"user-service/app/container/servicecontainer"

	"github.com/pkg/errors"
)

type UserService struct {
	container *servicecontainer.ServiceContainer
}

func main() {
	configFileName := fmt.Sprintf("/app/config/appConfig%s.yaml", os.Getenv("ENV"))
	container, err := buildContainer(configFileName)
	if err != nil {
		fmt.Printf("%+v\n", err)
		panic(err)
	}

	userService := &UserService{container: container}

	appPort := os.Getenv("USER_SERVICE_PORT")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", appPort),
		Handler: userService.routes(),
	}

	log.Println("Server listening on port " + appPort)
	err = srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
		fmt.Println(err)
	}
	
}



func buildContainer(filename string) (*servicecontainer.ServiceContainer, error) {
	container, err := app.InitApp(filename)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	sc := container.(*servicecontainer.ServiceContainer)
	return sc, nil
} 