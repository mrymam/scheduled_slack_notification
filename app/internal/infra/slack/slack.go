package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

type Client interface {
	PostMessage(url string, msg slack.WebhookMessage) error
}

type ClientImpl struct{}

func (c ClientImpl) PostMessage(url string, msg slack.WebhookMessage) error {
	fmt.Printf("message:%s to %s\n", "message", url)
	return nil
}
