package main

import "fmt"

func main() {
	emptyMap := map[string]string{}
	user := map[string]string{ //map[key type]value type
		"name":    "Renan",
		"surname": "Cavalcante",
	}
	fmt.Println(emptyMap, user, user["name"])

	delete(user, "surname") //Remove a key/value pair
	fmt.Println(user)

	user["zodiac"] = "Aquarius"
	fmt.Println(user)

	nestedUser := map[string]map[string]string{
		"name": {
			"first": "John",
			"last":  "Doe",
		},
	}
	getNestedValue := nestedUser["name"]
	_, hasMiddle := getNestedValue["middle"] //Check is map has key
	fmt.Println(nestedUser, nestedUser["name"],
		getNestedValue["first"], hasMiddle)
}
