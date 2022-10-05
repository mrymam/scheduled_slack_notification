package usecase

import (
	"context"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/domain/model"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/query"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/service/bigquery"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/service/slack"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/setting"
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/view"
)

type SendNotification interface {
	Do(context.Context, setting.Notification) error
}

func InitSendNotification() (SendNotification, error) {
	ssvc, err := slack.Init()
	if err != nil {
		return SendNotificationImple{}, err
	}

	bqsvc, err := bigquery.Init()
	if err != nil {
		return SendNotificationImple{}, err
	}

	return SendNotificationImple{
		slacksvc: ssvc,
		bqsvc:    bqsvc,
	}, nil
}

type SendNotificationImple struct {
	slacksvc slack.Service
	bqsvc    bigquery.Service
}

func (s SendNotificationImple) Do(ctx context.Context, sntf setting.Notification) error {
	ms := []model.Metric{}
	for _, m := range sntf.Metrics {
		querystr := query.LoadQuery(m.Query.Filename)
		v, err := s.bqsvc.Read(querystr, m.Query.Key)
		if err != nil {
			return err
		}
		ms = append(ms, model.Metric{
			Name:        m.Name,
			Description: m.Description,
			Value:       v,
		})
	}

	ntf := model.Notification{
		Name:       sntf.Name,
		WebhookURL: sntf.WebhookURL,
		Schedule:   sntf.Schedule,
		Metrics:    ms,
	}

	msg, err := view.GenMessage(ntf, sntf.Template)
	if err != nil {
		return err
	}

	s.slacksvc.PostMessage(ntf.WebhookURL, msg)
	return nil
}
