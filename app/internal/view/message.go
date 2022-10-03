package view

import (
	"fmt"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/domain/model"
	"github.com/slack-go/slack"
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
				Text:     fmt.Sprintf("Metric %s is %s.", m.Name, m.Value),
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

func (vmsg Message) GenWebHookMessage() slack.WebhookMessage {
	bs := []slack.Block{}
	for _, vb := range vmsg.Blocks {

		if vb.GetType() == BlockSection {
			b, ok := vb.(SectionBlock)
			if !ok {
				continue
			}
			obj := slack.NewTextBlockObject(slack.MarkdownType, b.Object.Text, false, false)
			ib := slack.NewSectionBlock(obj, nil, nil)
			bs = append(bs, ib)
			continue
		}

		if vb.GetType() == BlockDivider {
			ib := slack.NewDividerBlock()
			bs = append(bs, ib)
			continue
		}
	}

	m := slack.WebhookMessage{
		Blocks: &slack.Blocks{
			BlockSet: bs,
		},
	}
	return m
}
