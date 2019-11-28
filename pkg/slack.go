package notify

import (
	"os"

	"github.com/nlopes/slack"
)

func SendSlackMessage(channel string, msg string) {

	s := "a8db9ade47d8bc58cf524afe0cdf2032"

	os.Setenv("SLACK_TOKEN", s)
	api := slack.New(os.Getenv("SLACK_TOKEN"))
	_, _, err := api.PostMessage("dev", slack.MsgOptionText("Yes, hello.", false))

	if err != nil {
		panic(err)
	}

}
