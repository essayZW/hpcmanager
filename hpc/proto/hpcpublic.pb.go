// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: proto/hpcpublic.proto

package hpc

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

// HpcUser hpc_user表的消息映射
type HpcUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NodeUsername    string `protobuf:"bytes,2,opt,name=nodeUsername,proto3" json:"nodeUsername,omitempty"`
	NodeUID         int32  `protobuf:"varint,3,opt,name=nodeUID,proto3" json:"nodeUID,omitempty"`
	NodeMaxQuota    int32  `protobuf:"varint,4,opt,name=nodeMaxQuota,proto3" json:"nodeMaxQuota,omitempty"`
	QuotaStartTime  int64  `protobuf:"varint,5,opt,name=quotaStartTime,proto3" json:"quotaStartTime,omitempty"`
	QuotaEndTime    int64  `protobuf:"varint,6,opt,name=quotaEndTime,proto3" json:"quotaEndTime,omitempty"`
	ExtraAttributes string `protobuf:"bytes,7,opt,name=extraAttributes,proto3" json:"extraAttributes,omitempty"`
}

func (x *HpcUser) Reset() {
	*x = HpcUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpcpublic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HpcUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HpcUser) ProtoMessage() {}

func (x *HpcUser) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpcpublic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HpcUser.ProtoReflect.Descriptor instead.
func (*HpcUser) Descriptor() ([]byte, []int) {
	return file_proto_hpcpublic_proto_rawDescGZIP(), []int{0}
}

func (x *HpcUser) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HpcUser) GetNodeUsername() string {
	if x != nil {
		return x.NodeUsername
	}
	return ""
}

func (x *HpcUser) GetNodeUID() int32 {
	if x != nil {
		return x.NodeUID
	}
	return 0
}

func (x *HpcUser) GetNodeMaxQuota() int32 {
	if x != nil {
		return x.NodeMaxQuota
	}
	return 0
}

func (x *HpcUser) GetQuotaStartTime() int64 {
	if x != nil {
		return x.QuotaStartTime
	}
	return 0
}

func (x *HpcUser) GetQuotaEndTime() int64 {
	if x != nil {
		return x.QuotaEndTime
	}
	return 0
}

func (x *HpcUser) GetExtraAttributes() string {
	if x != nil {
		return x.ExtraAttributes
	}
	return ""
}

// HpcGroup hpc_group表的消息映射
type HpcGroup struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	QueueName       string `protobuf:"bytes,3,opt,name=queueName,proto3" json:"queueName,omitempty"`
	GID             int32  `protobuf:"varint,4,opt,name=gID,proto3" json:"gID,omitempty"`
	ExtraAttributes string `protobuf:"bytes,5,opt,name=extraAttributes,proto3" json:"extraAttributes,omitempty"`
}

func (x *HpcGroup) Reset() {
	*x = HpcGroup{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpcpublic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HpcGroup) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HpcGroup) ProtoMessage() {}

func (x *HpcGroup) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpcpublic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HpcGroup.ProtoReflect.Descriptor instead.
func (*HpcGroup) Descriptor() ([]byte, []int) {
	return file_proto_hpcpublic_proto_rawDescGZIP(), []int{1}
}

func (x *HpcGroup) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *HpcGroup) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *HpcGroup) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *HpcGroup) GetGID() int32 {
	if x != nil {
		return x.GID
	}
	return 0
}

func (x *HpcGroup) GetExtraAttributes() string {
	if x != nil {
		return x.ExtraAttributes
	}
	return ""
}

