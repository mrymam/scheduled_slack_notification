package slack

import (
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/infra/slack"
	"github.com/onyanko-pon/scheduled_slack_notification/app/pkg/config"
)

func Init(url string) (Service, error) {
	var clt slack.Client

	if config.GetEnv().IsProd() || config.GetEnv().IsDev() {
		clt = slack.ClientImpl{}
	} else {
		clt = slack.ClientDevImpl{}
	}

	return ServiceImpl{
		url: url,
		clt: clt,
	}, nil
}
