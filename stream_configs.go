package platformcontracts

import (
	"time"

	"github.com/nats-io/nats.go"
)

var DefaultReducerConfig = &nats.StreamConfig{
	Name:      PLATFORM_URL_RESPONSE_STREAM_NAME,
	MaxAge:    time.Hour * 24,
	Retention: nats.WorkQueuePolicy,
	Storage:   nats.FileStorage,
	Subjects:  []string{PLATFORM_URL_RESPONSE_STREAM_SUBJECT},
}

var DefaultPlatformServiceConfig = &nats.StreamConfig{
	Name:      PLATFORM_URL_REQUESTS_STREAM_NAME,
	MaxAge:    time.Hour,
	Retention: nats.WorkQueuePolicy,
	Storage:   nats.MemoryStorage,
	Subjects:  []string{PLATFORM_URL_REQUESTS_STREAM_SUBJECTS},
}
