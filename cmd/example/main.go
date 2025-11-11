package main

import (
	"encoding/json"
	"fmt"

	"github.com/veridian-sky/gocsf"
)

func main() {
	// Create an Account object
	account := gocsf.Account{
		Name:   "Production Account",
		Uid:    "acc-456",
		Type:   "Cloud",
		TypeId: 2,
	}

	// Create a User object
	user := gocsf.User{
		Name:      "john.doe",
		FullName:  "John Doe",
		EmailAddr: "john.doe@example.com",
		Uid:       "user-123",
		Type:      "Regular",
		TypeId:    1,
		Account:   &account, // Use pointer
	}

	// Create an Actor object
	actor := gocsf.Actor{
		User:    &user, // Use pointer
		AppName: "MyApp",
		AppUid:  "app-789",
	}

	// Serialize to JSON
	data, err := json.MarshalIndent(actor, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println("OCSF Actor JSON:")
	fmt.Println(string(data))
}
