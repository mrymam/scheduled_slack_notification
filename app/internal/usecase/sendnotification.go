package usecase

import (
	"context"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/domain/model"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/service/slack"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/setting"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/view"
)

type SendNotification interface {
	Do(context.Context, setting.Notification) error
}

func InitSendNotification() (SendNotification, error) {
	svc, err := slack.Init()
	if err != nil {
		return SendNotificationImple{}, err
	}

	return SendNotificationImple{
		slacksvc: svc,
	}, nil
}

type SendNotificationImple struct {
	slacksvc slack.Service
}

func (s SendNotificationImple) Do(ctx context.Context, sntf setting.Notification) error {
	ms := []model.Metric{}
	for _, m := range sntf.Metrics {
		ms = append(ms, model.Metric{
			Name:        m.Name,
			Description: m.Description,
			Value:       "100",
		})
	}

	ntf := model.Notification{
		Name:       sntf.Name,
		WebhookURL: sntf.WebhookURL,
		Schedule: model.Schedule{
			Name:        sntf.Schedule.Name,
			Description: sntf.Schedule.Description,
		},
		Metrics: ms,
	}

	msg, err := view.GenMessage(ntf)
	if err != nil {
		return err
	}

	s.slacksvc.PostMessage(ntf.WebhookURL, msg)
	return nil
}
