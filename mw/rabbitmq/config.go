package rabbitmq

type RetryBackoffFunction func(retryCount int) (string, error)

type QueueConfig struct {
	QueueKey              string
	DelayExpirationInMS   string
	RetryCount            int
	ConsumerPrefetchCount int
	ConsumerCount         int
	RetryBackoffFunction  RetryBackoffFunction
}

type Queues []QueueConfig
