// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/hpc.proto

package hpc

import (
	proto "github.com/essayZW/hpcmanager/gateway/proto"
	proto1 "github.com/essayZW/hpcmanager/proto"
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

type AddUserWithGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest   *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	TutorUsername string             `protobuf:"bytes,2,opt,name=tutorUsername,proto3" json:"tutorUsername,omitempty"`
	GroupName     string             `protobuf:"bytes,3,opt,name=groupName,proto3" json:"groupName,omitempty"`
	QueueName     string             `protobuf:"bytes,4,opt,name=queueName,proto3" json:"queueName,omitempty"`
}

func (x *AddUserWithGroupRequest) Reset() {
	*x = AddUserWithGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserWithGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserWithGroupRequest) ProtoMessage() {}

func (x *AddUserWithGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserWithGroupRequest.ProtoReflect.Descriptor instead.
func (*AddUserWithGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{0}
}

func (x *AddUserWithGroupRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *AddUserWithGroupRequest) GetTutorUsername() string {
	if x != nil {
		return x.TutorUsername
	}
	return ""
}

func (x *AddUserWithGroupRequest) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *AddUserWithGroupRequest) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

type AddUserWithGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupName  string `protobuf:"bytes,1,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Gid        int32  `protobuf:"varint,2,opt,name=gid,proto3" json:"gid,omitempty"`
	UserName   string `protobuf:"bytes,3,opt,name=userName,proto3" json:"userName,omitempty"`
	Uid        int32  `protobuf:"varint,4,opt,name=uid,proto3" json:"uid,omitempty"`
	HpcGroupID int32  `protobuf:"varint,5,opt,name=hpcGroupID,proto3" json:"hpcGroupID,omitempty"`
	HpcUserID  int32  `protobuf:"varint,6,opt,name=hpcUserID,proto3" json:"hpcUserID,omitempty"`
}

func (x *AddUserWithGroupResponse) Reset() {
	*x = AddUserWithGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserWithGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserWithGroupResponse) ProtoMessage() {}

func (x *AddUserWithGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserWithGroupResponse.ProtoReflect.Descriptor instead.
func (*AddUserWithGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserWithGroupResponse) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *AddUserWithGroupResponse) GetGid() int32 {
	if x != nil {
		return x.Gid
	}
	return 0
}

func (x *AddUserWithGroupResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AddUserWithGroupResponse) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *AddUserWithGroupResponse) GetHpcGroupID() int32 {
	if x != nil {
		return x.HpcGroupID
	}
	return 0
}

func (x *AddUserWithGroupResponse) GetHpcUserID() int32 {
	if x != nil {
		return x.HpcUserID
	}
	return 0
}

type AddUserToGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	UserName    string             `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	HpcGroupID  int32              `protobuf:"varint,3,opt,name=hpcGroupID,proto3" json:"hpcGroupID,omitempty"`
}

func (x *AddUserToGroupRequest) Reset() {
	*x = AddUserToGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToGroupRequest) ProtoMessage() {}

func (x *AddUserToGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToGroupRequest.ProtoReflect.Descriptor instead.
func (*AddUserToGroupRequest) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{2}
}

func (x *AddUserToGroupRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *AddUserToGroupRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AddUserToGroupRequest) GetHpcGroupID() int32 {
	if x != nil {
		return x.HpcGroupID
	}
	return 0
}

type AddUserToGroupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HpcUserID int32  `protobuf:"varint,1,opt,name=hpcUserID,proto3" json:"hpcUserID,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Uid       int32  `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *AddUserToGroupResponse) Reset() {
	*x = AddUserToGroupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddUserToGroupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserToGroupResponse) ProtoMessage() {}

func (x *AddUserToGroupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserToGroupResponse.ProtoReflect.Descriptor instead.
func (*AddUserToGroupResponse) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{3}
}

func (x *AddUserToGroupResponse) GetHpcUserID() int32 {
	if x != nil {
		return x.HpcUserID
	}
	return 0
}

func (x *AddUserToGroupResponse) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *AddUserToGroupResponse) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type GetUserInfoByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	HpcUserID   int32              `protobuf:"varint,2,opt,name=hpcUserID,proto3" json:"hpcUserID,omitempty"`
}

func (x *GetUserInfoByIDRequest) Reset() {
	*x = GetUserInfoByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoByIDRequest) ProtoMessage() {}

