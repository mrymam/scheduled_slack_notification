package bigquery

import (
	"fmt"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/infra/bigquery"
)

type Service interface {
	Read(query string) (string, error)
}

type ServiceImpl struct {
	clt bigquery.Client
}

func (c ServiceImpl) Read(query string) (string, error) {
	vs, err := c.clt.Read(query)
	if err != nil {
		return "", err
	}
	if len(vs) == 0 {
		return "", fmt.Errorf("lenght 0")
	}

	v, ok := vs[0].(string)
	if !ok {
		return "", fmt.Errorf("cast error")
	}
	return v, nil
}
