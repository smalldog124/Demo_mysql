package api

import (
	"Demo_mysql/model"
	"Demo_mysql/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type UserAPI struct {
	UserRepository repository.UserRepository
}

func (api UserAPI) CreatUserHandler(writer http.ResponseWriter, request *http.Request) {
	var newUser model.User
	body, err := ioutil.ReadAll(request.Body)
	if err = json.Unmarshal(body, &newUser); err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, err := api.UserRepository.CreateUser(newUser)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Context-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		log.Printf("error encodeing response %s", err.Error())
	}
}

func (api UserAPI) GetAllUserHandler(writer http.ResponseWriter, r *http.Request) {
	var user model.UserInfo
	users, err := api.UserRepository.GeatAllUser()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	user.User = users
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Context-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		log.Printf("error encodeing response %s", err.Error())
	}
}
