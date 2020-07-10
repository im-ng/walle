//WallE acts as bot and pushes stream of messages to Slack channel
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
)

const (
	usage       = `Run as : Walle <text.message>`
	botname     = "WallE"
	token       = ""
	fieldPrefix = ""
)

var (
	channel  = "welcome"
	color    = ""
	icon     = ""
	message  = "Hi there!"
	username = ""
)

func main() {
	if len(os.Args) == 1 {
		log.Println("Run as : walle <text.message>")
		os.Exit(2)
	}
	//
	apiClient := slack.New(token)
	//
	text := os.Args[1]
	//
	headerText := slack.NewTextBlockObject("mrkdwn", text, false, false)
	headerSection := slack.NewSectionBlock(headerText, nil, nil)
	blockOpt := slack.MsgOptionBlocks(headerSection)
	fallbackOpt := slack.MsgOptionText(text, false)
	msg := slack.MsgOptionCompose(fallbackOpt, blockOpt)
	//
	if _, _, err := apiClient.PostMessage(channel, msg); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to post a message to Slack: %s\n", err.Error())
		os.Exit(1)
	}
	log.Println("Message sent!")
}
