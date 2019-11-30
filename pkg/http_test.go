package http

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestHttp(t *testing.T) {

	channel := os.Getenv("SLACK_CHANNEL")

	date := time.Now().Format("2 Jan 2006 15:04:05")
	msg := fmt.Sprintf("%s-%s", "hello world", date)

	_, err := SendHttpMessage(channel, msg)
	if err != nil {
		panic(err)
	}
}