// HpcNodeUsage 机器节点机器时间使用情况
type HpcNodeUsage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username  string  `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	GroupName string  `protobuf:"bytes,2,opt,name=groupName,proto3" json:"groupName,omitempty"`
	QueueName string  `protobuf:"bytes,3,opt,name=queueName,proto3" json:"queueName,omitempty"`
	WallTime  float64 `protobuf:"fixed64,4,opt,name=wallTime,proto3" json:"wallTime,omitempty"`
	GwallTime float64 `protobuf:"fixed64,5,opt,name=gwallTime,proto3" json:"gwallTime,omitempty"`
}

func (x *HpcNodeUsage) Reset() {
	*x = HpcNodeUsage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_hpcpublic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HpcNodeUsage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HpcNodeUsage) ProtoMessage() {}

func (x *HpcNodeUsage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_hpcpublic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HpcNodeUsage.ProtoReflect.Descriptor instead.
func (*HpcNodeUsage) Descriptor() ([]byte, []int) {
	return file_proto_hpcpublic_proto_rawDescGZIP(), []int{2}
}

func (x *HpcNodeUsage) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *HpcNodeUsage) GetGroupName() string {
	if x != nil {
		return x.GroupName
	}
	return ""
}

func (x *HpcNodeUsage) GetQueueName() string {
	if x != nil {
		return x.QueueName
	}
	return ""
}

func (x *HpcNodeUsage) GetWallTime() float64 {
	if x != nil {
		return x.WallTime
	}
	return 0
}

func (x *HpcNodeUsage) GetGwallTime() float64 {
	if x != nil {
		return x.GwallTime
	}
	return 0
}

var File_proto_hpcpublic_proto protoreflect.FileDescriptor

var file_proto_hpcpublic_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x70, 0x63, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x68, 0x70, 0x63, 0x22, 0xf1, 0x01, 0x0a,
	0x07, 0x48, 0x70, 0x63, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6e, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6e, 0x6f, 0x64, 0x65, 0x55, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x6e,
	0x6f, 0x64, 0x65, 0x55, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x6e, 0x6f, 0x64, 0x65, 0x4d, 0x61,
	0x78, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x6e, 0x6f,
	0x64, 0x65, 0x4d, 0x61, 0x78, 0x51, 0x75, 0x6f, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0e, 0x71, 0x75,
	0x6f, 0x74, 0x61, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0e, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x45, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x71, 0x75, 0x6f, 0x74, 0x61, 0x45,
	0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41,
	0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73,
	0x22, 0x88, 0x01, 0x0a, 0x08, 0x48, 0x70, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x67, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x67, 0x49,
	0x44, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62,
	0x75, 0x74, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x0c,
	0x48, 0x70, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x55, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x67, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x67, 0x77, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x2d,
	0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x73, 0x73,
	0x61, 0x79, 0x5a, 0x57, 0x2f, 0x68, 0x70, 0x63, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2f,
	0x68, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x68, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_hpcpublic_proto_rawDescOnce sync.Once
	file_proto_hpcpublic_proto_rawDescData = file_proto_hpcpublic_proto_rawDesc
)

func file_proto_hpcpublic_proto_rawDescGZIP() []byte {
	file_proto_hpcpublic_proto_rawDescOnce.Do(func() {
		file_proto_hpcpublic_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_hpcpublic_proto_rawDescData)
	})
	return file_proto_hpcpublic_proto_rawDescData
}

var file_proto_hpcpublic_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_hpcpublic_proto_goTypes = []interface{}{
	(*HpcUser)(nil),      // 0: hpc.HpcUser
	(*HpcGroup)(nil),     // 1: hpc.HpcGroup
	(*HpcNodeUsage)(nil), // 2: hpc.HpcNodeUsage
}
var file_proto_hpcpublic_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_hpcpublic_proto_init() }
func file_proto_hpcpublic_proto_init() {
	if File_proto_hpcpublic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_hpcpublic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HpcUser); i {
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
		file_proto_hpcpublic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HpcGroup); i {
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
		file_proto_hpcpublic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HpcNodeUsage); i {
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
			RawDescriptor: file_proto_hpcpublic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_hpcpublic_proto_goTypes,
		DependencyIndexes: file_proto_hpcpublic_proto_depIdxs,
		MessageInfos:      file_proto_hpcpublic_proto_msgTypes,
	}.Build()
	File_proto_hpcpublic_proto = out.File
	file_proto_hpcpublic_proto_rawDesc = nil
	file_proto_hpcpublic_proto_goTypes = nil
	file_proto_hpcpublic_proto_depIdxs = nil
}
