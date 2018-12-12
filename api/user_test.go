package api_test

import (
	"Demo_mysql/api"
	"Demo_mysql/model"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Test_CreateUserHandler_Input_User_Should_Be_User(t *testing.T) {
	expectedUserID := `{"user_id":1}
`
	fixedTime, _ := time.Parse("2006-Jan-22", "2018-Oct-25")
	user := model.User{
		UserID:      "1",
		FristName:   "Smalldog",
		LastName:    "Adison",
		Address:     "123 californear",
		PhoneNumber: "092-3994-212",
		CreatedTime: fixedTime,
		UpdatedTime: fixedTime,
	}
	userJson, _ := json.Marshal(user)
	userAPI := api.UserAPI{
		UserRepository: &mockUserRepository{},
	}
	request := httptest.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(userJson))
	writer := httptest.NewRecorder()

	testRoute := mux.NewRouter()
	testRoute.HandleFunc("/api/v1/user", userAPI.CreatUserHandler).Methods("POST")
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, expectedUserID, string(actual))
}

func Test_GetAllUserHandler_Shoudl_Be_UserInfo(t *testing.T) {
	expected := `{"user":[{"user_id":"1","first_name":"Smalldog","last_name":"Adison","addess":"123 californear","phone_number":"092-3994-212","created_time":"0001-01-01T00:00:00Z","updated_time":"0001-01-01T00:00:00Z"}]}
`
	request := httptest.NewRequest("GET", "/api/v1/user", nil)
	writer := httptest.NewRecorder()
	userAPI := api.UserAPI{
		UserRepository: &mockUserRepository{},
	}

	testRoute := mux.NewRouter()
	testRoute.HandleFunc("/api/v1/user", userAPI.GetAllUserHandler).Methods("GET")
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, string(actual))
}

func Test_GetUserByIdHandler_Input_ID_1_Should_Be_User_Name_Smalldog(t *testing.T) {
	expected := `{"user_id":"1","first_name":"Smalldog","last_name":"Adison","addess":"123 californear","phone_number":"092-3994-212","created_time":"0001-01-01T00:00:00Z","updated_time":"0001-01-01T00:00:00Z"}
`
	request := httptest.NewRequest("GET", "/api/v1/user/1", nil)
	writer := httptest.NewRecorder()
	userAPI := api.UserAPI{
		UserRepository: &mockUserRepository{},
	}

	testRoute := mux.NewRouter()
	testRoute.HandleFunc("/api/v1/user/{id}", userAPI.GetUserByIdHandler)
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, string(actual))
}

func Test_EditUserHandler_Input_User_ID_1_And_Phone_Number_0984772211_Shoud_Be_User_Edited(t *testing.T) {
	expected := `{"user_id":"1","first_name":"Lek","last_name":"Adison","addess":"123 californear","phone_number":"098-4772-211","created_time":"0001-01-01T00:00:00Z","updated_time":"0001-01-01T00:00:00Z"}
`
	user := `{"user_id":"1","first_name":"Lek","last_name":"Adison","addess":"123 californear","phone_number":"098-4772-211","created_time":"0001-01-01T00:00:00Z","updated_time":"0001-01-01T00:00:00Z"}`
	request := httptest.NewRequest("PUT", "/api/v1/user/1", bytes.NewBufferString(user))
	writer := httptest.NewRecorder()
	userAPI := api.UserAPI{
		UserRepository: &mockUserRepository{},
	}

	testRoute := mux.NewRouter()
	testRoute.HandleFunc("/api/v1/user/{id}", userAPI.EditUserHandler).Methods("PUT")
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, expected, string(actual))
}

func Test_DeleteUserHandler_Input_User_ID_1_And_Phone_Number_0984772211_Shoud_Be_User_Edited(t *testing.T) {
	request := httptest.NewRequest("DELETE", "/api/v1/user/1", nil)
	writer := httptest.NewRecorder()
	userAPI := api.UserAPI{
		UserRepository: &mockUserRepository{},
	}

	testRoute := mux.NewRouter()
	testRoute.HandleFunc("/api/v1/user/{id}", userAPI.DeleteUserHandler).Methods("DELETE")
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()

	assert.Equal(t, http.StatusOK, response.StatusCode)
}
