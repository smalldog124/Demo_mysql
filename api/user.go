package api

import (
	"Demo_mysql/model"
	"Demo_mysql/repository"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
		log.Printf("error encodeing CreatUserHandler response %s", err.Error())
	}
}

func (api UserAPI) GetAllUserHandler(writer http.ResponseWriter, request *http.Request) {
	var user model.UserInfo
	users, err := api.UserRepository.GetAllUser()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user.User = users
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Context-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		log.Printf("error encodeing GetAllUserHandler response %s", err.Error())
	}
}

func (api UserAPI) GetUserByIdHandler(writer http.ResponseWriter, request *http.Request) {
	palamiter := mux.Vars(request)
	userID := palamiter["id"]
	user, err := api.UserRepository.GetUserById(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Context-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		log.Printf("error encodeing GetUserByIdHandler respondr %s", err.Error())
	}
}

func (api UserAPI) EditUserHandler(writer http.ResponseWriter, request *http.Request) {
	var dataUser model.User
	palamiter := mux.Vars(request)
	userID := palamiter["id"]
	body, err := ioutil.ReadAll(request.Body)
	if err = json.Unmarshal(body, &dataUser); err != nil {
		log.Printf("error decodeing EditeUserHandler %s", err.Error())
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	user, err := api.UserRepository.EditeUser(userID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Context-Type", "applcaion/json")
	if err := json.NewEncoder(writer).Encode(user); err != nil {
		log.Printf("error encodeing EditeUserHandler respondr %s", err.Error())
	}
}
