package main

import (
	"net/http"
	"time"
	"user-service/app/container/containerhelper"
	"user-service/app/logger"
	"user-service/domain/model"
)

func catchPanic() {
	if p := recover(); p != nil {
		logger.Log.Errorf("%+v\n", p)
	}
}

func (us *UserService) signUp(w http.ResponseWriter, r *http.Request) {
	defer catchPanic()
	var requestPayload struct {
		FullName string `json:"fullName"`
		Email string `json:"email"`
		Password string `json:"password"`
	}

	err := us.readJSON(w, r, &requestPayload)
	if err != nil {
		us.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	rcui, err := containerhelper.GetRegistrationUseCase(us.container)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return
	}

	hashedPassword, err := us.hashPassword(requestPayload.Password)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return
	}

	user := model.User{
		FullName: requestPayload.FullName,
		Email: requestPayload.Email,
		Password: hashedPassword,
		Created: time.Now(),
	}

	createdUser, err := rcui.SignUp(&user)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return
	}

	responsePayload := jsonResponse {
		Error: false,
		Data: *createdUser,
	}

	us.writeJSON(w, http.StatusAccepted, responsePayload)
}