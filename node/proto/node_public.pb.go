// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/node_public.proto

package node

import (
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

// NodeApply 数据库node_appl对应的消息映射
type NodeApply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreateTime             int64  `protobuf:"varint,2,opt,name=createTime,proto3" json:"createTime,omitempty"`
	CreaterID              int32  `protobuf:"varint,3,opt,name=createrID,proto3" json:"createrID,omitempty"`
	CreaterUsername        string `protobuf:"bytes,4,opt,name=createrUsername,proto3" json:"createrUsername,omitempty"`
	CreaterName            string `protobuf:"bytes,5,opt,name=createrName,proto3" json:"createrName,omitempty"`
	ProjectID              int32  `protobuf:"varint,6,opt,name=projectID,proto3" json:"projectID,omitempty"`
	TutorCheckStatus       int32  `protobuf:"varint,7,opt,name=tutorCheckStatus,proto3" json:"tutorCheckStatus,omitempty"`
	ManagerCheckStatus     int32  `protobuf:"varint,8,opt,name=managerCheckStatus,proto3" json:"managerCheckStatus,omitempty"`
	Status                 int32  `protobuf:"varint,9,opt,name=status,proto3" json:"status,omitempty"`
	MessageTutor           string `protobuf:"bytes,10,opt,name=messageTutor,proto3" json:"messageTutor,omitempty"`
	MessageManager         string `protobuf:"bytes,11,opt,name=messageManager,proto3" json:"messageManager,omitempty"`
	TutorCheckTime         int64  `protobuf:"varint,12,opt,name=tutorCheckTime,proto3" json:"tutorCheckTime,omitempty"`
	TutorID                int32  `protobuf:"varint,13,opt,name=tutorID,proto3" json:"tutorID,omitempty"`
	TutorUsername          string `protobuf:"bytes,14,opt,name=tutorUsername,proto3" json:"tutorUsername,omitempty"`
	TutorName              string `protobuf:"bytes,15,opt,name=tutorName,proto3" json:"tutorName,omitempty"`
	ManagerCheckTime       int64  `protobuf:"varint,16,opt,name=managerCheckTime,proto3" json:"managerCheckTime,omitempty"`
	ManagerCheckerID       int32  `protobuf:"varint,17,opt,name=managerCheckerID,proto3" json:"managerCheckerID,omitempty"`
	ManagerCheckerUsername string `protobuf:"bytes,18,opt,name=managerCheckerUsername,proto3" json:"managerCheckerUsername,omitempty"`
	ManagerCheckerName     string `protobuf:"bytes,19,opt,name=managerCheckerName,proto3" json:"managerCheckerName,omitempty"`
	ModifyTime             int64  `protobuf:"varint,20,opt,name=modifyTime,proto3" json:"modifyTime,omitempty"`
	ModifyUserID           int32  `protobuf:"varint,21,opt,name=modifyUserID,proto3" json:"modifyUserID,omitempty"`
	ModifyName             string `protobuf:"bytes,22,opt,name=modifyName,proto3" json:"modifyName,omitempty"`
	ModifyUsername         string `protobuf:"bytes,23,opt,name=modifyUsername,proto3" json:"modifyUsername,omitempty"`
	NodeType               string `protobuf:"bytes,24,opt,name=nodeType,proto3" json:"nodeType,omitempty"`
	NodeNum                int32  `protobuf:"varint,25,opt,name=nodeNum,proto3" json:"nodeNum,omitempty"`
	StartTime              int64  `protobuf:"varint,26,opt,name=startTime,proto3" json:"startTime,omitempty"`
	EndTime                int64  `protobuf:"varint,27,opt,name=endTime,proto3" json:"endTime,omitempty"`
	ExtraAttributes        string `protobuf:"bytes,28,opt,name=extraAttributes,proto3" json:"extraAttributes,omitempty"`
}

func (x *NodeApply) Reset() {
	*x = NodeApply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_public_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeApply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeApply) ProtoMessage() {}

func (x *NodeApply) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_public_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeApply.ProtoReflect.Descriptor instead.
func (*NodeApply) Descriptor() ([]byte, []int) {
	return file_proto_node_public_proto_rawDescGZIP(), []int{0}
}

