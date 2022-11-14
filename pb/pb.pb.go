// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: pb/pb.proto

package pb

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

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User     string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{0}
}

func (x *LoginRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	IsAdminToken bool   `protobuf:"varint,2,opt,name=isAdminToken,proto3" json:"isAdminToken,omitempty"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{1}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginReply) GetIsAdminToken() bool {
	if x != nil {
		return x.IsAdminToken
	}
	return false
}

type RegisterAreaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bearer      string `protobuf:"bytes,1,opt,name=bearer,proto3" json:"bearer,omitempty"`
	Area        string `protobuf:"bytes,2,opt,name=area,proto3" json:"area,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RegisterAreaRequest) Reset() {
	*x = RegisterAreaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAreaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAreaRequest) ProtoMessage() {}

func (x *RegisterAreaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAreaRequest.ProtoReflect.Descriptor instead.
func (*RegisterAreaRequest) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterAreaRequest) GetBearer() string {
	if x != nil {
		return x.Bearer
	}
	return ""
}

func (x *RegisterAreaRequest) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *RegisterAreaRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RegisterAreaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterAreaReply) Reset() {
	*x = RegisterAreaReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterAreaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterAreaReply) ProtoMessage() {}

func (x *RegisterAreaReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterAreaReply.ProtoReflect.Descriptor instead.
func (*RegisterAreaReply) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{3}
}

type RegisterServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bearer      string   `protobuf:"bytes,1,opt,name=bearer,proto3" json:"bearer,omitempty"`
	Service     string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Area        string   `protobuf:"bytes,3,opt,name=area,proto3" json:"area,omitempty"`
	Description string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Protocol    string   `protobuf:"bytes,5,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Host        string   `protobuf:"bytes,6,opt,name=host,proto3" json:"host,omitempty"`
	Port        string   `protobuf:"bytes,7,opt,name=port,proto3" json:"port,omitempty"`
	Tags        []string `protobuf:"bytes,8,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *RegisterServiceRequest) Reset() {
	*x = RegisterServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServiceRequest) ProtoMessage() {}

func (x *RegisterServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServiceRequest.ProtoReflect.Descriptor instead.
func (*RegisterServiceRequest) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{4}
}

func (x *RegisterServiceRequest) GetBearer() string {
	if x != nil {
		return x.Bearer
	}
	return ""
}

func (x *RegisterServiceRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *RegisterServiceRequest) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *RegisterServiceRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *RegisterServiceRequest) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *RegisterServiceRequest) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *RegisterServiceRequest) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *RegisterServiceRequest) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

type RegisterServiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterServiceReply) Reset() {
	*x = RegisterServiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterServiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterServiceReply) ProtoMessage() {}

func (x *RegisterServiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterServiceReply.ProtoReflect.Descriptor instead.
func (*RegisterServiceReply) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{5}
}

type DeleteAreaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bearer string `protobuf:"bytes,1,opt,name=bearer,proto3" json:"bearer,omitempty"`
	Area   string `protobuf:"bytes,2,opt,name=area,proto3" json:"area,omitempty"`
}

func (x *DeleteAreaRequest) Reset() {
	*x = DeleteAreaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAreaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAreaRequest) ProtoMessage() {}

func (x *DeleteAreaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAreaRequest.ProtoReflect.Descriptor instead.
func (*DeleteAreaRequest) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAreaRequest) GetBearer() string {
	if x != nil {
		return x.Bearer
	}
	return ""
}

func (x *DeleteAreaRequest) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

type DeleteAreaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAreaReply) Reset() {
	*x = DeleteAreaReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAreaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAreaReply) ProtoMessage() {}

func (x *DeleteAreaReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAreaReply.ProtoReflect.Descriptor instead.
func (*DeleteAreaReply) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{7}
}

type DeleteServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bearer  string `protobuf:"bytes,1,opt,name=bearer,proto3" json:"bearer,omitempty"`
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Area    string `protobuf:"bytes,3,opt,name=area,proto3" json:"area,omitempty"`
}

func (x *DeleteServiceRequest) Reset() {
	*x = DeleteServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiceRequest) ProtoMessage() {}

