// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OpsServiceClient is the client API for OpsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpsServiceClient interface {
	ListRepository(ctx context.Context, in *ListRepositoryReq, opts ...grpc.CallOption) (*ListRepositoryRes, error)
	GetRepository(ctx context.Context, in *RepositoryID, opts ...grpc.CallOption) (*Repository, error)
	DelRepository(ctx context.Context, in *RepositoryID, opts ...grpc.CallOption) (*Empty, error)
	PutRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*RepositoryID, error)
	UpdateRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Empty, error)
	ListVariable(ctx context.Context, in *ListVariableReq, opts ...grpc.CallOption) (*ListVariableRes, error)
	GetVariable(ctx context.Context, in *VariableID, opts ...grpc.CallOption) (*Variable, error)
	DelVariable(ctx context.Context, in *VariableID, opts ...grpc.CallOption) (*Empty, error)
	PutVariable(ctx context.Context, in *Variable, opts ...grpc.CallOption) (*VariableID, error)
	UpdateVariable(ctx context.Context, in *Variable, opts ...grpc.CallOption) (*Empty, error)
	ListJob(ctx context.Context, in *ListJobReq, opts ...grpc.CallOption) (*ListJobRes, error)
	GetJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*Job, error)
	DelJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*Empty, error)
	DescribeRepository(ctx context.Context, in *DescribeRepositoryReq, opts ...grpc.CallOption) (*Playbook, error)
	RunOps(ctx context.Context, in *RunOpsReq, opts ...grpc.CallOption) (*RunOpsRes, error)
}

type opsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOpsServiceClient(cc grpc.ClientConnInterface) OpsServiceClient {
	return &opsServiceClient{cc}
}

