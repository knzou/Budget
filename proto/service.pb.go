// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

type GetCategoriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Categories []*GetCategoriesResponse_Category `protobuf:"bytes,1,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *GetCategoriesResponse) Reset() {
	*x = GetCategoriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoriesResponse) ProtoMessage() {}

func (x *GetCategoriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoriesResponse.ProtoReflect.Descriptor instead.
func (*GetCategoriesResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetCategoriesResponse) GetCategories() []*GetCategoriesResponse_Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

type GetTransactionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*GetTransactionsResponse_Transaction `protobuf:"bytes,1,rep,name=transactions,proto3" json:"transactions,omitempty"`
}

func (x *GetTransactionsResponse) Reset() {
	*x = GetTransactionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsResponse) ProtoMessage() {}

func (x *GetTransactionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsResponse.ProtoReflect.Descriptor instead.
func (*GetTransactionsResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTransactionsResponse) GetTransactions() []*GetTransactionsResponse_Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type GetPeopleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPeopleRequest) Reset() {
	*x = GetPeopleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeopleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeopleRequest) ProtoMessage() {}

func (x *GetPeopleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeopleRequest.ProtoReflect.Descriptor instead.
func (*GetPeopleRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPeopleRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetPeopleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	People []*GetPeopleResponse_Person `protobuf:"bytes,1,rep,name=people,proto3" json:"people,omitempty"`
}

func (x *GetPeopleResponse) Reset() {
	*x = GetPeopleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeopleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeopleResponse) ProtoMessage() {}

func (x *GetPeopleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeopleResponse.ProtoReflect.Descriptor instead.
func (*GetPeopleResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetPeopleResponse) GetPeople() []*GetPeopleResponse_Person {
	if x != nil {
		return x.People
	}
	return nil
}

type GetTotalTransactionAmountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsParallel bool `protobuf:"varint,1,opt,name=isParallel,proto3" json:"isParallel,omitempty"`
}

func (x *GetTotalTransactionAmountRequest) Reset() {
	*x = GetTotalTransactionAmountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalTransactionAmountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalTransactionAmountRequest) ProtoMessage() {}

func (x *GetTotalTransactionAmountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalTransactionAmountRequest.ProtoReflect.Descriptor instead.
func (*GetTotalTransactionAmountRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{5}
}

func (x *GetTotalTransactionAmountRequest) GetIsParallel() bool {
	if x != nil {
		return x.IsParallel
	}
	return false
}

type GetTotalTransactionAmountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalAmount int64 `protobuf:"varint,1,opt,name=totalAmount,proto3" json:"totalAmount,omitempty"`
	TotalTime   int64 `protobuf:"varint,2,opt,name=totalTime,proto3" json:"totalTime,omitempty"`
}

func (x *GetTotalTransactionAmountResponse) Reset() {
	*x = GetTotalTransactionAmountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTotalTransactionAmountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTotalTransactionAmountResponse) ProtoMessage() {}

func (x *GetTotalTransactionAmountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTotalTransactionAmountResponse.ProtoReflect.Descriptor instead.
func (*GetTotalTransactionAmountResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetTotalTransactionAmountResponse) GetTotalAmount() int64 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

func (x *GetTotalTransactionAmountResponse) GetTotalTime() int64 {
	if x != nil {
		return x.TotalTime
	}
	return 0
}

type GetCategoriesResponse_Category struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CatId  int64  `protobuf:"varint,1,opt,name=catId,proto3" json:"catId,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	TypeId int64  `protobuf:"varint,3,opt,name=typeId,proto3" json:"typeId,omitempty"`
}

func (x *GetCategoriesResponse_Category) Reset() {
	*x = GetCategoriesResponse_Category{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCategoriesResponse_Category) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCategoriesResponse_Category) ProtoMessage() {}

func (x *GetCategoriesResponse_Category) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCategoriesResponse_Category.ProtoReflect.Descriptor instead.
func (*GetCategoriesResponse_Category) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1, 0}
}

func (x *GetCategoriesResponse_Category) GetCatId() int64 {
	if x != nil {
		return x.CatId
	}
	return 0
}

func (x *GetCategoriesResponse_Category) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCategoriesResponse_Category) GetTypeId() int64 {
	if x != nil {
		return x.TypeId
	}
	return 0
}

