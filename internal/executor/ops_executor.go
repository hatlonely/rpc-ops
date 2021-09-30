package executor

type OpsOptions struct {
}

func NewOpsExecutorWithOptions(options *Options) {

}

type OpsExecutor struct {
	executor Executor
}

func (e *OpsExecutor) Run() {
	e.executor.Run()
}

func (e *OpsExecutor) Stop() {
	e.executor.Stop()
}
