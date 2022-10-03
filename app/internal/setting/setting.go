package setting

type Setting struct {
	Notifications []Notification `yaml:"notifications"`
}

type Notification struct {
	WebhookURL string   `yaml:"webhook_url"`
	Schedule   Schedule `yaml:"schedule"`
	Metrics    []Metric `yaml:"metrics"`
}

type Schedule struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Cron        string `yaml:"cron"`
}

type Metric struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Query       Query  `yaml:"query"`
}

type Query struct {
	Filepath string `yaml:"filepath"`
}
