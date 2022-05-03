// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/notification_service.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type NotificationReceiver int32

const (
	NotificationReceiver_NOTIFICATION_RECEIVER_UNKNOWN NotificationReceiver = 0
	// Notification is received by collaborators of the entity.
	// If the collaborator is an organization, the notification is received by organization members.
	NotificationReceiver_NOTIFICATION_RECEIVER_COLLABORATOR NotificationReceiver = 1
	// Notification is received by administrative contact of the entity.
	// If this is an organization, the notification is received by organization members.
	NotificationReceiver_NOTIFICATION_RECEIVER_ADMINISTRATIVE_CONTACT NotificationReceiver = 3
	// Notification is received by technical contact of the entity.
	// If this is an organization, the notification is received by organization members.
	NotificationReceiver_NOTIFICATION_RECEIVER_TECHNICAL_CONTACT NotificationReceiver = 4
)

var NotificationReceiver_name = map[int32]string{
	0: "NOTIFICATION_RECEIVER_UNKNOWN",
	1: "NOTIFICATION_RECEIVER_COLLABORATOR",
	3: "NOTIFICATION_RECEIVER_ADMINISTRATIVE_CONTACT",
	4: "NOTIFICATION_RECEIVER_TECHNICAL_CONTACT",
}

var NotificationReceiver_value = map[string]int32{
	"NOTIFICATION_RECEIVER_UNKNOWN":                0,
	"NOTIFICATION_RECEIVER_COLLABORATOR":           1,
	"NOTIFICATION_RECEIVER_ADMINISTRATIVE_CONTACT": 3,
	"NOTIFICATION_RECEIVER_TECHNICAL_CONTACT":      4,
}

func (x NotificationReceiver) String() string {
	return proto.EnumName(NotificationReceiver_name, int32(x))
}

func (NotificationReceiver) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62db76852b39e787, []int{0}
}

type NotificationStatus int32

const (
	NotificationStatus_NOTIFICATION_STATUS_UNSEEN   NotificationStatus = 0
	NotificationStatus_NOTIFICATION_STATUS_SEEN     NotificationStatus = 1
	NotificationStatus_NOTIFICATION_STATUS_ARCHIVED NotificationStatus = 2
)

var NotificationStatus_name = map[int32]string{
	0: "NOTIFICATION_STATUS_UNSEEN",
	1: "NOTIFICATION_STATUS_SEEN",
	2: "NOTIFICATION_STATUS_ARCHIVED",
}

var NotificationStatus_value = map[string]int32{
	"NOTIFICATION_STATUS_UNSEEN":   0,
	"NOTIFICATION_STATUS_SEEN":     1,
	"NOTIFICATION_STATUS_ARCHIVED": 2,
}

func (x NotificationStatus) String() string {
	return proto.EnumName(NotificationStatus_name, int32(x))
}

func (NotificationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_62db76852b39e787, []int{1}
}

type Notification struct {
	// The immutable ID of the notification. Generated by the server.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The time when the notification was triggered.
	CreatedAt *types.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// The entity this notification is about.
	EntityIds *EntityIdentifiers `protobuf:"bytes,3,opt,name=entity_ids,json=entityIds,proto3" json:"entity_ids,omitempty"`
	// The type of this notification.
	NotificationType string `protobuf:"bytes,4,opt,name=notification_type,json=notificationType,proto3" json:"notification_type,omitempty"`
	// The data related to the notification.
	Data *types.Any `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
	// If the notification was triggered by a user action, this contains the identifiers of the user that triggered the notification.
	SenderIds *UserIdentifiers `protobuf:"bytes,6,opt,name=sender_ids,json=senderIds,proto3" json:"sender_ids,omitempty"`
	// Relation of the notification receiver to the entity.
	Receivers []NotificationReceiver `protobuf:"varint,8,rep,packed,name=receivers,proto3,enum=ttn.lorawan.v3.NotificationReceiver" json:"receivers,omitempty"`
	// Whether an email was sent for the notification.
	Email bool `protobuf:"varint,9,opt,name=email,proto3" json:"email,omitempty"`
	// The status of the notification.
	Status NotificationStatus `protobuf:"varint,10,opt,name=status,proto3,enum=ttn.lorawan.v3.NotificationStatus" json:"status,omitempty"`
	// The time when the notification status was updated.
	StatusUpdatedAt      *types.Timestamp `protobuf:"bytes,11,opt,name=status_updated_at,json=statusUpdatedAt,proto3" json:"status_updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_62db76852b39e787, []int{0}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (m *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(m, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Notification) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *Notification) GetEntityIds() *EntityIdentifiers {
	if m != nil {
		return m.EntityIds
	}
	return nil
}