func (x *NodeApply) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeApply) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *NodeApply) GetCreaterID() int32 {
	if x != nil {
		return x.CreaterID
	}
	return 0
}

func (x *NodeApply) GetCreaterUsername() string {
	if x != nil {
		return x.CreaterUsername
	}
	return ""
}

func (x *NodeApply) GetCreaterName() string {
	if x != nil {
		return x.CreaterName
	}
	return ""
}

func (x *NodeApply) GetProjectID() int32 {
	if x != nil {
		return x.ProjectID
	}
	return 0
}

func (x *NodeApply) GetTutorCheckStatus() int32 {
	if x != nil {
		return x.TutorCheckStatus
	}
	return 0
}

func (x *NodeApply) GetManagerCheckStatus() int32 {
	if x != nil {
		return x.ManagerCheckStatus
	}
	return 0
}

func (x *NodeApply) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *NodeApply) GetMessageTutor() string {
	if x != nil {
		return x.MessageTutor
	}
	return ""
}

func (x *NodeApply) GetMessageManager() string {
	if x != nil {
		return x.MessageManager
	}
	return ""
}

func (x *NodeApply) GetTutorCheckTime() int64 {
	if x != nil {
		return x.TutorCheckTime
	}
	return 0
}

func (x *NodeApply) GetTutorID() int32 {
	if x != nil {
		return x.TutorID
	}
	return 0
}

func (x *NodeApply) GetTutorUsername() string {
	if x != nil {
		return x.TutorUsername
	}
	return ""
}

func (x *NodeApply) GetTutorName() string {
	if x != nil {
		return x.TutorName
	}
	return ""
}

func (x *NodeApply) GetManagerCheckTime() int64 {
	if x != nil {
		return x.ManagerCheckTime
	}
	return 0
}

func (x *NodeApply) GetManagerCheckerID() int32 {
	if x != nil {
		return x.ManagerCheckerID
	}
	return 0
}

func (x *NodeApply) GetManagerCheckerUsername() string {
	if x != nil {
		return x.ManagerCheckerUsername
	}
	return ""
}

func (x *NodeApply) GetManagerCheckerName() string {
	if x != nil {
		return x.ManagerCheckerName
	}
	return ""
}

func (x *NodeApply) GetModifyTime() int64 {
	if x != nil {
		return x.ModifyTime
	}
	return 0
}

func (x *NodeApply) GetModifyUserID() int32 {
	if x != nil {
		return x.ModifyUserID
	}
	return 0
}

func (x *NodeApply) GetModifyName() string {
	if x != nil {
		return x.ModifyName
	}
	return ""
}

func (x *NodeApply) GetModifyUsername() string {
	if x != nil {
		return x.ModifyUsername
	}
	return ""
}

func (x *NodeApply) GetNodeType() string {
	if x != nil {
		return x.NodeType
	}
	return ""
}

func (x *NodeApply) GetNodeNum() int32 {
	if x != nil {
		return x.NodeNum
	}
	return 0
}

func (x *NodeApply) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *NodeApply) GetEndTime() int64 {
	if x != nil {
		return x.EndTime
	}
	return 0
}

func (x *NodeApply) GetExtraAttributes() string {
	if x != nil {
		return x.ExtraAttributes
	}
	return ""
}

// NodeDistribute 节点分配处理工单对应的消息映射
type NodeDistribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ApplyID          int32  `protobuf:"varint,2,opt,name=applyID,proto3" json:"applyID,omitempty"`
	HandlerFlag      int32  `protobuf:"varint,3,opt,name=handlerFlag,proto3" json:"handlerFlag,omitempty"`
	HandlerUserID    int32  `protobuf:"varint,4,opt,name=handlerUserID,proto3" json:"handlerUserID,omitempty"`
	HandlerUsername  string `protobuf:"bytes,5,opt,name=handlerUsername,proto3" json:"handlerUsername,omitempty"`
	HandlerName      string `protobuf:"bytes,6,opt,name=handlerName,proto3" json:"handlerName,omitempty"`
	DistributeBillID int32  `protobuf:"varint,7,opt,name=distributeBillID,proto3" json:"distributeBillID,omitempty"`
	CreateTime       int64  `protobuf:"varint,8,opt,name=createTime,proto3" json:"createTime,omitempty"`
	ExtraAttributes  string `protobuf:"bytes,9,opt,name=extraAttributes,proto3" json:"extraAttributes,omitempty"`
}

