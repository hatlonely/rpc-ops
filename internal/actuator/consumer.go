package actuator

type Consumer interface {
	Consume(interface{}) error
}

type consumerFunc struct {
	consumer func(interface{}) error
}

func (c consumerFunc) Consume(v interface{}) error {
	return c.consumer(v)
}

func ConsumerFunc(consumer func(interface{}) error) Consumer {
	return &consumerFunc{consumer: consumer}
}
