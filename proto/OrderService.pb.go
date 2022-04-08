// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: proto/OrderService.proto

package OrderService

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

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId      string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderTotal  float32 `protobuf:"fixed32,3,opt,name=order_total,json=orderTotal,proto3" json:"order_total,omitempty"`
	OrderPlaced int64   `protobuf:"varint,4,opt,name=order_placed,json=orderPlaced,proto3" json:"order_placed,omitempty"`
	OrderPaid   bool    `protobuf:"varint,5,opt,name=order_paid,json=orderPaid,proto3" json:"order_paid,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{0}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetOrderTotal() float32 {
	if x != nil {
		return x.OrderTotal
	}
	return 0
}

func (x *Order) GetOrderPlaced() int64 {
	if x != nil {
		return x.OrderPlaced
	}
	return 0
}

func (x *Order) GetOrderPaid() bool {
	if x != nil {
		return x.OrderPaid
	}
	return false
}

type OrdersForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *OrdersForUserRequest) Reset() {
	*x = OrdersForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersForUserRequest) ProtoMessage() {}

func (x *OrdersForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersForUserRequest.ProtoReflect.Descriptor instead.
func (*OrdersForUserRequest) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{1}
}

func (x *OrdersForUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type OrdersForUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order []*Order `protobuf:"bytes,1,rep,name=order,proto3" json:"order,omitempty"`
}

func (x *OrdersForUserResponse) Reset() {
	*x = OrdersForUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrdersForUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrdersForUserResponse) ProtoMessage() {}

func (x *OrdersForUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrdersForUserResponse.ProtoReflect.Descriptor instead.
func (*OrdersForUserResponse) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{2}
}

func (x *OrdersForUserResponse) GetOrder() []*Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type AddOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderTotal  float32 `protobuf:"fixed32,3,opt,name=order_total,json=orderTotal,proto3" json:"order_total,omitempty"`
	OrderPlaced int64   `protobuf:"varint,4,opt,name=order_placed,json=orderPlaced,proto3" json:"order_placed,omitempty"`
	OrderPaid   bool    `protobuf:"varint,5,opt,name=order_paid,json=orderPaid,proto3" json:"order_paid,omitempty"`
}

func (x *AddOrderRequest) Reset() {
	*x = AddOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderRequest) ProtoMessage() {}

func (x *AddOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderRequest.ProtoReflect.Descriptor instead.
func (*AddOrderRequest) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{3}
}

func (x *AddOrderRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddOrderRequest) GetOrderTotal() float32 {
	if x != nil {
		return x.OrderTotal
	}
	return 0
}

func (x *AddOrderRequest) GetOrderPlaced() int64 {
	if x != nil {
		return x.OrderPlaced
	}
	return 0
}

func (x *AddOrderRequest) GetOrderPaid() bool {
	if x != nil {
		return x.OrderPaid
	}
	return false
}

type AddOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *AddOrderResponse) Reset() {
	*x = AddOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrderResponse) ProtoMessage() {}

func (x *AddOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrderResponse.ProtoReflect.Descriptor instead.
func (*AddOrderResponse) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{4}
}

func (x *AddOrderResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *OrderByIdRequest) Reset() {
	*x = OrderByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderByIdRequest) ProtoMessage() {}

func (x *OrderByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderByIdRequest.ProtoReflect.Descriptor instead.
func (*OrderByIdRequest) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{5}
}

func (x *OrderByIdRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

type OrderByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Order *Order `protobuf:"bytes,1,opt,name=order,proto3" json:"order,omitempty"`
}

func (x *OrderByIdResponse) Reset() {
	*x = OrderByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderByIdResponse) ProtoMessage() {}

func (x *OrderByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderByIdResponse.ProtoReflect.Descriptor instead.
func (*OrderByIdResponse) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{6}
}

func (x *OrderByIdResponse) GetOrder() *Order {
	if x != nil {
		return x.Order
	}
	return nil
}

type OrderPaymentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Paid    bool   `protobuf:"varint,2,opt,name=paid,proto3" json:"paid,omitempty"`
}

func (x *OrderPaymentStatusRequest) Reset() {
	*x = OrderPaymentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPaymentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPaymentStatusRequest) ProtoMessage() {}

func (x *OrderPaymentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPaymentStatusRequest.ProtoReflect.Descriptor instead.
func (*OrderPaymentStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{7}
}

func (x *OrderPaymentStatusRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderPaymentStatusRequest) GetPaid() bool {
	if x != nil {
		return x.Paid
	}
	return false
}

type OrderPaymentStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *OrderPaymentStatusResponse) Reset() {
	*x = OrderPaymentStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_OrderService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderPaymentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderPaymentStatusResponse) ProtoMessage() {}

func (x *OrderPaymentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_OrderService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderPaymentStatusResponse.ProtoReflect.Descriptor instead.
func (*OrderPaymentStatusResponse) Descriptor() ([]byte, []int) {
	return file_proto_OrderService_proto_rawDescGZIP(), []int{8}
}

var File_proto_OrderService_proto protoreflect.FileDescriptor

var file_proto_OrderService_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x93, 0x01, 0x0a, 0x05, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x69, 0x64, 0x22, 0x39,
	0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x4a,
	0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x15, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x97, 0x01,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x6c, 0x61, 0x63, 0x65, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x69, 0x64, 0x4a, 0x04, 0x08,
	0x01, 0x10, 0x02, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x2d, 0x0a, 0x10, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x22, 0x3e, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x05, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x05,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x4a, 0x0a, 0x19, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x70, 0x61, 0x69,
	0x64, 0x22, 0x1c, 0x0a, 0x1a, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0xfe, 0x02, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x5d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x22, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x46, 0x6f,
	0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x08, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x1d, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1e, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x6f, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x27, 0x2e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x16, 0x5a, 0x14, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_OrderService_proto_rawDescOnce sync.Once
	file_proto_OrderService_proto_rawDescData = file_proto_OrderService_proto_rawDesc
)

func file_proto_OrderService_proto_rawDescGZIP() []byte {
	file_proto_OrderService_proto_rawDescOnce.Do(func() {
		file_proto_OrderService_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_OrderService_proto_rawDescData)
	})
	return file_proto_OrderService_proto_rawDescData
}

var file_proto_OrderService_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_OrderService_proto_goTypes = []interface{}{
	(*Order)(nil),                      // 0: OrderService.Order
	(*OrdersForUserRequest)(nil),       // 1: OrderService.OrdersForUserRequest
	(*OrdersForUserResponse)(nil),      // 2: OrderService.OrdersForUserResponse
	(*AddOrderRequest)(nil),            // 3: OrderService.AddOrderRequest
	(*AddOrderResponse)(nil),           // 4: OrderService.AddOrderResponse
	(*OrderByIdRequest)(nil),           // 5: OrderService.OrderByIdRequest
	(*OrderByIdResponse)(nil),          // 6: OrderService.OrderByIdResponse
	(*OrderPaymentStatusRequest)(nil),  // 7: OrderService.OrderPaymentStatusRequest
	(*OrderPaymentStatusResponse)(nil), // 8: OrderService.OrderPaymentStatusResponse
}
var file_proto_OrderService_proto_depIdxs = []int32{
	0, // 0: OrderService.OrdersForUserResponse.order:type_name -> OrderService.Order
	0, // 1: OrderService.AddOrderResponse.order:type_name -> OrderService.Order
	0, // 2: OrderService.OrderByIdResponse.order:type_name -> OrderService.Order
	1, // 3: OrderService.OrderService.GetOrdersForUser:input_type -> OrderService.OrdersForUserRequest
	3, // 4: OrderService.OrderService.AddOrder:input_type -> OrderService.AddOrderRequest
	5, // 5: OrderService.OrderService.GetOrderById:input_type -> OrderService.OrderByIdRequest
	7, // 6: OrderService.OrderService.UpdateOrderPaymentStatus:input_type -> OrderService.OrderPaymentStatusRequest
	2, // 7: OrderService.OrderService.GetOrdersForUser:output_type -> OrderService.OrdersForUserResponse
	4, // 8: OrderService.OrderService.AddOrder:output_type -> OrderService.AddOrderResponse
	6, // 9: OrderService.OrderService.GetOrderById:output_type -> OrderService.OrderByIdResponse
	8, // 10: OrderService.OrderService.UpdateOrderPaymentStatus:output_type -> OrderService.OrderPaymentStatusResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_OrderService_proto_init() }
func file_proto_OrderService_proto_init() {
	if File_proto_OrderService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_OrderService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_proto_OrderService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersForUserRequest); i {
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
		file_proto_OrderService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrdersForUserResponse); i {
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
		file_proto_OrderService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderRequest); i {
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
		file_proto_OrderService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrderResponse); i {
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
		file_proto_OrderService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderByIdRequest); i {
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
		file_proto_OrderService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderByIdResponse); i {
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
		file_proto_OrderService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPaymentStatusRequest); i {
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
		file_proto_OrderService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderPaymentStatusResponse); i {
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
			RawDescriptor: file_proto_OrderService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_OrderService_proto_goTypes,
		DependencyIndexes: file_proto_OrderService_proto_depIdxs,
		MessageInfos:      file_proto_OrderService_proto_msgTypes,
	}.Build()
	File_proto_OrderService_proto = out.File
	file_proto_OrderService_proto_rawDesc = nil
	file_proto_OrderService_proto_goTypes = nil
	file_proto_OrderService_proto_depIdxs = nil
}
