package ops

import (
	"context"

	"github.com/hatlonely/go-kit/logger"
	gops "github.com/hatlonely/go-kit/ops"
	"github.com/pkg/errors"
)

type ExecutorOptions struct {
	WorkerNum int
}

func NewExecutorWithOptions(manager *Manager, options *ExecutorOptions) (*Executor, error) {

	return &Executor{
		manager: manager,
		infoLog: logger.NewStdoutJsonLogger(),
		workLog: logger.NewStdoutJsonLogger(),
	}, nil
}

type Executor struct {
	manager *Manager

	infoLog *logger.Logger
	workLog *logger.Logger
}

func (e *Executor) SetLogger(infoLog *logger.Logger, workLog *logger.Logger) {
	e.infoLog = infoLog
	e.workLog = workLog
}

func (e *Executor) work(ctx context.Context) error {
	job, err := e.manager.FindOneWaitingJob(ctx)
	if err != nil {
		return errors.WithMessage(err, "m.manager.FindOneWaitingJob failed")
	}

	status, stdout, stderr, err := gops.ExecCommand(job.Command, nil, job.WorkDirectory)
	if err != nil {
		return errors.WithMessage(err, "ExecCommand failed")
	}
	_, _, _ = status, stdout, stderr

	return nil
}
