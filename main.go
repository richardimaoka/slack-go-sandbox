package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_TOKEN"))

	channels, _, err := api.GetConversations(&slack.GetConversationsParameters{})
	if err != nil {
		panic(err)
	}
	for _, c := range channels {
		bytes, err := json.MarshalIndent(c, "", "  ")
		if err != nil {
			fmt.Printf("JSON marshal error = %v\n", err)
			continue
		}
		fmt.Println(string(bytes))
	}

	email := os.Getenv("USER_EMAIL")
	if email == "" {
		panic("USER_EMAIL is not set")
	}

	user, err := api.GetUserByEmail(email)
	if err != nil {
		panic(err)
	}

	channel, timestamp, text, err := api.SendMessage(user.ID, slack.MsgOptionText(fmt.Sprintf("Hello %s! https://google.com", email), false))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Message successfully sent to channel %s at %s: %s\n", channel, timestamp, text)

}
