// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/usergroup.proto

package user

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

type GetGroupInfoByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	GroupID     int32              `protobuf:"varint,2,opt,name=groupID,proto3" json:"groupID,omitempty"`
}

func (x *GetGroupInfoByIDRequest) Reset() {
	*x = GetGroupInfoByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupInfoByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupInfoByIDRequest) ProtoMessage() {}

func (x *GetGroupInfoByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[0]
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
	return file_proto_usergroup_proto_rawDescGZIP(), []int{0}
}

func (x *GetGroupInfoByIDRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetGroupInfoByIDRequest) GetGroupID() int32 {
	if x != nil {
		return x.GroupID
	}
	return 0
}

type GetGroupInfoByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupInfo *GroupInfo `protobuf:"bytes,1,opt,name=groupInfo,proto3" json:"groupInfo,omitempty"`
}

func (x *GetGroupInfoByIDResponse) Reset() {
	*x = GetGroupInfoByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupInfoByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupInfoByIDResponse) ProtoMessage() {}

func (x *GetGroupInfoByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[1]
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
	return file_proto_usergroup_proto_rawDescGZIP(), []int{1}
}

func (x *GetGroupInfoByIDResponse) GetGroupInfo() *GroupInfo {
	if x != nil {
		return x.GroupInfo
	}
	return nil
}

type PaginationGetGroupInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	PageSize    int32              `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	PageIndex   int32              `protobuf:"varint,3,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
}

func (x *PaginationGetGroupInfoRequest) Reset() {
	*x = PaginationGetGroupInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationGetGroupInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationGetGroupInfoRequest) ProtoMessage() {}

func (x *PaginationGetGroupInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationGetGroupInfoRequest.ProtoReflect.Descriptor instead.
func (*PaginationGetGroupInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_usergroup_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationGetGroupInfoRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PaginationGetGroupInfoRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *PaginationGetGroupInfoRequest) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

type PaginationGetGroupInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GroupInfos []*GroupInfo `protobuf:"bytes,1,rep,name=groupInfos,proto3" json:"groupInfos,omitempty"`
	Count      int32        `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PaginationGetGroupInfoResponse) Reset() {
	*x = PaginationGetGroupInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationGetGroupInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationGetGroupInfoResponse) ProtoMessage() {}

func (x *PaginationGetGroupInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationGetGroupInfoResponse.ProtoReflect.Descriptor instead.
func (*PaginationGetGroupInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_usergroup_proto_rawDescGZIP(), []int{3}
}

func (x *PaginationGetGroupInfoResponse) GetGroupInfos() []*GroupInfo {
	if x != nil {
		return x.GroupInfos
	}
	return nil
}

func (x *PaginationGetGroupInfoResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CreateJoinGroupApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest  *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	ApplyGroupID int32              `protobuf:"varint,2,opt,name=applyGroupID,proto3" json:"applyGroupID,omitempty"`
}

func (x *CreateJoinGroupApplyRequest) Reset() {
	*x = CreateJoinGroupApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJoinGroupApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJoinGroupApplyRequest) ProtoMessage() {}

func (x *CreateJoinGroupApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJoinGroupApplyRequest.ProtoReflect.Descriptor instead.
func (*CreateJoinGroupApplyRequest) Descriptor() ([]byte, []int) {
	return file_proto_usergroup_proto_rawDescGZIP(), []int{4}
}

func (x *CreateJoinGroupApplyRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *CreateJoinGroupApplyRequest) GetApplyGroupID() int32 {
	if x != nil {
		return x.ApplyGroupID
	}
	return 0
}

type CreateJoinGroupApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool  `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ApplyID int32 `protobuf:"varint,2,opt,name=applyID,proto3" json:"applyID,omitempty"`
}

func (x *CreateJoinGroupApplyResponse) Reset() {
	*x = CreateJoinGroupApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateJoinGroupApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateJoinGroupApplyResponse) ProtoMessage() {}

