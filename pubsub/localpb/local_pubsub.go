package localpb

import (
	"context"
	"food-delivery/common"
	"food-delivery/pubsub"
	"log"
	"sync"
)

// A pb run locally (in-mem)
// It has a queue (buffer Topic) at it's core and many group of subscribers.
// Because we want to send a message with a specific topic for many subscribers in a group can handle.

type localPubSub struct {
	messageQueue chan *pubsub.Message
	mapTopic   map[pubsub.Topic][]chan *pubsub.Message
	locker       *sync.RWMutex
}

func NewPubSub() *localPubSub {
	pb := &localPubSub{
		messageQueue: make(chan *pubsub.Message, 10000),
		mapTopic:   make(map[pubsub.Topic][]chan *pubsub.Message),
		locker:       new(sync.RWMutex),
	}

	pb.run()

	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, Topic pubsub.Topic, data *pubsub.Message) error {
	data.SetTopic(Topic)

	go func() {
		defer common.AppRecover()
		ps.messageQueue <- data
		log.Println("New event published:", data.String())
	}()
	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, Topic pubsub.Topic) (ch <-chan *pubsub.Message, close func()) {
	c := make(chan *pubsub.Message)

	ps.locker.Lock()
	if val, ok := ps.mapTopic[Topic]; ok {
		val = append(ps.mapTopic[Topic], c)
		ps.mapTopic[Topic] = val
	} else {
		ps.mapTopic[Topic] = []chan *pubsub.Message{c}
	}
	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")

		if chans, ok := ps.mapTopic[Topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)

					ps.locker.Lock()
					ps.mapTopic[Topic] = chans
					ps.locker.Unlock()
					break
				}
			}
		}
	}

}

func (ps *localPubSub) run() error {
	log.Println("Pubsub started")

	go func() {
		for {
			mess := <-ps.messageQueue
			log.Println("Message dequeue:", mess)

			if subs, ok := ps.mapTopic[mess.Topic()]; ok {
				for i := range subs {
					go func(c chan *pubsub.Message) {
						c <- mess
					}(subs[i])
				}
			}
			//else {
			//	ps.messageQueue <- mess
			//}
		}
	}()

	return nil
}
