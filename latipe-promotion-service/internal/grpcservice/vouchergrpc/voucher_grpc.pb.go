// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.2
// source: internal/grpcservice/vouchergrpc/voucher_grpc.proto

package vouchergrpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckoutVoucherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderTotalAmount int64        `protobuf:"varint,1,opt,name=order_total_amount,json=orderTotalAmount,proto3" json:"order_total_amount,omitempty"`
	PaymentMethod    int32        `protobuf:"varint,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	VoucherData      *VoucherData `protobuf:"bytes,3,opt,name=voucher_data,json=voucherData,proto3" json:"voucher_data,omitempty"`
	UserId           string       `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CheckoutVoucherRequest) Reset() {
	*x = CheckoutVoucherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutVoucherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutVoucherRequest) ProtoMessage() {}

func (x *CheckoutVoucherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutVoucherRequest.ProtoReflect.Descriptor instead.
func (*CheckoutVoucherRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{0}
}

func (x *CheckoutVoucherRequest) GetOrderTotalAmount() int64 {
	if x != nil {
		return x.OrderTotalAmount
	}
	return 0
}

func (x *CheckoutVoucherRequest) GetPaymentMethod() int32 {
	if x != nil {
		return x.PaymentMethod
	}
	return 0
}

func (x *CheckoutVoucherRequest) GetVoucherData() *VoucherData {
	if x != nil {
		return x.VoucherData
	}
	return nil
}

func (x *CheckoutVoucherRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CheckoutVoucherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherDetail *VoucherDetail `protobuf:"bytes,1,opt,name=voucher_detail,json=voucherDetail,proto3" json:"voucher_detail,omitempty"`
}

func (x *CheckoutVoucherResponse) Reset() {
	*x = CheckoutVoucherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutVoucherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutVoucherResponse) ProtoMessage() {}

func (x *CheckoutVoucherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutVoucherResponse.ProtoReflect.Descriptor instead.
func (*CheckoutVoucherResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{1}
}

func (x *CheckoutVoucherResponse) GetVoucherDetail() *VoucherDetail {
	if x != nil {
		return x.VoucherDetail
	}
	return nil
}

type VoucherData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VoucherCode string `protobuf:"bytes,1,opt,name=voucher_code,json=voucherCode,proto3" json:"voucher_code,omitempty"`
}

func (x *VoucherData) Reset() {
	*x = VoucherData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoucherData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoucherData) ProtoMessage() {}

func (x *VoucherData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoucherData.ProtoReflect.Descriptor instead.
func (*VoucherData) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{2}
}

func (x *VoucherData) GetVoucherCode() string {
	if x != nil {
		return x.VoucherCode
	}
	return ""
}

type VoucherDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id             string                  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	VoucherCode    string                  `protobuf:"bytes,2,opt,name=voucher_code,json=voucherCode,proto3" json:"voucher_code,omitempty"`
	VoucherType    int32                   `protobuf:"varint,3,opt,name=voucher_type,json=voucherType,proto3" json:"voucher_type,omitempty"`
	VoucherCounts  int32                   `protobuf:"varint,4,opt,name=voucher_counts,json=voucherCounts,proto3" json:"voucher_counts,omitempty"`
	Detail         string                  `protobuf:"bytes,5,opt,name=detail,proto3" json:"detail,omitempty"`
	OwnerVoucher   string                  `protobuf:"bytes,6,opt,name=owner_voucher,json=ownerVoucher,proto3" json:"owner_voucher,omitempty"`
	Status         int32                   `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	DiscountData   *CheckingDiscountData   `protobuf:"bytes,8,opt,name=discount_data,json=discountData,proto3" json:"discount_data,omitempty"`
	VoucherRequire *CheckingVoucherRequire `protobuf:"bytes,9,opt,name=voucher_require,json=voucherRequire,proto3" json:"voucher_require,omitempty"`
	StatedTime     *timestamppb.Timestamp  `protobuf:"bytes,10,opt,name=stated_time,json=statedTime,proto3" json:"stated_time,omitempty"`
	EndedTime      *timestamppb.Timestamp  `protobuf:"bytes,11,opt,name=ended_time,json=endedTime,proto3" json:"ended_time,omitempty"`
	CreatedAt      *timestamppb.Timestamp  `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt      *timestamppb.Timestamp  `protobuf:"bytes,13,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *VoucherDetail) Reset() {
	*x = VoucherDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VoucherDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoucherDetail) ProtoMessage() {}

func (x *VoucherDetail) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoucherDetail.ProtoReflect.Descriptor instead.
func (*VoucherDetail) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{3}
}

func (x *VoucherDetail) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *VoucherDetail) GetVoucherCode() string {
	if x != nil {
		return x.VoucherCode
	}
	return ""
}

func (x *VoucherDetail) GetVoucherType() int32 {
	if x != nil {
		return x.VoucherType
	}
	return 0
}

func (x *VoucherDetail) GetVoucherCounts() int32 {
	if x != nil {
		return x.VoucherCounts
	}
	return 0
}

func (x *VoucherDetail) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

func (x *VoucherDetail) GetOwnerVoucher() string {
	if x != nil {
		return x.OwnerVoucher
	}
	return ""
}

func (x *VoucherDetail) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *VoucherDetail) GetDiscountData() *CheckingDiscountData {
	if x != nil {
		return x.DiscountData
	}
	return nil
}

func (x *VoucherDetail) GetVoucherRequire() *CheckingVoucherRequire {
	if x != nil {
		return x.VoucherRequire
	}
	return nil
}

func (x *VoucherDetail) GetStatedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StatedTime
	}
	return nil
}

func (x *VoucherDetail) GetEndedTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndedTime
	}
	return nil
}

func (x *VoucherDetail) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *VoucherDetail) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type CheckingVoucherRequire struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MinRequire        int64 `protobuf:"varint,1,opt,name=min_require,json=minRequire,proto3" json:"min_require,omitempty"`
	PaymentMethod     int32 `protobuf:"varint,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	MaxVoucherPerUser int32 `protobuf:"varint,3,opt,name=max_voucher_per_user,json=maxVoucherPerUser,proto3" json:"max_voucher_per_user,omitempty"`
}

