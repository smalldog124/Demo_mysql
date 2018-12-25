package route

import (
	"Demo_mysql/api"
	"Demo_mysql/repository"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func UserRouter(route *mux.Router, connectionDB *sqlx.DB) {
	repository := repository.RepositoryMysql{
		ConnectionDB: connectionDB,
	}
	userAPI := api.UserAPI{
		UserRepository: repository,
	}
	route.HandleFunc("/api/v1/user", userAPI.CreatUserHandler).Methods("POST")
	route.HandleFunc("/api/v1/user", userAPI.GetAllUserHandler).Methods("GET")
	route.HandleFunc("/api/v1/user/{id}", userAPI.GetUserByIdHandler).Methods("GET")
	route.HandleFunc("/api/v1/user/{id}", userAPI.GetUserByIdHandler).Methods("PUT")
	route.HandleFunc("/api/v1/user/{id}", userAPI.DeleteUserHandler).Methods("DELETE")
}
