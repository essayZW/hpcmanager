// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/award.proto

package award

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

type CreatePaperAwardRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest         *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	Title               string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Category            string             `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Partition           string             `protobuf:"bytes,4,opt,name=partition,proto3" json:"partition,omitempty"`
	FirstPageImageName  string             `protobuf:"bytes,5,opt,name=firstPageImageName,proto3" json:"firstPageImageName,omitempty"`
	ThanksPageImageName string             `protobuf:"bytes,6,opt,name=thanksPageImageName,proto3" json:"thanksPageImageName,omitempty"`
	RemarkMessage       string             `protobuf:"bytes,7,opt,name=remarkMessage,proto3" json:"remarkMessage,omitempty"`
}

func (x *CreatePaperAwardRequest) Reset() {
	*x = CreatePaperAwardRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaperAwardRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaperAwardRequest) ProtoMessage() {}

func (x *CreatePaperAwardRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaperAwardRequest.ProtoReflect.Descriptor instead.
func (*CreatePaperAwardRequest) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePaperAwardRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *CreatePaperAwardRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePaperAwardRequest) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CreatePaperAwardRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *CreatePaperAwardRequest) GetFirstPageImageName() string {
	if x != nil {
		return x.FirstPageImageName
	}
	return ""
}

func (x *CreatePaperAwardRequest) GetThanksPageImageName() string {
	if x != nil {
		return x.ThanksPageImageName
	}
	return ""
}

func (x *CreatePaperAwardRequest) GetRemarkMessage() string {
	if x != nil {
		return x.RemarkMessage
	}
	return ""
}

type CreatePaperAwardResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePaperAwardResponse) Reset() {
	*x = CreatePaperAwardResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePaperAwardResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePaperAwardResponse) ProtoMessage() {}

func (x *CreatePaperAwardResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePaperAwardResponse.ProtoReflect.Descriptor instead.
func (*CreatePaperAwardResponse) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePaperAwardResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PaginationGetPaperApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	PageIndex   int32              `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize    int32              `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PaginationGetPaperApplyRequest) Reset() {
	*x = PaginationGetPaperApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationGetPaperApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationGetPaperApplyRequest) ProtoMessage() {}

func (x *PaginationGetPaperApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationGetPaperApplyRequest.ProtoReflect.Descriptor instead.
func (*PaginationGetPaperApplyRequest) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationGetPaperApplyRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PaginationGetPaperApplyRequest) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *PaginationGetPaperApplyRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PaginationGetPaperApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int32         `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Applies []*PaperApply `protobuf:"bytes,2,rep,name=applies,proto3" json:"applies,omitempty"`
}

func (x *PaginationGetPaperApplyResponse) Reset() {
	*x = PaginationGetPaperApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationGetPaperApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationGetPaperApplyResponse) ProtoMessage() {}

func (x *PaginationGetPaperApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationGetPaperApplyResponse.ProtoReflect.Descriptor instead.
func (*PaginationGetPaperApplyResponse) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{3}
}

func (x *PaginationGetPaperApplyResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PaginationGetPaperApplyResponse) GetApplies() []*PaperApply {
	if x != nil {
		return x.Applies
	}
	return nil
}

type CheckPaperApplyByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest  *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	ApplyID      int32              `protobuf:"varint,2,opt,name=applyID,proto3" json:"applyID,omitempty"`
	Money        float64            `protobuf:"fixed64,3,opt,name=money,proto3" json:"money,omitempty"`
	CheckMessage string             `protobuf:"bytes,4,opt,name=checkMessage,proto3" json:"checkMessage,omitempty"`
	Accept       bool               `protobuf:"varint,5,opt,name=accept,proto3" json:"accept,omitempty"`
}

func (x *CheckPaperApplyByIDRequest) Reset() {
	*x = CheckPaperApplyByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPaperApplyByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPaperApplyByIDRequest) ProtoMessage() {}

