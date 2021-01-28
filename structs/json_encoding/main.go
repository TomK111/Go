package main

import (
	"encoding/json"
	"fmt"
)

type permissions map[string]bool

type user struct {
	Name        string      `json:"username"`
	Password    string      `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main() {
	//Encoding JSON is called marshalling.
	//Decoding JSON is called unmarshalling.
	users := []user{
		{"Tomislav", "1234", nil},
		{"Tommy", "5678", permissions{"admin": true}},
		{"John", "0001", permissions{"write": true}},
	}

	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(out))
}
