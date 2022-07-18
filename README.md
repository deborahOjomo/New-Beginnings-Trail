# New Beginnings

Implement a participant registry microservice with an API that supports adding, updating, removing and retrieving personal information about participants in the study.

# How to run the app

`go run api/main.go `

# Dependencies

`go get github.com/gorilla/mux`

`go get github.com/stretchr/testify`

# How to run tests

Run all test at once

`go mod vendor `

`go test -mod vendor -race -cover -coverprofile=coverage.txt -covermode=atomic ./...`

Run package test

`cd endpoint`

`go test -v`

# API

Below is a list of API endpoints with their respective input and output.

## Create User

```
POST
/api/user
```

input

```json
{
	"reference": "P95OE8",
	"name": "Deb",
	"dob": "2022-01-01",
	"number": "02085955858",
	"address": {
		"postcode": "NW63NU",
		"address": "21 parsloes Avenue"
	}
}
```

#### Output

```json
{
	"reference": "P95OE8",
	"name": "Deb",
	"dob": "2022-01-01",
	"number": "02085955858",
	"address": {
		"postcode": "NW63NU",
		"address": "21 parsloes Avenue"
	}
}
```

## Get a single user

```
GET
/api/user/{reference}
```

#### Output

```json
{
	"reference": "P95OE8",
	"name": "Deb",
	"dob": "2022-01-01",
	"number": "02085955858",
	"address": {
		"postcode": "NW63NU",
		"address": "21 parsloes Avenue"
	}
}
```

## Update user information

```
PUT
/api/user/{reference}
```

input

```json
{
	"name": "Deb",
	"dob": "2022-01-01",
	"number": "02085955858",
	"address": {
		"postcode": "SW9Q6P",
		"address": "9 Warren Street"
	}
}
```

#### Output

```json
{
	"name": "Deb",
	"dob": "2022-01-01",
	"number": "02085955858",
	"address": {
		"postcode": "SW9Q6P",
		"address": "9 Warren Street"
	}
}
```

## Delete user

```
DELETE
/api/user/{reference}
```

#### Output

```json
 The user with ID P95OE8 has been deleted successfully
```

## List all user accounts

```
GET
/api/user
```

#### Output

```json
[
	{
		"name": "Deb",
		"dob": "2022-01-01",
		"number": "02085955858",
		"address": {
			"postcode": "SW9Q6P",
			"address": "9 Warren Street"
		}
	},
	{
		"name": "Cole",
		"dob": "2022-01-01",
		"number": "07956233761",
		"address": {
			"postcode": "E154ER",
			"address": "92 Baker Street"
		}
	}
]
```

## Areas to be improved

- Add a database
- Field validations e.g valid postcodes, phone number can be improved to cater various formats
- Service could be dockerized
