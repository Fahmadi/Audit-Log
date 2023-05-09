// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: audit/v1/audit_service.proto

package auditv1

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

type ListAuditLogsByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Page   *Page  `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListAuditLogsByUserIDRequest) Reset() {
	*x = ListAuditLogsByUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditLogsByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditLogsByUserIDRequest) ProtoMessage() {}

func (x *ListAuditLogsByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditLogsByUserIDRequest.ProtoReflect.Descriptor instead.
func (*ListAuditLogsByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_service_proto_rawDescGZIP(), []int{0}
}

func (x *ListAuditLogsByUserIDRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ListAuditLogsByUserIDRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListAuditLogsByUserIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audits []*Audit `protobuf:"bytes,1,rep,name=audits,proto3" json:"audits,omitempty"`
	Page   *Page    `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListAuditLogsByUserIDResponse) Reset() {
	*x = ListAuditLogsByUserIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditLogsByUserIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditLogsByUserIDResponse) ProtoMessage() {}

func (x *ListAuditLogsByUserIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditLogsByUserIDResponse.ProtoReflect.Descriptor instead.
func (*ListAuditLogsByUserIDResponse) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_service_proto_rawDescGZIP(), []int{1}
}

func (x *ListAuditLogsByUserIDResponse) GetAudits() []*Audit {
	if x != nil {
		return x.Audits
	}
	return nil
}

func (x *ListAuditLogsByUserIDResponse) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListAuditLogsByEntityIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityId string `protobuf:"bytes,1,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Entity   string `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	Page     *Page  `protobuf:"bytes,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListAuditLogsByEntityIDRequest) Reset() {
	*x = ListAuditLogsByEntityIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditLogsByEntityIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditLogsByEntityIDRequest) ProtoMessage() {}

func (x *ListAuditLogsByEntityIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditLogsByEntityIDRequest.ProtoReflect.Descriptor instead.
func (*ListAuditLogsByEntityIDRequest) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListAuditLogsByEntityIDRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *ListAuditLogsByEntityIDRequest) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *ListAuditLogsByEntityIDRequest) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type ListAuditLogsByEntityIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Audits []*Audit `protobuf:"bytes,1,rep,name=audits,proto3" json:"audits,omitempty"`
	Page   *Page    `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *ListAuditLogsByEntityIDResponse) Reset() {
	*x = ListAuditLogsByEntityIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAuditLogsByEntityIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAuditLogsByEntityIDResponse) ProtoMessage() {}

func (x *ListAuditLogsByEntityIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAuditLogsByEntityIDResponse.ProtoReflect.Descriptor instead.
func (*ListAuditLogsByEntityIDResponse) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListAuditLogsByEntityIDResponse) GetAudits() []*Audit {
	if x != nil {
		return x.Audits
	}
	return nil
}

func (x *ListAuditLogsByEntityIDResponse) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

type Audit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action    string                 `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	UserId    string                 `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Entity    string                 `protobuf:"bytes,4,opt,name=entity,proto3" json:"entity,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Message   string                 `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Audit) Reset() {
	*x = Audit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Audit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Audit) ProtoMessage() {}

func (x *Audit) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Audit.ProtoReflect.Descriptor instead.
func (*Audit) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_service_proto_rawDescGZIP(), []int{4}
}

func (x *Audit) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *Audit) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Audit) GetEntity() string {
	if x != nil {
		return x.Entity
	}
	return ""
}

func (x *Audit) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Audit) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page  uint32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size  uint32 `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Total uint32 `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_audit_v1_audit_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_audit_v1_audit_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_audit_v1_audit_service_proto_rawDescGZIP(), []int{5}
}

func (x *Page) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *Page) GetSize() uint32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Page) GetTotal() uint32 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_audit_v1_audit_service_proto protoreflect.FileDescriptor

