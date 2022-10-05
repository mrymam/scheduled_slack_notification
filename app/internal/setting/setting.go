package setting

type Setting struct {
	Notifications []Notification `yaml:"notifications"`
}

type Notification struct {
	Name       string   `yaml:"name"`
	WebhookURL string   `yaml:"webhook_url"`
	Schedule   string   `yaml:"schedule"`
	Template   string   `yaml:"template"`
	Metrics    []Metric `yaml:"metrics"`
}

type Metric struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	Query       Query  `yaml:"query"`
}

type Query struct {
	Filename string `yaml:"filename"`
	Key      string `yaml:"key"`
}
