package http

import (
	"os"

	"github.com/nlopes/slack"
)

func SendSlackMessage(channel string, msg string) {

	s := ""

	os.Setenv("SLACK_TOKEN", s)
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	_, _, err := api.PostMessage("dev", slack.MsgOptionText("Yes, hello.", false))

	if err != nil {
		panic(err)
	}

}