func (x *DeleteServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiceRequest.ProtoReflect.Descriptor instead.
func (*DeleteServiceRequest) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteServiceRequest) GetBearer() string {
	if x != nil {
		return x.Bearer
	}
	return ""
}

func (x *DeleteServiceRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *DeleteServiceRequest) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

type DeleteServiceReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteServiceReply) Reset() {
	*x = DeleteServiceReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteServiceReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteServiceReply) ProtoMessage() {}

func (x *DeleteServiceReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteServiceReply.ProtoReflect.Descriptor instead.
func (*DeleteServiceReply) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{9}
}

type ListServiceInAreaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bearer  string `protobuf:"bytes,1,opt,name=bearer,proto3" json:"bearer,omitempty"`
	Service string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Area    string `protobuf:"bytes,3,opt,name=area,proto3" json:"area,omitempty"`
}

func (x *ListServiceInAreaRequest) Reset() {
	*x = ListServiceInAreaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceInAreaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceInAreaRequest) ProtoMessage() {}

func (x *ListServiceInAreaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceInAreaRequest.ProtoReflect.Descriptor instead.
func (*ListServiceInAreaRequest) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{10}
}

func (x *ListServiceInAreaRequest) GetBearer() string {
	if x != nil {
		return x.Bearer
	}
	return ""
}

func (x *ListServiceInAreaRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ListServiceInAreaRequest) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

type ListServiceInAreaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service     string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Area        string   `protobuf:"bytes,2,opt,name=area,proto3" json:"area,omitempty"`
	Description string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Protocol    string   `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	Host        string   `protobuf:"bytes,5,opt,name=host,proto3" json:"host,omitempty"`
	Port        string   `protobuf:"bytes,6,opt,name=port,proto3" json:"port,omitempty"`
	Tags        []string `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
}

func (x *ListServiceInAreaReply) Reset() {
	*x = ListServiceInAreaReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_pb_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListServiceInAreaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListServiceInAreaReply) ProtoMessage() {}

func (x *ListServiceInAreaReply) ProtoReflect() protoreflect.Message {
	mi := &file_pb_pb_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListServiceInAreaReply.ProtoReflect.Descriptor instead.
func (*ListServiceInAreaReply) Descriptor() ([]byte, []int) {
	return file_pb_pb_proto_rawDescGZIP(), []int{11}
}

func (x *ListServiceInAreaReply) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

func (x *ListServiceInAreaReply) GetArea() string {
	if x != nil {
		return x.Area
	}
	return ""
}

func (x *ListServiceInAreaReply) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ListServiceInAreaReply) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *ListServiceInAreaReply) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *ListServiceInAreaReply) GetPort() string {
	if x != nil {
		return x.Port
	}
	return ""
}

func (x *ListServiceInAreaReply) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

var File_pb_pb_proto protoreflect.FileDescriptor

var file_pb_pb_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x70, 0x62, 0x2f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0x3e, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x46, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x69, 0x73, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x63, 0x0a, 0x13, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x13,
	0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0xd8, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x61, 0x72, 0x65, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x22, 0x16,
	0x0a, 0x14, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x3f, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62,
	0x65, 0x61, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x61,
	0x72, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x22, 0x11, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x5c, 0x0a, 0x14, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x61, 0x72, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x60,
	0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x41,
	0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x65,
	0x61, 0x72, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x65, 0x61, 0x72,
	0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x61, 0x72, 0x65, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61,
	0x22, 0xc0, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x65, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x72, 0x65, 0x61, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x74,
	0x61, 0x67, 0x73, 0x32, 0x34, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x2b, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x50, 0x0a, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x12, 0x40, 0x0a, 0x0c, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x12, 0x17, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x5c, 0x0a, 0x0f, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49,
	0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x48, 0x0a, 0x0a, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x12, 0x3a, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x32, 0x54, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x32, 0x64, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x12, 0x4f,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x41,
	0x72, 0x65, 0x61, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x6e, 0x41, 0x72, 0x65, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x69,
	0x63, 0x63, 0x6f, 0x62, 0x69, 0x74, 0x2f, 0x73, 0x69, 0x73, 0x63, 0x6f, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_pb_proto_rawDescOnce sync.Once
	file_pb_pb_proto_rawDescData = file_pb_pb_proto_rawDesc
)

