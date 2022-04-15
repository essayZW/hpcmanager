// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/fee.proto

package fee

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

type CreateNodeDistributeBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest      *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	NodeDistributeID int32              `protobuf:"varint,2,opt,name=nodeDistributeID,proto3" json:"nodeDistributeID,omitempty"`
}

func (x *CreateNodeDistributeBillRequest) Reset() {
	*x = CreateNodeDistributeBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeDistributeBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeDistributeBillRequest) ProtoMessage() {}

func (x *CreateNodeDistributeBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeDistributeBillRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeDistributeBillRequest) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{0}
}

func (x *CreateNodeDistributeBillRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *CreateNodeDistributeBillRequest) GetNodeDistributeID() int32 {
	if x != nil {
		return x.NodeDistributeID
	}
	return 0
}

type CreateNodeDistributeBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateNodeDistributeBillResponse) Reset() {
	*x = CreateNodeDistributeBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeDistributeBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeDistributeBillResponse) ProtoMessage() {}

func (x *CreateNodeDistributeBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeDistributeBillResponse.ProtoReflect.Descriptor instead.
func (*CreateNodeDistributeBillResponse) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{1}
}

func (x *CreateNodeDistributeBillResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type PaginationGetNodeDistributeBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	PageIndex   int32              `protobuf:"varint,2,opt,name=pageIndex,proto3" json:"pageIndex,omitempty"`
	PageSize    int32              `protobuf:"varint,3,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *PaginationGetNodeDistributeBillRequest) Reset() {
	*x = PaginationGetNodeDistributeBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationGetNodeDistributeBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationGetNodeDistributeBillRequest) ProtoMessage() {}

func (x *PaginationGetNodeDistributeBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationGetNodeDistributeBillRequest.ProtoReflect.Descriptor instead.
func (*PaginationGetNodeDistributeBillRequest) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{2}
}

func (x *PaginationGetNodeDistributeBillRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PaginationGetNodeDistributeBillRequest) GetPageIndex() int32 {
	if x != nil {
		return x.PageIndex
	}
	return 0
}

func (x *PaginationGetNodeDistributeBillRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type PaginationGetNodeDistributeBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bills []*NodeDistributeBill `protobuf:"bytes,1,rep,name=bills,proto3" json:"bills,omitempty"`
	Count int32                 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *PaginationGetNodeDistributeBillResponse) Reset() {
	*x = PaginationGetNodeDistributeBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginationGetNodeDistributeBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationGetNodeDistributeBillResponse) ProtoMessage() {}

func (x *PaginationGetNodeDistributeBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationGetNodeDistributeBillResponse.ProtoReflect.Descriptor instead.
func (*PaginationGetNodeDistributeBillResponse) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{3}
}

func (x *PaginationGetNodeDistributeBillResponse) GetBills() []*NodeDistributeBill {
	if x != nil {
		return x.Bills
	}
	return nil
}

func (x *PaginationGetNodeDistributeBillResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type PayNodeDistributeBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	Id          int32              `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	PayMoney    float64            `protobuf:"fixed64,3,opt,name=payMoney,proto3" json:"payMoney,omitempty"`
	PayType     int32              `protobuf:"varint,4,opt,name=payType,proto3" json:"payType,omitempty"`
	PayMessage  string             `protobuf:"bytes,5,opt,name=payMessage,proto3" json:"payMessage,omitempty"`
}

func (x *PayNodeDistributeBillRequest) Reset() {
	*x = PayNodeDistributeBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNodeDistributeBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNodeDistributeBillRequest) ProtoMessage() {}

func (x *PayNodeDistributeBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNodeDistributeBillRequest.ProtoReflect.Descriptor instead.
func (*PayNodeDistributeBillRequest) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{4}
}

func (x *PayNodeDistributeBillRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *PayNodeDistributeBillRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PayNodeDistributeBillRequest) GetPayMoney() float64 {
	if x != nil {
		return x.PayMoney
	}
	return 0
}

func (x *PayNodeDistributeBillRequest) GetPayType() int32 {
	if x != nil {
		return x.PayType
	}
	return 0
}

func (x *PayNodeDistributeBillRequest) GetPayMessage() string {
	if x != nil {
		return x.PayMessage
	}
	return ""
}

type PayNodeDistributeBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *PayNodeDistributeBillResponse) Reset() {
	*x = PayNodeDistributeBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayNodeDistributeBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayNodeDistributeBillResponse) ProtoMessage() {}

