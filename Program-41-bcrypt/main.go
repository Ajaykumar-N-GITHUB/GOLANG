package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "@j@yAK1234567"

	bs, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(password)
	fmt.Println(bs)

	password = "@j@yAK12345"

	err = bcrypt.CompareHashAndPassword(bs, []byte(password))
	if err != nil {
		fmt.Println("Wrong password!!!!")
	} else {
		fmt.Println("Welocme!!!")
	}

}
