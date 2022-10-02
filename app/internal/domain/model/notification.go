package domain

type Notification struct {
	Name       string
	WebhookURL string
	Schedule   Schedule
	Metrics    []Metric
}

type Schedule struct {
	Name        string
	Description string
}
