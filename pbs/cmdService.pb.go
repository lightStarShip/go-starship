// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: cmdService.proto

package pbs

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

type EmptyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRequest) Reset() {
	*x = EmptyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRequest) ProtoMessage() {}

func (x *EmptyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRequest.ProtoReflect.Descriptor instead.
func (*EmptyRequest) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{0}
}

type UserCounterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
}

func (x *UserCounterReq) Reset() {
	*x = UserCounterReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserCounterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserCounterReq) ProtoMessage() {}

func (x *UserCounterReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserCounterReq.ProtoReflect.Descriptor instead.
func (*UserCounterReq) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{1}
}

func (x *UserCounterReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type CounterResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bucket int32 `protobuf:"varint,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
}

func (x *CounterResult) Reset() {
	*x = CounterResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CounterResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CounterResult) ProtoMessage() {}

func (x *CounterResult) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CounterResult.ProtoReflect.Descriptor instead.
func (*CounterResult) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{2}
}

func (x *CounterResult) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CounterResult) GetBucket() int32 {
	if x != nil {
		return x.Bucket
	}
	return 0
}

type EmptyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyReq) Reset() {
	*x = EmptyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyReq) ProtoMessage() {}

func (x *EmptyReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyReq.ProtoReflect.Descriptor instead.
func (*EmptyReq) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{3}
}

type ReceiptReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Report int32  `protobuf:"varint,2,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *ReceiptReq) Reset() {
	*x = ReceiptReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptReq) ProtoMessage() {}

func (x *ReceiptReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptReq.ProtoReflect.Descriptor instead.
func (*ReceiptReq) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{4}
}

func (x *ReceiptReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReceiptReq) GetReport() int32 {
	if x != nil {
		return x.Report
	}
	return 0
}

type ReceiptOneReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User   string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Credit string `protobuf:"bytes,2,opt,name=credit,proto3" json:"credit,omitempty"`
	Report int32  `protobuf:"varint,3,opt,name=report,proto3" json:"report,omitempty"`
}

func (x *ReceiptOneReq) Reset() {
	*x = ReceiptOneReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiptOneReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiptOneReq) ProtoMessage() {}

func (x *ReceiptOneReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiptOneReq.ProtoReflect.Descriptor instead.
func (*ReceiptOneReq) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{5}
}

func (x *ReceiptOneReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ReceiptOneReq) GetCredit() string {
	if x != nil {
		return x.Credit
	}
	return ""
}

func (x *ReceiptOneReq) GetReport() int32 {
	if x != nil {
		return x.Report
	}
	return 0
}

type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{6}
}

func (x *UserInfoReq) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{7}
}

func (x *CommonResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LogLevel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Module string `protobuf:"bytes,1,opt,name=Module,proto3" json:"Module,omitempty"`
	Level  int32  `protobuf:"varint,2,opt,name=Level,proto3" json:"Level,omitempty"`
}

func (x *LogLevel) Reset() {
	*x = LogLevel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogLevel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogLevel) ProtoMessage() {}

func (x *LogLevel) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogLevel.ProtoReflect.Descriptor instead.
func (*LogLevel) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{8}
}

func (x *LogLevel) GetModule() string {
	if x != nil {
		return x.Module
	}
	return ""
}

func (x *LogLevel) GetLevel() int32 {
	if x != nil {
		return x.Level
	}
	return 0
}

type AccessAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Adddr string `protobuf:"bytes,1,opt,name=Adddr,proto3" json:"Adddr,omitempty"`
	Op    int32  `protobuf:"varint,2,opt,name=op,proto3" json:"op,omitempty"`
}

func (x *AccessAddress) Reset() {
	*x = AccessAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessAddress) ProtoMessage() {}

func (x *AccessAddress) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessAddress.ProtoReflect.Descriptor instead.
func (*AccessAddress) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{9}
}

func (x *AccessAddress) GetAdddr() string {
	if x != nil {
		return x.Adddr
	}
	return ""
}

func (x *AccessAddress) GetOp() int32 {
	if x != nil {
		return x.Op
	}
	return 0
}

type WebPort struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Port int32 `protobuf:"varint,1,opt,name=Port,proto3" json:"Port,omitempty"`
}

