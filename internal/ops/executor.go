package ops

import (
	"context"
	"sync"

	"github.com/hatlonely/go-kit/logger"
	gops "github.com/hatlonely/go-kit/ops"
	"github.com/pkg/errors"
)

type ExecutorOptions struct {
	WorkerNum int
}

func NewExecutorWithOptions(manager *Manager, options *ExecutorOptions) (*Executor, error) {
	ctx, cancel := context.WithCancel(context.Background())

	return &Executor{
		options: options,
		manager: manager,
		infoLog: logger.NewStdoutJsonLogger(),
		workLog: logger.NewStdoutJsonLogger(),
		ctx:     ctx,
		cancel:  cancel,
	}, nil
}

type Executor struct {
	options *ExecutorOptions
	manager *Manager

	infoLog *logger.Logger
	workLog *logger.Logger

	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

func (e *Executor) SetLogger(infoLog *logger.Logger, workLog *logger.Logger) {
	e.infoLog = infoLog
	e.workLog = workLog
}

func (e *Executor) Run() {
	for i := 0; i < e.options.WorkerNum; i++ {
		e.wg.Add(1)
		go func() {
		out:
			for {
				select {
				case <-e.ctx.Done():
					break out
				default:
				}
				_ = e.work(e.ctx)
			}
			e.wg.Done()
		}()
	}
}

func (e *Executor) Stop() {
	e.cancel()
	e.wg.Wait()
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
