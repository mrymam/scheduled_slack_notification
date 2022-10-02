package main

import (
	"github.com/onyanko-pon/scheduled_slack_notification/app/pkg/infra/slack"
)

func main() {

	clt := slack.ClientImpl{}
	msg := slack.Message{
		Text: "hello world",
	}
	clt.PostMessage("url", msg)
}
