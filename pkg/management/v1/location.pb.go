// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: location.proto

package management

import (
	v1 "github.com/lvlBA/online_shop/pkg/api/v1"
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

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{0}
}

func (x *Location) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Location) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateLocationRequest) Reset() {
	*x = CreateLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLocationRequest) ProtoMessage() {}

func (x *CreateLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLocationRequest.ProtoReflect.Descriptor instead.
func (*CreateLocationRequest) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{1}
}

func (x *CreateLocationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *CreateLocationResponse) Reset() {
	*x = CreateLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLocationResponse) ProtoMessage() {}

func (x *CreateLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLocationResponse.ProtoReflect.Descriptor instead.
func (*CreateLocationResponse) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{2}
}

func (x *CreateLocationResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type GetLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetLocationRequest) Reset() {
	*x = GetLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationRequest) ProtoMessage() {}

func (x *GetLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationRequest.ProtoReflect.Descriptor instead.
func (*GetLocationRequest) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{3}
}

func (x *GetLocationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location *Location `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
}

func (x *GetLocationResponse) Reset() {
	*x = GetLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocationResponse) ProtoMessage() {}

func (x *GetLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocationResponse.ProtoReflect.Descriptor instead.
func (*GetLocationResponse) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{4}
}

func (x *GetLocationResponse) GetLocation() *Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type DeleteLocationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteLocationRequest) Reset() {
	*x = DeleteLocationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLocationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLocationRequest) ProtoMessage() {}

func (x *DeleteLocationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLocationRequest.ProtoReflect.Descriptor instead.
func (*DeleteLocationRequest) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteLocationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteLocationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteLocationResponse) Reset() {
	*x = DeleteLocationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteLocationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteLocationResponse) ProtoMessage() {}

func (x *DeleteLocationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteLocationResponse.ProtoReflect.Descriptor instead.
func (*DeleteLocationResponse) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{6}
}

type ListLocationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *v1.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListLocationsRequest) Reset() {
	*x = ListLocationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocationsRequest) ProtoMessage() {}

func (x *ListLocationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocationsRequest.ProtoReflect.Descriptor instead.
func (*ListLocationsRequest) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{7}
}

func (x *ListLocationsRequest) GetPagination() *v1.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListLocationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location []*Location `protobuf:"bytes,1,rep,name=location,proto3" json:"location,omitempty"`
}

func (x *ListLocationsResponse) Reset() {
	*x = ListLocationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_location_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListLocationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListLocationsResponse) ProtoMessage() {}

func (x *ListLocationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_location_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListLocationsResponse.ProtoReflect.Descriptor instead.
func (*ListLocationsResponse) Descriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{8}
}

func (x *ListLocationsResponse) GetLocation() []*Location {
	if x != nil {
		return x.Location
	}
	return nil
}

var File_location_proto protoreflect.FileDescriptor

var file_location_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x19, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x34, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x76, 0x6c, 0x42, 0x41, 0x2f, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31,
	0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x2e, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x2b, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x59,
	0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x56, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x27, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x53, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x58, 0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xe8, 0x03, 0x0a, 0x0f, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x30, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73,
	0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x77, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x73, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2f, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74,
	0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x30, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6c, 0x76, 0x6c, 0x42, 0x41, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x2f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3b,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_location_proto_rawDescOnce sync.Once
	file_location_proto_rawDescData = file_location_proto_rawDesc
)

func file_location_proto_rawDescGZIP() []byte {
	file_location_proto_rawDescOnce.Do(func() {
		file_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_location_proto_rawDescData)
	})
	return file_location_proto_rawDescData
}

var file_location_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_location_proto_goTypes = []interface{}{
	(*Location)(nil),               // 0: online_shop.management.v1.Location
	(*CreateLocationRequest)(nil),  // 1: online_shop.management.v1.CreateLocationRequest
	(*CreateLocationResponse)(nil), // 2: online_shop.management.v1.CreateLocationResponse
	(*GetLocationRequest)(nil),     // 3: online_shop.management.v1.GetLocationRequest
	(*GetLocationResponse)(nil),    // 4: online_shop.management.v1.GetLocationResponse
	(*DeleteLocationRequest)(nil),  // 5: online_shop.management.v1.DeleteLocationRequest
	(*DeleteLocationResponse)(nil), // 6: online_shop.management.v1.DeleteLocationResponse
	(*ListLocationsRequest)(nil),   // 7: online_shop.management.v1.ListLocationsRequest
	(*ListLocationsResponse)(nil),  // 8: online_shop.management.v1.ListLocationsResponse
	(*v1.Pagination)(nil),          // 9: online_shop.api.Pagination
}
var file_location_proto_depIdxs = []int32{
	0, // 0: online_shop.management.v1.CreateLocationResponse.location:type_name -> online_shop.management.v1.Location
	0, // 1: online_shop.management.v1.GetLocationResponse.location:type_name -> online_shop.management.v1.Location
	9, // 2: online_shop.management.v1.ListLocationsRequest.pagination:type_name -> online_shop.api.Pagination
	0, // 3: online_shop.management.v1.ListLocationsResponse.location:type_name -> online_shop.management.v1.Location
	1, // 4: online_shop.management.v1.LocationService.CreateLocation:input_type -> online_shop.management.v1.CreateLocationRequest
	3, // 5: online_shop.management.v1.LocationService.GetLocation:input_type -> online_shop.management.v1.GetLocationRequest
	5, // 6: online_shop.management.v1.LocationService.DeleteLocation:input_type -> online_shop.management.v1.DeleteLocationRequest
	7, // 7: online_shop.management.v1.LocationService.ListLocation:input_type -> online_shop.management.v1.ListLocationsRequest
	2, // 8: online_shop.management.v1.LocationService.CreateLocation:output_type -> online_shop.management.v1.CreateLocationResponse
	4, // 9: online_shop.management.v1.LocationService.GetLocation:output_type -> online_shop.management.v1.GetLocationResponse
	6, // 10: online_shop.management.v1.LocationService.DeleteLocation:output_type -> online_shop.management.v1.DeleteLocationResponse
	8, // 11: online_shop.management.v1.LocationService.ListLocation:output_type -> online_shop.management.v1.ListLocationsResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_location_proto_init() }
func file_location_proto_init() {
	if File_location_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_location_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_location_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLocationRequest); i {
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
		file_location_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateLocationResponse); i {
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
		file_location_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationRequest); i {
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
		file_location_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLocationResponse); i {
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
		file_location_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLocationRequest); i {
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
		file_location_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteLocationResponse); i {
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
		file_location_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocationsRequest); i {
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
		file_location_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListLocationsResponse); i {
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
			RawDescriptor: file_location_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_location_proto_goTypes,
		DependencyIndexes: file_location_proto_depIdxs,
		MessageInfos:      file_location_proto_msgTypes,
	}.Build()
	File_location_proto = out.File
	file_location_proto_rawDesc = nil
	file_location_proto_goTypes = nil
	file_location_proto_depIdxs = nil
}
