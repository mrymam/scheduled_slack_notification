package view

import (
	"github.com/slack-go/slack"
)

type Message struct {
	Blocks []Block
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
