package notify

import (
	"testing"
)

func TestSlack(t *testing.T) {
	SendSlackMessage("dev", "hello world")
}
