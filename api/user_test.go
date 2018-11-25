package api_test

import (
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
	expected := `{"user_id":"1","first_name":"Smalldog","last_name":"Adison","addess":"123 californear","phone_number":"092-3994-212"}`
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
	userAPI := UserAPI{
		UserRepository: &mockUserRepository{},
	}
	request := httptest.NewRequest("POST", "/api/v1/user", bytes.NewBuffer(userJson))
	writer := httptest.NewRecorder()

	testRoute := mux.NewRouter()
	testRoute.HandleFunc("/api/v1/user", api.CreateUserHandler).Methods("POST")
	testRoute.ServeHTTP(writer, request)

	response := writer.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, expected, string(actual))
}
