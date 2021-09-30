package executor

type Producer interface {
	Produce() (interface{}, error)
}

type producerFunc struct {
	producer func() (interface{}, error)
}

func (p producerFunc) Produce() (interface{}, error) {
	return p.producer()
}

func ProducerFunc(producer func() (interface{}, error)) Producer {
	return &producerFunc{producer: producer}
}
