// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: services.proto

package protobuf

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

type ReturnCode int32

const (
	ReturnCode_SUCCESS ReturnCode = 0
	ReturnCode_ERROR   ReturnCode = 1
)

// Enum value maps for ReturnCode.
var (
	ReturnCode_name = map[int32]string{
		0: "SUCCESS",
		1: "ERROR",
	}
	ReturnCode_value = map[string]int32{
		"SUCCESS": 0,
		"ERROR":   1,
	}
)

func (x ReturnCode) Enum() *ReturnCode {
	p := new(ReturnCode)
	*p = x
	return p
}

func (x ReturnCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReturnCode) Descriptor() protoreflect.EnumDescriptor {
	return file_services_proto_enumTypes[0].Descriptor()
}

func (ReturnCode) Type() protoreflect.EnumType {
	return &file_services_proto_enumTypes[0]
}

func (x ReturnCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReturnCode.Descriptor instead.
func (ReturnCode) EnumDescriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{0}
}

type AddressInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Address       string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *AddressInfoReq) Reset() {
	*x = AddressInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressInfoReq) ProtoMessage() {}

func (x *AddressInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressInfoReq.ProtoReflect.Descriptor instead.
func (*AddressInfoReq) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{0}
}

func (x *AddressInfoReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *AddressInfoReq) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type AddressInfoRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=fe_service.protobuf.ReturnCode" json:"code,omitempty"`
	Msg  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AddressInfoRep) Reset() {
	*x = AddressInfoRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddressInfoRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddressInfoRep) ProtoMessage() {}

func (x *AddressInfoRep) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddressInfoRep.ProtoReflect.Descriptor instead.
func (*AddressInfoRep) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{1}
}

func (x *AddressInfoRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *AddressInfoRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type NftCollectionsInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	TokenAddress  string `protobuf:"bytes,2,opt,name=token_address,json=tokenAddress,proto3" json:"token_address,omitempty"`
}

func (x *NftCollectionsInfoReq) Reset() {
	*x = NftCollectionsInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftCollectionsInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftCollectionsInfoReq) ProtoMessage() {}

func (x *NftCollectionsInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftCollectionsInfoReq.ProtoReflect.Descriptor instead.
func (*NftCollectionsInfoReq) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{2}
}

func (x *NftCollectionsInfoReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *NftCollectionsInfoReq) GetTokenAddress() string {
	if x != nil {
		return x.TokenAddress
	}
	return ""
}

type NftCollectionsInfoRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=fe_service.protobuf.ReturnCode" json:"code,omitempty"`
	Msg  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *NftCollectionsInfoRep) Reset() {
	*x = NftCollectionsInfoRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftCollectionsInfoRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftCollectionsInfoRep) ProtoMessage() {}

func (x *NftCollectionsInfoRep) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftCollectionsInfoRep.ProtoReflect.Descriptor instead.
func (*NftCollectionsInfoRep) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{3}
}

func (x *NftCollectionsInfoRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *NftCollectionsInfoRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type NftInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	TokenId       string `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
}

func (x *NftInfoReq) Reset() {
	*x = NftInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftInfoReq) ProtoMessage() {}

func (x *NftInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftInfoReq.ProtoReflect.Descriptor instead.
func (*NftInfoReq) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{4}
}

func (x *NftInfoReq) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *NftInfoReq) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

type NftInfoRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code ReturnCode `protobuf:"varint,1,opt,name=code,proto3,enum=fe_service.protobuf.ReturnCode" json:"code,omitempty"`
	Msg  string     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *NftInfoRep) Reset() {
	*x = NftInfoRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftInfoRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftInfoRep) ProtoMessage() {}

func (x *NftInfoRep) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftInfoRep.ProtoReflect.Descriptor instead.
func (*NftInfoRep) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{5}
}

func (x *NftInfoRep) GetCode() ReturnCode {
	if x != nil {
		return x.Code
	}
	return ReturnCode_SUCCESS
}

func (x *NftInfoRep) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x66, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x51, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x57, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x66, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x63, 0x0a, 0x15, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x5e, 0x0a, 0x15, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x12,
	0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1f, 0x2e,
	0x66, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x4e, 0x0a, 0x0a, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f,
	0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x0a, 0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x70, 0x12, 0x33, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x66, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x43,
	0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x24, 0x0a, 0x0a, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43,
	0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10,
	0x01, 0x32, 0xb7, 0x02, 0x0a, 0x12, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x23, 0x2e, 0x66, 0x65, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a,
	0x23, 0x2e, 0x66, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x71, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4e, 0x66, 0x74,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x2a, 0x2e, 0x66, 0x65, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x66, 0x65,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4e, 0x66, 0x74, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x4e, 0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x2e, 0x66, 0x65, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e, 0x66,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x66, 0x65, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4e,
	0x66, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x2e,
	0x2e, 0x2f, 0x66, 0x65, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_proto_rawDescOnce sync.Once
	file_services_proto_rawDescData = file_services_proto_rawDesc
)

func file_services_proto_rawDescGZIP() []byte {
	file_services_proto_rawDescOnce.Do(func() {
		file_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_proto_rawDescData)
	})
	return file_services_proto_rawDescData
}

var file_services_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_services_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_services_proto_goTypes = []interface{}{
	(ReturnCode)(0),               // 0: fe_service.protobuf.ReturnCode
	(*AddressInfoReq)(nil),        // 1: fe_service.protobuf.AddressInfoReq
	(*AddressInfoRep)(nil),        // 2: fe_service.protobuf.AddressInfoRep
	(*NftCollectionsInfoReq)(nil), // 3: fe_service.protobuf.NftCollectionsInfoReq
	(*NftCollectionsInfoRep)(nil), // 4: fe_service.protobuf.NftCollectionsInfoRep
	(*NftInfoReq)(nil),            // 5: fe_service.protobuf.NftInfoReq
	(*NftInfoRep)(nil),            // 6: fe_service.protobuf.NftInfoRep
}
var file_services_proto_depIdxs = []int32{
	0, // 0: fe_service.protobuf.AddressInfoRep.code:type_name -> fe_service.protobuf.ReturnCode
	0, // 1: fe_service.protobuf.NftCollectionsInfoRep.code:type_name -> fe_service.protobuf.ReturnCode
	0, // 2: fe_service.protobuf.NftInfoRep.code:type_name -> fe_service.protobuf.ReturnCode
	1, // 3: fe_service.protobuf.InternalRpcService.GetAddressInfo:input_type -> fe_service.protobuf.AddressInfoReq
	3, // 4: fe_service.protobuf.InternalRpcService.GetNftCollectionsInfo:input_type -> fe_service.protobuf.NftCollectionsInfoReq
	5, // 5: fe_service.protobuf.InternalRpcService.GetNftInfo:input_type -> fe_service.protobuf.NftInfoReq
	2, // 6: fe_service.protobuf.InternalRpcService.GetAddressInfo:output_type -> fe_service.protobuf.AddressInfoRep
	4, // 7: fe_service.protobuf.InternalRpcService.GetNftCollectionsInfo:output_type -> fe_service.protobuf.NftCollectionsInfoRep
	6, // 8: fe_service.protobuf.InternalRpcService.GetNftInfo:output_type -> fe_service.protobuf.NftInfoRep
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressInfoReq); i {
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
		file_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddressInfoRep); i {
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
		file_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftCollectionsInfoReq); i {
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
		file_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftCollectionsInfoRep); i {
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
		file_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftInfoReq); i {
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
		file_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftInfoRep); i {
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
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
		EnumInfos:         file_services_proto_enumTypes,
		MessageInfos:      file_services_proto_msgTypes,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}
