package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NotKatsu/Discord-Mass-Reporter/helpers"
)

var base_url string = "https://discord.com/api/v9"

func user(authentication string) {
	req, err := http.NewRequest("GET", base_url+"/users/@me", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", authentication)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
	} else {

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(string(body))
	}
}

func main() {
	var user_id string
	var authentication string

	fmt.Printf("Authentication: ")
	fmt.Scanln(&authentication)

	fmt.Printf("Discord ID: ")
	fmt.Scanln(&user_id)

	if helpers.IntCheck(user_id) == true {
		user(authentication)
	}
}
