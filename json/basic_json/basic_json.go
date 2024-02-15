package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func main() {
	jsonStr := `{"name":"Marco", "surname":"Introini"}`
	var user User
	err := json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user.Name)

	jsonStr = `{"name":5, "surname":"Introini"}`
	var user2 User
	err = json.Unmarshal([]byte(jsonStr), &user2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(user2.Name)

	u2 := User{
		Name:    "Test Name",
		Surname: "My Surname",
	}

	jsonBytes, err := json.Marshal(u2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}
