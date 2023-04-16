package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/NotKatsu/Discord-Mass-Reporter/helpers"
)

type userBase struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
}

type reportPayload struct {
	ChannelID string `json:"channel_id"`
	GuildID   string `json:"guild_id"`
	MessageID string `json:"message_id"`
	Reason    string `json:"reason"`
}

var baseURL = "https://discord.com/api/v9"

func getUser(authentication string) {
	req, err := http.NewRequest("GET", baseURL+"/users/@me", nil)
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

			var response userBase

			err = json.Unmarshal(body, &response)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("[LOGGED IN] " + response.Username + "#" + response.Discriminator + " (" + response.ID + ")")
			}
		}

	}
}

func sendReport(channelID, guildID, messageID, reason, authentication string) {
	rPayload := reportPayload{
		ChannelID: channelID,
		GuildID:   guildID,
		MessageID: messageID,
		Reason:    reason,
	}

	for count := 1; count == 0; count++ {
		jsonPayload, err := json.Marshal(rPayload)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		req, err := http.NewRequest("POST", baseURL+"/report", bytes.NewBuffer(jsonPayload))
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		req.Header.Set("Authorization", authentication)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusCreated {
			fmt.Printf("[REPORT] Successfully sent report #%d\n", count)
		}

		time.Sleep(3 * time.Second)
	}
}

func main() {
	var authentication, userID, channelID, guildID, messageID, reason string

	fmt.Printf("Authentication: ")
	fmt.Scanln(&authentication)

	fmt.Printf("Discord ID: ")
	fmt.Scanln(&userID)

	fmt.Printf("Channel ID: ")
	fmt.Scanln(&channelID)

	fmt.Printf("Guild ID: ")
	fmt.Scanln(&guildID)

	fmt.Printf("Message ID: ")
	fmt.Scanln(&messageID)

	fmt.Printf("Reason: ")
	fmt.Scanln(&reason)

	if helpers.IntCheck(userID) == true {
		getUser(authentication)
		sendReport(channelID, guildID, messageID, reason, authentication)
	}
}
