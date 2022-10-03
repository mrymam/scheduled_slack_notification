package slack

import (
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/infra/slack"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/view"
)

type Service interface {
	PostMessage(url string, msg view.Message) error
}

type ServiceImpl struct {
	clt slack.Client
}

func (c ServiceImpl) PostMessage(url string, vmsg view.Message) error {
	msg := vmsg.GenWebHookMessage()
	return c.clt.PostMessage(url, msg)
}
