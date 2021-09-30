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

func NewExecutorWithOptions(infoLogger, execLogger *logger.Logger, options *Options) *Executor {
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

func (a *Executor) SetProducerFunc(producerNum int, producer func() (interface{}, error)) {
	for i := 0; i < producerNum; i++ {
		a.producers = append(a.producers, ProducerFunc(producer))
	}
}

func (a *Executor) SetConsumerFunc(consumerNum int, consumer func(interface{}) error) {
	for i := 0; i < consumerNum; i++ {
		a.consumers = append(a.consumers, ConsumerFunc(consumer))
	}
}

func (a *Executor) AddProducer(producer Producer) {
	a.producers = append(a.producers, producer)
}

func (a *Executor) AddCustomer(consumer Consumer) {
	a.consumers = append(a.consumers, consumer)
}

func (a *Executor) Run() {
	a.ctx, a.cancel = context.WithCancel(context.Background())

	for _, producer := range a.producers {
		a.wgp.Add(1)
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
					a.errQueue <- err
				case product := <-chp:
					a.taskQueue <- product
				}
			}
			a.wgp.Done()
		}(a.ctx, producer)
	}

	for _, consumer := range a.consumers {
		a.wgc.Add(1)
		go func(ctx context.Context, consumer Consumer) {
			for task := range a.taskQueue {
				now := time.Now()
				err := consumer.Consume(task)
				if err != nil {
					a.errQueue <- err
					a.execLogger.With("task", task).With("err", err.Error()).With("resTimeMs", time.Now().Sub(now).Milliseconds()).Info("")
				} else {
					a.execLogger.With("task", task).With("errCode", "OK").With("resTimeMs", time.Now().Sub(now).Milliseconds()).Info("")
				}
			}
			a.wgc.Done()
		}(a.ctx, consumer)
	}

	a.wge.Add(1)
	go func() {
		for err := range a.errQueue {
			a.infoLogger.With("err", err.Error()).With("errStack", fmt.Sprintf("%+v", err))
		}
		a.wge.Done()
	}()
}

func (a *Executor) Stop() {
	a.cancel()
	a.wgp.Wait()
	close(a.taskQueue)
	a.wgc.Wait()
	close(a.errQueue)
	a.wge.Wait()
}
