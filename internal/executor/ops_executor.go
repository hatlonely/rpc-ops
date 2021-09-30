package executor

import "github.com/hatlonely/go-kit/logger"

type OpsOptions struct {
	Executor Options

	ProducerNum int `dft:"10"`
	ConsumerNum int `dft:"10"`
}

func NewOpsExecutorWithOptions(options *OpsOptions, infoLogger, execLogger *logger.Logger) *OpsExecutor {
	executor := NewExecutorWithOptions(&options.Executor, infoLogger, execLogger)

	executor.SetProducerFunc(options.ProducerNum, func() (interface{}, error) {
		return nil, nil
	})
	executor.SetConsumerFunc(options.ConsumerNum, func(i interface{}) error {
		return nil
	})

	return &OpsExecutor{
		executor: executor,
	}
}

type OpsExecutor struct {
	executor *Executor
}

func (e *OpsExecutor) Run() {
	e.executor.Run()
}

func (e *OpsExecutor) Stop() {
	e.executor.Stop()
}
