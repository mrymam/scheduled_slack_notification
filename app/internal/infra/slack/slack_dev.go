package slack

import (
	"fmt"

	"github.com/slack-go/slack"
)

type ClientDevImpl struct{}

func (c ClientDevImpl) PostMessage(url string, msg slack.WebhookMessage) error {
	fmt.Printf("message to %s.\n", url)
	return nil
}
