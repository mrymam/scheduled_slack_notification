package setting_test

import (
	"testing"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/setting"
	"gotest.tools/assert"
)

func Test_LoadSetting(t *testing.T) {

	t.Run("loadSettingTest", func(t *testing.T) {
		s := setting.Get()

		assert.Equal(t, len(s.Notifications), 1)

		ntf := s.Notifications[0]
		assert.Equal(t, ntf.Name, "weekly_report_fourkeys")
		assert.Equal(t, ntf.WebhookURL, "url")
		assert.Equal(t, ntf.Schedule, "weekly")

		m := ntf.Metrics[0]
		assert.Equal(t, m.Name, "frequency")
		assert.Equal(t, m.Description, "deployments frequency")
		assert.Equal(t, m.Query.Filename, "test.sql")
	})
}

// notifications:
//   - name: "weekly_report_fourkeys"
//     webhook_url: "url"
//     schedule: "weekly"
//     metrics:
//       - name: "frequency"
//         description: "deployments frequency"
//         query:
//           filename: "frequency.sql"
//           key: "frequency"

// schedules:
//   - name: "weekly"
//     description: "every Monday at 9:00"
//     cron: "0 9 * * 1"
