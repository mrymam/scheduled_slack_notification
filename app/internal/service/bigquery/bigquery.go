package bigquery

import (
	"fmt"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/infra/bigquery"
)

type Service interface {
	Read(query string, key string) (interface{}, error)
}

type ServiceImpl struct {
	clt bigquery.Client
}

func (c ServiceImpl) Read(query string, key string) (interface{}, error) {
	m, err := c.clt.Read(query)
	if err != nil {
		return "", err
	}
	v, ok := m[key]
	if !ok {
		return "", fmt.Errorf("key not found: %s", key)
	}
	return v, nil
}
