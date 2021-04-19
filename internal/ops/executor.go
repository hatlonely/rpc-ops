package ops

import "github.com/hatlonely/go-kit/logger"

type ExecutorOptions struct {
}

func NewExecutorWithOptions(options *ExecutorOptions) (*Executor, error) {
	return &Executor{
		infoLog: logger.NewStdoutJsonLogger(),
		workLog: logger.NewStdoutJsonLogger(),
	}, nil
}

type Executor struct {
	infoLog *logger.Logger
	workLog *logger.Logger
}

func (e *Executor) SetLogger(infoLog *logger.Logger, workLog *logger.Logger) {
	e.infoLog = infoLog
	e.workLog = workLog
}