func (x *GetUserInfoByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoByIDRequest.ProtoReflect.Descriptor instead.
func (*GetUserInfoByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserInfoByIDRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetUserInfoByIDRequest) GetHpcUserID() int32 {
	if x != nil {
		return x.HpcUserID
	}
	return 0
}

type GetUserInfoByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *HpcUser `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserInfoByIDResponse) Reset() {
	*x = GetUserInfoByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoByIDResponse) ProtoMessage() {}

func (x *GetUserInfoByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoByIDResponse.ProtoReflect.Descriptor instead.
func (*GetUserInfoByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserInfoByIDResponse) GetUser() *HpcUser {
	if x != nil {
		return x.User
	}
	return nil
}

type GetGroupInfoByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	HpcGroupID  int32              `protobuf:"varint,2,opt,name=hpcGroupID,proto3" json:"hpcGroupID,omitempty"`
}

func (x *GetGroupInfoByIDRequest) Reset() {
	*x = GetGroupInfoByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupInfoByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupInfoByIDRequest) ProtoMessage() {}

func (x *GetGroupInfoByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupInfoByIDRequest.ProtoReflect.Descriptor instead.
func (*GetGroupInfoByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{6}
}

func (x *GetGroupInfoByIDRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetGroupInfoByIDRequest) GetHpcGroupID() int32 {
	if x != nil {
		return x.HpcGroupID
	}
	return 0
}

type GetGroupInfoByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *HpcGroup `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *GetGroupInfoByIDResponse) Reset() {
	*x = GetGroupInfoByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupInfoByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupInfoByIDResponse) ProtoMessage() {}

func (x *GetGroupInfoByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupInfoByIDResponse.ProtoReflect.Descriptor instead.
func (*GetGroupInfoByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{7}
}

func (x *GetGroupInfoByIDResponse) GetGroup() *HpcGroup {
	if x != nil {
		return x.Group
	}
	return nil
}

type GetNodeUsageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest   *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	StartTimeUnix int64              `protobuf:"varint,2,opt,name=startTimeUnix,proto3" json:"startTimeUnix,omitempty"`
	EndTimeUnix   int64              `protobuf:"varint,3,opt,name=endTimeUnix,proto3" json:"endTimeUnix,omitempty"`
}

func (x *GetNodeUsageRequest) Reset() {
	*x = GetNodeUsageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeUsageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeUsageRequest) ProtoMessage() {}

func (x *GetNodeUsageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeUsageRequest.ProtoReflect.Descriptor instead.
func (*GetNodeUsageRequest) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{8}
}

func (x *GetNodeUsageRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetNodeUsageRequest) GetStartTimeUnix() int64 {
	if x != nil {
		return x.StartTimeUnix
	}
	return 0
}

func (x *GetNodeUsageRequest) GetEndTimeUnix() int64 {
	if x != nil {
		return x.EndTimeUnix
	}
	return 0
}

type GetNodeUsageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Usages []*HpcNodeUsage `protobuf:"bytes,1,rep,name=usages,proto3" json:"usages,omitempty"`
}

func (x *GetNodeUsageResponse) Reset() {
	*x = GetNodeUsageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpc_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeUsageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeUsageResponse) ProtoMessage() {}

func (x *GetNodeUsageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpc_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeUsageResponse.ProtoReflect.Descriptor instead.
func (*GetNodeUsageResponse) Descriptor() ([]byte, []int) {
	return file_proto_hpc_proto_rawDescGZIP(), []int{9}
}

func (x *GetNodeUsageResponse) GetUsages() []*HpcNodeUsage {
	if x != nil {
		return x.Usages
	}
	return nil
}

var File_proto_hpc_proto protoreflect.FileDescriptor

var file_proto_hpc_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x68, 0x70, 0x63, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x70, 0x63, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x57, 0x69, 0x74, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x75, 0x74, 0x6f,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c,
	0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x18, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x67, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x70, 0x63, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x70, 0x63, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x70, 0x63, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x22, 0x8b, 0x01, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x70, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x70, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x44, 0x22, 0x64, 0x0a, 0x16, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x68,
	0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09,
	0x68, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x6e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x70, 0x63,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x68, 0x70,
	0x63, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x3b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0c, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x48, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x71, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x70, 0x63, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x68, 0x70, 0x63,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22, 0x3f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x48, 0x70, 0x63, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x95, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x4e, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0d, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x12, 0x20,
	0x0a, 0x0b, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x78,
	0x22, 0x41, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x75, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x48,
	0x70, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x75, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x32, 0xc8, 0x03, 0x0a, 0x03, 0x48, 0x70, 0x63, 0x12, 0x37, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x57,
	0x69, 0x74, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1c, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x57, 0x69, 0x74, 0x68, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x1a, 0x2e, 0x68, 0x70, 0x63, 0x2e,
	0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x41, 0x64, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x54, 0x6f, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1c, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4e, 0x6f,
	0x64, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x68, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x55,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x73,
	0x61, 0x79, 0x5a, 0x57, 0x2f, 0x68, 0x70, 0x63, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x68, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hpc_proto_rawDescOnce sync.Once
	file_proto_hpc_proto_rawDescData = file_proto_hpc_proto_rawDesc
)

func file_proto_hpc_proto_rawDescGZIP() []byte {
	file_proto_hpc_proto_rawDescOnce.Do(func() {
		file_proto_hpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hpc_proto_rawDescData)
	})
	return file_proto_hpc_proto_rawDescData
}