func (x *CreateJoinGroupApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateJoinGroupApplyResponse.ProtoReflect.Descriptor instead.
func (*CreateJoinGroupApplyResponse) Descriptor() ([]byte, []int) {
	return file_proto_usergroup_proto_rawDescGZIP(), []int{5}
}

func (x *CreateJoinGroupApplyResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CreateJoinGroupApplyResponse) GetApplyID() int32 {
	if x != nil {
		return x.ApplyID
	}
	return 0
}

type SearchTutorInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	Username    string             `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *SearchTutorInfoRequest) Reset() {
	*x = SearchTutorInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTutorInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTutorInfoRequest) ProtoMessage() {}

func (x *SearchTutorInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTutorInfoRequest.ProtoReflect.Descriptor instead.
func (*SearchTutorInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_usergroup_proto_rawDescGZIP(), []int{6}
}

func (x *SearchTutorInfoRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *SearchTutorInfoRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type SearchTutorInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TutorID       int32  `protobuf:"varint,1,opt,name=tutorID,proto3" json:"tutorID,omitempty"`
	TutorUsername string `protobuf:"bytes,2,opt,name=tutorUsername,proto3" json:"tutorUsername,omitempty"`
	TutorName     string `protobuf:"bytes,3,opt,name=tutorName,proto3" json:"tutorName,omitempty"`
	GroupID       int32  `protobuf:"varint,4,opt,name=groupID,proto3" json:"groupID,omitempty"`
	GroupName     string `protobuf:"bytes,5,opt,name=groupName,proto3" json:"groupName,omitempty"`
}

func (x *SearchTutorInfoResponse) Reset() {
	*x = SearchTutorInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchTutorInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchTutorInfoResponse) ProtoMessage() {}

func (x *SearchTutorInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchTutorInfoResponse.ProtoReflect.Descriptor instead.
func (*SearchTutorInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_usergroup_proto_rawDescGZIP(), []int{7}
}

func (x *SearchTutorInfoResponse) GetTutorID() int32 {
	if x != nil {
		return x.TutorID
	}
	return 0
}

func (x *SearchTutorInfoResponse) GetTutorUsername() string {
	if x != nil {
		return x.TutorUsername
	}
	return ""
}

func (x *SearchTutorInfoResponse) GetTutorName() string {
	if x != nil {
		return x.TutorName
	}
	return ""
}

func (x *SearchTutorInfoResponse) GetGroupID() int32 {
	if x != nil {
		return x.GroupID
	}
	return 0
}

func (x *SearchTutorInfoResponse) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

type PageGetApplyGroupInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	PageIndex   int32              `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize    int32              `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PageGetApplyGroupInfoRequest) Reset() {
	*x = PageGetApplyGroupInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGetApplyGroupInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGetApplyGroupInfoRequest) ProtoMessage() {}

func (x *PageGetApplyGroupInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGetApplyGroupInfoRequest.ProtoReflect.Descriptor instead.
func (*PageGetApplyGroupInfoRequest) Descriptor() ([]byte, []int) {
	return file_proto_usergroup_proto_rawDescGZIP(), []int{8}
}

func (x *PageGetApplyGroupInfoRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PageGetApplyGroupInfoRequest) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *PageGetApplyGroupInfoRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PageGetApplyGroupInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Applies []*UserGroupApply `protobuf:"bytes,1,rep,name=applies,proto3" json:"applies,omitempty"`
	Count   int32             `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PageGetApplyGroupInfoResponse) Reset() {
	*x = PageGetApplyGroupInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_usergroup_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageGetApplyGroupInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageGetApplyGroupInfoResponse) ProtoMessage() {}

func (x *PageGetApplyGroupInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_usergroup_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageGetApplyGroupInfoResponse.ProtoReflect.Descriptor instead.
func (*PageGetApplyGroupInfoResponse) Descriptor() ([]byte, []int) {
	return file_proto_usergroup_proto_rawDescGZIP(), []int{9}
}

func (x *PageGetApplyGroupInfoResponse) GetApplies() []*UserGroupApply {
	if x != nil {
		return x.Applies
	}
	return nil
}

func (x *PageGetApplyGroupInfoResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_proto_usergroup_proto protoreflect.FileDescriptor

var file_proto_usergroup_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x10, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22, 0x49, 0x0a, 0x18, 0x47, 0x65, 0x74,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x91, 0x01, 0x0a, 0x1d, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x67, 0x0a, 0x1e, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x0a, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0a, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x79, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x70, 0x70, 0x6c,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x44, 0x22, 0x52, 0x0a, 0x1c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49,
	0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x44,
	0x22, 0x6c, 0x0a, 0x16, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xaf,
	0x01, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x75, 0x74,
	0x6f, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x75, 0x74,
	0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x90, 0x01, 0x0a, 0x1c, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x65, 0x0a, 0x1d, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x9a, 0x04, 0x0a, 0x0c, 0x47,
	0x72, 0x6f, 0x75, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47,
	0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x16, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x5f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x21, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4a, 0x6f, 0x69, 0x6e, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x0f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x75, 0x74, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x54, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x15, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x22, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x47, 0x65, 0x74, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x73, 0x61, 0x79, 0x5a, 0x57, 0x2f, 0x68, 0x70,
	0x63, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_usergroup_proto_rawDescOnce sync.Once
	file_proto_usergroup_proto_rawDescData = file_proto_usergroup_proto_rawDesc
)

func file_proto_usergroup_proto_rawDescGZIP() []byte {
	file_proto_usergroup_proto_rawDescOnce.Do(func() {
		file_proto_usergroup_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_usergroup_proto_rawDescData)
	})
	return file_proto_usergroup_proto_rawDescData
}