func (x *PayNodeDistributeBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayNodeDistributeBillResponse.ProtoReflect.Descriptor instead.
func (*PayNodeDistributeBillResponse) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{5}
}

func (x *PayNodeDistributeBillResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetNodeDistributeFeeRateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
}

func (x *GetNodeDistributeFeeRateRequest) Reset() {
	*x = GetNodeDistributeFeeRateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeDistributeFeeRateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeDistributeFeeRateRequest) ProtoMessage() {}

func (x *GetNodeDistributeFeeRateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeDistributeFeeRateRequest.ProtoReflect.Descriptor instead.
func (*GetNodeDistributeFeeRateRequest) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{6}
}

func (x *GetNodeDistributeFeeRateRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

type GetNodeDistributeFeeRateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rate36CPU float64 `protobuf:"fixed64,1,opt,name=rate36CPU,proto3" json:"rate36CPU,omitempty"`
	Rate4GPU  float64 `protobuf:"fixed64,2,opt,name=rate4GPU,proto3" json:"rate4GPU,omitempty"`
	Rate8GPU  float64 `protobuf:"fixed64,3,opt,name=rate8GPU,proto3" json:"rate8GPU,omitempty"`
}

func (x *GetNodeDistributeFeeRateResponse) Reset() {
	*x = GetNodeDistributeFeeRateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNodeDistributeFeeRateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNodeDistributeFeeRateResponse) ProtoMessage() {}

func (x *GetNodeDistributeFeeRateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNodeDistributeFeeRateResponse.ProtoReflect.Descriptor instead.
func (*GetNodeDistributeFeeRateResponse) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{7}
}

func (x *GetNodeDistributeFeeRateResponse) GetRate36CPU() float64 {
	if x != nil {
		return x.Rate36CPU
	}
	return 0
}

func (x *GetNodeDistributeFeeRateResponse) GetRate4GPU() float64 {
	if x != nil {
		return x.Rate4GPU
	}
	return 0
}

func (x *GetNodeDistributeFeeRateResponse) GetRate8GPU() float64 {
	if x != nil {
		return x.Rate8GPU
	}
	return 0
}

type CreateNodeWeekUsageBillRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest *proto.BaseRequest `protobuf:"bytes,1,opt,name=baseRequest,proto3" json:"baseRequest,omitempty"`
	UserID      int32              `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	WallTime    int32              `protobuf:"varint,3,opt,name=wallTime,proto3" json:"wallTime,omitempty"`
	GwallTime   int32              `protobuf:"varint,4,opt,name=gwallTime,proto3" json:"gwallTime,omitempty"`
	StartTime   int64              `protobuf:"varint,5,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime     int64              `protobuf:"varint,6,opt,name=endTime,proto3" json:"endTime,omitempty"`
}

func (x *CreateNodeWeekUsageBillRequest) Reset() {
	*x = CreateNodeWeekUsageBillRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeWeekUsageBillRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeWeekUsageBillRequest) ProtoMessage() {}

func (x *CreateNodeWeekUsageBillRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeWeekUsageBillRequest.ProtoReflect.Descriptor instead.
func (*CreateNodeWeekUsageBillRequest) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{8}
}

func (x *CreateNodeWeekUsageBillRequest) GetBaseRequest() *proto.BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *CreateNodeWeekUsageBillRequest) GetUserID() int32 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *CreateNodeWeekUsageBillRequest) GetWallTime() int32 {
	if x != nil {
		return x.WallTime
	}
	return 0
}

func (x *CreateNodeWeekUsageBillRequest) GetGwallTime() int32 {
	if x != nil {
		return x.GwallTime
	}
	return 0
}

func (x *CreateNodeWeekUsageBillRequest) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *CreateNodeWeekUsageBillRequest) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

type CreateNodeWeekUsageBillResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateNodeWeekUsageBillResponse) Reset() {
	*x = CreateNodeWeekUsageBillResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_fee_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNodeWeekUsageBillResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNodeWeekUsageBillResponse) ProtoMessage() {}

func (x *CreateNodeWeekUsageBillResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_fee_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNodeWeekUsageBillResponse.ProtoReflect.Descriptor instead.
func (*CreateNodeWeekUsageBillResponse) Descriptor() ([]byte, []int) {
	return file_proto_fee_proto_rawDescGZIP(), []int{9}
}

func (x *CreateNodeWeekUsageBillResponse) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_proto_fee_proto protoreflect.FileDescriptor

var file_proto_fee_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x66, 0x65, 0x65, 0x1a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x69,
	0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x65, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69,
	0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2a, 0x0a, 0x10, 0x6e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x6e, 0x6f, 0x64,
	0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x49, 0x44, 0x22, 0x32, 0x0a,
	0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x9a, 0x01, 0x0a, 0x26, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x64,
	0x65, 0x78, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6e,
	0x0a, 0x27, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x62, 0x69, 0x6c,
	0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x4e,
	0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c,
	0x6c, 0x52, 0x05, 0x62, 0x69, 0x6c, 0x6c, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xbc,
	0x01, 0x0a, 0x1c, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x79, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x61, 0x79, 0x4d, 0x6f,
	0x6e, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x61, 0x79, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a,
	0x0a, 0x70, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x61, 0x79, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a,
	0x1d, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x59, 0x0a, 0x1f, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x46, 0x65, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x0b, 0x62,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x78, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x61, 0x74, 0x65, 0x33,
	0x36, 0x43, 0x50, 0x55, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x72, 0x61, 0x74, 0x65,
	0x33, 0x36, 0x43, 0x50, 0x55, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61, 0x74, 0x65, 0x34, 0x47, 0x50,
	0x55, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x61, 0x74, 0x65, 0x34, 0x47, 0x50,
	0x55, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x61, 0x74, 0x65, 0x38, 0x47, 0x50, 0x55, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x08, 0x72, 0x61, 0x74, 0x65, 0x38, 0x47, 0x50, 0x55, 0x22, 0xe2, 0x01,
	0x0a, 0x1e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x65, 0x65, 0x6b,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x67, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x67, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x31, 0x0a, 0x1f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65,
	0x57, 0x65, 0x65, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xde, 0x04, 0x0a, 0x03, 0x46, 0x65, 0x65, 0x12, 0x37, 0x0a,
	0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x19, 0x2e, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69,
	0x6c, 0x6c, 0x12, 0x24, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c,
	0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x7e, 0x0a, 0x1f, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65,
	0x42, 0x69, 0x6c, 0x6c, 0x12, 0x2b, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2c, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x60, 0x0a, 0x15, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x21, 0x2e, 0x66, 0x65, 0x65,
	0x2e, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x66, 0x65, 0x65, 0x2e, 0x50, 0x61, 0x79, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x69, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69,
	0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x12,
	0x24, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x46, 0x65, 0x65, 0x52, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x66, 0x65, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x4e,
	0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x46, 0x65, 0x65,
	0x52, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x66,
	0x0a, 0x17, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x65, 0x65, 0x6b,
	0x55, 0x73, 0x61, 0x67, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x23, 0x2e, 0x66, 0x65, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x57, 0x65, 0x65, 0x6b, 0x55, 0x73,
	0x61, 0x67, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24,
	0x2e, 0x66, 0x65, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x57,
	0x65, 0x65, 0x6b, 0x55, 0x73, 0x61, 0x67, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x73, 0x61, 0x79, 0x5a, 0x57, 0x2f, 0x68, 0x70, 0x63,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x66, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x3b, 0x66, 0x65, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_fee_proto_rawDescOnce sync.Once
	file_proto_fee_proto_rawDescData = file_proto_fee_proto_rawDesc
)

func file_proto_fee_proto_rawDescGZIP() []byte {
	file_proto_fee_proto_rawDescOnce.Do(func() {
		file_proto_fee_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_fee_proto_rawDescData)
	})
	return file_proto_fee_proto_rawDescData
}

