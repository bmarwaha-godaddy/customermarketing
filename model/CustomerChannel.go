package model

import "fmt"

type CustomerChannel struct {
	Name         string `json:"name"`
	Identifier   string `json:"identifier"`
	LoggedInFrom string `json:"loggedInFrom"`
	CustomerId   string `json:"customerId"`
}
type Channel struct {
	Name         string `dynamodbav:"name"`
	Identifier   string `dynamodbav:"identifier"`
	LoggedInFrom string `dynamodbav:"loggedInFrom"`
	CustomerId   string `dynamodbav:"customerId"`
}

func (customer CustomerChannel) printCustomerChannel() {

	fmt.Printf(" %s is CustomerId lastLogged in %s", customer.CustomerId, customer.LoggedInFrom)
}
