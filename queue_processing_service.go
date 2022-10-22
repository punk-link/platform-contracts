package platformcontracts

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/punk-link/logger"
)

type QueueProcessingService struct {
	logger         logger.Logger
	natsConnection *nats.Conn
}

func NewQueueProcessingService(logger logger.Logger, natsConnection *nats.Conn) *QueueProcessingService {
	return &QueueProcessingService{
		logger:         logger,
		natsConnection: natsConnection,
	}
}

func (t *QueueProcessingService) Process(platformer Platformer) {
	jetStreamContext, err := t.natsConnection.JetStream()
	subscription, err := t.getSubscription(err, jetStreamContext, platformer.GetPlatformName())
	if err != nil {
		t.logger.LogError(err, err.Error())
		return
	}

	for {
		messages, _ := subscription.Fetch(platformer.GetBatchSize())
		containers := make([]UpcContainer, len(messages))
		for i, message := range messages {
			message.Ack()

			var container UpcContainer
			_ = json.Unmarshal(message.Data, &container)

			containers[i] = container
		}

		urlResults := platformer.GetReleaseUrlsByUpc(containers)
		err = t.createJstStreamIfNotExist(nil, jetStreamContext)
		_ = t.publishUrlResults(err, jetStreamContext, urlResults)
	}
}

func (t *QueueProcessingService) createJstStreamIfNotExist(err error, jetStreamContext nats.JetStreamContext) error {
	if err != nil {
		return err
	}

	stream, _ := jetStreamContext.StreamInfo(PLATFORM_URL_RESPONSE_STREAM_NAME)
	if stream == nil {
		t.logger.LogInfo("Creating Nats stream %s and subjects %s", PLATFORM_URL_RESPONSE_STREAM_NAME, PLATFORM_URL_RESPONSE_STREAM_SUBJECT)
		_, err = jetStreamContext.AddStream(DefaultReducerConfig)
	}

	return err
}

func (t *QueueProcessingService) getSubscription(err error, jetStreamContext nats.JetStreamContext, platformName string) (*nats.Subscription, error) {
	if err != nil {
		return nil, err
	}

	subject := GetRequestStreamSubject(platformName)
	consumerName := GetRequestConsumerName(platformName)
	return jetStreamContext.PullSubscribe(subject, consumerName)
}

func (t *QueueProcessingService) publishUrlResults(err error, jetStreamContext nats.JetStreamContext, urlResults []UrlResultContainer) error {
	if err != nil {
		return err
	}

	for _, urlResult := range urlResults {
		json, _ := json.Marshal(urlResult)
		jetStreamContext.Publish(PLATFORM_URL_RESPONSE_STREAM_SUBJECT, json)
	}

	return err
}
