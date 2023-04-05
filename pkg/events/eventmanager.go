package events

type EventManager interface {
	Publish(topic string, message interface{}) error
	Subscribe(topic string, handler func(message interface{})) error
	Enqueue(queue string, message interface{}) error
	Dequeue(queue string, handler func(message interface{})) error
}

type RedisEventManager struct {
	// Redis client and configuration
}

type SQSEventManager struct {
	// AWS SDK clients and configuration
}