func (x *NodeDistribute) Reset() {
	*x = NodeDistribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_node_public_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeDistribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeDistribute) ProtoMessage() {}

func (x *NodeDistribute) ProtoReflect() protoreflect.Message {
	mi := &file_proto_node_public_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeDistribute.ProtoReflect.Descriptor instead.
func (*NodeDistribute) Descriptor() ([]byte, []int) {
	return file_proto_node_public_proto_rawDescGZIP(), []int{1}
}

func (x *NodeDistribute) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NodeDistribute) GetApplyID() int32 {
	if x != nil {
		return x.ApplyID
	}
	return 0
}

func (x *NodeDistribute) GetHandlerFlag() int32 {
	if x != nil {
		return x.HandlerFlag
	}
	return 0
}

func (x *NodeDistribute) GetHandlerUserID() int32 {
	if x != nil {
		return x.HandlerUserID
	}
	return 0
}

func (x *NodeDistribute) GetHandlerUsername() string {
	if x != nil {
		return x.HandlerUsername
	}
	return ""
}

func (x *NodeDistribute) GetHandlerName() string {
	if x != nil {
		return x.HandlerName
	}
	return ""
}

func (x *NodeDistribute) GetDistributeBillID() int32 {
	if x != nil {
		return x.DistributeBillID
	}
	return 0
}

func (x *NodeDistribute) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *NodeDistribute) GetExtraAttributes() string {
	if x != nil {
		return x.ExtraAttributes
	}
	return ""
}

var File_proto_node_public_proto protoreflect.FileDescriptor

var file_proto_node_public_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x6f, 0x64, 0x65, 0x22,
	0xed, 0x07, 0x0a, 0x09, 0x4e, 0x6f, 0x64, 0x65, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x10, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2e, 0x0a, 0x12, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x12, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x75, 0x74, 0x6f, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x49, 0x44, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x75, 0x74, 0x6f, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x74, 0x75, 0x74, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x18, 0x11, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x10, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x36, 0x0a, 0x16, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x15, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0c, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x1e, 0x0a, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x26, 0x0a, 0x0e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x19,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x1a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x1b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x74,
	0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22,
	0xc4, 0x02, 0x0a, 0x0e, 0x4e, 0x6f, 0x64, 0x65, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b,
	0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x46, 0x6c, 0x61, 0x67, 0x12, 0x24,
	0x0a, 0x0d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x0f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x55,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69,
	0x6c, 0x6c, 0x49, 0x44, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x42, 0x69, 0x6c, 0x6c, 0x49, 0x44, 0x12, 0x1e, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x74, 0x74, 0x72,
	0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x73, 0x61, 0x79, 0x5a, 0x57, 0x2f, 0x68, 0x70, 0x63,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f, 0x6e, 0x6f, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x3b, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_node_public_proto_rawDescOnce sync.Once
	file_proto_node_public_proto_rawDescData = file_proto_node_public_proto_rawDesc
)

func file_proto_node_public_proto_rawDescGZIP() []byte {
	file_proto_node_public_proto_rawDescOnce.Do(func() {
		file_proto_node_public_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_node_public_proto_rawDescData)
	})
	return file_proto_node_public_proto_rawDescData
}

var file_proto_node_public_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_node_public_proto_goTypes = []interface{}{
	(*NodeApply)(nil),      // 0: node.NodeApply
	(*NodeDistribute)(nil), // 1: node.NodeDistribute
}
var file_proto_node_public_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_node_public_proto_init() }
func file_proto_node_public_proto_init() {
	if File_proto_node_public_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_node_public_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeApply); i {
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
		file_proto_node_public_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeDistribute); i {
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
			RawDescriptor: file_proto_node_public_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_node_public_proto_goTypes,
		DependencyIndexes: file_proto_node_public_proto_depIdxs,
		MessageInfos:      file_proto_node_public_proto_msgTypes,
	}.Build()
	File_proto_node_public_proto = out.File
	file_proto_node_public_proto_rawDesc = nil
	file_proto_node_public_proto_goTypes = nil
	file_proto_node_public_proto_depIdxs = nil
}
