// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Authorisation.proto

package authorisation

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Operation int32

const (
	Operation_OPERATION_INVALID  Operation = 0
	Operation_OPERATION_READ     Operation = 1
	Operation_OPERATION_TRANSACT Operation = 2
	Operation_OPERATION_APPROVE  Operation = 3
)

var Operation_name = map[int32]string{
	0: "OPERATION_INVALID",
	1: "OPERATION_READ",
	2: "OPERATION_TRANSACT",
	3: "OPERATION_APPROVE",
}

var Operation_value = map[string]int32{
	"OPERATION_INVALID":  0,
	"OPERATION_READ":     1,
	"OPERATION_TRANSACT": 2,
	"OPERATION_APPROVE":  3,
}

func (x Operation) String() string {
	return proto.EnumName(Operation_name, int32(x))
}

func (Operation) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ea208e9d0855f7f8, []int{0}
}

// types for Authorisation
type AccountPermissions struct {
	Permissions          []int64  `protobuf:"varint,2,rep,packed,name=permissions,proto3" json:"permissions,omitempty"`
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AccountPermissions) Reset()         { *m = AccountPermissions{} }
func (m *AccountPermissions) String() string { return proto.CompactTextString(m) }
func (*AccountPermissions) ProtoMessage()    {}
func (*AccountPermissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea208e9d0855f7f8, []int{0}
}

func (m *AccountPermissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountPermissions.Unmarshal(m, b)
}
func (m *AccountPermissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountPermissions.Marshal(b, m, deterministic)
}
func (m *AccountPermissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountPermissions.Merge(m, src)
}
func (m *AccountPermissions) XXX_Size() int {
	return xxx_messageInfo_AccountPermissions.Size(m)
}
func (m *AccountPermissions) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountPermissions.DiscardUnknown(m)
}

var xxx_messageInfo_AccountPermissions proto.InternalMessageInfo

func (m *AccountPermissions) GetPermissions() []int64 {
	if m != nil {
		return m.Permissions
	}
	return nil
}

func (m *AccountPermissions) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

type ListMeRequest struct {
	Operation            Operation `protobuf:"varint,2,opt,name=operation,proto3,enum=authorisation.Operation" json:"operation,omitempty"`
	UserId               string    `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListMeRequest) Reset()         { *m = ListMeRequest{} }
func (m *ListMeRequest) String() string { return proto.CompactTextString(m) }
func (*ListMeRequest) ProtoMessage()    {}
func (*ListMeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea208e9d0855f7f8, []int{1}
}

func (m *ListMeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeRequest.Unmarshal(m, b)
}
func (m *ListMeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeRequest.Marshal(b, m, deterministic)
}
func (m *ListMeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeRequest.Merge(m, src)
}
func (m *ListMeRequest) XXX_Size() int {
	return xxx_messageInfo_ListMeRequest.Size(m)
}
func (m *ListMeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeRequest proto.InternalMessageInfo

func (m *ListMeRequest) GetOperation() Operation {
	if m != nil {
		return m.Operation
	}
	return Operation_OPERATION_INVALID
}

func (m *ListMeRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type ListMeResponse struct {
	Permissions          []*AccountPermissions `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListMeResponse) Reset()         { *m = ListMeResponse{} }
func (m *ListMeResponse) String() string { return proto.CompactTextString(m) }
func (*ListMeResponse) ProtoMessage()    {}
func (*ListMeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea208e9d0855f7f8, []int{2}
}

