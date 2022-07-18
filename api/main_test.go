package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"grail-api/endpoint"

	"github.com/stretchr/testify/assert"
)

func TestCreateUserStatusCreated(t *testing.T) {

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

	req.Header.Add("Content-Type", "application/json")

	//send request
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(endpoint.CreateUser)
	handler.ServeHTTP(responseRecorder, req)

	assert.Equal(t, http.StatusCreated, responseRecorder.Code)
}