type GetTransactionsResponse_Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TranId    int64 `protobuf:"varint,1,opt,name=tranId,proto3" json:"tranId,omitempty"`
	CatId     int64 `protobuf:"varint,2,opt,name=catId,proto3" json:"catId,omitempty"`
	TransDate *Date `protobuf:"bytes,3,opt,name=transDate,proto3" json:"transDate,omitempty"`
	Amount    int64 `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *GetTransactionsResponse_Transaction) Reset() {
	*x = GetTransactionsResponse_Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTransactionsResponse_Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransactionsResponse_Transaction) ProtoMessage() {}

func (x *GetTransactionsResponse_Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransactionsResponse_Transaction.ProtoReflect.Descriptor instead.
func (*GetTransactionsResponse_Transaction) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2, 0}
}

func (x *GetTransactionsResponse_Transaction) GetTranId() int64 {
	if x != nil {
		return x.TranId
	}
	return 0
}

func (x *GetTransactionsResponse_Transaction) GetCatId() int64 {
	if x != nil {
		return x.CatId
	}
	return 0
}

func (x *GetTransactionsResponse_Transaction) GetTransDate() *Date {
	if x != nil {
		return x.TransDate
	}
	return nil
}

func (x *GetTransactionsResponse_Transaction) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type GetPeopleResponse_Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid  int64  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetPeopleResponse_Person) Reset() {
	*x = GetPeopleResponse_Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPeopleResponse_Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPeopleResponse_Person) ProtoMessage() {}

func (x *GetPeopleResponse_Person) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPeopleResponse_Person.ProtoReflect.Descriptor instead.
func (*GetPeopleResponse_Person) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetPeopleResponse_Person) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *GetPeopleResponse_Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x0a, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0xb0, 0x01, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x0a, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x4c, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x74, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x79,
	0x70, 0x65, 0x49, 0x64, 0x22, 0xec, 0x01, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x52, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x7d, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x72, 0x61, 0x6e, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x74, 0x72, 0x61, 0x6e, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x61, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x61, 0x74, 0x49,
	0x64, 0x12, 0x28, 0x0a, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x44, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x61, 0x74, 0x65,
	0x52, 0x09, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x44, 0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x22, 0x26, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x80, 0x01, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x3b, 0x0a, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x23, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x1a, 0x2e,
	0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42,
	0x0a, 0x20, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c, 0x65, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x50, 0x61, 0x72, 0x61, 0x6c, 0x6c,
	0x65, 0x6c, 0x22, 0x63, 0x0a, 0x21, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x32, 0xde, 0x02, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x12, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x42, 0x75,
	0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a,
	0x0f, 0x67, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x12, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x67, 0x65, 0x74, 0x50,
	0x65, 0x6f, 0x70, 0x6c, 0x65, 0x12, 0x1b, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x65, 0x6f, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x76, 0x0a, 0x19, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2b, 0x2e,
	0x42, 0x75, 0x64, 0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x74,
	0x61, 0x6c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x67, 0x65, 0x74, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1c, 0x5a, 0x1a, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x6e, 0x7a, 0x6f, 0x75, 0x2f, 0x42, 0x75, 0x64,
	0x67, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_service_proto_goTypes = []interface{}{
	(*Request)(nil),                             // 0: Budget.v1.Request
	(*GetCategoriesResponse)(nil),               // 1: Budget.v1.GetCategoriesResponse
	(*GetTransactionsResponse)(nil),             // 2: Budget.v1.GetTransactionsResponse
	(*GetPeopleRequest)(nil),                    // 3: Budget.v1.GetPeopleRequest
	(*GetPeopleResponse)(nil),                   // 4: Budget.v1.GetPeopleResponse
	(*GetTotalTransactionAmountRequest)(nil),    // 5: Budget.v1.getTotalTransactionAmountRequest
	(*GetTotalTransactionAmountResponse)(nil),   // 6: Budget.v1.getTotalTransactionAmountResponse
	(*GetCategoriesResponse_Category)(nil),      // 7: Budget.v1.GetCategoriesResponse.Category
	(*GetTransactionsResponse_Transaction)(nil), // 8: Budget.v1.GetTransactionsResponse.Transaction
	(*GetPeopleResponse_Person)(nil),            // 9: Budget.v1.GetPeopleResponse.Person
	(*Date)(nil),                                // 10: date.Date
}
var file_service_proto_depIdxs = []int32{
	7,  // 0: Budget.v1.GetCategoriesResponse.categories:type_name -> Budget.v1.GetCategoriesResponse.Category
	8,  // 1: Budget.v1.GetTransactionsResponse.transactions:type_name -> Budget.v1.GetTransactionsResponse.Transaction
	9,  // 2: Budget.v1.GetPeopleResponse.people:type_name -> Budget.v1.GetPeopleResponse.Person
	10, // 3: Budget.v1.GetTransactionsResponse.Transaction.transDate:type_name -> date.Date
	0,  // 4: Budget.v1.AddService.getCategories:input_type -> Budget.v1.Request
	0,  // 5: Budget.v1.AddService.getTransactions:input_type -> Budget.v1.Request
	3,  // 6: Budget.v1.AddService.getPeople:input_type -> Budget.v1.GetPeopleRequest
	5,  // 7: Budget.v1.AddService.getTotalTransactionAmount:input_type -> Budget.v1.getTotalTransactionAmountRequest
	1,  // 8: Budget.v1.AddService.getCategories:output_type -> Budget.v1.GetCategoriesResponse
	2,  // 9: Budget.v1.AddService.getTransactions:output_type -> Budget.v1.GetTransactionsResponse
	4,  // 10: Budget.v1.AddService.getPeople:output_type -> Budget.v1.GetPeopleResponse
	6,  // 11: Budget.v1.AddService.getTotalTransactionAmount:output_type -> Budget.v1.getTotalTransactionAmountResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	file_date_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoriesResponse); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsResponse); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeopleRequest); i {
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
		file_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeopleResponse); i {
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
		file_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalTransactionAmountRequest); i {
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
		file_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTotalTransactionAmountResponse); i {
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
		file_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCategoriesResponse_Category); i {
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
		file_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTransactionsResponse_Transaction); i {
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
		file_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPeopleResponse_Person); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AddServiceClient is the client API for AddService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddServiceClient interface {
	GetCategories(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetCategoriesResponse, error)
	GetTransactions(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetTransactionsResponse, error)
	GetPeople(ctx context.Context, in *GetPeopleRequest, opts ...grpc.CallOption) (*GetPeopleResponse, error)
	GetTotalTransactionAmount(ctx context.Context, in *GetTotalTransactionAmountRequest, opts ...grpc.CallOption) (*GetTotalTransactionAmountResponse, error)
}

type addServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAddServiceClient(cc grpc.ClientConnInterface) AddServiceClient {
	return &addServiceClient{cc}
}

func (c *addServiceClient) GetCategories(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetCategoriesResponse, error) {
	out := new(GetCategoriesResponse)
	err := c.cc.Invoke(ctx, "/Budget.v1.AddService/getCategories", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) GetTransactions(ctx context.Context, in *Request, opts ...grpc.CallOption) (*GetTransactionsResponse, error) {
	out := new(GetTransactionsResponse)
	err := c.cc.Invoke(ctx, "/Budget.v1.AddService/getTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) GetPeople(ctx context.Context, in *GetPeopleRequest, opts ...grpc.CallOption) (*GetPeopleResponse, error) {
	out := new(GetPeopleResponse)
	err := c.cc.Invoke(ctx, "/Budget.v1.AddService/getPeople", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addServiceClient) GetTotalTransactionAmount(ctx context.Context, in *GetTotalTransactionAmountRequest, opts ...grpc.CallOption) (*GetTotalTransactionAmountResponse, error) {
	out := new(GetTotalTransactionAmountResponse)
	err := c.cc.Invoke(ctx, "/Budget.v1.AddService/getTotalTransactionAmount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddServiceServer is the server API for AddService service.
type AddServiceServer interface {
	GetCategories(context.Context, *Request) (*GetCategoriesResponse, error)
	GetTransactions(context.Context, *Request) (*GetTransactionsResponse, error)
	GetPeople(context.Context, *GetPeopleRequest) (*GetPeopleResponse, error)
	GetTotalTransactionAmount(context.Context, *GetTotalTransactionAmountRequest) (*GetTotalTransactionAmountResponse, error)
}

// UnimplementedAddServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAddServiceServer struct {
}

func (*UnimplementedAddServiceServer) GetCategories(context.Context, *Request) (*GetCategoriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCategories not implemented")
}
func (*UnimplementedAddServiceServer) GetTransactions(context.Context, *Request) (*GetTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransactions not implemented")
}
func (*UnimplementedAddServiceServer) GetPeople(context.Context, *GetPeopleRequest) (*GetPeopleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPeople not implemented")
}
func (*UnimplementedAddServiceServer) GetTotalTransactionAmount(context.Context, *GetTotalTransactionAmountRequest) (*GetTotalTransactionAmountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTotalTransactionAmount not implemented")
}

func RegisterAddServiceServer(s *grpc.Server, srv AddServiceServer) {
	s.RegisterService(&_AddService_serviceDesc, srv)
}

func _AddService_GetCategories_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).GetCategories(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Budget.v1.AddService/GetCategories",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).GetCategories(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_GetTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).GetTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Budget.v1.AddService/GetTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).GetTransactions(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_GetPeople_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPeopleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).GetPeople(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Budget.v1.AddService/GetPeople",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).GetPeople(ctx, req.(*GetPeopleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddService_GetTotalTransactionAmount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTotalTransactionAmountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServiceServer).GetTotalTransactionAmount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Budget.v1.AddService/GetTotalTransactionAmount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServiceServer).GetTotalTransactionAmount(ctx, req.(*GetTotalTransactionAmountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Budget.v1.AddService",
	HandlerType: (*AddServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getCategories",
			Handler:    _AddService_GetCategories_Handler,
		},
		{
			MethodName: "getTransactions",
			Handler:    _AddService_GetTransactions_Handler,
		},
		{
			MethodName: "getPeople",
			Handler:    _AddService_GetPeople_Handler,
		},
		{
			MethodName: "getTotalTransactionAmount",
			Handler:    _AddService_GetTotalTransactionAmount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
