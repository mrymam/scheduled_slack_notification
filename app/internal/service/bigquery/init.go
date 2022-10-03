package bigquery

import (
	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/infra/bigquery"
)

func Init() (Service, error) {
	var clt bigquery.Client

	clt, err := bigquery.InitClient()
	if err != nil {
		return nil, err
	}

	return ServiceImpl{
		clt: clt,
	}, nil
}
