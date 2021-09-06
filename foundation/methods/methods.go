package main

import "fmt"

type user struct {
	name string
	age  uint8
}

func (saveUser user) save() {
	fmt.Printf("Saving user %s on database\n", saveUser.name)
}

func (userAge user) isOfAge() bool {
	return userAge.age >= 18
}

func (userBirthday *user) itsYourBirthday() {
	userBirthday.age++
}

func main() {
	user1 := user{"Renan", 31}
	user1.save()
	println(user1.isOfAge())
	user1.itsYourBirthday()
	println(user1.age)
}
