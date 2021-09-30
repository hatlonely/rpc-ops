package executor

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/hatlonely/go-kit/logger"
)

type Options struct {
	TaskLen int `dft:"100"`
	ErrLen  int `dft:"100"`
}

func NewExecutorWithOptions(options *Options, infoLogger, execLogger *logger.Logger) *Executor {
	return &Executor{
		taskQueue:  make(chan interface{}, options.TaskLen),
		errQueue:   make(chan error, options.ErrLen),
		execLogger: execLogger,
		infoLogger: infoLogger,
	}
}

type Executor struct {
	taskQueue chan interface{}
	errQueue  chan error
	producers []Producer
	consumers []Consumer

	wgp    sync.WaitGroup
	wgc    sync.WaitGroup
	wge    sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc

	execLogger *logger.Logger
	infoLogger *logger.Logger
}

func (e *Executor) SetProducerFunc(producerNum int, producer func() (interface{}, error)) {
	for i := 0; i < producerNum; i++ {
		e.producers = append(e.producers, ProducerFunc(producer))
	}
}

func (e *Executor) SetConsumerFunc(consumerNum int, consumer func(interface{}) error) {
	for i := 0; i < consumerNum; i++ {
		e.consumers = append(e.consumers, ConsumerFunc(consumer))
	}
}

func (e *Executor) AddProducer(producer Producer) {
	e.producers = append(e.producers, producer)
}

func (e *Executor) AddCustomer(consumer Consumer) {
	e.consumers = append(e.consumers, consumer)
}

func (e *Executor) Run() {
	e.ctx, e.cancel = context.WithCancel(context.Background())

	for _, producer := range e.producers {
		e.wgp.Add(1)
		go func(ctx context.Context, producer Producer) {
		out:
			for {
				chp := make(chan interface{})
				che := make(chan error)
				go func() {
					product, err := producer.Produce()
					if err != nil {
						che <- err
					} else {
						chp <- product
					}
				}()

				select {
				case <-ctx.Done():
					break out
				case err := <-che:
					e.errQueue <- err
				case product := <-chp:
					e.taskQueue <- product
				}
			}
			e.wgp.Done()
		}(e.ctx, producer)
	}

	for _, consumer := range e.consumers {
		e.wgc.Add(1)
		go func(ctx context.Context, consumer Consumer) {
			for task := range e.taskQueue {
				now := time.Now()
				err := consumer.Consume(task)
				if err != nil {
					e.errQueue <- err
					e.execLogger.With("task", task).With("err", err.Error()).With("resTimeMs", time.Now().Sub(now).Milliseconds()).Info("")
				} else {
					e.execLogger.With("task", task).With("errCode", "OK").With("resTimeMs", time.Now().Sub(now).Milliseconds()).Info("")
				}
			}
			e.wgc.Done()
		}(e.ctx, consumer)
	}

	e.wge.Add(1)
	go func() {
		for err := range e.errQueue {
			e.infoLogger.With("err", err.Error()).With("errStack", fmt.Sprintf("%+v", err))
		}
		e.wge.Done()
	}()
}

func (e *Executor) Stop() {
	e.cancel()
	e.wgp.Wait()
	close(e.taskQueue)
	e.wgc.Wait()
	close(e.errQueue)
	e.wge.Wait()
}
