package view

import (
	"embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/domain/model"
)

var (
	//go:embed templates/*.md
	files embed.FS
)

type data struct {
	Metrics map[string]string
}

func loadMessageBody(ntf model.Notification, templateName string) (string, error) {
	metrics := map[string]string{}
	for _, m := range ntf.Metrics {
		v, err := m.GetValueStr()
		if err != nil {
			continue
		}
		metrics[m.Name] = v
	}
	d := data{
		Metrics: metrics,
	}
	filepath := fmt.Sprintf("templates/%s", templateName)
	buf, err := files.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	str := string(buf)

	w := new(strings.Builder)
	t, err := template.New("").Parse(str)
	if err != nil {
		return "", err
	}
	err = t.Execute(w, d)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

func GenMessage(ntf model.Notification, templateName string) (Message, error) {
	body, err := loadMessageBody(ntf, templateName)
	if err != nil {
		return Message{}, err
	}
	b := SectionBlock{
		Object: TextBlockObject{
			TextType: TextBlockMarkdown,
			Text:     body,
		},
	}

	m := Message{
		Blocks: []Block{b},
	}
	return m, nil
}