func (x *CheckPaperApplyByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPaperApplyByIDRequest.ProtoReflect.Descriptor instead.
func (*CheckPaperApplyByIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{4}
}

func (x *CheckPaperApplyByIDRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *CheckPaperApplyByIDRequest) GetApplyID() int32 {
	if x != nil {
		return x.ApplyID
	}
	return 0
}

func (x *CheckPaperApplyByIDRequest) GetMoney() float64 {
	if x != nil {
		return x.Money
	}
	return 0
}

func (x *CheckPaperApplyByIDRequest) GetCheckMessage() string {
	if x != nil {
		return x.CheckMessage
	}
	return ""
}

func (x *CheckPaperApplyByIDRequest) GetAccept() bool {
	if x != nil {
		return x.Accept
	}
	return false
}

type CheckPaperApplyByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *CheckPaperApplyByIDResponse) Reset() {
	*x = CheckPaperApplyByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPaperApplyByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPaperApplyByIDResponse) ProtoMessage() {}

func (x *CheckPaperApplyByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPaperApplyByIDResponse.ProtoReflect.Descriptor instead.
func (*CheckPaperApplyByIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{5}
}

func (x *CheckPaperApplyByIDResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type CreateTechnologyAwardApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest    *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	ProjectID      int32              `protobuf:"varint,2,opt,name=projectID,proto3" json:"projectID,omitempty"`
	PrizeLevel     string             `protobuf:"bytes,3,opt,name=prizeLevel,proto3" json:"prizeLevel,omitempty"`
	PrizeImageName string             `protobuf:"bytes,4,opt,name=prizeImageName,proto3" json:"prizeImageName,omitempty"`
	RemarkMessage  string             `protobuf:"bytes,5,opt,name=remarkMessage,proto3" json:"remarkMessage,omitempty"`
}

func (x *CreateTechnologyAwardApplyRequest) Reset() {
	*x = CreateTechnologyAwardApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTechnologyAwardApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTechnologyAwardApplyRequest) ProtoMessage() {}

func (x *CreateTechnologyAwardApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTechnologyAwardApplyRequest.ProtoReflect.Descriptor instead.
func (*CreateTechnologyAwardApplyRequest) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{6}
}

func (x *CreateTechnologyAwardApplyRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *CreateTechnologyAwardApplyRequest) GetProjectID() int32 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *CreateTechnologyAwardApplyRequest) GetPrizeLevel() string {
	if x != nil {
		return x.PrizeLevel
	}
	return ""
}

func (x *CreateTechnologyAwardApplyRequest) GetPrizeImageName() string {
	if x != nil {
		return x.PrizeImageName
	}
	return ""
}

func (x *CreateTechnologyAwardApplyRequest) GetRemarkMessage() string {
	if x != nil {
		return x.RemarkMessage
	}
	return ""
}

type CreateTechnologyAwardApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateTechnologyAwardApplyResponse) Reset() {
	*x = CreateTechnologyAwardApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTechnologyAwardApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTechnologyAwardApplyResponse) ProtoMessage() {}

func (x *CreateTechnologyAwardApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTechnologyAwardApplyResponse.ProtoReflect.Descriptor instead.
func (*CreateTechnologyAwardApplyResponse) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{7}
}

func (x *CreateTechnologyAwardApplyResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PaginationGetTechnologyApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	PageIndex   int32              `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize    int32              `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PaginationGetTechnologyApplyRequest) Reset() {
	*x = PaginationGetTechnologyApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationGetTechnologyApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationGetTechnologyApplyRequest) ProtoMessage() {}

func (x *PaginationGetTechnologyApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationGetTechnologyApplyRequest.ProtoReflect.Descriptor instead.
func (*PaginationGetTechnologyApplyRequest) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{8}
}

func (x *PaginationGetTechnologyApplyRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PaginationGetTechnologyApplyRequest) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *PaginationGetTechnologyApplyRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PaginationGetTechnologyApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count   int32              `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	Applies []*TechnologyApply `protobuf:"bytes,2,rep,name=applies,proto3" json:"applies,omitempty"`
}

func (x *PaginationGetTechnologyApplyResponse) Reset() {
	*x = PaginationGetTechnologyApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_award_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationGetTechnologyApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationGetTechnologyApplyResponse) ProtoMessage() {}

func (x *PaginationGetTechnologyApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_award_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationGetTechnologyApplyResponse.ProtoReflect.Descriptor instead.
func (*PaginationGetTechnologyApplyResponse) Descriptor() ([]byte, []int) {
	return file_proto_award_proto_rawDescGZIP(), []int{9}
}

func (x *PaginationGetTechnologyApplyResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PaginationGetTechnologyApplyResponse) GetApplies() []*TechnologyApply {
	if x != nil {
		return x.Applies
	}
	return nil
}

var File_proto_award_proto protoreflect.FileDescriptor

var file_proto_award_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61, 0x77, 0x61, 0x72, 0x64, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x70, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x62, 0x61, 0x73, 0x65, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x77, 0x61, 0x72, 0x64, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x02, 0x0a, 0x17,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x77, 0x61, 0x72, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x12, 0x66, 0x69, 0x72, 0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x30, 0x0a, 0x13, 0x74, 0x68, 0x61, 0x6e, 0x6b, 0x73, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6d, 0x61,
	0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x74, 0x68,
	0x61, 0x6e, 0x6b, 0x73, 0x50, 0x61, 0x67, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x24, 0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x1e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x64, 0x0a, 0x1f, 0x50, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x2b, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x61, 0x70, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0xc0,
	0x01, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a,
	0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x44, 0x12,
	0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05,
	0x6d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63,
	0x65, 0x70, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x70,
	0x74, 0x22, 0x37, 0x0a, 0x1b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0xe7, 0x01, 0x0a, 0x21, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x41,
	0x77, 0x61, 0x72, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x69, 0x7a,
	0x65, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x26, 0x0a, 0x0e, 0x70, 0x72, 0x69, 0x7a, 0x65, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x70, 0x72, 0x69, 0x7a, 0x65, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24,
	0x0a, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x22, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x41, 0x77, 0x61, 0x72, 0x64, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x97, 0x01, 0x0a, 0x23, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x63, 0x68,
	0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61,
	0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70,
	0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x6e, 0x0a, 0x24, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x30, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x54, 0x65, 0x63, 0x68,
	0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x07, 0x61, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x73, 0x32, 0xda, 0x04, 0x0a, 0x0c, 0x41, 0x77, 0x61, 0x72, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x37, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e,
	0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x77, 0x61,
	0x72, 0x64, 0x12, 0x1e, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x77, 0x61, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x17, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x12, 0x25, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x70,
	0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x5e, 0x0a, 0x13, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41,
	0x70, 0x70, 0x6c, 0x79, 0x42, 0x79, 0x49, 0x44, 0x12, 0x21, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x61, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x61, 0x70, 0x65, 0x72, 0x41, 0x70,
	0x70, 0x6c, 0x79, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x73, 0x0a, 0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x63, 0x68, 0x6e,
	0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x41, 0x77, 0x61, 0x72, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12,
	0x28, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65,
	0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x41, 0x77, 0x61, 0x72, 0x64, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29, 0x2e, 0x61, 0x77, 0x61, 0x72,
	0x64, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x41, 0x77, 0x61, 0x72, 0x64, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x79, 0x0a, 0x1c, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f, 0x67,
	0x79, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x2a, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x63, 0x68,
	0x6e, 0x6f, 0x6c, 0x6f, 0x67, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x54, 0x65, 0x63, 0x68, 0x6e, 0x6f, 0x6c, 0x6f,
	0x67, 0x79, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x65, 0x73, 0x73, 0x61, 0x79, 0x5a, 0x57, 0x2f, 0x68, 0x70, 0x63, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x2f, 0x61, 0x77, 0x61, 0x72, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x61,
	0x77, 0x61, 0x72, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_award_proto_rawDescOnce sync.Once
	file_proto_award_proto_rawDescData = file_proto_award_proto_rawDesc
)

func file_proto_award_proto_rawDescGZIP() []byte {
	file_proto_award_proto_rawDescOnce.Do(func() {
		file_proto_award_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_award_proto_rawDescData)
	})
	return file_proto_award_proto_rawDescData
}

var file_proto_award_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_award_proto_goTypes = []interface{}{
	(*CreatePaperAwardRequest)(nil),              // 0: award.CreatePaperAwardRequest
	(*CreatePaperAwardResponse)(nil),             // 1: award.CreatePaperAwardResponse
	(*PaginationGetPaperApplyRequest)(nil),       // 2: award.PaginationGetPaperApplyRequest
	(*PaginationGetPaperApplyResponse)(nil),      // 3: award.PaginationGetPaperApplyResponse
	(*CheckPaperApplyByIDRequest)(nil),           // 4: award.CheckPaperApplyByIDRequest
	(*CheckPaperApplyByIDResponse)(nil),          // 5: award.CheckPaperApplyByIDResponse
	(*CreateTechnologyAwardApplyRequest)(nil),    // 6: award.CreateTechnologyAwardApplyRequest
	(*CreateTechnologyAwardApplyResponse)(nil),   // 7: award.CreateTechnologyAwardApplyResponse
	(*PaginationGetTechnologyApplyRequest)(nil),  // 8: award.PaginationGetTechnologyApplyRequest
	(*PaginationGetTechnologyApplyResponse)(nil), // 9: award.PaginationGetTechnologyApplyResponse
	(*proto.BaseRequest)(nil),                    // 10: request.BaseRequest
	(*PaperApply)(nil),                           // 11: award.PaperApply
	(*TechnologyApply)(nil),                      // 12: award.TechnologyApply
	(*proto1.Empty)(nil),                         // 13: publicproto.Empty
	(*proto1.PingResponse)(nil),                  // 14: publicproto.PingResponse
}
var file_proto_award_proto_depIdxs = []int32{
	10, // 0: award.CreatePaperAwardRequest.baseRequest:type_name -> request.BaseRequest
	10, // 1: award.PaginationGetPaperApplyRequest.baseRequest:type_name -> request.BaseRequest
	11, // 2: award.PaginationGetPaperApplyResponse.applies:type_name -> award.PaperApply
	10, // 3: award.CheckPaperApplyByIDRequest.baseRequest:type_name -> request.BaseRequest
	10, // 4: award.CreateTechnologyAwardApplyRequest.baseRequest:type_name -> request.BaseRequest
	10, // 5: award.PaginationGetTechnologyApplyRequest.baseRequest:type_name -> request.BaseRequest
	12, // 6: award.PaginationGetTechnologyApplyResponse.applies:type_name -> award.TechnologyApply
	13, // 7: award.AwardService.Ping:input_type -> publicproto.Empty
	0,  // 8: award.AwardService.CreatePaperAward:input_type -> award.CreatePaperAwardRequest
	2,  // 9: award.AwardService.PaginationGetPaperApply:input_type -> award.PaginationGetPaperApplyRequest
	4,  // 10: award.AwardService.CheckPaperApplyByID:input_type -> award.CheckPaperApplyByIDRequest
	6,  // 11: award.AwardService.CreateTechnologyAwardApply:input_type -> award.CreateTechnologyAwardApplyRequest
	8,  // 12: award.AwardService.PaginationGetTechnologyApply:input_type -> award.PaginationGetTechnologyApplyRequest
	14, // 13: award.AwardService.Ping:output_type -> publicproto.PingResponse
	1,  // 14: award.AwardService.CreatePaperAward:output_type -> award.CreatePaperAwardResponse
	3,  // 15: award.AwardService.PaginationGetPaperApply:output_type -> award.PaginationGetPaperApplyResponse
	5,  // 16: award.AwardService.CheckPaperApplyByID:output_type -> award.CheckPaperApplyByIDResponse
	7,  // 17: award.AwardService.CreateTechnologyAwardApply:output_type -> award.CreateTechnologyAwardApplyResponse
	9,  // 18: award.AwardService.PaginationGetTechnologyApply:output_type -> award.PaginationGetTechnologyApplyResponse
	13, // [13:19] is the sub-list for method output_type
	7,  // [7:13] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_award_proto_init() }
func file_proto_award_proto_init() {
	if File_proto_award_proto != nil {
		return
	}
	file_proto_award_public_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_award_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaperAwardRequest); i {
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
		file_proto_award_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePaperAwardResponse); i {
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
		file_proto_award_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationGetPaperApplyRequest); i {
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
		file_proto_award_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationGetPaperApplyResponse); i {
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
		file_proto_award_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPaperApplyByIDRequest); i {
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
		file_proto_award_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPaperApplyByIDResponse); i {
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
		file_proto_award_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTechnologyAwardApplyRequest); i {
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
		file_proto_award_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTechnologyAwardApplyResponse); i {
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
		file_proto_award_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationGetTechnologyApplyRequest); i {
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
		file_proto_award_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationGetTechnologyApplyResponse); i {
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
			RawDescriptor: file_proto_award_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_award_proto_goTypes,
		DependencyIndexes: file_proto_award_proto_depIdxs,
		MessageInfos:      file_proto_award_proto_msgTypes,
	}.Build()
	File_proto_award_proto = out.File
	file_proto_award_proto_rawDesc = nil
	file_proto_award_proto_goTypes = nil
	file_proto_award_proto_depIdxs = nil
}
