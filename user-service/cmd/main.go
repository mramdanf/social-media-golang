package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"user-service/app"
)

type Config struct {}

func main() {
	configFileName := fmt.Sprintf("/app/config/appConfig%s.yaml", os.Getenv("ENV"))
	app.InitApp(configFileName)

	web := Config{}

	appPort := os.Getenv("USER_SERVICE_PORT")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", appPort),
		Handler: web.routes(),
	}

	log.Println("Server listening on port " + appPort)
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
		fmt.Println(err)
	}
	
}