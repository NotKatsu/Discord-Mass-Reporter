package main

import (
	"fmt"
	"regexp"
)

var digitCheck = regexp.MustCompile(`^[0-9]+$`)

func main() {
	var user_id int

	fmt.Printf("Discord ID: ")
	fmt.Scanln(&user_id)

}
