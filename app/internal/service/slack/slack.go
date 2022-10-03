package slack

import (
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/infra/slack"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/view"
)

type Service interface {
	PostMessage(msg view.Message) error
}

type ServiceImpl struct {
	url string
	clt slack.Client
}

func (c ServiceImpl) PostMessage(vmsg view.Message) error {
	msg := vmsg.GenWebHookMessage()
	return c.clt.PostMessage(c.url, msg)
}
