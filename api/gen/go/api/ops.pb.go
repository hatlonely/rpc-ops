// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.15.5
// source: api/ops.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_api_ops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_api_ops_proto_rawDescGZIP(), []int{0}
}

type Repository struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Username   string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password   string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Endpoint   string `protobuf:"bytes,4,opt,name=endpoint,proto3" json:"endpoint,omitempty"`
	Repository string `protobuf:"bytes,5,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *Repository) Reset() {
	*x = Repository{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Repository) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Repository) ProtoMessage() {}

func (x *Repository) ProtoReflect() protoreflect.Message {
	mi := &file_api_ops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Repository.ProtoReflect.Descriptor instead.
func (*Repository) Descriptor() ([]byte, []int) {
	return file_api_ops_proto_rawDescGZIP(), []int{1}
}

func (x *Repository) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Repository) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Repository) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Repository) GetEndpoint() string {
	if x != nil {
		return x.Endpoint
	}
	return ""
}

func (x *Repository) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

type ListRepositoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *ListRepositoryReq) Reset() {
	*x = ListRepositoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepositoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepositoryReq) ProtoMessage() {}

func (x *ListRepositoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_ops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepositoryReq.ProtoReflect.Descriptor instead.
func (*ListRepositoryReq) Descriptor() ([]byte, []int) {
	return file_api_ops_proto_rawDescGZIP(), []int{2}
}

func (x *ListRepositoryReq) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListRepositoryReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListRepositoryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total        int32         `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Repositories []*Repository `protobuf:"bytes,2,rep,name=repositories,proto3" json:"repositories,omitempty"`
}

func (x *ListRepositoryRes) Reset() {
	*x = ListRepositoryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRepositoryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRepositoryRes) ProtoMessage() {}

func (x *ListRepositoryRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_ops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRepositoryRes.ProtoReflect.Descriptor instead.
func (*ListRepositoryRes) Descriptor() ([]byte, []int) {
	return file_api_ops_proto_rawDescGZIP(), []int{3}
}

func (x *ListRepositoryRes) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListRepositoryRes) GetRepositories() []*Repository {
	if x != nil {
		return x.Repositories
	}
	return nil
}

type GetRepositoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRepositoryReq) Reset() {
	*x = GetRepositoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRepositoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRepositoryReq) ProtoMessage() {}

func (x *GetRepositoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_ops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRepositoryReq.ProtoReflect.Descriptor instead.
func (*GetRepositoryReq) Descriptor() ([]byte, []int) {
	return file_api_ops_proto_rawDescGZIP(), []int{4}
}

func (x *GetRepositoryReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DelRepositoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DelRepositoryReq) Reset() {
	*x = DelRepositoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelRepositoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelRepositoryReq) ProtoMessage() {}

func (x *DelRepositoryReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_ops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelRepositoryReq.ProtoReflect.Descriptor instead.
func (*DelRepositoryReq) Descriptor() ([]byte, []int) {
	return file_api_ops_proto_rawDescGZIP(), []int{5}
}

func (x *DelRepositoryReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PutRepositoryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *PutRepositoryRes) Reset() {
	*x = PutRepositoryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_ops_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutRepositoryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutRepositoryRes) ProtoMessage() {}

func (x *PutRepositoryRes) ProtoReflect() protoreflect.Message {
	mi := &file_api_ops_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutRepositoryRes.ProtoReflect.Descriptor instead.
func (*PutRepositoryRes) Descriptor() ([]byte, []int) {
	return file_api_ops_proto_rawDescGZIP(), []int{6}
}

func (x *PutRepositoryRes) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_api_ops_proto protoreflect.FileDescriptor

var file_api_ops_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x90, 0x01, 0x0a, 0x0a,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1e,
	0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x41,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x5e, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x33, 0x0a, 0x0c,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x50, 0x75, 0x74,
	0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32, 0xb2, 0x03,
	0x0a, 0x0a, 0x4f, 0x70, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x0e,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x16,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x52, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x1a, 0x0f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x22,
	0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4f, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x15, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72,
	0x79, 0x52, 0x65, 0x71, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x2a, 0x13, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x52, 0x0a,
	0x0d, 0x50, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x1a,
	0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e,
	0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x3a, 0x01,
	0x2a, 0x12, 0x4f, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x0a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x22, 0x1e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x18, 0x1a, 0x13, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a,
	0x01, 0x2a, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x68, 0x61, 0x74, 0x6c, 0x6f, 0x6e, 0x65, 0x6c, 0x79, 0x2f, 0x72, 0x70, 0x63, 0x2d, 0x6f,
	0x70, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_ops_proto_rawDescOnce sync.Once
	file_api_ops_proto_rawDescData = file_api_ops_proto_rawDesc
)

func file_api_ops_proto_rawDescGZIP() []byte {
	file_api_ops_proto_rawDescOnce.Do(func() {
		file_api_ops_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_ops_proto_rawDescData)
	})
	return file_api_ops_proto_rawDescData
}

var file_api_ops_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_ops_proto_goTypes = []interface{}{
	(*Empty)(nil),             // 0: api.Empty
	(*Repository)(nil),        // 1: api.Repository
	(*ListRepositoryReq)(nil), // 2: api.ListRepositoryReq
	(*ListRepositoryRes)(nil), // 3: api.ListRepositoryRes
	(*GetRepositoryReq)(nil),  // 4: api.GetRepositoryReq
	(*DelRepositoryReq)(nil),  // 5: api.DelRepositoryReq
	(*PutRepositoryRes)(nil),  // 6: api.PutRepositoryRes
}
var file_api_ops_proto_depIdxs = []int32{
	1, // 0: api.ListRepositoryRes.repositories:type_name -> api.Repository
	2, // 1: api.OpsService.ListRepository:input_type -> api.ListRepositoryReq
	4, // 2: api.OpsService.GetRepository:input_type -> api.GetRepositoryReq
	5, // 3: api.OpsService.DelRepository:input_type -> api.DelRepositoryReq
	1, // 4: api.OpsService.PutRepository:input_type -> api.Repository
	1, // 5: api.OpsService.UpdateRepository:input_type -> api.Repository
	3, // 6: api.OpsService.ListRepository:output_type -> api.ListRepositoryRes
	1, // 7: api.OpsService.GetRepository:output_type -> api.Repository
	0, // 8: api.OpsService.DelRepository:output_type -> api.Empty
	6, // 9: api.OpsService.PutRepository:output_type -> api.PutRepositoryRes
	0, // 10: api.OpsService.UpdateRepository:output_type -> api.Empty
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_ops_proto_init() }
func file_api_ops_proto_init() {
	if File_api_ops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_ops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Repository); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepositoryReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRepositoryRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRepositoryReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelRepositoryReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_ops_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutRepositoryRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_ops_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_ops_proto_goTypes,
		DependencyIndexes: file_api_ops_proto_depIdxs,
		MessageInfos:      file_api_ops_proto_msgTypes,
	}.Build()
	File_api_ops_proto = out.File
	file_api_ops_proto_rawDesc = nil
	file_api_ops_proto_goTypes = nil
	file_api_ops_proto_depIdxs = nil
}