func (x *CheckingVoucherRequire) Reset() {
	*x = CheckingVoucherRequire{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckingVoucherRequire) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckingVoucherRequire) ProtoMessage() {}

func (x *CheckingVoucherRequire) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckingVoucherRequire.ProtoReflect.Descriptor instead.
func (*CheckingVoucherRequire) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{4}
}

func (x *CheckingVoucherRequire) GetMinRequire() int64 {
	if x != nil {
		return x.MinRequire
	}
	return 0
}

func (x *CheckingVoucherRequire) GetPaymentMethod() int32 {
	if x != nil {
		return x.PaymentMethod
	}
	return 0
}

func (x *CheckingVoucherRequire) GetMaxVoucherPerUser() int32 {
	if x != nil {
		return x.MaxVoucherPerUser
	}
	return 0
}

type CheckingDiscountData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DiscountType    int32   `protobuf:"varint,1,opt,name=discount_type,json=discountType,proto3" json:"discount_type,omitempty"`
	ShippingValue   uint64  `protobuf:"varint,2,opt,name=shipping_value,json=shippingValue,proto3" json:"shipping_value,omitempty"`
	DiscountValue   uint64  `protobuf:"varint,3,opt,name=discount_value,json=discountValue,proto3" json:"discount_value,omitempty"`
	DiscountPercent float32 `protobuf:"fixed32,4,opt,name=discount_percent,json=discountPercent,proto3" json:"discount_percent,omitempty"`
	MaximumValue    uint64  `protobuf:"varint,5,opt,name=maximum_value,json=maximumValue,proto3" json:"maximum_value,omitempty"`
}

func (x *CheckingDiscountData) Reset() {
	*x = CheckingDiscountData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckingDiscountData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckingDiscountData) ProtoMessage() {}

