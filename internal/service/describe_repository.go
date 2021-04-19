package service

import (
	"context"
	"fmt"
	"os"
	"sort"

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
		status, stdout, stderr, err := gops.ExecCommand(command, nil, workDir)
		if err != nil {
			return nil, errors.Wrapf(err, "ExecCommand failed. command: [%v]", command)
		}
		if status != 0 {
			return nil, errors.Errorf("ExecCommand status not ok. status: [%v], stdout: [%v], stderr: [%v]", status, stdout, stderr)
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
