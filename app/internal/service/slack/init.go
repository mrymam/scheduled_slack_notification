package slack

import "github.com/onyanko-pon/scheduled_slack_notification/app/internal/infra/slack"

func Init(url string) (Service, error) {
	var clt slack.Client
	if true {
		clt = slack.ClientImpl{}
	} else {
		clt = slack.ClientDevImpl{}
	}

	return ServiceImpl{
		url: url,
		clt: clt,
	}, nil
}
