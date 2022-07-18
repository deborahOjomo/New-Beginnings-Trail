package endpoint

import (
	"bytes"
	"encoding/json"
	"grail-api/model"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {

	var jsonStr = []byte(`{
    "name": "Sue",
    "dob": "01-02-2002",
    "number": "0123456789",
    "address": {
        "postcode": "E159N",
        "address": "56 Warren Street"}}`)

	req, err := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	//send request
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}

func TestGetUser(t *testing.T) {

	vars := map[string]string{
		"reference": "6FBBD7",
	}

	// send request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/user/6FBBD7", nil)
	req = mux.SetURLVars(req, vars)
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(responseRecorder, req)

	expected, _ := json.Marshal(userDetails())

	result := responseRecorder.Result()

	actual, _ := ioutil.ReadAll(result.Body)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, string(expected), strings.TrimSpace(string(actual)))
}

func TestGetUserError(t *testing.T) {

	vars := map[string]string{
		"reference": "ASDF35",
	}

	// send request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/user/ASDF35", nil)
	req = mux.SetURLVars(req, vars)
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(result.Body)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	assert.ObjectsAreEqual("User with reference ASDF35 not found", string(actual))

}

func TestUpdateUser(t *testing.T) {

	var jsonStr = []byte(`{
    "name": "Sue",
    "dob": "2022-01-01",
    "number": "0987654321",
    "address": {
        "postcode": "W26FN",
        "address": "456 Grace Street"
    }}`)
	vars := map[string]string{
		"reference": "6FBBD7",
	}

	req, err := http.NewRequest(http.MethodPut, "/api/user/6FBBD7", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, vars)

	//send request
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateUser)
	handler.ServeHTTP(responseRecorder, req)

	var expected = []byte(`{
	"reference": "6FBBD7",
    "name": "Sue",
    "dob": "2022-01-01",
    "number": "0987654321",
    "address": {
        "postcode": "W26FN",
        "address": "456 Grace Street"}}`)

	result := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(result.Body)

	assert.ObjectsAreEqual(strings.TrimSpace(string(expected)), strings.TrimSpace(string(actual)))
	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestDeleteUser(t *testing.T) {

	vars := map[string]string{
		"reference": "6FBBD7",
	}

	// send request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/user/6FBBD7", nil)
	req = mux.SetURLVars(req, vars)
	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(result.Body)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.ObjectsAreEqual("The user with ID 6FBBD7 has been deleted successfully", string(actual))
}

func TestDeleteUserError(t *testing.T) {

	vars := map[string]string{
		"reference": "9FCCD9",
	}

	// send request
	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/user/9FCCD9", nil)
	req = mux.SetURLVars(req, vars)
	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(responseRecorder, req)

	result := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(result.Body)

	assert.Equal(t, http.StatusNotFound, responseRecorder.Code)
	assert.ObjectsAreEqual("User with reference 9FCCD7 not found", string(actual))
}

func userDetails() model.Account {
	address := model.Address{
		Postcode: "E159N",
		Address:  "56 Warren Street",
	}
	expectedMessage := model.Account{
		Reference: "6FBBD7",
		Name:      "Sue",
		Dob:       "01-02-2002",
		Number:    "0123456789",
		Address:   address,
	}
	return expectedMessage
}
