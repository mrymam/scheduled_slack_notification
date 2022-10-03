package bigquery

import (
	"context"

	"cloud.google.com/go/bigquery"
	"github.com/onyanko-pon/scheduled_slack_notification/app/pkg/config"
)

type Client interface {
	Read(query string) (map[string]bigquery.Value, error)
}

type ClientImpl struct {
	clt *bigquery.Client
}

func InitClient() (Client, error) {
	ctx := context.Background()
	gcp := config.GetGCP()
	clt, err := bigquery.NewClient(ctx, gcp.ProjectID)
	if err != nil {
		return ClientImpl{}, err
	}
	return ClientImpl{
		clt: clt,
	}, nil
}

func (c ClientImpl) Read(query string) (map[string]bigquery.Value, error) {
	ctx := context.Background()

	it, err := c.clt.Query(query).Read(ctx)
	if err != nil {
		return nil, err
	}

	var m map[string]bigquery.Value
	err = it.Next(&m)
	if err != nil {
		return nil, err
	}

	return m, nil
}
