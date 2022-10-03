package slack

import (
	"github.com/slack-go/slack"
)

type Client interface {
	PostMessage(url string, msg slack.WebhookMessage) error
}

type ClientImpl struct{}

func (c ClientImpl) PostMessage(url string, msg slack.WebhookMessage) error {
	return slack.PostWebhook(url, &msg)
}
