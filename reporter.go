package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NotKatsu/Discord-Mass-Reporter/helpers"
)

type user_base struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}

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
		} else {

			var response user_base

			err = json.Unmarshal(body, &response)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("[LOGGED IN] " + response.Username + "#" + response.Discriminator + " (" + response.ID + ")")
			}
		}

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
