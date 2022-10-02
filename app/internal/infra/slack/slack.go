package slack

import (
	"fmt"

	"github.com/onyanko-pon/scheduled_slack_notification/app/internal/view"
	"github.com/slack-go/slack"
)

func GenMessage(vmsg view.Message) slack.WebhookMessage {
	bs := []slack.Block{}
	for _, vb := range vmsg.Blocks {

		if vb.GetType() == view.BlockSection {
			b, ok := vb.(view.SectionBlock)
			if !ok {
				continue
			}
			obj := slack.NewTextBlockObject(slack.MarkdownType, b.Object.Text, false, false)
			ib := slack.NewSectionBlock(obj, nil, nil)
			bs = append(bs, ib)
			continue
		}

		if vb.GetType() == view.BlockDivider {
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

type Client interface {
	PostMessage(url string, msg slack.WebhookMessage) error
}

type ClientImpl struct{}

func (c ClientImpl) PostMessage(url string, msg slack.WebhookMessage) error {
	fmt.Printf("message:%s to %s\n", "message", url)
	return nil
}
