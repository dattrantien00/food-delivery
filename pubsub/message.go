package pubsub

import (
	"fmt"
	"time"
)

type Message struct {
	id        string
	topic     Topic // can be ignore
	data      interface{}
	createdAt time.Time
}

func NewMessage(data interface{}) *Message {
	now := time.Now().UTC()

	return &Message{
		id:        fmt.Sprintf("%d", now.Nanosecond()),
		data:      data,
		createdAt: now,
	}
}

func (evt *Message) String() string {
	return fmt.Sprintf("Message %s", evt.Topic())
}

func (evt *Message) Topic() Topic {
	return evt.topic
}

func (evt *Message) SetTopic(topic Topic) {
	evt.topic = topic
}

func (evt *Message) Data() interface{} {
	return evt.data
}
