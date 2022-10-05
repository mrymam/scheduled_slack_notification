package view

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/domain/model"
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
	w := new(strings.Builder)
	t := template.Must(template.New(filepath).ParseFiles(filepath))
	err := t.Execute(w, d)
	if err != nil {
		return "", err
	}
	return w.String(), nil
}

func GenMessage(ntf model.Notification) (Message, error) {
	bs := []Block{}
	for _, m := range ntf.Metrics {
		v, err := m.GetValueStr()
		if err != nil {
			continue
		}
		b := SectionBlock{
			Object: TextBlockObject{
				TextType: TextBlockMarkdown,
				Text:     fmt.Sprintf("Metric %s is %s.", m.Name, v),
			},
		}
		bs = append(bs, b)
		bs = append(bs, DeviderBlock{})
	}

	m := Message{
		Blocks: bs,
	}
	return m, nil
}
