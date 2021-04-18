package service

import (
	"context"
	"time"

	"github.com/hatlonely/rpc-ops/api/gen/go/api"
	"github.com/hatlonely/rpc-ops/internal/ops"
)

func (s *Service) ListVariable(ctx context.Context, req *api.ListVariableReq) (*api.ListVariableRes, error) {
	variables, err := s.manager.ListVariable(ctx, req.Offset, req.Limit)
	if err != nil {
		return nil, err
	}
	var res api.ListVariableRes
	for _, variable := range variables {
		res.Variables = append(res.Variables, &api.Variable{
			Id:       variable.ID,
			Name:     variable.Name,
			Body:     variable.Body,
			CreateAt: variable.CreateAt.Format(time.RFC3339),
			UpdateAt: variable.UpdateAt.Format(time.RFC3339),
		})
	}
	return &res, nil
}

func (s *Service) GetVariable(ctx context.Context, req *api.VariableID) (*api.Variable, error) {
	variable, err := s.manager.GetVariable(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Variable{
		Id:       variable.ID,
		Name:     variable.Name,
		Body:     variable.Body,
		CreateAt: variable.CreateAt.Format(time.RFC3339),
		UpdateAt: variable.UpdateAt.Format(time.RFC3339),
	}, nil
}

func (s *Service) DelVariable(ctx context.Context, req *api.VariableID) (*api.Empty, error) {
	err := s.manager.DelVariable(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}

func (s *Service) PutVariable(ctx context.Context, req *api.Variable) (*api.VariableID, error) {
	id, err := s.manager.PutVariable(ctx, &ops.Variable{
		Name: req.Name,
		Body: req.Body,
	})
	if err != nil {
		return nil, err
	}
	return &api.VariableID{Id: id}, nil
}

func (s *Service) UpdateVariable(ctx context.Context, req *api.Variable) (*api.Empty, error) {
	err := s.manager.UpdateVariable(ctx, &ops.Variable{
		Name: req.Name,
		Body: req.Body,
	})
	if err != nil {
		return nil, err
	}
	return &api.Empty{}, nil
}