var file_proto_hpc_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_hpc_proto_goTypes = []interface{}{
	(*AddUserWithGroupRequest)(nil),  // 0: hpc.AddUserWithGroupRequest
	(*AddUserWithGroupResponse)(nil), // 1: hpc.AddUserWithGroupResponse
	(*AddUserToGroupRequest)(nil),    // 2: hpc.AddUserToGroupRequest
	(*AddUserToGroupResponse)(nil),   // 3: hpc.AddUserToGroupResponse
	(*GetUserInfoByIDRequest)(nil),   // 4: hpc.GetUserInfoByIDRequest
	(*GetUserInfoByIDResponse)(nil),  // 5: hpc.GetUserInfoByIDResponse
	(*GetGroupInfoByIDRequest)(nil),  // 6: hpc.GetGroupInfoByIDRequest
	(*GetGroupInfoByIDResponse)(nil), // 7: hpc.GetGroupInfoByIDResponse
	(*GetNodeUsageRequest)(nil),      // 8: hpc.GetNodeUsageRequest
	(*GetNodeUsageResponse)(nil),     // 9: hpc.GetNodeUsageResponse
	(*proto.BaseRequest)(nil),        // 10: request.BaseRequest
	(*HpcUser)(nil),                  // 11: hpc.HpcUser
	(*HpcGroup)(nil),                 // 12: hpc.HpcGroup
	(*HpcNodeUsage)(nil),             // 13: hpc.HpcNodeUsage
	(*proto1.Empty)(nil),             // 14: publicproto.Empty
	(*proto1.PingResponse)(nil),      // 15: publicproto.PingResponse
}
var file_proto_hpc_proto_depIdxs = []int32{
	10, // 0: hpc.AddUserWithGroupRequest.baseRequest:type_name -> request.BaseRequest
	10, // 1: hpc.AddUserToGroupRequest.baseRequest:type_name -> request.BaseRequest
	10, // 2: hpc.GetUserInfoByIDRequest.baseRequest:type_name -> request.BaseRequest
	11, // 3: hpc.GetUserInfoByIDResponse.user:type_name -> hpc.HpcUser
	10, // 4: hpc.GetGroupInfoByIDRequest.baseRequest:type_name -> request.BaseRequest
	12, // 5: hpc.GetGroupInfoByIDResponse.group:type_name -> hpc.HpcGroup
	10, // 6: hpc.GetNodeUsageRequest.baseRequest:type_name -> request.BaseRequest
	13, // 7: hpc.GetNodeUsageResponse.usages:type_name -> hpc.HpcNodeUsage
	14, // 8: hpc.Hpc.Ping:input_type -> publicproto.Empty
	0,  // 9: hpc.Hpc.AddUserWithGroup:input_type -> hpc.AddUserWithGroupRequest
	2,  // 10: hpc.Hpc.AddUserToGroup:input_type -> hpc.AddUserToGroupRequest
	4,  // 11: hpc.Hpc.GetUserInfoByID:input_type -> hpc.GetUserInfoByIDRequest
	6,  // 12: hpc.Hpc.GetGroupInfoByID:input_type -> hpc.GetGroupInfoByIDRequest
	8,  // 13: hpc.Hpc.GetNodeUsage:input_type -> hpc.GetNodeUsageRequest
	15, // 14: hpc.Hpc.Ping:output_type -> publicproto.PingResponse
	1,  // 15: hpc.Hpc.AddUserWithGroup:output_type -> hpc.AddUserWithGroupResponse
	3,  // 16: hpc.Hpc.AddUserToGroup:output_type -> hpc.AddUserToGroupResponse
	5,  // 17: hpc.Hpc.GetUserInfoByID:output_type -> hpc.GetUserInfoByIDResponse
	7,  // 18: hpc.Hpc.GetGroupInfoByID:output_type -> hpc.GetGroupInfoByIDResponse
	9,  // 19: hpc.Hpc.GetNodeUsage:output_type -> hpc.GetNodeUsageResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_hpc_proto_init() }
func file_proto_hpc_proto_init() {
	if File_proto_hpc_proto != nil {
		return
	}
	file_proto_hpcpublic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_hpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserWithGroupRequest); i {
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
		file_proto_hpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserWithGroupResponse); i {
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
		file_proto_hpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToGroupRequest); i {
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
		file_proto_hpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddUserToGroupResponse); i {
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
		file_proto_hpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoByIDRequest); i {
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
		file_proto_hpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoByIDResponse); i {
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
		file_proto_hpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupInfoByIDRequest); i {
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
		file_proto_hpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupInfoByIDResponse); i {
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
		file_proto_hpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeUsageRequest); i {
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
		file_proto_hpc_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeUsageResponse); i {
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
			RawDescriptor: file_proto_hpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_hpc_proto_goTypes,
		DependencyIndexes: file_proto_hpc_proto_depIdxs,
		MessageInfos:      file_proto_hpc_proto_msgTypes,
	}.Build()
	File_proto_hpc_proto = out.File
	file_proto_hpc_proto_rawDesc = nil
	file_proto_hpc_proto_goTypes = nil
	file_proto_hpc_proto_depIdxs = nil
}
