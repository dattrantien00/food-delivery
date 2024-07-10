package subscriber

import (
	"context"
	"food-delivery/common"
	"food-delivery/component/appctx"
	"food-delivery/component/asyncjob"
	"food-delivery/pubsub"
	"log"
)

type consumerJob struct {
	Title string
	Hld   func(ctx context.Context, message *pubsub.Message) error
}



type consumerEngine struct {
	appCtx appctx.AppContext
}

func NewEngine(appContext appctx.AppContext) *consumerEngine {
	return &consumerEngine{appCtx: appContext}
}
func (engine *consumerEngine) Start() error {

	engine.startSubTopic(
		common.TopicUserLikeRestaurant,
		false,
		IncreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
		EmitRealtimeAfterUserLikeRestaurant(engine.appCtx),
	)
	engine.startSubTopic(
		common.TopicUserUnLikeRestaurant,
		false,
		DecreaseLikeCountAfterUserLikeRestaurant(engine.appCtx),
	)

	return nil
}

type GroupJob interface {
	Run(ctx context.Context) error
}

func (engine *consumerEngine) startSubTopic(topic pubsub.Topic, isParallel bool, hdls ...consumerJob) error {
	// forever: listen message and execute group
	// hdls => []job
	// new group ([]job)

	c, _ := engine.appCtx.GetPubsub().Subscribe(context.Background(), topic)

	for _, item := range hdls {
		log.Println("Setup consumer for:", item.Title)
	}

	getHld := func(job *consumerJob, message *pubsub.Message) func(ctx context.Context) error {
		return func(ctx context.Context) error {
			log.Println("running job for ", job.Title, ". Value: ", message.Data())
			return job.Hld(ctx, message)
		}
	}

	go func() {
		for {
			msg := <-c

			jobHdlArr := make([]asyncjob.Job, len(hdls))

			for i := range hdls {

				jobHdlArr[i] = asyncjob.NewJob(getHld(&hdls[i], msg))
			}

			group := asyncjob.NewGroup(isParallel, jobHdlArr...)

			if err := group.Run(context.Background()); err != nil {
				log.Println(err)
			}
		}
	}()

	return nil
}
