package executor

import (
	"github.com/hatlonely/go-kit/logger"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-ops/internal/ops"
)

type OpsOptions struct {
	Executor Options

	ProducerNum int `dft:"10"`
	ConsumerNum int `dft:"10"`
}

func NewOpsExecutorWithOptions(options *OpsOptions, manager *ops.Manager, infoLogger, execLogger *logger.Logger) *OpsExecutor {
	executor := NewExecutorWithOptions(&options.Executor, infoLogger, execLogger)
	oe := &OpsExecutor{
		executor: executor,
		manager:  manager,
	}

	executor.SetProducerFunc(options.ProducerNum, oe.Produce)
	executor.SetConsumerFunc(options.ConsumerNum, oe.Consume)

	return oe
}

type OpsExecutor struct {
	executor *Executor
	manager  *ops.Manager
}

func (e *OpsExecutor) Produce() (interface{}, error) {
	job, err := e.manager.FindOneWaitingJob(e.executor.ctx)
	if err != nil {
		return nil, errors.WithMessage(err, "m.manager.FindOneWaitingJob failed")
	}

	return job, nil
}

func (e *OpsExecutor) Consume(interface{}) error {
	return nil
}

func (e *OpsExecutor) Run() {
	e.executor.Run()
}

func (e *OpsExecutor) Stop() {
	e.executor.Stop()
}
