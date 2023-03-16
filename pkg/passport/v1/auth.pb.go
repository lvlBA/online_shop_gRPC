// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: auth.proto

package passport

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

type Auth struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Login string `protobuf:"bytes,2,opt,name=login,proto3" json:"login,omitempty"`
	Token string `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Auth) Reset() {
	*x = Auth{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Auth) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Auth) ProtoMessage() {}

func (x *Auth) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Auth.ProtoReflect.Descriptor instead.
func (*Auth) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Auth) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Auth) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *Auth) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SetUserAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *SetUserAccessRequest) Reset() {
	*x = SetUserAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserAccessRequest) ProtoMessage() {}

func (x *SetUserAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserAccessRequest.ProtoReflect.Descriptor instead.
func (*SetUserAccessRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *SetUserAccessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *SetUserAccessRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type SetUserAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetUserAccessResponse) Reset() {
	*x = SetUserAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserAccessResponse) ProtoMessage() {}

func (x *SetUserAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserAccessResponse.ProtoReflect.Descriptor instead.
func (*SetUserAccessResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

type DeleteUserAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ResourceId string `protobuf:"bytes,2,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
}

func (x *DeleteUserAccessRequest) Reset() {
	*x = DeleteUserAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserAccessRequest) ProtoMessage() {}

func (x *DeleteUserAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserAccessRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserAccessRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteUserAccessRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeleteUserAccessRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

type DeleteUserAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserAccessResponse) Reset() {
	*x = DeleteUserAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserAccessResponse) ProtoMessage() {}

func (x *DeleteUserAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserAccessResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserAccessResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{4}
}

type GetUserTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GetUserTokenRequest) Reset() {
	*x = GetUserTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenRequest) ProtoMessage() {}

func (x *GetUserTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenRequest.ProtoReflect.Descriptor instead.
func (*GetUserTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserTokenRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *GetUserTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetUserTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token []byte `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetUserTokenResponse) Reset() {
	*x = GetUserTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserTokenResponse) ProtoMessage() {}

func (x *GetUserTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserTokenResponse.ProtoReflect.Descriptor instead.
func (*GetUserTokenResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserTokenResponse) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type CheckUserAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	Token      []byte `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *CheckUserAccessRequest) Reset() {
	*x = CheckUserAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAccessRequest) ProtoMessage() {}

func (x *CheckUserAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAccessRequest.ProtoReflect.Descriptor instead.
func (*CheckUserAccessRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{7}
}

func (x *CheckUserAccessRequest) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *CheckUserAccessRequest) GetToken() []byte {
	if x != nil {
		return x.Token
	}
	return nil
}

type CheckUserAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CheckUserAccessResponse) Reset() {
	*x = CheckUserAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckUserAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckUserAccessResponse) ProtoMessage() {}

func (x *CheckUserAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckUserAccessResponse.ProtoReflect.Descriptor instead.
func (*CheckUserAccessResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{8}
}

func (x *CheckUserAccessResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type DeleteUserTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *DeleteUserTokenRequest) Reset() {
	*x = DeleteUserTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTokenRequest) ProtoMessage() {}

func (x *DeleteUserTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTokenRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserTokenRequest) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteUserTokenRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *DeleteUserTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type DeleteUserTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteUserTokenResponse) Reset() {
	*x = DeleteUserTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteUserTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserTokenResponse) ProtoMessage() {}

func (x *DeleteUserTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserTokenResponse.ProtoReflect.Descriptor instead.
func (*DeleteUserTokenResponse) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{10}
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x42, 0x0a, 0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x50, 0x0a, 0x14, 0x53, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x47, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x2c,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4f, 0x0a, 0x16,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x29, 0x0a,
	0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x4a, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xd9, 0x04, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12,
	0x2c, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61,
	0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70,
	0x0a, 0x0d, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x2d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61,
	0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e,
	0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x76, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x2f, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x76, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x2f, 0x2e, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x55, 0x73, 0x65, 0x72,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x79, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x12, 0x30, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x76, 0x6c, 0x42, 0x41, 0x2f,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x70, 0x61, 0x73, 0x73,
	0x70, 0x6f, 0x72, 0x74, 0x3b, 0x70, 0x61, 0x73, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_auth_proto_goTypes = []interface{}{
	(*Auth)(nil),                     // 0: online_shop.passport.v1.Auth
	(*SetUserAccessRequest)(nil),     // 1: online_shop.passport.v1.SetUserAccessRequest
	(*SetUserAccessResponse)(nil),    // 2: online_shop.passport.v1.SetUserAccessResponse
	(*DeleteUserAccessRequest)(nil),  // 3: online_shop.passport.v1.DeleteUserAccessRequest
	(*DeleteUserAccessResponse)(nil), // 4: online_shop.passport.v1.DeleteUserAccessResponse
	(*GetUserTokenRequest)(nil),      // 5: online_shop.passport.v1.GetUserTokenRequest
	(*GetUserTokenResponse)(nil),     // 6: online_shop.passport.v1.GetUserTokenResponse
	(*CheckUserAccessRequest)(nil),   // 7: online_shop.passport.v1.CheckUserAccessRequest
	(*CheckUserAccessResponse)(nil),  // 8: online_shop.passport.v1.CheckUserAccessResponse
	(*DeleteUserTokenRequest)(nil),   // 9: online_shop.passport.v1.DeleteUserTokenRequest
	(*DeleteUserTokenResponse)(nil),  // 10: online_shop.passport.v1.DeleteUserTokenResponse
}
var file_auth_proto_depIdxs = []int32{
	5,  // 0: online_shop.passport.v1.AuthService.GetUserToken:input_type -> online_shop.passport.v1.GetUserTokenRequest
	1,  // 1: online_shop.passport.v1.AuthService.SetUserAccess:input_type -> online_shop.passport.v1.SetUserAccessRequest
	9,  // 2: online_shop.passport.v1.AuthService.DeleteUserToken:input_type -> online_shop.passport.v1.DeleteUserTokenRequest
	7,  // 3: online_shop.passport.v1.AuthService.CheckUserAccess:input_type -> online_shop.passport.v1.CheckUserAccessRequest
	3,  // 4: online_shop.passport.v1.AuthService.DeleteUserAccess:input_type -> online_shop.passport.v1.DeleteUserAccessRequest
	6,  // 5: online_shop.passport.v1.AuthService.GetUserToken:output_type -> online_shop.passport.v1.GetUserTokenResponse
	2,  // 6: online_shop.passport.v1.AuthService.SetUserAccess:output_type -> online_shop.passport.v1.SetUserAccessResponse
	10, // 7: online_shop.passport.v1.AuthService.DeleteUserToken:output_type -> online_shop.passport.v1.DeleteUserTokenResponse
	8,  // 8: online_shop.passport.v1.AuthService.CheckUserAccess:output_type -> online_shop.passport.v1.CheckUserAccessResponse
	4,  // 9: online_shop.passport.v1.AuthService.DeleteUserAccess:output_type -> online_shop.passport.v1.DeleteUserAccessResponse
	5,  // [5:10] is the sub-list for method output_type
	0,  // [0:5] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Auth); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserAccessRequest); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserAccessResponse); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserAccessRequest); i {
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
		file_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserAccessResponse); i {
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
		file_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenRequest); i {
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
		file_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserTokenResponse); i {
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
		file_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAccessRequest); i {
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
		file_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckUserAccessResponse); i {
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
		file_auth_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTokenRequest); i {
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
		file_auth_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteUserTokenResponse); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
