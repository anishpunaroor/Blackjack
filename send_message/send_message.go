package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("xoxb-2497620755441-2470350747111-Xr9vBE5SOPtkKIxHB3oensn8"))

	channelID, timestamp, err := api.PostMessage(
		"C02E5V0H7NZ",
		slack.MsgOptionText("hello world!", false),
	)

	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("Message sent successfully to channel %s at %s", channelID, timestamp)
}
