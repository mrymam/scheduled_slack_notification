package slack

import "fmt"

type Message struct {
	Text string
}

type Client interface {
	PostMessage(url string, msg Message) error
}

type ClientImpl struct{}

func (c ClientImpl) PostMessage(url string, msg Message) error {
	fmt.Printf("message:%s to %s\n", msg.Text, url)
	return nil
}
