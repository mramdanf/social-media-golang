package main

import (
	"net/http"
	"time"
	"user-service/app/container/containerhelper"
	"user-service/app/logger"
	"user-service/app/utils"
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

	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	rcui, err := containerhelper.GetRegistrationUseCase(us.container)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		return
	}

	hashedPassword, err := utils.HashPassword(requestPayload.Password)
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
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := utils.JsonResponse {
		Error: false,
		Data: *createdUser,
	}

	utils.WriteJSON(w, http.StatusAccepted, responsePayload)
}

func (us *UserService) login(w http.ResponseWriter, r *http.Request) {
	var requestPayload struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}
	err := utils.ReadJSON(w, r, &requestPayload)
	if err != nil {
		utils.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	credentialUseCase, err := containerhelper.GetCredentialUseCase(us.container)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	user, err := credentialUseCase.Login(requestPayload.Email, requestPayload.Password)
	if err != nil {
		logger.Log.Errorf("%+v\n", err)
		utils.ErrorJSON(w, err, http.StatusInternalServerError)
		return
	}

	responsePayload := utils.JsonResponse {
		Error: false,
		Data: user,
	}

	utils.WriteJSON(w, http.StatusOK, responsePayload)

}