func file_pb_pb_proto_rawDescGZIP() []byte {
	file_pb_pb_proto_rawDescOnce.Do(func() {
		file_pb_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_pb_proto_rawDescData)
	})
	return file_pb_pb_proto_rawDescData
}

var file_pb_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_pb_pb_proto_goTypes = []interface{}{
	(*LoginRequest)(nil),             // 0: pb.LoginRequest
	(*LoginReply)(nil),               // 1: pb.LoginReply
	(*RegisterAreaRequest)(nil),      // 2: pb.RegisterAreaRequest
	(*RegisterAreaReply)(nil),        // 3: pb.RegisterAreaReply
	(*RegisterServiceRequest)(nil),   // 4: pb.RegisterServiceRequest
	(*RegisterServiceReply)(nil),     // 5: pb.RegisterServiceReply
	(*DeleteAreaRequest)(nil),        // 6: pb.DeleteAreaRequest
	(*DeleteAreaReply)(nil),          // 7: pb.DeleteAreaReply
	(*DeleteServiceRequest)(nil),     // 8: pb.DeleteServiceRequest
	(*DeleteServiceReply)(nil),       // 9: pb.DeleteServiceReply
	(*ListServiceInAreaRequest)(nil), // 10: pb.ListServiceInAreaRequest
	(*ListServiceInAreaReply)(nil),   // 11: pb.ListServiceInAreaReply
}
var file_pb_pb_proto_depIdxs = []int32{
	0,  // 0: pb.Login.Login:input_type -> pb.LoginRequest
	2,  // 1: pb.RegisterArea.RegisterArea:input_type -> pb.RegisterAreaRequest
	4,  // 2: pb.RegisterService.RegisterService:input_type -> pb.RegisterServiceRequest
	6,  // 3: pb.DeleteArea.DeleteArea:input_type -> pb.DeleteAreaRequest
	8,  // 4: pb.DeleteService.DeleteService:input_type -> pb.DeleteServiceRequest
	10, // 5: pb.ListServiceInArea.ListServiceInArea:input_type -> pb.ListServiceInAreaRequest
	1,  // 6: pb.Login.Login:output_type -> pb.LoginReply
	3,  // 7: pb.RegisterArea.RegisterArea:output_type -> pb.RegisterAreaReply
	5,  // 8: pb.RegisterService.RegisterService:output_type -> pb.RegisterServiceReply
	7,  // 9: pb.DeleteArea.DeleteArea:output_type -> pb.DeleteAreaReply
	9,  // 10: pb.DeleteService.DeleteService:output_type -> pb.DeleteServiceReply
	11, // 11: pb.ListServiceInArea.ListServiceInArea:output_type -> pb.ListServiceInAreaReply
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_pb_pb_proto_init() }
func file_pb_pb_proto_init() {
	if File_pb_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginRequest); i {
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
		file_pb_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReply); i {
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
		file_pb_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAreaRequest); i {
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
		file_pb_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterAreaReply); i {
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
		file_pb_pb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterServiceRequest); i {
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
		file_pb_pb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterServiceReply); i {
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
		file_pb_pb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAreaRequest); i {
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
		file_pb_pb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAreaReply); i {
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
		file_pb_pb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServiceRequest); i {
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
		file_pb_pb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteServiceReply); i {
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
		file_pb_pb_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceInAreaRequest); i {
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
		file_pb_pb_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListServiceInAreaReply); i {
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
			RawDescriptor: file_pb_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   6,
		},
		GoTypes:           file_pb_pb_proto_goTypes,
		DependencyIndexes: file_pb_pb_proto_depIdxs,
		MessageInfos:      file_pb_pb_proto_msgTypes,
	}.Build()
	File_pb_pb_proto = out.File
	file_pb_pb_proto_rawDesc = nil
	file_pb_pb_proto_goTypes = nil
	file_pb_pb_proto_depIdxs = nil
}