func (c *opsServiceClient) ListRepository(ctx context.Context, in *ListRepositoryReq, opts ...grpc.CallOption) (*ListRepositoryRes, error) {
	out := new(ListRepositoryRes)
	err := c.cc.Invoke(ctx, "/api.OpsService/ListRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) GetRepository(ctx context.Context, in *RepositoryID, opts ...grpc.CallOption) (*Repository, error) {
	out := new(Repository)
	err := c.cc.Invoke(ctx, "/api.OpsService/GetRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) DelRepository(ctx context.Context, in *RepositoryID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.OpsService/DelRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) PutRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*RepositoryID, error) {
	out := new(RepositoryID)
	err := c.cc.Invoke(ctx, "/api.OpsService/PutRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) UpdateRepository(ctx context.Context, in *Repository, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.OpsService/UpdateRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) ListVariable(ctx context.Context, in *ListVariableReq, opts ...grpc.CallOption) (*ListVariableRes, error) {
	out := new(ListVariableRes)
	err := c.cc.Invoke(ctx, "/api.OpsService/ListVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) GetVariable(ctx context.Context, in *VariableID, opts ...grpc.CallOption) (*Variable, error) {
	out := new(Variable)
	err := c.cc.Invoke(ctx, "/api.OpsService/GetVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) DelVariable(ctx context.Context, in *VariableID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.OpsService/DelVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) PutVariable(ctx context.Context, in *Variable, opts ...grpc.CallOption) (*VariableID, error) {
	out := new(VariableID)
	err := c.cc.Invoke(ctx, "/api.OpsService/PutVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) UpdateVariable(ctx context.Context, in *Variable, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.OpsService/UpdateVariable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) ListJob(ctx context.Context, in *ListJobReq, opts ...grpc.CallOption) (*ListJobRes, error) {
	out := new(ListJobRes)
	err := c.cc.Invoke(ctx, "/api.OpsService/ListJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) GetJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*Job, error) {
	out := new(Job)
	err := c.cc.Invoke(ctx, "/api.OpsService/GetJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) DelJob(ctx context.Context, in *JobID, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/api.OpsService/DelJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) DescribeRepository(ctx context.Context, in *DescribeRepositoryReq, opts ...grpc.CallOption) (*Playbook, error) {
	out := new(Playbook)
	err := c.cc.Invoke(ctx, "/api.OpsService/DescribeRepository", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsServiceClient) RunOps(ctx context.Context, in *RunOpsReq, opts ...grpc.CallOption) (*RunOpsRes, error) {
	out := new(RunOpsRes)
	err := c.cc.Invoke(ctx, "/api.OpsService/RunOps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpsServiceServer is the server API for OpsService service.
// All implementations must embed UnimplementedOpsServiceServer
// for forward compatibility
type OpsServiceServer interface {
	ListRepository(context.Context, *ListRepositoryReq) (*ListRepositoryRes, error)
	GetRepository(context.Context, *RepositoryID) (*Repository, error)
	DelRepository(context.Context, *RepositoryID) (*Empty, error)
	PutRepository(context.Context, *Repository) (*RepositoryID, error)
	UpdateRepository(context.Context, *Repository) (*Empty, error)
	ListVariable(context.Context, *ListVariableReq) (*ListVariableRes, error)
	GetVariable(context.Context, *VariableID) (*Variable, error)
	DelVariable(context.Context, *VariableID) (*Empty, error)
	PutVariable(context.Context, *Variable) (*VariableID, error)
	UpdateVariable(context.Context, *Variable) (*Empty, error)
	ListJob(context.Context, *ListJobReq) (*ListJobRes, error)
	GetJob(context.Context, *JobID) (*Job, error)
	DelJob(context.Context, *JobID) (*Empty, error)
	DescribeRepository(context.Context, *DescribeRepositoryReq) (*Playbook, error)
	RunOps(context.Context, *RunOpsReq) (*RunOpsRes, error)
	mustEmbedUnimplementedOpsServiceServer()
}

// UnimplementedOpsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOpsServiceServer struct {
}

func (UnimplementedOpsServiceServer) ListRepository(context.Context, *ListRepositoryReq) (*ListRepositoryRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRepository not implemented")
}
func (UnimplementedOpsServiceServer) GetRepository(context.Context, *RepositoryID) (*Repository, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRepository not implemented")
}
func (UnimplementedOpsServiceServer) DelRepository(context.Context, *RepositoryID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelRepository not implemented")
}
func (UnimplementedOpsServiceServer) PutRepository(context.Context, *Repository) (*RepositoryID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutRepository not implemented")
}
func (UnimplementedOpsServiceServer) UpdateRepository(context.Context, *Repository) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepository not implemented")
}
func (UnimplementedOpsServiceServer) ListVariable(context.Context, *ListVariableReq) (*ListVariableRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVariable not implemented")
}
func (UnimplementedOpsServiceServer) GetVariable(context.Context, *VariableID) (*Variable, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVariable not implemented")
}
func (UnimplementedOpsServiceServer) DelVariable(context.Context, *VariableID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelVariable not implemented")
}
func (UnimplementedOpsServiceServer) PutVariable(context.Context, *Variable) (*VariableID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutVariable not implemented")
}
func (UnimplementedOpsServiceServer) UpdateVariable(context.Context, *Variable) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVariable not implemented")
}
func (UnimplementedOpsServiceServer) ListJob(context.Context, *ListJobReq) (*ListJobRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListJob not implemented")
}
func (UnimplementedOpsServiceServer) GetJob(context.Context, *JobID) (*Job, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetJob not implemented")
}
func (UnimplementedOpsServiceServer) DelJob(context.Context, *JobID) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelJob not implemented")
}
func (UnimplementedOpsServiceServer) DescribeRepository(context.Context, *DescribeRepositoryReq) (*Playbook, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeRepository not implemented")
}
func (UnimplementedOpsServiceServer) RunOps(context.Context, *RunOpsReq) (*RunOpsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunOps not implemented")
}
func (UnimplementedOpsServiceServer) mustEmbedUnimplementedOpsServiceServer() {}

// UnsafeOpsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpsServiceServer will
// result in compilation errors.
type UnsafeOpsServiceServer interface {
	mustEmbedUnimplementedOpsServiceServer()
}

func RegisterOpsServiceServer(s grpc.ServiceRegistrar, srv OpsServiceServer) {
	s.RegisterService(&OpsService_ServiceDesc, srv)
}

func _OpsService_ListRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRepositoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).ListRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/ListRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).ListRepository(ctx, req.(*ListRepositoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_GetRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).GetRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/GetRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).GetRepository(ctx, req.(*RepositoryID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_DelRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).DelRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/DelRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).DelRepository(ctx, req.(*RepositoryID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_PutRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).PutRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/PutRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).PutRepository(ctx, req.(*Repository))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_UpdateRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Repository)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).UpdateRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/UpdateRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).UpdateRepository(ctx, req.(*Repository))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_ListVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVariableReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).ListVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/ListVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).ListVariable(ctx, req.(*ListVariableReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_GetVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VariableID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).GetVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/GetVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).GetVariable(ctx, req.(*VariableID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_DelVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VariableID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).DelVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/DelVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).DelVariable(ctx, req.(*VariableID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_PutVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Variable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).PutVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/PutVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).PutVariable(ctx, req.(*Variable))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_UpdateVariable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Variable)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).UpdateVariable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/UpdateVariable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).UpdateVariable(ctx, req.(*Variable))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_ListJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListJobReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).ListJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/ListJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).ListJob(ctx, req.(*ListJobReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_GetJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).GetJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/GetJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).GetJob(ctx, req.(*JobID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_DelJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JobID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).DelJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/DelJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).DelJob(ctx, req.(*JobID))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_DescribeRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeRepositoryReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).DescribeRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/DescribeRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).DescribeRepository(ctx, req.(*DescribeRepositoryReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OpsService_RunOps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunOpsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServiceServer).RunOps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.OpsService/RunOps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServiceServer).RunOps(ctx, req.(*RunOpsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OpsService_ServiceDesc is the grpc.ServiceDesc for OpsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OpsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.OpsService",
	HandlerType: (*OpsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRepository",
			Handler:    _OpsService_ListRepository_Handler,
		},
		{
			MethodName: "GetRepository",
			Handler:    _OpsService_GetRepository_Handler,
		},
		{
			MethodName: "DelRepository",
			Handler:    _OpsService_DelRepository_Handler,
		},
		{
			MethodName: "PutRepository",
			Handler:    _OpsService_PutRepository_Handler,
		},
		{
			MethodName: "UpdateRepository",
			Handler:    _OpsService_UpdateRepository_Handler,
		},
		{
			MethodName: "ListVariable",
			Handler:    _OpsService_ListVariable_Handler,
		},
		{
			MethodName: "GetVariable",
			Handler:    _OpsService_GetVariable_Handler,
		},
		{
			MethodName: "DelVariable",
			Handler:    _OpsService_DelVariable_Handler,
		},
		{
			MethodName: "PutVariable",
			Handler:    _OpsService_PutVariable_Handler,
		},
		{
			MethodName: "UpdateVariable",
			Handler:    _OpsService_UpdateVariable_Handler,
		},
		{
			MethodName: "ListJob",
			Handler:    _OpsService_ListJob_Handler,
		},
		{
			MethodName: "GetJob",
			Handler:    _OpsService_GetJob_Handler,
		},
		{
			MethodName: "DelJob",
			Handler:    _OpsService_DelJob_Handler,
		},
		{
			MethodName: "DescribeRepository",
			Handler:    _OpsService_DescribeRepository_Handler,
		},
		{
			MethodName: "RunOps",
			Handler:    _OpsService_RunOps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ops.proto",
}