func (x *CheckingDiscountData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckingDiscountData.ProtoReflect.Descriptor instead.
func (*CheckingDiscountData) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{5}
}

func (x *CheckingDiscountData) GetDiscountType() int32 {
	if x != nil {
		return x.DiscountType
	}
	return 0
}

func (x *CheckingDiscountData) GetShippingValue() uint64 {
	if x != nil {
		return x.ShippingValue
	}
	return 0
}

func (x *CheckingDiscountData) GetDiscountValue() uint64 {
	if x != nil {
		return x.DiscountValue
	}
	return 0
}

func (x *CheckingDiscountData) GetDiscountPercent() float32 {
	if x != nil {
		return x.DiscountPercent
	}
	return 0
}

func (x *CheckingDiscountData) GetMaximumValue() uint64 {
	if x != nil {
		return x.MaximumValue
	}
	return 0
}

type ApplyVoucherRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CheckoutData *CheckoutRequest `protobuf:"bytes,2,opt,name=checkout_data,json=checkoutData,proto3" json:"checkout_data,omitempty"`
	Vouchers     []string         `protobuf:"bytes,3,rep,name=vouchers,proto3" json:"vouchers,omitempty"`
}

func (x *ApplyVoucherRequest) Reset() {
	*x = ApplyVoucherRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyVoucherRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyVoucherRequest) ProtoMessage() {}

func (x *ApplyVoucherRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyVoucherRequest.ProtoReflect.Descriptor instead.
func (*ApplyVoucherRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{6}
}

func (x *ApplyVoucherRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ApplyVoucherRequest) GetCheckoutData() *CheckoutRequest {
	if x != nil {
		return x.CheckoutData
	}
	return nil
}

func (x *ApplyVoucherRequest) GetVouchers() []string {
	if x != nil {
		return x.Vouchers
	}
	return nil
}

type CheckoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckoutId string   `protobuf:"bytes,1,opt,name=checkout_id,json=checkoutId,proto3" json:"checkout_id,omitempty"`
	OrderIds   []string `protobuf:"bytes,2,rep,name=order_ids,json=orderIds,proto3" json:"order_ids,omitempty"`
}

func (x *CheckoutRequest) Reset() {
	*x = CheckoutRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutRequest) ProtoMessage() {}

func (x *CheckoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutRequest.ProtoReflect.Descriptor instead.
func (*CheckoutRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{7}
}

func (x *CheckoutRequest) GetCheckoutId() string {
	if x != nil {
		return x.CheckoutId
	}
	return ""
}

func (x *CheckoutRequest) GetOrderIds() []string {
	if x != nil {
		return x.OrderIds
	}
	return nil
}

type ApplyVoucherResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *ApplyVoucherResponse) Reset() {
	*x = ApplyVoucherResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyVoucherResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyVoucherResponse) ProtoMessage() {}