var file_proto_usergroup_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_usergroup_proto_goTypes = []interface{}{
	(*GetGroupInfoByIDRequest)(nil),        // 0: user.GetGroupInfoByIDRequest
	(*GetGroupInfoByIDResponse)(nil),       // 1: user.GetGroupInfoByIDResponse
	(*PaginationGetGroupInfoRequest)(nil),  // 2: user.PaginationGetGroupInfoRequest
	(*PaginationGetGroupInfoResponse)(nil), // 3: user.PaginationGetGroupInfoResponse
	(*CreateJoinGroupApplyRequest)(nil),    // 4: user.CreateJoinGroupApplyRequest
	(*CreateJoinGroupApplyResponse)(nil),   // 5: user.CreateJoinGroupApplyResponse
	(*SearchTutorInfoRequest)(nil),         // 6: user.SearchTutorInfoRequest
	(*SearchTutorInfoResponse)(nil),        // 7: user.SearchTutorInfoResponse
	(*PageGetApplyGroupInfoRequest)(nil),   // 8: user.PageGetApplyGroupInfoRequest
	(*PageGetApplyGroupInfoResponse)(nil),  // 9: user.PageGetApplyGroupInfoResponse
	(*proto.BaseRequest)(nil),              // 10: request.BaseRequest
	(*GroupInfo)(nil),                      // 11: user.GroupInfo
	(*UserGroupApply)(nil),                 // 12: user.UserGroupApply
	(*proto1.Empty)(nil),                   // 13: publicproto.Empty
	(*proto1.PingResponse)(nil),            // 14: publicproto.PingResponse
}
var file_proto_usergroup_proto_depIdxs = []int32{
	10, // 0: user.GetGroupInfoByIDRequest.baseRequest:type_name -> request.BaseRequest
	11, // 1: user.GetGroupInfoByIDResponse.groupInfo:type_name -> user.GroupInfo
	10, // 2: user.PaginationGetGroupInfoRequest.baseRequest:type_name -> request.BaseRequest
	11, // 3: user.PaginationGetGroupInfoResponse.groupInfos:type_name -> user.GroupInfo
	10, // 4: user.CreateJoinGroupApplyRequest.baseRequest:type_name -> request.BaseRequest
	10, // 5: user.SearchTutorInfoRequest.baseRequest:type_name -> request.BaseRequest
	10, // 6: user.PageGetApplyGroupInfoRequest.baseRequest:type_name -> request.BaseRequest
	12, // 7: user.PageGetApplyGroupInfoResponse.applies:type_name -> user.UserGroupApply
	13, // 8: user.GroupService.Ping:input_type -> publicproto.Empty
	0,  // 9: user.GroupService.GetGroupInfoByID:input_type -> user.GetGroupInfoByIDRequest
	2,  // 10: user.GroupService.PaginationGetGroupInfo:input_type -> user.PaginationGetGroupInfoRequest
	4,  // 11: user.GroupService.CreateJoinGroupApply:input_type -> user.CreateJoinGroupApplyRequest
	6,  // 12: user.GroupService.SearchTutorInfo:input_type -> user.SearchTutorInfoRequest
	8,  // 13: user.GroupService.PageGetApplyGroupInfo:input_type -> user.PageGetApplyGroupInfoRequest
	14, // 14: user.GroupService.Ping:output_type -> publicproto.PingResponse
	1,  // 15: user.GroupService.GetGroupInfoByID:output_type -> user.GetGroupInfoByIDResponse
	3,  // 16: user.GroupService.PaginationGetGroupInfo:output_type -> user.PaginationGetGroupInfoResponse
	5,  // 17: user.GroupService.CreateJoinGroupApply:output_type -> user.CreateJoinGroupApplyResponse
	7,  // 18: user.GroupService.SearchTutorInfo:output_type -> user.SearchTutorInfoResponse
	9,  // 19: user.GroupService.PageGetApplyGroupInfo:output_type -> user.PageGetApplyGroupInfoResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_usergroup_proto_init() }
func file_proto_usergroup_proto_init() {
	if File_proto_usergroup_proto != nil {
		return
	}
	file_proto_userpublic_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_usergroup_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_usergroup_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_usergroup_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationGetGroupInfoRequest); i {
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
		file_proto_usergroup_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationGetGroupInfoResponse); i {
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
		file_proto_usergroup_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJoinGroupApplyRequest); i {
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
		file_proto_usergroup_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateJoinGroupApplyResponse); i {
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
		file_proto_usergroup_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTutorInfoRequest); i {
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
		file_proto_usergroup_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchTutorInfoResponse); i {
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
		file_proto_usergroup_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGetApplyGroupInfoRequest); i {
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
		file_proto_usergroup_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageGetApplyGroupInfoResponse); i {
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
			RawDescriptor: file_proto_usergroup_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_usergroup_proto_goTypes,
		DependencyIndexes: file_proto_usergroup_proto_depIdxs,
		MessageInfos:      file_proto_usergroup_proto_msgTypes,
	}.Build()
	File_proto_usergroup_proto = out.File
	file_proto_usergroup_proto_rawDesc = nil
	file_proto_usergroup_proto_goTypes = nil
	file_proto_usergroup_proto_depIdxs = nil
}
