// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: api/srv_marketplace.proto

package api

import (
	api "gitlab.ozon.dev/lvjonok/homework-3/pkg/common/api"
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

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=Desc,proto3" json:"Desc,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{0}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type CreateProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *CreateProductResponse) Reset() {
	*x = CreateProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductResponse) ProtoMessage() {}

func (x *CreateProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductResponse.ProtoReflect.Descriptor instead.
func (*CreateProductResponse) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProductResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{2}
}

func (x *GetProductRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=Desc,proto3" json:"Desc,omitempty"`
}

func (x *GetProductResponse) Reset() {
	*x = GetProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductResponse) ProtoMessage() {}

func (x *GetProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductResponse.ProtoReflect.Descriptor instead.
func (*GetProductResponse) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{3}
}

func (x *GetProductResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetProductResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductResponse) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type AddReviewRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID uint64 `protobuf:"varint,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Text      string `protobuf:"bytes,2,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *AddReviewRequest) Reset() {
	*x = AddReviewRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReviewRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReviewRequest) ProtoMessage() {}

func (x *AddReviewRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReviewRequest.ProtoReflect.Descriptor instead.
func (*AddReviewRequest) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{4}
}

func (x *AddReviewRequest) GetProductID() uint64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *AddReviewRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type AddReviewResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *AddReviewResponse) Reset() {
	*x = AddReviewResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReviewResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReviewResponse) ProtoMessage() {}

func (x *AddReviewResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReviewResponse.ProtoReflect.Descriptor instead.
func (*AddReviewResponse) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{5}
}

func (x *AddReviewResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetReviewsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductID uint64 `protobuf:"varint,1,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
}

func (x *GetReviewsRequest) Reset() {
	*x = GetReviewsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReviewsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsRequest) ProtoMessage() {}

func (x *GetReviewsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsRequest.ProtoReflect.Descriptor instead.
func (*GetReviewsRequest) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{6}
}

func (x *GetReviewsRequest) GetProductID() uint64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

type Review struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	ProductID uint64 `protobuf:"varint,2,opt,name=ProductID,proto3" json:"ProductID,omitempty"`
	Text      string `protobuf:"bytes,3,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *Review) Reset() {
	*x = Review{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Review) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Review) ProtoMessage() {}

func (x *Review) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Review.ProtoReflect.Descriptor instead.
func (*Review) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{7}
}

func (x *Review) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Review) GetProductID() uint64 {
	if x != nil {
		return x.ProductID
	}
	return 0
}

func (x *Review) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetReviewsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reviews []*Review `protobuf:"bytes,1,rep,name=Reviews,proto3" json:"Reviews,omitempty"`
}

func (x *GetReviewsResponse) Reset() {
	*x = GetReviewsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReviewsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReviewsResponse) ProtoMessage() {}

func (x *GetReviewsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReviewsResponse.ProtoReflect.Descriptor instead.
func (*GetReviewsResponse) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{8}
}

func (x *GetReviewsResponse) GetReviews() []*Review {
	if x != nil {
		return x.Reviews
	}
	return nil
}

type UpdateCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64             `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Products []*api.ProductUnit `protobuf:"bytes,2,rep,name=Products,proto3" json:"Products,omitempty"`
}

func (x *UpdateCartRequest) Reset() {
	*x = UpdateCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartRequest) ProtoMessage() {}

func (x *UpdateCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartRequest.ProtoReflect.Descriptor instead.
func (*UpdateCartRequest) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateCartRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UpdateCartRequest) GetProducts() []*api.ProductUnit {
	if x != nil {
		return x.Products
	}
	return nil
}

type UpdateCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *UpdateCartResponse) Reset() {
	*x = UpdateCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCartResponse) ProtoMessage() {}

func (x *UpdateCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCartResponse.ProtoReflect.Descriptor instead.
func (*UpdateCartResponse) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateCartResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID uint64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *GetCartRequest) Reset() {
	*x = GetCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartRequest) ProtoMessage() {}

func (x *GetCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartRequest.ProtoReflect.Descriptor instead.
func (*GetCartRequest) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{11}
}

func (x *GetCartRequest) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type GetCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       uint64             `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Products []*api.ProductUnit `protobuf:"bytes,2,rep,name=Products,proto3" json:"Products,omitempty"`
}

func (x *GetCartResponse) Reset() {
	*x = GetCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_srv_marketplace_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartResponse) ProtoMessage() {}

func (x *GetCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_srv_marketplace_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartResponse.ProtoReflect.Descriptor instead.
func (*GetCartResponse) Descriptor() ([]byte, []int) {
	return file_api_srv_marketplace_proto_rawDescGZIP(), []int{12}
}

func (x *GetCartResponse) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *GetCartResponse) GetProducts() []*api.ProductUnit {
	if x != nil {
		return x.Products
	}
	return nil
}

var File_api_srv_marketplace_proto protoreflect.FileDescriptor

var file_api_srv_marketplace_proto_rawDesc = []byte{
	0x0a, 0x19, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x72, 0x76, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x61, 0x70, 0x69,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x1a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x44, 0x65, 0x73, 0x63, 0x22, 0x27, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22,
	0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x02, 0x49, 0x44, 0x22, 0x4c, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65,
	0x73, 0x63, 0x22, 0x44, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x23, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x31, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44,
	0x22, 0x4a, 0x0a, 0x06, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x4f, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x07, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x07, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22, 0x54, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x2f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x22, 0x52, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x44, 0x12, 0x2f,
	0x0a, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x55, 0x6e, 0x69, 0x74, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32,
	0xf4, 0x04, 0x0a, 0x0b, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x12,
	0x6e, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x2e, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x65, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x2a, 0x2e,
	0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c,
	0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x62, 0x0a, 0x09, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76,
	0x69, 0x65, 0x77, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2a,
	0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72,
	0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x65, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x2a, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61,
	0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x65, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x12,
	0x2a, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x61, 0x70,
	0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x74, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x61,
	0x70, 0x69, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x40, 0x5a, 0x3e, 0x67, 0x69, 0x74, 0x6c, 0x61, 0x62,
	0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x6c, 0x76, 0x6a, 0x6f, 0x6e, 0x6f,
	0x6b, 0x2f, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x2d, 0x33, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_srv_marketplace_proto_rawDescOnce sync.Once
	file_api_srv_marketplace_proto_rawDescData = file_api_srv_marketplace_proto_rawDesc
)

