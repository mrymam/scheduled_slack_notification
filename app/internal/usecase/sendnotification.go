package usecase

import (
	"context"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/service/slack"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/setting"
)

type SendNotification interface{}

type SendNotificationImple struct {
	slacksvc slack.Service
}

func (s SendNotificationImple) Do(ctx context.Context, ntf setting.Notification) error {
	return nil
}
