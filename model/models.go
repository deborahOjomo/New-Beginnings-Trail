package model

type Account struct {
	Reference string  `json:"reference"`
	Name      string  `json:"name" `
	Dob       string  `json:"dob"`
	Number    string  `json:"number"`
	Address   Address `json:"address"`
}

type Address struct {
	Postcode string `json:"postcode"`
	Address  string `json:"address"`
}