func (m *Notification) GetNotificationType() string {
	if m != nil {
		return m.NotificationType
	}
	return ""
}

func (m *Notification) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Notification) GetSenderIds() *UserIdentifiers {
	if m != nil {
		return m.SenderIds
	}
	return nil
}

func (m *Notification) GetReceivers() []NotificationReceiver {
	if m != nil {
		return m.Receivers
	}
	return nil
}

func (m *Notification) GetEmail() bool {
	if m != nil {
		return m.Email
	}
	return false
}

func (m *Notification) GetStatus() NotificationStatus {
	if m != nil {
		return m.Status
	}
	return NotificationStatus_NOTIFICATION_STATUS_UNSEEN
}

func (m *Notification) GetStatusUpdatedAt() *types.Timestamp {
	if m != nil {
		return m.StatusUpdatedAt
	}
	return nil
}

type CreateNotificationRequest struct {
	// The entity this notification is about.
	EntityIds *EntityIdentifiers `protobuf:"bytes,1,opt,name=entity_ids,json=entityIds,proto3" json:"entity_ids,omitempty"`
	// The type of this notification.
	NotificationType string `protobuf:"bytes,2,opt,name=notification_type,json=notificationType,proto3" json:"notification_type,omitempty"`
	// The data related to the notification.
	Data *types.Any `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	// If the notification was triggered by a user action, this contains the identifiers of the user that triggered the notification.
	SenderIds *UserIdentifiers `protobuf:"bytes,4,opt,name=sender_ids,json=senderIds,proto3" json:"sender_ids,omitempty"`
	// Receivers of the notification.
	Receivers []NotificationReceiver `protobuf:"varint,5,rep,packed,name=receivers,proto3,enum=ttn.lorawan.v3.NotificationReceiver" json:"receivers,omitempty"`
	// Whether an email should be sent for the notification.
	Email                bool     `protobuf:"varint,6,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNotificationRequest) Reset()         { *m = CreateNotificationRequest{} }
func (m *CreateNotificationRequest) String() string { return proto.CompactTextString(m) }
func (*CreateNotificationRequest) ProtoMessage()    {}
func (*CreateNotificationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62db76852b39e787, []int{1}
}
func (m *CreateNotificationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNotificationRequest.Unmarshal(m, b)
}
func (m *CreateNotificationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNotificationRequest.Marshal(b, m, deterministic)
}
func (m *CreateNotificationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNotificationRequest.Merge(m, src)
}
func (m *CreateNotificationRequest) XXX_Size() int {
	return xxx_messageInfo_CreateNotificationRequest.Size(m)
}
func (m *CreateNotificationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNotificationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNotificationRequest proto.InternalMessageInfo

func (m *CreateNotificationRequest) GetEntityIds() *EntityIdentifiers {
	if m != nil {
		return m.EntityIds
	}
	return nil
}

func (m *CreateNotificationRequest) GetNotificationType() string {
	if m != nil {
		return m.NotificationType
	}
	return ""
}

func (m *CreateNotificationRequest) GetData() *types.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *CreateNotificationRequest) GetSenderIds() *UserIdentifiers {
	if m != nil {
		return m.SenderIds
	}
	return nil
}

func (m *CreateNotificationRequest) GetReceivers() []NotificationReceiver {
	if m != nil {
		return m.Receivers
	}
	return nil
}

func (m *CreateNotificationRequest) GetEmail() bool {
	if m != nil {
		return m.Email
	}
	return false
}

type CreateNotificationResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateNotificationResponse) Reset()         { *m = CreateNotificationResponse{} }
func (m *CreateNotificationResponse) String() string { return proto.CompactTextString(m) }
func (*CreateNotificationResponse) ProtoMessage()    {}
func (*CreateNotificationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62db76852b39e787, []int{2}
}
func (m *CreateNotificationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateNotificationResponse.Unmarshal(m, b)
}
func (m *CreateNotificationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateNotificationResponse.Marshal(b, m, deterministic)
}
func (m *CreateNotificationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateNotificationResponse.Merge(m, src)
}
func (m *CreateNotificationResponse) XXX_Size() int {
	return xxx_messageInfo_CreateNotificationResponse.Size(m)
}
func (m *CreateNotificationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateNotificationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateNotificationResponse proto.InternalMessageInfo

func (m *CreateNotificationResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListNotificationsRequest struct {
	// The IDs of the receiving user.
	ReceiverIds *UserIdentifiers `protobuf:"bytes,1,opt,name=receiver_ids,json=receiverIds,proto3" json:"receiver_ids,omitempty"`
	// Select notifications with these statuses.
	// An empty list is interpreted as "all".
	Status []NotificationStatus `protobuf:"varint,2,rep,packed,name=status,proto3,enum=ttn.lorawan.v3.NotificationStatus" json:"status,omitempty"`
	// Limit the number of results per page.
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// Page number for pagination. 0 is interpreted as 1.
	Page                 uint32   `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListNotificationsRequest) Reset()         { *m = ListNotificationsRequest{} }
func (m *ListNotificationsRequest) String() string { return proto.CompactTextString(m) }
func (*ListNotificationsRequest) ProtoMessage()    {}
func (*ListNotificationsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62db76852b39e787, []int{3}
}
func (m *ListNotificationsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNotificationsRequest.Unmarshal(m, b)
}
func (m *ListNotificationsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNotificationsRequest.Marshal(b, m, deterministic)
}
func (m *ListNotificationsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNotificationsRequest.Merge(m, src)
}
func (m *ListNotificationsRequest) XXX_Size() int {
	return xxx_messageInfo_ListNotificationsRequest.Size(m)
}
func (m *ListNotificationsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNotificationsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListNotificationsRequest proto.InternalMessageInfo

func (m *ListNotificationsRequest) GetReceiverIds() *UserIdentifiers {
	if m != nil {
		return m.ReceiverIds
	}
	return nil
}

func (m *ListNotificationsRequest) GetStatus() []NotificationStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListNotificationsRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListNotificationsRequest) GetPage() uint32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ListNotificationsResponse struct {
	Notifications        []*Notification `protobuf:"bytes,1,rep,name=notifications,proto3" json:"notifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListNotificationsResponse) Reset()         { *m = ListNotificationsResponse{} }
func (m *ListNotificationsResponse) String() string { return proto.CompactTextString(m) }
func (*ListNotificationsResponse) ProtoMessage()    {}
func (*ListNotificationsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_62db76852b39e787, []int{4}
}
func (m *ListNotificationsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListNotificationsResponse.Unmarshal(m, b)
}
func (m *ListNotificationsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListNotificationsResponse.Marshal(b, m, deterministic)
}
func (m *ListNotificationsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListNotificationsResponse.Merge(m, src)
}
func (m *ListNotificationsResponse) XXX_Size() int {
	return xxx_messageInfo_ListNotificationsResponse.Size(m)
}
func (m *ListNotificationsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListNotificationsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListNotificationsResponse proto.InternalMessageInfo

func (m *ListNotificationsResponse) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

type UpdateNotificationStatusRequest struct {
	// The IDs of the receiving user.
	ReceiverIds *UserIdentifiers `protobuf:"bytes,1,opt,name=receiver_ids,json=receiverIds,proto3" json:"receiver_ids,omitempty"`
	// The IDs of the notifications to update the status of.
	Ids []string `protobuf:"bytes,2,rep,name=ids,proto3" json:"ids,omitempty"`
	// The status to set on the notifications.
	Status               NotificationStatus `protobuf:"varint,3,opt,name=status,proto3,enum=ttn.lorawan.v3.NotificationStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateNotificationStatusRequest) Reset()         { *m = UpdateNotificationStatusRequest{} }
func (m *UpdateNotificationStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateNotificationStatusRequest) ProtoMessage()    {}
func (*UpdateNotificationStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_62db76852b39e787, []int{5}
}
func (m *UpdateNotificationStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNotificationStatusRequest.Unmarshal(m, b)
}
func (m *UpdateNotificationStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNotificationStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateNotificationStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNotificationStatusRequest.Merge(m, src)
}
func (m *UpdateNotificationStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateNotificationStatusRequest.Size(m)
}
func (m *UpdateNotificationStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNotificationStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNotificationStatusRequest proto.InternalMessageInfo

func (m *UpdateNotificationStatusRequest) GetReceiverIds() *UserIdentifiers {
	if m != nil {
		return m.ReceiverIds
	}
	return nil
}

func (m *UpdateNotificationStatusRequest) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

func (m *UpdateNotificationStatusRequest) GetStatus() NotificationStatus {
	if m != nil {
		return m.Status
	}
	return NotificationStatus_NOTIFICATION_STATUS_UNSEEN
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.NotificationReceiver", NotificationReceiver_name, NotificationReceiver_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.NotificationReceiver", NotificationReceiver_name, NotificationReceiver_value)
	proto.RegisterEnum("ttn.lorawan.v3.NotificationStatus", NotificationStatus_name, NotificationStatus_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.NotificationStatus", NotificationStatus_name, NotificationStatus_value)
	proto.RegisterType((*Notification)(nil), "ttn.lorawan.v3.Notification")
	golang_proto.RegisterType((*Notification)(nil), "ttn.lorawan.v3.Notification")
	proto.RegisterType((*CreateNotificationRequest)(nil), "ttn.lorawan.v3.CreateNotificationRequest")
	golang_proto.RegisterType((*CreateNotificationRequest)(nil), "ttn.lorawan.v3.CreateNotificationRequest")
	proto.RegisterType((*CreateNotificationResponse)(nil), "ttn.lorawan.v3.CreateNotificationResponse")
	golang_proto.RegisterType((*CreateNotificationResponse)(nil), "ttn.lorawan.v3.CreateNotificationResponse")
	proto.RegisterType((*ListNotificationsRequest)(nil), "ttn.lorawan.v3.ListNotificationsRequest")
	golang_proto.RegisterType((*ListNotificationsRequest)(nil), "ttn.lorawan.v3.ListNotificationsRequest")
	proto.RegisterType((*ListNotificationsResponse)(nil), "ttn.lorawan.v3.ListNotificationsResponse")
	golang_proto.RegisterType((*ListNotificationsResponse)(nil), "ttn.lorawan.v3.ListNotificationsResponse")
	proto.RegisterType((*UpdateNotificationStatusRequest)(nil), "ttn.lorawan.v3.UpdateNotificationStatusRequest")
	golang_proto.RegisterType((*UpdateNotificationStatusRequest)(nil), "ttn.lorawan.v3.UpdateNotificationStatusRequest")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/notification_service.proto", fileDescriptor_62db76852b39e787)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/notification_service.proto", fileDescriptor_62db76852b39e787)
}

var fileDescriptor_62db76852b39e787 = []byte{
	// 1096 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0x4d, 0x73, 0xdb, 0x44,
	0x18, 0x66, 0xfd, 0x15, 0x7b, 0x93, 0xb8, 0xce, 0x36, 0x80, 0xec, 0xa6, 0x89, 0x6b, 0x3a, 0xc5,
	0x75, 0x62, 0xa9, 0xe3, 0x4c, 0x3b, 0x43, 0x0f, 0x0c, 0x92, 0xe3, 0x4e, 0xdd, 0x06, 0x7b, 0x46,
	0x91, 0xc3, 0x4c, 0x2f, 0x1a, 0xc5, 0xda, 0x3a, 0x4b, 0x6c, 0x49, 0x68, 0xd7, 0x2e, 0x1e, 0x86,
	0x4b, 0x8f, 0x5c, 0x18, 0xc2, 0x85, 0x0b, 0x7f, 0x20, 0xbf, 0x82, 0xff, 0xc0, 0x99, 0xe1, 0x00,
	0x87, 0xd2, 0x1f, 0xc0, 0x21, 0x27, 0x46, 0x2b, 0x39, 0x91, 0x6d, 0x35, 0x75, 0x0e, 0xdc, 0x56,
	0x7a, 0x9f, 0xf7, 0xf3, 0x79, 0xde, 0x9d, 0x85, 0x3b, 0x7d, 0xdb, 0x35, 0x5e, 0x19, 0x56, 0x95,
	0x32, 0xa3, 0x7b, 0x22, 0x19, 0x0e, 0x91, 0x2c, 0x9b, 0x91, 0x97, 0xa4, 0x6b, 0x30, 0x62, 0x5b,
	0x3a, 0xc5, 0xee, 0x88, 0x74, 0xb1, 0xe8, 0xb8, 0x36, 0xb3, 0x51, 0x96, 0x31, 0x4b, 0x0c, 0x3c,
	0xc4, 0xd1, 0x6e, 0x41, 0xee, 0x11, 0x76, 0x3c, 0x3c, 0x12, 0xbb, 0xf6, 0x40, 0xc2, 0xd6, 0xc8,
	0x1e, 0x3b, 0xae, 0xfd, 0xed, 0x58, 0xe2, 0xe0, 0x6e, 0xb5, 0x87, 0xad, 0xea, 0xc8, 0xe8, 0x13,
	0xd3, 0x60, 0x58, 0x9a, 0x3b, 0xf8, 0x21, 0x0b, 0xd5, 0x50, 0x88, 0x9e, 0xdd, 0xb3, 0x7d, 0xe7,
	0xa3, 0xe1, 0x4b, 0xfe, 0xc5, 0x3f, 0xf8, 0x29, 0x80, 0xd7, 0x43, 0x70, 0xed, 0x18, 0x6b, 0xc7,
	0xc4, 0xea, 0xd1, 0xa6, 0x65, 0x0e, 0x29, 0x73, 0x09, 0xa6, 0xe1, 0xd4, 0x3d, 0xbb, 0xfa, 0x35,
	0xb5, 0x2d, 0xc9, 0xb0, 0x2c, 0x9b, 0xf1, 0x6e, 0x68, 0x10, 0x64, 0xa3, 0x67, 0xdb, 0xbd, 0x3e,
	0xe6, 0xdd, 0xce, 0x5b, 0xf3, 0x81, 0xf5, 0xa2, 0x10, 0xc3, 0x1a, 0x07, 0xa6, 0x5b, 0xb3, 0x26,
	0x3c, 0x70, 0xd8, 0xc4, 0xb8, 0x35, 0x6b, 0x64, 0x64, 0x80, 0x29, 0x33, 0x06, 0x4e, 0x00, 0xf8,
	0x64, 0x7e, 0xd6, 0xc4, 0xc4, 0x96, 0x37, 0x6d, 0xec, 0x06, 0xd9, 0x4b, 0x67, 0x09, 0xb8, 0xd2,
	0x0a, 0x31, 0x80, 0xb2, 0x30, 0x46, 0x4c, 0x01, 0x14, 0x41, 0x39, 0xa3, 0xc6, 0x88, 0x89, 0x3e,
	0x83, 0xb0, 0xeb, 0x62, 0x83, 0x61, 0x53, 0x37, 0x98, 0x10, 0x2b, 0x82, 0xf2, 0x72, 0xad, 0x20,
	0xfa, 0xb9, 0xc5, 0x49, 0x6e, 0x51, 0x9b, 0xe4, 0x56, 0x33, 0x01, 0x5a, 0x66, 0xe8, 0x0b, 0x08,
	0xbd, 0x74, 0x6c, 0xac, 0x13, 0x93, 0x0a, 0x71, 0xee, 0x7a, 0x47, 0x9c, 0xe6, 0x54, 0x6c, 0x70,
	0x44, 0xf3, 0xb2, 0x30, 0x35, 0x83, 0x83, 0x5f, 0x14, 0x6d, 0xc3, 0xb5, 0x29, 0x79, 0xb0, 0xb1,
	0x83, 0x85, 0x04, 0xaf, 0x2d, 0x17, 0x36, 0x68, 0x63, 0x07, 0xa3, 0x32, 0x4c, 0x98, 0x06, 0x33,
	0x84, 0x24, 0x4f, 0xb4, 0x3e, 0x57, 0xa3, 0x6c, 0x8d, 0x55, 0x8e, 0x40, 0x9f, 0x43, 0x48, 0xb1,
	0x65, 0x62, 0x97, 0x17, 0x96, 0xe2, 0xf8, 0xad, 0xd9, 0xc2, 0x3a, 0x14, 0xbb, 0x53, 0x65, 0xf9,
	0x2e, 0x5e, 0x59, 0x0a, 0xcc, 0xb8, 0xb8, 0x8b, 0xc9, 0x08, 0xbb, 0x54, 0x48, 0x17, 0xe3, 0xe5,
	0x6c, 0xed, 0xee, 0xac, 0x7b, 0x78, 0xa8, 0x6a, 0x00, 0x56, 0x2f, 0xdd, 0xd0, 0x3a, 0x4c, 0xe2,
	0x81, 0x41, 0xfa, 0x42, 0xa6, 0x08, 0xca, 0x69, 0xd5, 0xff, 0x40, 0x8f, 0x61, 0x8a, 0x32, 0x83,
	0x0d, 0xa9, 0x00, 0x8b, 0xa0, 0x9c, 0xad, 0x95, 0xae, 0x0a, 0x7b, 0xc0, 0x91, 0x6a, 0xe0, 0x81,
	0x9e, 0xc0, 0x35, 0xff, 0xa4, 0x0f, 0x1d, 0x73, 0x42, 0xd8, 0xf2, 0x7b, 0x09, 0xbb, 0xe1, 0x3b,
	0x75, 0x7c, 0x1f, 0x99, 0x3d, 0x4b, 0xa4, 0x97, 0x72, 0x69, 0x15, 0xe1, 0x11, 0xb6, 0x98, 0xee,
	0x10, 0x07, 0xf7, 0x89, 0x85, 0xbd, 0x49, 0x95, 0xfe, 0x8d, 0xc1, 0x7c, 0x9d, 0xd3, 0x3b, 0xdd,
	0xdd, 0x37, 0x43, 0x4c, 0x19, 0x7a, 0x36, 0x45, 0x37, 0x58, 0x90, 0x6e, 0x25, 0x7d, 0xae, 0x24,
	0x7f, 0x00, 0xb1, 0x1c, 0x08, 0x13, 0xff, 0x28, 0x8a, 0x78, 0x4f, 0x7c, 0x19, 0x25, 0x73, 0xae,
	0xa4, 0xdc, 0x44, 0x0e, 0x08, 0xe6, 0x15, 0x1a, 0x88, 0x5f, 0x53, 0x03, 0x89, 0x6b, 0x6b, 0xa0,
	0x13, 0xd6, 0x40, 0x72, 0x71, 0x0d, 0x28, 0x6b, 0xe7, 0x4a, 0xf6, 0x14, 0x2c, 0xa7, 0x81, 0x00,
	0x4a, 0xc9, 0xd7, 0x7e, 0xe3, 0x11, 0xb2, 0x48, 0x85, 0x64, 0x51, 0xda, 0x81, 0x85, 0xa8, 0xb9,
	0x53, 0xc7, 0xb6, 0x28, 0x9e, 0x5d, 0xd9, 0xd2, 0x3f, 0x00, 0x0a, 0xfb, 0x84, 0xb2, 0x30, 0x98,
	0x4e, 0x58, 0xda, 0x87, 0x2b, 0x93, 0x6c, 0x21, 0x9e, 0xde, 0xd7, 0x79, 0x88, 0xa5, 0xe5, 0x89,
	0xbb, 0x37, 0x85, 0xe7, 0x17, 0x7a, 0x8d, 0xf1, 0x11, 0x2c, 0xa0, 0x57, 0xe5, 0xc6, 0xb9, 0xb2,
	0x72, 0x0a, 0x32, 0x97, 0xed, 0x4f, 0x04, 0xbc, 0x09, 0x93, 0x7d, 0x32, 0x20, 0x8c, 0xb3, 0xb7,
	0xca, 0x53, 0x56, 0xe2, 0xc2, 0x9b, 0x25, 0xd5, 0xff, 0x8d, 0x10, 0x4c, 0x38, 0x46, 0xcf, 0xbf,
	0x00, 0x56, 0x55, 0x7e, 0x2e, 0xe9, 0x30, 0x1f, 0xd1, 0x6a, 0x30, 0x18, 0x05, 0xae, 0x86, 0x15,
	0xe2, 0x35, 0x1b, 0x2f, 0x2f, 0xd7, 0x36, 0xae, 0xe4, 0x69, 0xda, 0xa5, 0xf4, 0x07, 0x80, 0x5b,
	0xfe, 0x6e, 0x44, 0xac, 0xde, 0xff, 0x32, 0xd3, 0x7b, 0x30, 0xee, 0x05, 0xf1, 0x06, 0x9a, 0x51,
	0xd6, 0xcf, 0x95, 0xb5, 0x53, 0x90, 0x4d, 0x83, 0xdc, 0x9b, 0x25, 0x6f, 0x62, 0x6e, 0xfc, 0x17,
	0x70, 0x57, 0xf5, 0x00, 0x68, 0xef, 0x62, 0xf6, 0xf1, 0x45, 0xef, 0x0a, 0x9e, 0x72, 0x6a, 0xe8,
	0x95, 0x3f, 0x01, 0x5c, 0x8f, 0xd2, 0x29, 0xba, 0x03, 0x6f, 0xb7, 0xda, 0x5a, 0xf3, 0x49, 0xb3,
	0x2e, 0x6b, 0xcd, 0x76, 0x4b, 0x57, 0x1b, 0xf5, 0x46, 0xf3, 0xb0, 0xa1, 0xea, 0x9d, 0xd6, 0xf3,
	0x56, 0xfb, 0xab, 0x56, 0xee, 0x03, 0x74, 0x0f, 0x96, 0xa2, 0x21, 0xf5, 0xf6, 0xfe, 0xbe, 0xac,
	0xb4, 0x55, 0x59, 0x6b, 0xab, 0x39, 0x80, 0x1e, 0xc0, 0x9d, 0x68, 0x9c, 0xbc, 0xf7, 0x65, 0xb3,
	0xd5, 0x3c, 0xd0, 0x54, 0x59, 0x6b, 0x1e, 0x36, 0xf4, 0x7a, 0xbb, 0xa5, 0xc9, 0x75, 0x2d, 0x17,
	0x47, 0xdb, 0xf0, 0xd3, 0x68, 0x0f, 0xad, 0x51, 0x7f, 0xda, 0x6a, 0xd6, 0xe5, 0xfd, 0x0b, 0x70,
	0xa2, 0x70, 0xfb, 0xed, 0x59, 0x3e, 0x2f, 0x80, 0xca, 0x87, 0x91, 0x2e, 0x95, 0x9f, 0x00, 0x44,
	0xf3, 0xa3, 0x40, 0x9b, 0xb0, 0x30, 0x85, 0x3f, 0xd0, 0x64, 0xad, 0x73, 0xa0, 0x77, 0x5a, 0x07,
	0x8d, 0x86, 0xd7, 0xdc, 0x06, 0x14, 0xa2, 0xec, 0xdc, 0x0a, 0x50, 0x11, 0x6e, 0x44, 0x59, 0x65,
	0xb5, 0xfe, 0xb4, 0x79, 0xd8, 0xd8, 0xcb, 0xc5, 0x0a, 0xb7, 0xde, 0x9e, 0xe5, 0x3f, 0x16, 0x40,
	0xe5, 0x66, 0x04, 0xae, 0xf6, 0x6b, 0x1c, 0xde, 0x9c, 0xaa, 0xc9, 0x7f, 0xf7, 0x20, 0x1d, 0xa6,
	0xfc, 0x45, 0x47, 0xf7, 0x67, 0xd9, 0x7c, 0xe7, 0xc5, 0x5b, 0xa8, 0x2c, 0x02, 0x0d, 0x56, 0xe2,
	0x14, 0xc0, 0x84, 0xb7, 0x30, 0xa8, 0x3c, 0xeb, 0xf4, 0xae, 0x1b, 0xa3, 0x70, 0x7f, 0x01, 0xa4,
	0x1f, 0xbd, 0xf4, 0xf0, 0xf5, 0xef, 0x7f, 0xff, 0x1c, 0x93, 0x5e, 0x54, 0xd1, 0xb6, 0x34, 0xa4,
	0xd8, 0xa5, 0xd2, 0x77, 0xe1, 0xc5, 0x10, 0xbd, 0x7f, 0x3a, 0x31, 0xbf, 0x9f, 0x7a, 0xf4, 0x51,
	0xf4, 0x23, 0x80, 0x2b, 0xfe, 0x8e, 0x05, 0xdc, 0x48, 0x73, 0xab, 0x73, 0xf5, 0x06, 0x16, 0x3e,
	0x9a, 0xbb, 0xe9, 0x1b, 0xde, 0x53, 0xa9, 0xf4, 0x88, 0x17, 0xf4, 0xa0, 0x76, 0x9d, 0x72, 0x1e,
	0x83, 0x8a, 0xf2, 0xf0, 0xb7, 0xbf, 0x36, 0xc1, 0x0b, 0xa9, 0x67, 0x8b, 0xec, 0x18, 0x33, 0xfe,
	0xea, 0x13, 0x2d, 0xcc, 0x5e, 0xd9, 0xee, 0x89, 0x34, 0xfd, 0xac, 0x1a, 0xed, 0x4a, 0xce, 0x49,
	0x4f, 0x62, 0xcc, 0x72, 0x8e, 0x8e, 0x52, 0x3c, 0xfd, 0xee, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x88, 0xeb, 0xc7, 0x89, 0xe7, 0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NotificationServiceClient is the client API for NotificationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NotificationServiceClient interface {
	// Create a new notification. Can only be called by internal services using cluster auth.
	Create(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error)
	// List the notifications for a user or an organization.
	// When called with user credentials and empty receiver_ids, this will list
	// notifications for the current user and its organizations.
	List(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error)
	// Batch-update multiple notifications to the same status.
	UpdateStatus(ctx context.Context, in *UpdateNotificationStatusRequest, opts ...grpc.CallOption) (*types.Empty, error)
}

type notificationServiceClient struct {
	cc *grpc.ClientConn
}

func NewNotificationServiceClient(cc *grpc.ClientConn) NotificationServiceClient {
	return &notificationServiceClient{cc}
}

func (c *notificationServiceClient) Create(ctx context.Context, in *CreateNotificationRequest, opts ...grpc.CallOption) (*CreateNotificationResponse, error) {
	out := new(CreateNotificationResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NotificationService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) List(ctx context.Context, in *ListNotificationsRequest, opts ...grpc.CallOption) (*ListNotificationsResponse, error) {
	out := new(ListNotificationsResponse)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NotificationService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationServiceClient) UpdateStatus(ctx context.Context, in *UpdateNotificationStatusRequest, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.NotificationService/UpdateStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationServiceServer is the server API for NotificationService service.
type NotificationServiceServer interface {
	// Create a new notification. Can only be called by internal services using cluster auth.
	Create(context.Context, *CreateNotificationRequest) (*CreateNotificationResponse, error)
	// List the notifications for a user or an organization.
	// When called with user credentials and empty receiver_ids, this will list
	// notifications for the current user and its organizations.
	List(context.Context, *ListNotificationsRequest) (*ListNotificationsResponse, error)
	// Batch-update multiple notifications to the same status.
	UpdateStatus(context.Context, *UpdateNotificationStatusRequest) (*types.Empty, error)
}

// UnimplementedNotificationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNotificationServiceServer struct {
}

func (*UnimplementedNotificationServiceServer) Create(ctx context.Context, req *CreateNotificationRequest) (*CreateNotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedNotificationServiceServer) List(ctx context.Context, req *ListNotificationsRequest) (*ListNotificationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedNotificationServiceServer) UpdateStatus(ctx context.Context, req *UpdateNotificationStatusRequest) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}

func RegisterNotificationServiceServer(s *grpc.Server, srv NotificationServiceServer) {
	s.RegisterService(&_NotificationService_serviceDesc, srv)
}

func _NotificationService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNotificationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NotificationService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).Create(ctx, req.(*CreateNotificationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNotificationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NotificationService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).List(ctx, req.(*ListNotificationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NotificationService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNotificationStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.NotificationService/UpdateStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationServiceServer).UpdateStatus(ctx, req.(*UpdateNotificationStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NotificationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NotificationService",
	HandlerType: (*NotificationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _NotificationService_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _NotificationService_List_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _NotificationService_UpdateStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/notification_service.proto",
}