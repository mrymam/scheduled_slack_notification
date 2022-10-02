package view

import (
	"fmt"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/domain/model"
)

type Message struct {
	Blocks []Block
}

func InitMessage(ntf model.Notification) (Message, error) {
	bs := []Block{}
	for _, m := range ntf.Metrics {
		b := SectionBlock{
			Object: TextBlockObject{
				TextType: TextBlockMarkdown,
				Text:     fmt.Sprintf("name: %s", m.Name),
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
