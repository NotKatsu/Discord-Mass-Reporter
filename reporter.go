package main

import (
	"fmt"
	"regexp"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func main() {
	var user_id string

	fmt.Printf("Discord ID: ")
	fmt.Scanln(&user_id)

	if digitCheck.MatchString(user_id) == true {
		fmt.Println("Is Number.")
	}

}