func (x *WebPort) Reset() {
	*x = WebPort{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmdService_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebPort) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebPort) ProtoMessage() {}

func (x *WebPort) ProtoReflect() protoreflect.Message {
	mi := &file_cmdService_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebPort.ProtoReflect.Descriptor instead.
func (*WebPort) Descriptor() ([]byte, []int) {
	return file_cmdService_proto_rawDescGZIP(), []int{10}
}

func (x *WebPort) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

var File_cmdService_proto protoreflect.FileDescriptor

var file_cmdService_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x6d, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x70, 0x62, 0x73, 0x22, 0x0e, 0x0a, 0x0c, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x73, 0x65, 0x72, 0x22, 0x37, 0x0a,
	0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x22, 0x38, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x53, 0x0a, 0x0d,
	0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x4f, 0x6e, 0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x22, 0x21, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x22, 0x22, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x38, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x4c, 0x65, 0x76,
	0x65, 0x6c, 0x22, 0x35, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x41, 0x64, 0x64, 0x64, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x41, 0x64, 0x64, 0x64, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x6f, 0x70, 0x22, 0x1d, 0x0a, 0x07, 0x57, 0x65, 0x62,
	0x50, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x50, 0x6f, 0x72, 0x74, 0x32, 0x41, 0x0a, 0x0a, 0x43, 0x6d, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x33, 0x0a, 0x0b, 0x53, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0d, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x4c, 0x6f, 0x67, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x64, 0x65, 0x73, 0x6c,
	0x61, 0x62, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x62, 0x73,
	0x3b, 0x70, 0x62, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cmdService_proto_rawDescOnce sync.Once
	file_cmdService_proto_rawDescData = file_cmdService_proto_rawDesc
)

func file_cmdService_proto_rawDescGZIP() []byte {
	file_cmdService_proto_rawDescOnce.Do(func() {
		file_cmdService_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmdService_proto_rawDescData)
	})
	return file_cmdService_proto_rawDescData
}

var file_cmdService_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_cmdService_proto_goTypes = []interface{}{
	(*EmptyRequest)(nil),   // 0: pbs.EmptyRequest
	(*UserCounterReq)(nil), // 1: pbs.UserCounterReq
	(*CounterResult)(nil),  // 2: pbs.CounterResult
	(*EmptyReq)(nil),       // 3: pbs.EmptyReq
	(*ReceiptReq)(nil),     // 4: pbs.ReceiptReq
	(*ReceiptOneReq)(nil),  // 5: pbs.ReceiptOneReq
	(*UserInfoReq)(nil),    // 6: pbs.UserInfoReq
	(*CommonResponse)(nil), // 7: pbs.CommonResponse
	(*LogLevel)(nil),       // 8: pbs.LogLevel
	(*AccessAddress)(nil),  // 9: pbs.AccessAddress
	(*WebPort)(nil),        // 10: pbs.WebPort
}
var file_cmdService_proto_depIdxs = []int32{
	8, // 0: pbs.CmdService.SetLogLevel:input_type -> pbs.LogLevel
	7, // 1: pbs.CmdService.SetLogLevel:output_type -> pbs.CommonResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_cmdService_proto_init() }
func file_cmdService_proto_init() {
	if File_cmdService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmdService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRequest); i {
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
		file_cmdService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserCounterReq); i {
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
		file_cmdService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CounterResult); i {
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
		file_cmdService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyReq); i {
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
		file_cmdService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiptReq); i {
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
		file_cmdService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiptOneReq); i {
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
		file_cmdService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoReq); i {
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
		file_cmdService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
		file_cmdService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogLevel); i {
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
		file_cmdService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessAddress); i {
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
		file_cmdService_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebPort); i {
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
			RawDescriptor: file_cmdService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cmdService_proto_goTypes,
		DependencyIndexes: file_cmdService_proto_depIdxs,
		MessageInfos:      file_cmdService_proto_msgTypes,
	}.Build()
	File_cmdService_proto = out.File
	file_cmdService_proto_rawDesc = nil
	file_cmdService_proto_goTypes = nil
	file_cmdService_proto_depIdxs = nil
}
