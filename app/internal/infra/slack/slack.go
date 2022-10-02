package slack

import "github.com/onyanko-pon/scheduled_slack_notification/app/pkg/infra/slack"

type Client interface {
	PostMessage(msg Message) error
}

type ClientImpl struct {
	url string
	clt slack.Client
}

func (c ClientImpl) PostMessage(msg Message) error {
	c.clt.PostMessage(c.url, slack.Message(msg))
	return nil
}
