package http

import (
	"fmt"
	"testing"
	"time"
)

func TestHttp(t *testing.T) {

	channel := "https://hooks.slack.com/services/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"

	date := time.Now().Format("2 Jan 2006 15:04:05")
	msg := fmt.Sprintf("%s-%s", "hello world", date)

	_, err := SendHttpMessage(channel, msg)
	if err != nil {
		panic(err)
	}
}
