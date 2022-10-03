package query_test

import (
	"fmt"
	"testing"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/query"
	"github.com/stretchr/testify/assert"
)

func Test_LoadQuery(t *testing.T) {

	t.Run("loadQueryTest", func(t *testing.T) {

		querystr := query.LoadQuery("test.sql")

		assert.NotZero(t, len(querystr))
		fmt.Println(querystr)
	})
}
