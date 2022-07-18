package endpoint

import (
	"encoding/json"
	"fmt"
	"grail-api/model"
	"io/ioutil"
	"net/http"

	"math/rand"

	"github.com/gorilla/mux"
)

type account []model.Account

var address = model.Address{
	Postcode: "E159N",
	Address:  "56 Warren Street",
}
var userAccounts = account{
	{
		Reference: "6FBBD7",
		Name:      "Sue",
		Dob:       "01-02-2002",
		Number:    "0123456789",
		Address:   address,
	},
}

func GenerateRefID(n int) string {
	const charset = "1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

// Create user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.Account

	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Reference = GenerateRefID(6)

	userAccounts = append(userAccounts, user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Get all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(userAccounts)
}

// Get single user
func GetUser(w http.ResponseWriter, r *http.Request) {

	// Get the ref ID from the url
	ref := mux.Vars(r)["reference"]

	for _, singleUser := range userAccounts {
		if singleUser.Reference == ref {
			json.NewEncoder(w).Encode(singleUser)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User with reference " + ref + " not found")
}

// Updating account
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Get the ref ID from the url
	ref := mux.Vars(r)["reference"]

	var updatedUser model.Account
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedUser)

	for i, singleUser := range userAccounts {
		if singleUser.Reference == ref {
			singleUser.Dob = updatedUser.Dob
			singleUser.Number = updatedUser.Number
			singleUser.Address.Address = updatedUser.Address.Address
			singleUser.Address.Postcode = updatedUser.Address.Postcode

			userAccounts = append(userAccounts[:i], singleUser)
			json.NewEncoder(w).Encode(singleUser)
			w.WriteHeader(http.StatusOK)

		}
	}
}

// Delete user
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ref := mux.Vars(r)["reference"]

	for i, singleUser := range userAccounts {
		if singleUser.Reference == ref {
			userAccounts = append(userAccounts[:i], userAccounts[i+1:]...)
			fmt.Fprintf(w, "The user with ID %v has been deleted successfully", ref)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("User with reference " + ref + " not found")
}
