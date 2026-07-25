package main

import (
	"fmt"

	"github.com/SiddharthAJMore/golang-tutorial/auth"
	"github.com/SiddharthAJMore/golang-tutorial/user"
	//"github.com/fatih/color"
)

func main() {
	auth.LoginWithCredentials("siddharth", "pas123")
	session := auth.GetSession()
	fmt.Println("session", session)

	user1 := user.User{
		Email: "Test@gmail.com",
		Name:  "Siddharth",
	}

	fmt.Println(user1)
	//color.red("hi")
}
