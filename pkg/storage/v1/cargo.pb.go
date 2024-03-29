// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: cargo.proto

package storage

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

type Carrier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string  `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Capacity     uint64  `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Price        float32 `protobuf:"fixed32,4,opt,name=price,proto3" json:"price,omitempty"`
	Availability bool    `protobuf:"varint,5,opt,name=availability,proto3" json:"availability,omitempty"`
}

func (x *Carrier) Reset() {
	*x = Carrier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Carrier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Carrier) ProtoMessage() {}

func (x *Carrier) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Carrier.ProtoReflect.Descriptor instead.
func (*Carrier) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{0}
}

func (x *Carrier) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Carrier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Carrier) GetCapacity() uint64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *Carrier) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Carrier) GetAvailability() bool {
	if x != nil {
		return x.Availability
	}
	return false
}

type CreateCarrierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity     uint64  `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	Price        float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Availability bool    `protobuf:"varint,4,opt,name=availability,proto3" json:"availability,omitempty"`
}

func (x *CreateCarrierRequest) Reset() {
	*x = CreateCarrierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarrierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarrierRequest) ProtoMessage() {}

func (x *CreateCarrierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarrierRequest.ProtoReflect.Descriptor instead.
func (*CreateCarrierRequest) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCarrierRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateCarrierRequest) GetCapacity() uint64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

func (x *CreateCarrierRequest) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CreateCarrierRequest) GetAvailability() bool {
	if x != nil {
		return x.Availability
	}
	return false
}

type CreateCarrierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Carrier *Carrier `protobuf:"bytes,1,opt,name=carrier,proto3" json:"carrier,omitempty"`
}

func (x *CreateCarrierResponse) Reset() {
	*x = CreateCarrierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarrierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarrierResponse) ProtoMessage() {}

func (x *CreateCarrierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarrierResponse.ProtoReflect.Descriptor instead.
func (*CreateCarrierResponse) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCarrierResponse) GetCarrier() *Carrier {
	if x != nil {
		return x.Carrier
	}
	return nil
}

type GetCarrierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetCarrierRequest) Reset() {
	*x = GetCarrierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarrierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarrierRequest) ProtoMessage() {}

func (x *GetCarrierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarrierRequest.ProtoReflect.Descriptor instead.
func (*GetCarrierRequest) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{3}
}

func (x *GetCarrierRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetCarrierRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetCarrierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Carrier *Carrier `protobuf:"bytes,1,opt,name=carrier,proto3" json:"carrier,omitempty"`
}

func (x *GetCarrierResponse) Reset() {
	*x = GetCarrierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarrierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarrierResponse) ProtoMessage() {}

func (x *GetCarrierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarrierResponse.ProtoReflect.Descriptor instead.
func (*GetCarrierResponse) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{4}
}

func (x *GetCarrierResponse) GetCarrier() *Carrier {
	if x != nil {
		return x.Carrier
	}
	return nil
}

type DeleteCarrierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCarrierRequest) Reset() {
	*x = DeleteCarrierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarrierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarrierRequest) ProtoMessage() {}

func (x *DeleteCarrierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarrierRequest.ProtoReflect.Descriptor instead.
func (*DeleteCarrierRequest) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCarrierRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCarrierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCarrierResponse) Reset() {
	*x = DeleteCarrierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarrierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarrierResponse) ProtoMessage() {}

func (x *DeleteCarrierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarrierResponse.ProtoReflect.Descriptor instead.
func (*DeleteCarrierResponse) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{6}
}

type ListCarrierRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *v1.Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListCarrierRequest) Reset() {
	*x = ListCarrierRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarrierRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarrierRequest) ProtoMessage() {}

func (x *ListCarrierRequest) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarrierRequest.ProtoReflect.Descriptor instead.
func (*ListCarrierRequest) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{7}
}

func (x *ListCarrierRequest) GetPagination() *v1.Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListCarrierResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Carrier []*Carrier `protobuf:"bytes,1,rep,name=carrier,proto3" json:"carrier,omitempty"`
}

func (x *ListCarrierResponse) Reset() {
	*x = ListCarrierResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cargo_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarrierResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarrierResponse) ProtoMessage() {}

func (x *ListCarrierResponse) ProtoReflect() protoreflect.Message {
	mi := &file_cargo_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarrierResponse.ProtoReflect.Descriptor instead.
func (*ListCarrierResponse) Descriptor() ([]byte, []int) {
	return file_cargo_proto_rawDescGZIP(), []int{8}
}

func (x *ListCarrierResponse) GetCarrier() []*Carrier {
	if x != nil {
		return x.Carrier
	}
	return nil
}

var File_cargo_proto protoreflect.FileDescriptor

var file_cargo_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6c, 0x76, 0x6c, 0x42, 0x41, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73,
	0x68, 0x6f, 0x70, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x07,
	0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74,
	0x79, 0x22, 0x80, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x08, 0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x22, 0x52, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52,
	0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x07, 0x63, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72,
	0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x51, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x72, 0x69,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x0a, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69,
	0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x50, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x39, 0x0a,
	0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52,
	0x07, 0x63, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x32, 0xc0, 0x03, 0x0a, 0x0c, 0x43, 0x61, 0x72,
	0x67, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6e, 0x0a, 0x0d, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x2c, 0x2e, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x65, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x29, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65,
	0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x6e, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x12, 0x2c, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x69, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x73,
	0x12, 0x2a, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61,
	0x72, 0x72, 0x69, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x76, 0x6c, 0x42, 0x41, 0x2f,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x3b, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_cargo_proto_rawDescOnce sync.Once
	file_cargo_proto_rawDescData = file_cargo_proto_rawDesc
)

func file_cargo_proto_rawDescGZIP() []byte {
	file_cargo_proto_rawDescOnce.Do(func() {
		file_cargo_proto_rawDescData = protoimpl.X.CompressGZIP(file_cargo_proto_rawDescData)
	})
	return file_cargo_proto_rawDescData
}

var file_cargo_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cargo_proto_goTypes = []interface{}{
	(*Carrier)(nil),               // 0: online_shop.storage.v1.Carrier
	(*CreateCarrierRequest)(nil),  // 1: online_shop.storage.v1.CreateCarrierRequest
	(*CreateCarrierResponse)(nil), // 2: online_shop.storage.v1.CreateCarrierResponse
	(*GetCarrierRequest)(nil),     // 3: online_shop.storage.v1.GetCarrierRequest
	(*GetCarrierResponse)(nil),    // 4: online_shop.storage.v1.GetCarrierResponse
	(*DeleteCarrierRequest)(nil),  // 5: online_shop.storage.v1.DeleteCarrierRequest
	(*DeleteCarrierResponse)(nil), // 6: online_shop.storage.v1.DeleteCarrierResponse
	(*ListCarrierRequest)(nil),    // 7: online_shop.storage.v1.ListCarrierRequest
	(*ListCarrierResponse)(nil),   // 8: online_shop.storage.v1.ListCarrierResponse
	(*v1.Pagination)(nil),         // 9: online_shop.api.Pagination
}
var file_cargo_proto_depIdxs = []int32{
	0, // 0: online_shop.storage.v1.CreateCarrierResponse.carrier:type_name -> online_shop.storage.v1.Carrier
	0, // 1: online_shop.storage.v1.GetCarrierResponse.carrier:type_name -> online_shop.storage.v1.Carrier
	9, // 2: online_shop.storage.v1.ListCarrierRequest.pagination:type_name -> online_shop.api.Pagination
	0, // 3: online_shop.storage.v1.ListCarrierResponse.carrier:type_name -> online_shop.storage.v1.Carrier
	1, // 4: online_shop.storage.v1.CargoService.CreateCarrier:input_type -> online_shop.storage.v1.CreateCarrierRequest
	3, // 5: online_shop.storage.v1.CargoService.GetCarrier:input_type -> online_shop.storage.v1.GetCarrierRequest
	5, // 6: online_shop.storage.v1.CargoService.DeleteCarrier:input_type -> online_shop.storage.v1.DeleteCarrierRequest
	7, // 7: online_shop.storage.v1.CargoService.ListCarriers:input_type -> online_shop.storage.v1.ListCarrierRequest
	2, // 8: online_shop.storage.v1.CargoService.CreateCarrier:output_type -> online_shop.storage.v1.CreateCarrierResponse
	4, // 9: online_shop.storage.v1.CargoService.GetCarrier:output_type -> online_shop.storage.v1.GetCarrierResponse
	6, // 10: online_shop.storage.v1.CargoService.DeleteCarrier:output_type -> online_shop.storage.v1.DeleteCarrierResponse
	8, // 11: online_shop.storage.v1.CargoService.ListCarriers:output_type -> online_shop.storage.v1.ListCarrierResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_cargo_proto_init() }
func file_cargo_proto_init() {
	if File_cargo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cargo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Carrier); i {
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
		file_cargo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarrierRequest); i {
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
		file_cargo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarrierResponse); i {
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
		file_cargo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarrierRequest); i {
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
		file_cargo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarrierResponse); i {
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
		file_cargo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarrierRequest); i {
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
		file_cargo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarrierResponse); i {
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
		file_cargo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarrierRequest); i {
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
		file_cargo_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarrierResponse); i {
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
			RawDescriptor: file_cargo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cargo_proto_goTypes,
		DependencyIndexes: file_cargo_proto_depIdxs,
		MessageInfos:      file_cargo_proto_msgTypes,
	}.Build()
	File_cargo_proto = out.File
	file_cargo_proto_rawDesc = nil
	file_cargo_proto_goTypes = nil
	file_cargo_proto_depIdxs = nil
}
