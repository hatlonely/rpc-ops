package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"syscall"

	"github.com/hatlonely/go-kit/config"
	gops "github.com/hatlonely/go-kit/ops"
	"github.com/hatlonely/go-kit/refx"
	"github.com/pkg/errors"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
	"github.com/hatlonely/rpc-ops/internal/ops"
)

func (s *Service) DescribeRepository(ctx context.Context, req *api.DescribeRepositoryReq) (*api.Playbook, error) {
	repo, err := s.manager.GetRepository(ctx, req.Id)
	if err != nil {
		return nil, errors.WithMessage(err, "s.manager.GetRepository failed")
	}
	command := s.generateGitCloneCommand(repo, req.Version)
	workDir := fmt.Sprintf("%s/%s/%s/%s", s.options.WorkRoot, repo.Endpoint, repo.Username, repo.Name)

	if _, err := os.Stat(fmt.Sprintf("%s/%s", workDir, req.Version)); os.IsNotExist(err) {
		status, stdin, stdout, err := ExecCommand(command, nil, workDir)
		if err != nil {
			return nil, errors.Wrapf(err, "ExecCommand failed. command: [%v]", command)
		}
		if status != 0 {
			return nil, errors.Errorf("ExecCommand status not ok. status: [%v], stdin: [%v], stdout: [%v]", status, stdin, stdout)
		}
	}

	var playbook gops.Playbook
	cfg, err := config.NewConfigWithSimpleFile(repo.Playbook)
	if err != nil {
		return nil, errors.WithMessagef(err, "config.NewConfigWithSimpleFile failed. playbook: [%v]", repo.Playbook)
	}
	if err := cfg.Unmarshal(&playbook, refx.WithCamelName()); err != nil {
		return nil, errors.Wrap(err, "cfg.Unmarshal failed")
	}

	var res api.Playbook
	res.Name = playbook.Name
	for key, val := range playbook.Env {
		if key != "default" {
			res.Envs = append(res.Envs, &api.Playbook_Env{
				Key: key,
				Val: val,
			})
		}
	}
	sort.Slice(res.Envs, func(i, j int) bool {
		return res.Envs[i].Key < res.Envs[j].Key
	})
	for key, val := range playbook.Env {
		if key == "default" {
			res.Envs = append([]*api.Playbook_Env{{
				Key: key,
				Val: val,
			}}, res.Envs...)
		}
	}
	for key, val := range playbook.Task {
		args := map[string]*api.Playbook_Task_Args{}
		for k, arg := range val.Args {
			args[k] = &api.Playbook_Task_Args{
				Type:       arg.Type,
				Default:    arg.Default,
				Validation: arg.Validation,
			}
		}
		res.Tasks[key] = &api.Playbook_Task{
			Args:  args,
			Const: val.Const,
			Step:  val.Step,
		}
	}

	return &res, nil
}

func (s *Service) generateGitCloneCommand(repo *ops.Repository, version string) string {
	var command string
	if version != "" {
		command = fmt.Sprintf(
			"git clone --depth 1 --branch %s https://%s:%s@%s/%s/%s.git %s",
			version, repo.Username, repo.Password, repo.Endpoint, repo.Username, repo.Name, version,
		)
	} else {
		command = fmt.Sprintf(
			"git clone --depth 1 --branch master https://%s:%s@%s/%s/%s.git %s",
			repo.Username, repo.Password, repo.Endpoint, repo.Username, repo.Name, version,
		)
	}
	return command
}

func ExecCommandWithOutput(command string, environment []string, workDir string, stdout io.Writer, stderr io.Writer) (int, error) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, environment...)

	if workDir != "" {
		cmd.Dir = workDir
	}
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Start(); err != nil {
		return -1, err
	}

	if err := cmd.Wait(); err != nil {
		if e, ok := err.(*exec.ExitError); ok {
			if status, ok := e.Sys().(syscall.WaitStatus); ok {
				return status.ExitStatus(), nil
			}
		}

		return -1, err
	}

	return 0, nil
}

func ExecCommand(command string, environment []string, workDir string) (int, string, string, error) {
	var stdout = &bytes.Buffer{}
	var stderr = &bytes.Buffer{}
	status, err := ExecCommandWithOutput(command, environment, workDir, stdout, stderr)
	return status, stdout.String(), stderr.String(), err
}