var file_audit_v1_audit_service_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x1c, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x6c, 0x0a, 0x1d, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x74, 0x52, 0x06, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73,
	0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x22, 0x79, 0x0a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0x6e, 0x0a, 0x1f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73,
	0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x52, 0x06, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x12, 0x22, 0x0a, 0x04, 0x70,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22,
	0xa4, 0x01, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x44, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xec, 0x01, 0x0a,
	0x0c, 0x41, 0x75, 0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6a, 0x0a,
	0x15, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x26, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x70, 0x0a, 0x17, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x49, 0x44, 0x12, 0x28, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x79, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x29,
	0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x42, 0x79, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x8c, 0x01, 0x0a, 0x0c,
	0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x11, 0x41, 0x75,
	0x64, 0x69, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x28, 0x61, 0x75, 0x64, 0x69, 0x74, 0x4c, 0x6f, 0x67, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74,
	0x2f, 0x76, 0x31, 0x3b, 0x61, 0x75, 0x64, 0x69, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x58,
	0x58, 0xaa, 0x02, 0x08, 0x41, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x41,
	0x75, 0x64, 0x69, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x41, 0x75, 0x64, 0x69, 0x74, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x09, 0x41, 0x75, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_audit_v1_audit_service_proto_rawDescOnce sync.Once
	file_audit_v1_audit_service_proto_rawDescData = file_audit_v1_audit_service_proto_rawDesc
)

func file_audit_v1_audit_service_proto_rawDescGZIP() []byte {
	file_audit_v1_audit_service_proto_rawDescOnce.Do(func() {
		file_audit_v1_audit_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_audit_v1_audit_service_proto_rawDescData)
	})
	return file_audit_v1_audit_service_proto_rawDescData
}

var file_audit_v1_audit_service_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_audit_v1_audit_service_proto_goTypes = []interface{}{
	(*ListAuditLogsByUserIDRequest)(nil),    // 0: audit.v1.ListAuditLogsByUserIDRequest
	(*ListAuditLogsByUserIDResponse)(nil),   // 1: audit.v1.ListAuditLogsByUserIDResponse
	(*ListAuditLogsByEntityIDRequest)(nil),  // 2: audit.v1.ListAuditLogsByEntityIDRequest
	(*ListAuditLogsByEntityIDResponse)(nil), // 3: audit.v1.ListAuditLogsByEntityIDResponse
	(*Audit)(nil),                           // 4: audit.v1.Audit
	(*Page)(nil),                            // 5: audit.v1.Page
	(*timestamppb.Timestamp)(nil),           // 6: google.protobuf.Timestamp
}
var file_audit_v1_audit_service_proto_depIdxs = []int32{
	5, // 0: audit.v1.ListAuditLogsByUserIDRequest.page:type_name -> audit.v1.Page
	4, // 1: audit.v1.ListAuditLogsByUserIDResponse.audits:type_name -> audit.v1.Audit
	5, // 2: audit.v1.ListAuditLogsByUserIDResponse.page:type_name -> audit.v1.Page
	5, // 3: audit.v1.ListAuditLogsByEntityIDRequest.page:type_name -> audit.v1.Page
	4, // 4: audit.v1.ListAuditLogsByEntityIDResponse.audits:type_name -> audit.v1.Audit
	5, // 5: audit.v1.ListAuditLogsByEntityIDResponse.page:type_name -> audit.v1.Page
	6, // 6: audit.v1.Audit.timestamp:type_name -> google.protobuf.Timestamp
	0, // 7: audit.v1.AuditService.ListAuditLogsByUserID:input_type -> audit.v1.ListAuditLogsByUserIDRequest
	2, // 8: audit.v1.AuditService.ListAuditLogsByEntityID:input_type -> audit.v1.ListAuditLogsByEntityIDRequest
	1, // 9: audit.v1.AuditService.ListAuditLogsByUserID:output_type -> audit.v1.ListAuditLogsByUserIDResponse
	3, // 10: audit.v1.AuditService.ListAuditLogsByEntityID:output_type -> audit.v1.ListAuditLogsByEntityIDResponse
	9, // [9:11] is the sub-list for method output_type
	7, // [7:9] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_audit_v1_audit_service_proto_init() }
func file_audit_v1_audit_service_proto_init() {
	if File_audit_v1_audit_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_audit_v1_audit_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditLogsByUserIDRequest); i {
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
		file_audit_v1_audit_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditLogsByUserIDResponse); i {
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
		file_audit_v1_audit_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditLogsByEntityIDRequest); i {
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
		file_audit_v1_audit_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAuditLogsByEntityIDResponse); i {
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
		file_audit_v1_audit_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Audit); i {
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
		file_audit_v1_audit_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
			RawDescriptor: file_audit_v1_audit_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_audit_v1_audit_service_proto_goTypes,
		DependencyIndexes: file_audit_v1_audit_service_proto_depIdxs,
		MessageInfos:      file_audit_v1_audit_service_proto_msgTypes,
	}.Build()
	File_audit_v1_audit_service_proto = out.File
	file_audit_v1_audit_service_proto_rawDesc = nil
	file_audit_v1_audit_service_proto_goTypes = nil
	file_audit_v1_audit_service_proto_depIdxs = nil
}