var file_proto_fee_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_fee_proto_goTypes = []interface{}{
	(*CreateNodeDistributeBillRequest)(nil),         // 0: fee.CreateNodeDistributeBillRequest
	(*CreateNodeDistributeBillResponse)(nil),        // 1: fee.CreateNodeDistributeBillResponse
	(*PaginationGetNodeDistributeBillRequest)(nil),  // 2: fee.PaginationGetNodeDistributeBillRequest
	(*PaginationGetNodeDistributeBillResponse)(nil), // 3: fee.PaginationGetNodeDistributeBillResponse
	(*PayNodeDistributeBillRequest)(nil),            // 4: fee.PayNodeDistributeBillRequest
	(*PayNodeDistributeBillResponse)(nil),           // 5: fee.PayNodeDistributeBillResponse
	(*GetNodeDistributeFeeRateRequest)(nil),         // 6: fee.GetNodeDistributeFeeRateRequest
	(*GetNodeDistributeFeeRateResponse)(nil),        // 7: fee.GetNodeDistributeFeeRateResponse
	(*CreateNodeWeekUsageBillRequest)(nil),          // 8: fee.CreateNodeWeekUsageBillRequest
	(*CreateNodeWeekUsageBillResponse)(nil),         // 9: fee.CreateNodeWeekUsageBillResponse
	(*proto.BaseRequest)(nil),                       // 10: request.BaseRequest
	(*NodeDistributeBill)(nil),                      // 11: fee.NodeDistributeBill
	(*proto1.Empty)(nil),                            // 12: publicproto.Empty
	(*proto1.PingResponse)(nil),                     // 13: publicproto.PingResponse
}
var file_proto_fee_proto_depIdxs = []int32{
	10, // 0: fee.CreateNodeDistributeBillRequest.baseRequest:type_name -> request.BaseRequest
	10, // 1: fee.PaginationGetNodeDistributeBillRequest.baseRequest:type_name -> request.BaseRequest
	11, // 2: fee.PaginationGetNodeDistributeBillResponse.bills:type_name -> fee.NodeDistributeBill
	10, // 3: fee.PayNodeDistributeBillRequest.baseRequest:type_name -> request.BaseRequest
	10, // 4: fee.GetNodeDistributeFeeRateRequest.baseRequest:type_name -> request.BaseRequest
	10, // 5: fee.CreateNodeWeekUsageBillRequest.baseRequest:type_name -> request.BaseRequest
	12, // 6: fee.Fee.Ping:input_type -> publicproto.Empty
	0,  // 7: fee.Fee.CreateNodeDistributeBill:input_type -> fee.CreateNodeDistributeBillRequest
	2,  // 8: fee.Fee.PaginationGetNodeDistributeBill:input_type -> fee.PaginationGetNodeDistributeBillRequest
	4,  // 9: fee.Fee.PayNodeDistributeBill:input_type -> fee.PayNodeDistributeBillRequest
	6,  // 10: fee.Fee.GetNodeDistributeFeeRate:input_type -> fee.GetNodeDistributeFeeRateRequest
	8,  // 11: fee.Fee.CreateNodeWeekUsageBill:input_type -> fee.CreateNodeWeekUsageBillRequest
	13, // 12: fee.Fee.Ping:output_type -> publicproto.PingResponse
	1,  // 13: fee.Fee.CreateNodeDistributeBill:output_type -> fee.CreateNodeDistributeBillResponse
	3,  // 14: fee.Fee.PaginationGetNodeDistributeBill:output_type -> fee.PaginationGetNodeDistributeBillResponse
	5,  // 15: fee.Fee.PayNodeDistributeBill:output_type -> fee.PayNodeDistributeBillResponse
	7,  // 16: fee.Fee.GetNodeDistributeFeeRate:output_type -> fee.GetNodeDistributeFeeRateResponse
	9,  // 17: fee.Fee.CreateNodeWeekUsageBill:output_type -> fee.CreateNodeWeekUsageBillResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_proto_fee_proto_init() }
func file_proto_fee_proto_init() {
	if File_proto_fee_proto != nil {
		return
	}
	file_proto_fee_public_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_fee_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeDistributeBillRequest); i {
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
		file_proto_fee_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeDistributeBillResponse); i {
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
		file_proto_fee_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationGetNodeDistributeBillRequest); i {
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
		file_proto_fee_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginationGetNodeDistributeBillResponse); i {
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
		file_proto_fee_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNodeDistributeBillRequest); i {
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
		file_proto_fee_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayNodeDistributeBillResponse); i {
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
		file_proto_fee_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeDistributeFeeRateRequest); i {
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
		file_proto_fee_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNodeDistributeFeeRateResponse); i {
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
		file_proto_fee_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeWeekUsageBillRequest); i {
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
		file_proto_fee_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNodeWeekUsageBillResponse); i {
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
			RawDescriptor: file_proto_fee_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_fee_proto_goTypes,
		DependencyIndexes: file_proto_fee_proto_depIdxs,
		MessageInfos:      file_proto_fee_proto_msgTypes,
	}.Build()
	File_proto_fee_proto = out.File
	file_proto_fee_proto_rawDesc = nil
	file_proto_fee_proto_goTypes = nil
	file_proto_fee_proto_depIdxs = nil
}