func file_api_srv_marketplace_proto_rawDescGZIP() []byte {
	file_api_srv_marketplace_proto_rawDescOnce.Do(func() {
		file_api_srv_marketplace_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_srv_marketplace_proto_rawDescData)
	})
	return file_api_srv_marketplace_proto_rawDescData
}

var file_api_srv_marketplace_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_srv_marketplace_proto_goTypes = []interface{}{
	(*CreateProductRequest)(nil),  // 0: api_service_marketplace.CreateProductRequest
	(*CreateProductResponse)(nil), // 1: api_service_marketplace.CreateProductResponse
	(*GetProductRequest)(nil),     // 2: api_service_marketplace.GetProductRequest
	(*GetProductResponse)(nil),    // 3: api_service_marketplace.GetProductResponse
	(*AddReviewRequest)(nil),      // 4: api_service_marketplace.AddReviewRequest
	(*AddReviewResponse)(nil),     // 5: api_service_marketplace.AddReviewResponse
	(*GetReviewsRequest)(nil),     // 6: api_service_marketplace.GetReviewsRequest
	(*Review)(nil),                // 7: api_service_marketplace.Review
	(*GetReviewsResponse)(nil),    // 8: api_service_marketplace.GetReviewsResponse
	(*UpdateCartRequest)(nil),     // 9: api_service_marketplace.UpdateCartRequest
	(*UpdateCartResponse)(nil),    // 10: api_service_marketplace.UpdateCartResponse
	(*GetCartRequest)(nil),        // 11: api_service_marketplace.GetCartRequest
	(*GetCartResponse)(nil),       // 12: api_service_marketplace.GetCartResponse
	(*api.ProductUnit)(nil),       // 13: common.ProductUnit
}
var file_api_srv_marketplace_proto_depIdxs = []int32{
	7,  // 0: api_service_marketplace.GetReviewsResponse.Reviews:type_name -> api_service_marketplace.Review
	13, // 1: api_service_marketplace.UpdateCartRequest.Products:type_name -> common.ProductUnit
	13, // 2: api_service_marketplace.GetCartResponse.Products:type_name -> common.ProductUnit
	0,  // 3: api_service_marketplace.Marketplace.CreateProduct:input_type -> api_service_marketplace.CreateProductRequest
	2,  // 4: api_service_marketplace.Marketplace.GetProduct:input_type -> api_service_marketplace.GetProductRequest
	4,  // 5: api_service_marketplace.Marketplace.AddReview:input_type -> api_service_marketplace.AddReviewRequest
	6,  // 6: api_service_marketplace.Marketplace.GetReviews:input_type -> api_service_marketplace.GetReviewsRequest
	9,  // 7: api_service_marketplace.Marketplace.UpdateCart:input_type -> api_service_marketplace.UpdateCartRequest
	11, // 8: api_service_marketplace.Marketplace.GetCart:input_type -> api_service_marketplace.GetCartRequest
	1,  // 9: api_service_marketplace.Marketplace.CreateProduct:output_type -> api_service_marketplace.CreateProductResponse
	3,  // 10: api_service_marketplace.Marketplace.GetProduct:output_type -> api_service_marketplace.GetProductResponse
	5,  // 11: api_service_marketplace.Marketplace.AddReview:output_type -> api_service_marketplace.AddReviewResponse
	8,  // 12: api_service_marketplace.Marketplace.GetReviews:output_type -> api_service_marketplace.GetReviewsResponse
	10, // 13: api_service_marketplace.Marketplace.UpdateCart:output_type -> api_service_marketplace.UpdateCartResponse
	12, // 14: api_service_marketplace.Marketplace.GetCart:output_type -> api_service_marketplace.GetCartResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_api_srv_marketplace_proto_init() }
func file_api_srv_marketplace_proto_init() {
	if File_api_srv_marketplace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_srv_marketplace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_api_srv_marketplace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductResponse); i {
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
		file_api_srv_marketplace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
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
		file_api_srv_marketplace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductResponse); i {
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
		file_api_srv_marketplace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReviewRequest); i {
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
		file_api_srv_marketplace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReviewResponse); i {
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
		file_api_srv_marketplace_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReviewsRequest); i {
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
		file_api_srv_marketplace_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Review); i {
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
		file_api_srv_marketplace_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReviewsResponse); i {
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
		file_api_srv_marketplace_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartRequest); i {
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
		file_api_srv_marketplace_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCartResponse); i {
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
		file_api_srv_marketplace_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartRequest); i {
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
		file_api_srv_marketplace_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartResponse); i {
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
			RawDescriptor: file_api_srv_marketplace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_srv_marketplace_proto_goTypes,
		DependencyIndexes: file_api_srv_marketplace_proto_depIdxs,
		MessageInfos:      file_api_srv_marketplace_proto_msgTypes,
	}.Build()
	File_api_srv_marketplace_proto = out.File
	file_api_srv_marketplace_proto_rawDesc = nil
	file_api_srv_marketplace_proto_goTypes = nil
	file_api_srv_marketplace_proto_depIdxs = nil
}
