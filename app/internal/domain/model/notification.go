package model

type Notification struct {
	Name       string
	WebhookURL string
	Schedule   string
	Metrics    []Metric
}