func (m *ListMeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListMeResponse.Unmarshal(m, b)
}
func (m *ListMeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListMeResponse.Marshal(b, m, deterministic)
}
func (m *ListMeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListMeResponse.Merge(m, src)
}
func (m *ListMeResponse) XXX_Size() int {
	return xxx_messageInfo_ListMeResponse.Size(m)
}
func (m *ListMeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListMeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListMeResponse proto.InternalMessageInfo

func (m *ListMeResponse) GetPermissions() []*AccountPermissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type MayIRequest struct {
	Operation            *OpCode              `protobuf:"bytes,2,opt,name=operation,proto3" json:"operation,omitempty"`
	CurrentTime          *timestamp.Timestamp `protobuf:"bytes,4,opt,name=current_time,json=currentTime,proto3" json:"current_time,omitempty"`
	AccountId            string               `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	UserId               string               `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *MayIRequest) Reset()         { *m = MayIRequest{} }
func (m *MayIRequest) String() string { return proto.CompactTextString(m) }
func (*MayIRequest) ProtoMessage()    {}
func (*MayIRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea208e9d0855f7f8, []int{3}
}

func (m *MayIRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MayIRequest.Unmarshal(m, b)
}
func (m *MayIRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MayIRequest.Marshal(b, m, deterministic)
}
func (m *MayIRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MayIRequest.Merge(m, src)
}
func (m *MayIRequest) XXX_Size() int {
	return xxx_messageInfo_MayIRequest.Size(m)
}
func (m *MayIRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MayIRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MayIRequest proto.InternalMessageInfo

func (m *MayIRequest) GetOperation() *OpCode {
	if m != nil {
		return m.Operation
	}
	return nil
}

func (m *MayIRequest) GetCurrentTime() *timestamp.Timestamp {
	if m != nil {
		return m.CurrentTime
	}
	return nil
}

func (m *MayIRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *MayIRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type MayIResponse struct {
	YouMay               bool     `protobuf:"varint,1,opt,name=you_may,json=youMay,proto3" json:"you_may,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MayIResponse) Reset()         { *m = MayIResponse{} }
func (m *MayIResponse) String() string { return proto.CompactTextString(m) }
func (*MayIResponse) ProtoMessage()    {}
func (*MayIResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea208e9d0855f7f8, []int{4}
}

func (m *MayIResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MayIResponse.Unmarshal(m, b)
}
func (m *MayIResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MayIResponse.Marshal(b, m, deterministic)
}
func (m *MayIResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MayIResponse.Merge(m, src)
}
func (m *MayIResponse) XXX_Size() int {
	return xxx_messageInfo_MayIResponse.Size(m)
}
func (m *MayIResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MayIResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MayIResponse proto.InternalMessageInfo

func (m *MayIResponse) GetYouMay() bool {
	if m != nil {
		return m.YouMay
	}
	return false
}

func (m *MayIResponse) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

type OpCode struct {
	AliasOpCode          Operation `protobuf:"varint,1,opt,name=aliasOpCode,proto3,enum=authorisation.Operation" json:"aliasOpCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *OpCode) Reset()         { *m = OpCode{} }
func (m *OpCode) String() string { return proto.CompactTextString(m) }
func (*OpCode) ProtoMessage()    {}
func (*OpCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea208e9d0855f7f8, []int{5}
}

func (m *OpCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OpCode.Unmarshal(m, b)
}
func (m *OpCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OpCode.Marshal(b, m, deterministic)
}
func (m *OpCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OpCode.Merge(m, src)
}
func (m *OpCode) XXX_Size() int {
	return xxx_messageInfo_OpCode.Size(m)
}
func (m *OpCode) XXX_DiscardUnknown() {
	xxx_messageInfo_OpCode.DiscardUnknown(m)
}

var xxx_messageInfo_OpCode proto.InternalMessageInfo

func (m *OpCode) GetAliasOpCode() Operation {
	if m != nil {
		return m.AliasOpCode
	}
	return Operation_OPERATION_INVALID
}

func init() {
	proto.RegisterEnum("authorisation.Operation", Operation_name, Operation_value)
	proto.RegisterType((*AccountPermissions)(nil), "authorisation.AccountPermissions")
	proto.RegisterType((*ListMeRequest)(nil), "authorisation.ListMeRequest")
	proto.RegisterType((*ListMeResponse)(nil), "authorisation.ListMeResponse")
	proto.RegisterType((*MayIRequest)(nil), "authorisation.MayIRequest")
	proto.RegisterType((*MayIResponse)(nil), "authorisation.MayIResponse")
	proto.RegisterType((*OpCode)(nil), "authorisation.OpCode")
}

func init() { proto.RegisterFile("Authorisation.proto", fileDescriptor_ea208e9d0855f7f8) }

var fileDescriptor_ea208e9d0855f7f8 = []byte{
	// 466 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xad, 0xe3, 0xca, 0xc5, 0xe3, 0x26, 0x84, 0x41, 0x6d, 0x2d, 0x43, 0x85, 0xf1, 0x29, 0xe2,
	0xe0, 0x4a, 0xa9, 0xc4, 0x01, 0x09, 0x21, 0x93, 0x44, 0xc8, 0x52, 0xf3, 0xa1, 0x25, 0xed, 0x81,
	0x4b, 0xd8, 0x26, 0x4b, 0xb1, 0x54, 0x7b, 0x8d, 0x77, 0x7d, 0xf0, 0x6f, 0xe1, 0x97, 0xf0, 0xef,
	0x50, 0xfc, 0x81, 0x3f, 0x4a, 0x7b, 0x9c, 0x37, 0x6f, 0x76, 0xde, 0xbc, 0x99, 0x85, 0x97, 0x5e,
	0x2a, 0x7f, 0xf2, 0x24, 0x10, 0x54, 0x06, 0x3c, 0x72, 0xe3, 0x84, 0x4b, 0x8e, 0x7d, 0xda, 0x04,
	0xad, 0x37, 0x77, 0x9c, 0xdf, 0xdd, 0xb3, 0x8b, 0x3c, 0x79, 0x9b, 0xfe, 0xb8, 0x90, 0x41, 0xc8,
	0x84, 0xa4, 0x61, 0x5c, 0xf0, 0x9d, 0x6b, 0x40, 0x6f, 0xbb, 0xe5, 0x69, 0x24, 0x57, 0x2c, 0x09,
	0x03, 0x21, 0x02, 0x1e, 0x09, 0xb4, 0xc1, 0x88, 0xeb, 0xd0, 0xec, 0xd9, 0xea, 0x48, 0x25, 0x4d,
	0x08, 0xcf, 0x01, 0x68, 0x51, 0xb7, 0x09, 0x76, 0xa6, 0x62, 0x2b, 0x23, 0x9d, 0xe8, 0x25, 0xe2,
	0xef, 0x9c, 0xef, 0xd0, 0xbf, 0x0a, 0x84, 0x9c, 0x33, 0xc2, 0x7e, 0xa5, 0x4c, 0x48, 0x7c, 0x0f,
	0x3a, 0x8f, 0x59, 0x92, 0xab, 0x32, 0x7b, 0xb6, 0x32, 0x1a, 0x8c, 0x4d, 0xb7, 0xa5, 0xd5, 0x5d,
	0x56, 0x79, 0x52, 0x53, 0xf1, 0x0c, 0x8e, 0x52, 0xc1, 0x92, 0xba, 0x89, 0xb6, 0x0f, 0xfd, 0x9d,
	0x73, 0x0d, 0x83, 0xaa, 0x83, 0x88, 0x79, 0x24, 0x18, 0x4e, 0xda, 0xa2, 0x15, 0x5b, 0x1d, 0x19,
	0xe3, 0xb7, 0x9d, 0x26, 0x0f, 0x87, 0x6d, 0xcd, 0xe5, 0xfc, 0x51, 0xc0, 0x98, 0xd3, 0xcc, 0xaf,
	0x74, 0x5f, 0x76, 0x75, 0x1b, 0xe3, 0x93, 0x07, 0xba, 0x27, 0x7c, 0xc7, 0x9a, 0xa2, 0x3f, 0xc2,
	0xf1, 0x36, 0x4d, 0x12, 0x16, 0xc9, 0xcd, 0xde, 0x6f, 0xf3, 0x30, 0xaf, 0xb3, 0xdc, 0x62, 0x19,
	0x6e, 0xb5, 0x0c, 0x77, 0x5d, 0x2d, 0x83, 0x18, 0x25, 0x7f, 0x8f, 0x74, 0xbc, 0x55, 0x3b, 0xde,
	0x3e, 0x6e, 0xc9, 0x27, 0x38, 0x2e, 0xa4, 0x97, 0x86, 0x9c, 0xc1, 0x51, 0xc6, 0xd3, 0x4d, 0x48,
	0xb3, 0x9c, 0xf8, 0x8c, 0x68, 0x19, 0x4f, 0xe7, 0x34, 0xc3, 0x53, 0xd0, 0x12, 0x46, 0x45, 0x39,
	0x91, 0x4e, 0xca, 0xc8, 0x99, 0x82, 0x56, 0x0c, 0x83, 0x1f, 0xc0, 0xa0, 0xf7, 0x01, 0x15, 0x45,
	0x98, 0x97, 0x3f, 0xb5, 0xb0, 0x26, 0xf9, 0x1d, 0x03, 0xfd, 0x5f, 0x06, 0x4f, 0xe0, 0xc5, 0x72,
	0x35, 0x23, 0xde, 0xda, 0x5f, 0x2e, 0x36, 0xfe, 0xe2, 0xc6, 0xbb, 0xf2, 0xa7, 0xc3, 0x03, 0x44,
	0x18, 0xd4, 0x30, 0x99, 0x79, 0xd3, 0xa1, 0x82, 0xa7, 0x80, 0x35, 0xb6, 0x26, 0xde, 0xe2, 0xab,
	0x37, 0x59, 0x0f, 0x7b, 0xed, 0x27, 0xbc, 0xd5, 0x8a, 0x2c, 0x6f, 0x66, 0x43, 0x75, 0xfc, 0x5b,
	0x81, 0x7e, 0xeb, 0x07, 0xe0, 0x17, 0xd0, 0xf6, 0x27, 0x11, 0x32, 0x7c, 0xdd, 0x51, 0xda, 0xba,
	0x45, 0xeb, 0xfc, 0x91, 0x6c, 0x61, 0x9b, 0x73, 0x80, 0x1e, 0x1c, 0xce, 0x69, 0x16, 0xa0, 0xd5,
	0x21, 0x36, 0x0e, 0xc3, 0x7a, 0xf5, 0xdf, 0x5c, 0xf5, 0xc4, 0xe7, 0xe7, 0xdf, 0xda, 0x3f, 0xf1,
	0x56, 0xcb, 0xb7, 0x7e, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x68, 0x12, 0x07, 0x12, 0xb6, 0x03,
	0x00, 0x00,
}
