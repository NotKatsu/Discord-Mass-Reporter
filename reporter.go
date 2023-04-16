package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NotKatsu/Discord-Mass-Reporter/helpers"
)

var base_url string = "https://discord.com/api/v9"

func user(user_token string) {
	req, err := http.NewRequest("GET", base_url+"/users/@me", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authorization", "MTA1Mjk4MjcyMTU5ODczODUyMg.GuquGC.tqwUdU7Hyn4EFom2QH31_T2_ZAOu2m1m1q7QxI")
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
