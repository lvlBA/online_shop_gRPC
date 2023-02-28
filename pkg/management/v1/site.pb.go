// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: site.proto

package management

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

type Site struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Site) Reset() {
	*x = Site{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Site) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Site) ProtoMessage() {}

func (x *Site) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Site.ProtoReflect.Descriptor instead.
func (*Site) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{0}
}

func (x *Site) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Site) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateSideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CreateSideRequest) Reset() {
	*x = CreateSideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSideRequest) ProtoMessage() {}

func (x *CreateSideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSideRequest.ProtoReflect.Descriptor instead.
func (*CreateSideRequest) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSideRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateSideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site *Site `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
}

func (x *CreateSideResponse) Reset() {
	*x = CreateSideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSideResponse) ProtoMessage() {}

func (x *CreateSideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSideResponse.ProtoReflect.Descriptor instead.
func (*CreateSideResponse) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSideResponse) GetSite() *Site {
	if x != nil {
		return x.Site
	}
	return nil
}

type GetSideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetSideRequest) Reset() {
	*x = GetSideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSideRequest) ProtoMessage() {}

func (x *GetSideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSideRequest.ProtoReflect.Descriptor instead.
func (*GetSideRequest) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{3}
}

func (x *GetSideRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site *Site `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
}

func (x *GetSideResponse) Reset() {
	*x = GetSideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSideResponse) ProtoMessage() {}

func (x *GetSideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSideResponse.ProtoReflect.Descriptor instead.
func (*GetSideResponse) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{4}
}

func (x *GetSideResponse) GetSite() *Site {
	if x != nil {
		return x.Site
	}
	return nil
}

type DeleteSiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSiteRequest) Reset() {
	*x = DeleteSiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSiteRequest) ProtoMessage() {}

func (x *DeleteSiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSiteRequest.ProtoReflect.Descriptor instead.
func (*DeleteSiteRequest) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteSiteRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteSiteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSiteResponse) Reset() {
	*x = DeleteSiteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSiteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSiteResponse) ProtoMessage() {}

func (x *DeleteSiteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSiteResponse.ProtoReflect.Descriptor instead.
func (*DeleteSiteResponse) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{6}
}

type ListSitesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pagination *Pagination `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *ListSitesRequest) Reset() {
	*x = ListSitesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSitesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSitesRequest) ProtoMessage() {}

func (x *ListSitesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSitesRequest.ProtoReflect.Descriptor instead.
func (*ListSitesRequest) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{7}
}

func (x *ListSitesRequest) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type ListSitesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sites []*Site `protobuf:"bytes,1,rep,name=sites,proto3" json:"sites,omitempty"`
}

func (x *ListSitesResponse) Reset() {
	*x = ListSitesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_site_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSitesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSitesResponse) ProtoMessage() {}

func (x *ListSitesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_site_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSitesResponse.ProtoReflect.Descriptor instead.
func (*ListSitesResponse) Descriptor() ([]byte, []int) {
	return file_site_proto_rawDescGZIP(), []int{8}
}

func (x *ListSitesResponse) GetSites() []*Site {
	if x != nil {
		return x.Sites
	}
	return nil
}

var File_site_proto protoreflect.FileDescriptor

var file_site_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x69, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x10, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x04, 0x53, 0x69, 0x74,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x27, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x69, 0x74, 0x65, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x46, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x33,
	0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6f,
	0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x52, 0x04, 0x73,
	0x69, 0x74, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x45, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f,
	0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4a, 0x0a, 0x11, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35,
	0x0a, 0x05, 0x73, 0x69, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x52, 0x05,
	0x73, 0x69, 0x74, 0x65, 0x73, 0x32, 0xb5, 0x03, 0x0a, 0x0b, 0x53, 0x69, 0x74, 0x65, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6b, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x69, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f,
	0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x62, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x69, 0x64, 0x65, 0x12, 0x29, 0x2e,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x64,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e,
	0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x69, 0x74, 0x65, 0x12, 0x2c, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68,
	0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70,
	0x2e, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x68, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x74, 0x65, 0x73,
	0x12, 0x2b, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x34, 0x5a,
	0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x76, 0x6c, 0x42,
	0x41, 0x2f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x5f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x3b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_site_proto_rawDescOnce sync.Once
	file_site_proto_rawDescData = file_site_proto_rawDesc
)

func file_site_proto_rawDescGZIP() []byte {
	file_site_proto_rawDescOnce.Do(func() {
		file_site_proto_rawDescData = protoimpl.X.CompressGZIP(file_site_proto_rawDescData)
	})
	return file_site_proto_rawDescData
}

var file_site_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_site_proto_goTypes = []interface{}{
	(*Site)(nil),               // 0: online_shop.management.v1.Site
	(*CreateSideRequest)(nil),  // 1: online_shop.management.v1.CreateSideRequest
	(*CreateSideResponse)(nil), // 2: online_shop.management.v1.CreateSideResponse
	(*GetSideRequest)(nil),     // 3: online_shop.management.v1.GetSideRequest
	(*GetSideResponse)(nil),    // 4: online_shop.management.v1.GetSideResponse
	(*DeleteSiteRequest)(nil),  // 5: online_shop.management.v1.DeleteSiteRequest
	(*DeleteSiteResponse)(nil), // 6: online_shop.management.v1.DeleteSiteResponse
	(*ListSitesRequest)(nil),   // 7: online_shop.management.v1.ListSitesRequest
	(*ListSitesResponse)(nil),  // 8: online_shop.management.v1.ListSitesResponse
	(*Pagination)(nil),         // 9: online_shop.management.v1.Pagination
}
var file_site_proto_depIdxs = []int32{
	0, // 0: online_shop.management.v1.CreateSideResponse.site:type_name -> online_shop.management.v1.Site
	0, // 1: online_shop.management.v1.GetSideResponse.site:type_name -> online_shop.management.v1.Site
	9, // 2: online_shop.management.v1.ListSitesRequest.pagination:type_name -> online_shop.management.v1.Pagination
	0, // 3: online_shop.management.v1.ListSitesResponse.sites:type_name -> online_shop.management.v1.Site
	1, // 4: online_shop.management.v1.SiteService.CreateSite:input_type -> online_shop.management.v1.CreateSideRequest
	3, // 5: online_shop.management.v1.SiteService.GetSide:input_type -> online_shop.management.v1.GetSideRequest
	5, // 6: online_shop.management.v1.SiteService.DeleteSite:input_type -> online_shop.management.v1.DeleteSiteRequest
	7, // 7: online_shop.management.v1.SiteService.ListSites:input_type -> online_shop.management.v1.ListSitesRequest
	2, // 8: online_shop.management.v1.SiteService.CreateSite:output_type -> online_shop.management.v1.CreateSideResponse
	4, // 9: online_shop.management.v1.SiteService.GetSide:output_type -> online_shop.management.v1.GetSideResponse
	6, // 10: online_shop.management.v1.SiteService.DeleteSite:output_type -> online_shop.management.v1.DeleteSiteResponse
	8, // 11: online_shop.management.v1.SiteService.ListSites:output_type -> online_shop.management.v1.ListSitesResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_site_proto_init() }
func file_site_proto_init() {
	if File_site_proto != nil {
		return
	}
	file_pagination_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_site_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Site); i {
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
		file_site_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSideRequest); i {
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
		file_site_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSideResponse); i {
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
		file_site_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSideRequest); i {
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
		file_site_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSideResponse); i {
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
		file_site_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSiteRequest); i {
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
		file_site_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSiteResponse); i {
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
		file_site_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSitesRequest); i {
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
		file_site_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSitesResponse); i {
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
			RawDescriptor: file_site_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_site_proto_goTypes,
		DependencyIndexes: file_site_proto_depIdxs,
		MessageInfos:      file_site_proto_msgTypes,
	}.Build()
	File_site_proto = out.File
	file_site_proto_rawDesc = nil
	file_site_proto_goTypes = nil
	file_site_proto_depIdxs = nil
}