func (x *ApplyVoucherResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyVoucherResponse.ProtoReflect.Descriptor instead.
func (*ApplyVoucherResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP(), []int{8}
}

func (x *ApplyVoucherResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

var File_internal_grpcservice_vouchergrpc_voucher_grpc_proto protoreflect.FileDescriptor

var file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDesc = []byte{
	0x0a, 0x33, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb7, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x6f, 0x75, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2f, 0x0a, 0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x56,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b, 0x76, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x50, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x56, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0e, 0x76,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x52, 0x0d, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x22, 0x30, 0x0a, 0x0b, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x22, 0xcd, 0x04, 0x0a, 0x0d, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x6f, 0x75,
	0x63, 0x68, 0x65, 0x72, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e,
	0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x40, 0x0a, 0x0f, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x52, 0x0e, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x65, 0x64, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x65, 0x6e, 0x64, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0x91, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x69, 0x6e,
	0x67, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x6d, 0x69, 0x6e, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x2f, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x5f, 0x76,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x11, 0x6d, 0x61, 0x78, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65,
	0x72, 0x50, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x22, 0xd9, 0x01, 0x0a, 0x14, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d,
	0x73, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0d, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0f,
	0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x12,
	0x23, 0x0a, 0x0d, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0c, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x13, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x56, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x35, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08,
	0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x73, 0x22, 0x4f, 0x0a, 0x0f, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x14, 0x41, 0x70, 0x70,
	0x6c, 0x79, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x32, 0xac, 0x01, 0x0a, 0x0e, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x1a, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x56,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73,
	0x65, 0x12, 0x17, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x56, 0x6f, 0x75, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x6f, 0x75, 0x74, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x16, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x56,
	0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x54, 0x6f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x12, 0x14, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x56, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x56, 0x6f,
	0x75, 0x63, 0x68, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0e, 0x5a, 0x0c, 0x76, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescOnce sync.Once
	file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescData = file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDesc
)

func file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescGZIP() []byte {
	file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescOnce.Do(func() {
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescData)
	})
	return file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDescData
}

var file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_goTypes = []interface{}{
	(*CheckoutVoucherRequest)(nil),  // 0: CheckoutVoucherRequest
	(*CheckoutVoucherResponse)(nil), // 1: CheckoutVoucherResponse
	(*VoucherData)(nil),             // 2: VoucherData
	(*VoucherDetail)(nil),           // 3: VoucherDetail
	(*CheckingVoucherRequire)(nil),  // 4: CheckingVoucherRequire
	(*CheckingDiscountData)(nil),    // 5: CheckingDiscountData
	(*ApplyVoucherRequest)(nil),     // 6: ApplyVoucherRequest
	(*CheckoutRequest)(nil),         // 7: CheckoutRequest
	(*ApplyVoucherResponse)(nil),    // 8: ApplyVoucherResponse
	(*timestamppb.Timestamp)(nil),   // 9: google.protobuf.Timestamp
}
var file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_depIdxs = []int32{
	2,  // 0: CheckoutVoucherRequest.voucher_data:type_name -> VoucherData
	3,  // 1: CheckoutVoucherResponse.voucher_detail:type_name -> VoucherDetail
	5,  // 2: VoucherDetail.discount_data:type_name -> CheckingDiscountData
	4,  // 3: VoucherDetail.voucher_require:type_name -> CheckingVoucherRequire
	9,  // 4: VoucherDetail.stated_time:type_name -> google.protobuf.Timestamp
	9,  // 5: VoucherDetail.ended_time:type_name -> google.protobuf.Timestamp
	9,  // 6: VoucherDetail.created_at:type_name -> google.protobuf.Timestamp
	9,  // 7: VoucherDetail.updated_at:type_name -> google.protobuf.Timestamp
	7,  // 8: ApplyVoucherRequest.checkout_data:type_name -> CheckoutRequest
	0,  // 9: VoucherService.CheckoutVoucherForPurchase:input_type -> CheckoutVoucherRequest
	6,  // 10: VoucherService.ApplyVoucherToPurchase:input_type -> ApplyVoucherRequest
	1,  // 11: VoucherService.CheckoutVoucherForPurchase:output_type -> CheckoutVoucherResponse
	8,  // 12: VoucherService.ApplyVoucherToPurchase:output_type -> ApplyVoucherResponse
	11, // [11:13] is the sub-list for method output_type
	9,  // [9:11] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_init() }
func file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_init() {
	if File_internal_grpcservice_vouchergrpc_voucher_grpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutVoucherRequest); i {
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
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutVoucherResponse); i {
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
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoucherData); i {
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
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VoucherDetail); i {
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
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckingVoucherRequire); i {
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
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckingDiscountData); i {
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
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyVoucherRequest); i {
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
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutRequest); i {
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
		file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyVoucherResponse); i {
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
			RawDescriptor: file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_goTypes,
		DependencyIndexes: file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_depIdxs,
		MessageInfos:      file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_msgTypes,
	}.Build()
	File_internal_grpcservice_vouchergrpc_voucher_grpc_proto = out.File
	file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_rawDesc = nil
	file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_goTypes = nil
	file_internal_grpcservice_vouchergrpc_voucher_grpc_proto_depIdxs = nil
}
