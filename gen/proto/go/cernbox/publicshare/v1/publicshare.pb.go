// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cernbox/publicshare/v1/publicshare.proto

package publicsharev1pb

import (
	context "context"
	fmt "fmt"
	rpc "github.com/cs3/cs3apis/gen/proto/go/cernbox/rpc"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreatePublicShareRequest struct {
	Filename             string     `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	Permission           Permission `protobuf:"varint,2,opt,name=permission,proto3,enum=cernbox.publicsharev1.Permission" json:"permission,omitempty"`
	Password             string     `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Expiration           uint64     `protobuf:"varint,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *CreatePublicShareRequest) Reset()         { *m = CreatePublicShareRequest{} }
func (m *CreatePublicShareRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePublicShareRequest) ProtoMessage()    {}
func (*CreatePublicShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{0}
}

func (m *CreatePublicShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePublicShareRequest.Unmarshal(m, b)
}
func (m *CreatePublicShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePublicShareRequest.Marshal(b, m, deterministic)
}
func (m *CreatePublicShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePublicShareRequest.Merge(m, src)
}
func (m *CreatePublicShareRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePublicShareRequest.Size(m)
}
func (m *CreatePublicShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePublicShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePublicShareRequest proto.InternalMessageInfo

func (m *CreatePublicShareRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *CreatePublicShareRequest) GetPermission() Permission {
	if m != nil {
		return m.Permission
	}
	return Permission_PERMISSION_INVALID
}

func (m *CreatePublicShareRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CreatePublicShareRequest) GetExpiration() uint64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type CreatePublicShareResponse struct {
	Status               *rpc.Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PublicShare          *PublicShare `protobuf:"bytes,2,opt,name=public_share,json=publicShare,proto3" json:"public_share,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CreatePublicShareResponse) Reset()         { *m = CreatePublicShareResponse{} }
func (m *CreatePublicShareResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePublicShareResponse) ProtoMessage()    {}
func (*CreatePublicShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{1}
}

func (m *CreatePublicShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePublicShareResponse.Unmarshal(m, b)
}
func (m *CreatePublicShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePublicShareResponse.Marshal(b, m, deterministic)
}
func (m *CreatePublicShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePublicShareResponse.Merge(m, src)
}
func (m *CreatePublicShareResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePublicShareResponse.Size(m)
}
func (m *CreatePublicShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePublicShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePublicShareResponse proto.InternalMessageInfo

func (m *CreatePublicShareResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CreatePublicShareResponse) GetPublicShare() *PublicShare {
	if m != nil {
		return m.PublicShare
	}
	return nil
}

type UpdatePublicShareRequest struct {
	PublicShareId        string        `protobuf:"bytes,1,opt,name=public_share_id,json=publicShareId,proto3" json:"public_share_id,omitempty"`
	Policy               *UpdatePolicy `protobuf:"bytes,2,opt,name=policy,proto3" json:"policy,omitempty"`
	Password             string        `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Expiration           string        `protobuf:"bytes,4,opt,name=expiration,proto3" json:"expiration,omitempty"`
	Permission           Permission    `protobuf:"varint,5,opt,name=permission,proto3,enum=cernbox.publicsharev1.Permission" json:"permission,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *UpdatePublicShareRequest) Reset()         { *m = UpdatePublicShareRequest{} }
func (m *UpdatePublicShareRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePublicShareRequest) ProtoMessage()    {}
func (*UpdatePublicShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{2}
}

func (m *UpdatePublicShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePublicShareRequest.Unmarshal(m, b)
}
func (m *UpdatePublicShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePublicShareRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePublicShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePublicShareRequest.Merge(m, src)
}
func (m *UpdatePublicShareRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePublicShareRequest.Size(m)
}
func (m *UpdatePublicShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePublicShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePublicShareRequest proto.InternalMessageInfo

func (m *UpdatePublicShareRequest) GetPublicShareId() string {
	if m != nil {
		return m.PublicShareId
	}
	return ""
}

func (m *UpdatePublicShareRequest) GetPolicy() *UpdatePolicy {
	if m != nil {
		return m.Policy
	}
	return nil
}

func (m *UpdatePublicShareRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UpdatePublicShareRequest) GetExpiration() string {
	if m != nil {
		return m.Expiration
	}
	return ""
}

func (m *UpdatePublicShareRequest) GetPermission() Permission {
	if m != nil {
		return m.Permission
	}
	return Permission_PERMISSION_INVALID
}

type UpdatePublicShareResponse struct {
	Status               *rpc.Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PublicShare          *PublicShare `protobuf:"bytes,2,opt,name=public_share,json=publicShare,proto3" json:"public_share,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *UpdatePublicShareResponse) Reset()         { *m = UpdatePublicShareResponse{} }
func (m *UpdatePublicShareResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePublicShareResponse) ProtoMessage()    {}
func (*UpdatePublicShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{3}
}

func (m *UpdatePublicShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePublicShareResponse.Unmarshal(m, b)
}
func (m *UpdatePublicShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePublicShareResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePublicShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePublicShareResponse.Merge(m, src)
}
func (m *UpdatePublicShareResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePublicShareResponse.Size(m)
}
func (m *UpdatePublicShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePublicShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePublicShareResponse proto.InternalMessageInfo

func (m *UpdatePublicShareResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *UpdatePublicShareResponse) GetPublicShare() *PublicShare {
	if m != nil {
		return m.PublicShare
	}
	return nil
}

type ListPublicSharesRequest struct {
	Filename             string   `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPublicSharesRequest) Reset()         { *m = ListPublicSharesRequest{} }
func (m *ListPublicSharesRequest) String() string { return proto.CompactTextString(m) }
func (*ListPublicSharesRequest) ProtoMessage()    {}
func (*ListPublicSharesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{4}
}

func (m *ListPublicSharesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPublicSharesRequest.Unmarshal(m, b)
}
func (m *ListPublicSharesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPublicSharesRequest.Marshal(b, m, deterministic)
}
func (m *ListPublicSharesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPublicSharesRequest.Merge(m, src)
}
func (m *ListPublicSharesRequest) XXX_Size() int {
	return xxx_messageInfo_ListPublicSharesRequest.Size(m)
}
func (m *ListPublicSharesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPublicSharesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPublicSharesRequest proto.InternalMessageInfo

func (m *ListPublicSharesRequest) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

type ListPublicSharesResponse struct {
	Status               *rpc.Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PublicShare          *PublicShare `protobuf:"bytes,2,opt,name=public_share,json=publicShare,proto3" json:"public_share,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ListPublicSharesResponse) Reset()         { *m = ListPublicSharesResponse{} }
func (m *ListPublicSharesResponse) String() string { return proto.CompactTextString(m) }
func (*ListPublicSharesResponse) ProtoMessage()    {}
func (*ListPublicSharesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{5}
}

func (m *ListPublicSharesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPublicSharesResponse.Unmarshal(m, b)
}
func (m *ListPublicSharesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPublicSharesResponse.Marshal(b, m, deterministic)
}
func (m *ListPublicSharesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPublicSharesResponse.Merge(m, src)
}
func (m *ListPublicSharesResponse) XXX_Size() int {
	return xxx_messageInfo_ListPublicSharesResponse.Size(m)
}
func (m *ListPublicSharesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPublicSharesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPublicSharesResponse proto.InternalMessageInfo

func (m *ListPublicSharesResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *ListPublicSharesResponse) GetPublicShare() *PublicShare {
	if m != nil {
		return m.PublicShare
	}
	return nil
}

type RevokePublicShareRequest struct {
	PublicShareId        string   `protobuf:"bytes,1,opt,name=public_share_id,json=publicShareId,proto3" json:"public_share_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokePublicShareRequest) Reset()         { *m = RevokePublicShareRequest{} }
func (m *RevokePublicShareRequest) String() string { return proto.CompactTextString(m) }
func (*RevokePublicShareRequest) ProtoMessage()    {}
func (*RevokePublicShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{6}
}

func (m *RevokePublicShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokePublicShareRequest.Unmarshal(m, b)
}
func (m *RevokePublicShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokePublicShareRequest.Marshal(b, m, deterministic)
}
func (m *RevokePublicShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokePublicShareRequest.Merge(m, src)
}
func (m *RevokePublicShareRequest) XXX_Size() int {
	return xxx_messageInfo_RevokePublicShareRequest.Size(m)
}
func (m *RevokePublicShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokePublicShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokePublicShareRequest proto.InternalMessageInfo

func (m *RevokePublicShareRequest) GetPublicShareId() string {
	if m != nil {
		return m.PublicShareId
	}
	return ""
}

type RevokePublicShareResponse struct {
	Status               *rpc.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *RevokePublicShareResponse) Reset()         { *m = RevokePublicShareResponse{} }
func (m *RevokePublicShareResponse) String() string { return proto.CompactTextString(m) }
func (*RevokePublicShareResponse) ProtoMessage()    {}
func (*RevokePublicShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{7}
}

func (m *RevokePublicShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokePublicShareResponse.Unmarshal(m, b)
}
func (m *RevokePublicShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokePublicShareResponse.Marshal(b, m, deterministic)
}
func (m *RevokePublicShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokePublicShareResponse.Merge(m, src)
}
func (m *RevokePublicShareResponse) XXX_Size() int {
	return xxx_messageInfo_RevokePublicShareResponse.Size(m)
}
func (m *RevokePublicShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokePublicShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokePublicShareResponse proto.InternalMessageInfo

func (m *RevokePublicShareResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type GetPublicShareRequest struct {
	PublicShareId        string   `protobuf:"bytes,1,opt,name=public_share_id,json=publicShareId,proto3" json:"public_share_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPublicShareRequest) Reset()         { *m = GetPublicShareRequest{} }
func (m *GetPublicShareRequest) String() string { return proto.CompactTextString(m) }
func (*GetPublicShareRequest) ProtoMessage()    {}
func (*GetPublicShareRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{8}
}

func (m *GetPublicShareRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicShareRequest.Unmarshal(m, b)
}
func (m *GetPublicShareRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicShareRequest.Marshal(b, m, deterministic)
}
func (m *GetPublicShareRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicShareRequest.Merge(m, src)
}
func (m *GetPublicShareRequest) XXX_Size() int {
	return xxx_messageInfo_GetPublicShareRequest.Size(m)
}
func (m *GetPublicShareRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicShareRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicShareRequest proto.InternalMessageInfo

func (m *GetPublicShareRequest) GetPublicShareId() string {
	if m != nil {
		return m.PublicShareId
	}
	return ""
}

type GetPublicShareResponse struct {
	Status               *rpc.Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PublicShare          *PublicShare `protobuf:"bytes,2,opt,name=public_share,json=publicShare,proto3" json:"public_share,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetPublicShareResponse) Reset()         { *m = GetPublicShareResponse{} }
func (m *GetPublicShareResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicShareResponse) ProtoMessage()    {}
func (*GetPublicShareResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{9}
}

func (m *GetPublicShareResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicShareResponse.Unmarshal(m, b)
}
func (m *GetPublicShareResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicShareResponse.Marshal(b, m, deterministic)
}
func (m *GetPublicShareResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicShareResponse.Merge(m, src)
}
func (m *GetPublicShareResponse) XXX_Size() int {
	return xxx_messageInfo_GetPublicShareResponse.Size(m)
}
func (m *GetPublicShareResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicShareResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicShareResponse proto.InternalMessageInfo

func (m *GetPublicShareResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetPublicShareResponse) GetPublicShare() *PublicShare {
	if m != nil {
		return m.PublicShare
	}
	return nil
}

type GetPublicShareByTokenRequest struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPublicShareByTokenRequest) Reset()         { *m = GetPublicShareByTokenRequest{} }
func (m *GetPublicShareByTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetPublicShareByTokenRequest) ProtoMessage()    {}
func (*GetPublicShareByTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{10}
}

func (m *GetPublicShareByTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicShareByTokenRequest.Unmarshal(m, b)
}
func (m *GetPublicShareByTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicShareByTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetPublicShareByTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicShareByTokenRequest.Merge(m, src)
}
func (m *GetPublicShareByTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetPublicShareByTokenRequest.Size(m)
}
func (m *GetPublicShareByTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicShareByTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicShareByTokenRequest proto.InternalMessageInfo

func (m *GetPublicShareByTokenRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetPublicShareByTokenResponse struct {
	Status               *rpc.Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	PublicShare          *PublicShare `protobuf:"bytes,2,opt,name=public_share,json=publicShare,proto3" json:"public_share,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *GetPublicShareByTokenResponse) Reset()         { *m = GetPublicShareByTokenResponse{} }
func (m *GetPublicShareByTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetPublicShareByTokenResponse) ProtoMessage()    {}
func (*GetPublicShareByTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_508404bc649b24b2, []int{11}
}

func (m *GetPublicShareByTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPublicShareByTokenResponse.Unmarshal(m, b)
}
func (m *GetPublicShareByTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPublicShareByTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetPublicShareByTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPublicShareByTokenResponse.Merge(m, src)
}
func (m *GetPublicShareByTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetPublicShareByTokenResponse.Size(m)
}
func (m *GetPublicShareByTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPublicShareByTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPublicShareByTokenResponse proto.InternalMessageInfo

func (m *GetPublicShareByTokenResponse) GetStatus() *rpc.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *GetPublicShareByTokenResponse) GetPublicShare() *PublicShare {
	if m != nil {
		return m.PublicShare
	}
	return nil
}

func init() {
	proto.RegisterType((*CreatePublicShareRequest)(nil), "cernbox.publicsharev1.CreatePublicShareRequest")
	proto.RegisterType((*CreatePublicShareResponse)(nil), "cernbox.publicsharev1.CreatePublicShareResponse")
	proto.RegisterType((*UpdatePublicShareRequest)(nil), "cernbox.publicsharev1.UpdatePublicShareRequest")
	proto.RegisterType((*UpdatePublicShareResponse)(nil), "cernbox.publicsharev1.UpdatePublicShareResponse")
	proto.RegisterType((*ListPublicSharesRequest)(nil), "cernbox.publicsharev1.ListPublicSharesRequest")
	proto.RegisterType((*ListPublicSharesResponse)(nil), "cernbox.publicsharev1.ListPublicSharesResponse")
	proto.RegisterType((*RevokePublicShareRequest)(nil), "cernbox.publicsharev1.RevokePublicShareRequest")
	proto.RegisterType((*RevokePublicShareResponse)(nil), "cernbox.publicsharev1.RevokePublicShareResponse")
	proto.RegisterType((*GetPublicShareRequest)(nil), "cernbox.publicsharev1.GetPublicShareRequest")
	proto.RegisterType((*GetPublicShareResponse)(nil), "cernbox.publicsharev1.GetPublicShareResponse")
	proto.RegisterType((*GetPublicShareByTokenRequest)(nil), "cernbox.publicsharev1.GetPublicShareByTokenRequest")
	proto.RegisterType((*GetPublicShareByTokenResponse)(nil), "cernbox.publicsharev1.GetPublicShareByTokenResponse")
}

func init() {
	proto.RegisterFile("cernbox/publicshare/v1/publicshare.proto", fileDescriptor_508404bc649b24b2)
}

var fileDescriptor_508404bc649b24b2 = []byte{
	// 631 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x56, 0xc1, 0x6e, 0xd3, 0x4c,
	0x10, 0xd6, 0xe6, 0x6f, 0xa3, 0xbf, 0x13, 0x68, 0x9b, 0x85, 0x80, 0x63, 0x15, 0x14, 0x8c, 0x54,
	0x45, 0x02, 0xec, 0x26, 0x2d, 0x27, 0x0e, 0xa8, 0x8e, 0x2a, 0x40, 0x42, 0x10, 0x39, 0x80, 0x10,
	0x42, 0xaa, 0x1c, 0x67, 0x29, 0x56, 0x13, 0xef, 0xb2, 0x6b, 0x9b, 0xe4, 0xca, 0x19, 0xc1, 0x81,
	0x37, 0xe0, 0xc8, 0x9d, 0x97, 0xe0, 0x41, 0x78, 0x0c, 0x84, 0x62, 0x6f, 0x82, 0x93, 0x78, 0xa5,
	0xb4, 0x3d, 0xe4, 0x38, 0x9e, 0x99, 0x6f, 0xbe, 0x6f, 0x76, 0x67, 0xbc, 0x50, 0xf7, 0x08, 0x0f,
	0xba, 0x74, 0x68, 0xb1, 0xa8, 0xdb, 0xf7, 0x3d, 0xf1, 0xde, 0xe5, 0xc4, 0x8a, 0x1b, 0x59, 0xd3,
	0x64, 0x9c, 0x86, 0x14, 0x57, 0x64, 0xa4, 0x99, 0x71, 0xc5, 0x0d, 0x7d, 0x57, 0x01, 0xc0, 0x89,
	0xa0, 0x11, 0xf7, 0x88, 0x48, 0xd3, 0x75, 0x6d, 0x12, 0xc7, 0x99, 0x67, 0x89, 0xd0, 0x0d, 0xa3,
	0x89, 0x67, 0xe7, 0x84, 0xd2, 0x93, 0x3e, 0xb1, 0x5c, 0xe6, 0x5b, 0x6e, 0x10, 0xd0, 0xd0, 0x0d,
	0x7d, 0x1a, 0x48, 0xaf, 0xf1, 0x13, 0x81, 0xd6, 0xe2, 0xc4, 0x0d, 0x49, 0x3b, 0x29, 0xd0, 0x19,
	0x17, 0x70, 0xc8, 0x87, 0x88, 0x88, 0x10, 0xeb, 0xf0, 0xff, 0x3b, 0xbf, 0x4f, 0x02, 0x77, 0x40,
	0x34, 0x54, 0x43, 0xf5, 0x0d, 0x67, 0x6a, 0xe3, 0x43, 0x00, 0x46, 0xf8, 0xc0, 0x17, 0xc2, 0xa7,
	0x81, 0x56, 0xa8, 0xa1, 0xfa, 0x66, 0xf3, 0x96, 0x99, 0x2b, 0xc2, 0x6c, 0x4f, 0x03, 0x9d, 0x4c,
	0xd2, 0x18, 0x9e, 0xb9, 0x42, 0x7c, 0xa4, 0xbc, 0xa7, 0xfd, 0x97, 0xc2, 0x4f, 0x6c, 0x7c, 0x13,
	0x80, 0x0c, 0x99, 0xcf, 0x13, 0xb2, 0xda, 0x5a, 0x0d, 0xd5, 0xd7, 0x9c, 0xcc, 0x17, 0xe3, 0x2b,
	0x82, 0x6a, 0x0e, 0x6f, 0xc1, 0x68, 0x20, 0x08, 0xbe, 0x03, 0xc5, 0xb4, 0x07, 0x09, 0xed, 0x52,
	0xf3, 0xca, 0x94, 0x18, 0x67, 0x9e, 0xd9, 0x49, 0x5c, 0x8e, 0x0c, 0xc1, 0x47, 0x70, 0x29, 0xa5,
	0x7b, 0x9c, 0xf0, 0x4d, 0xb4, 0x94, 0x9a, 0x86, 0x4a, 0x4b, 0xa6, 0x5c, 0x89, 0xfd, 0x33, 0x8c,
	0x3f, 0x08, 0xb4, 0x97, 0xac, 0x97, 0xdf, 0xc9, 0x5d, 0xd8, 0xca, 0xd6, 0x38, 0xf6, 0x7b, 0xb2,
	0xa1, 0x97, 0x33, 0x10, 0x4f, 0x7a, 0xf8, 0x01, 0x14, 0x19, 0xed, 0xfb, 0xde, 0x48, 0xb2, 0xb8,
	0xad, 0x60, 0x21, 0x0b, 0x25, 0xa1, 0x8e, 0x4c, 0x39, 0x63, 0x3f, 0x37, 0xb2, 0xfd, 0x9c, 0x3b,
	0xce, 0xf5, 0x73, 0x1c, 0x67, 0x72, 0x24, 0x39, 0x0d, 0x58, 0xe1, 0x91, 0xdc, 0x87, 0xeb, 0x4f,
	0x7d, 0x11, 0x66, 0xfc, 0x62, 0x89, 0xab, 0x6d, 0x7c, 0x41, 0xa0, 0x2d, 0xe6, 0xad, 0x50, 0x87,
	0x0d, 0x9a, 0x43, 0x62, 0x7a, 0x7a, 0x81, 0x9b, 0x65, 0x3c, 0x86, 0x6a, 0x0e, 0xc6, 0x39, 0x44,
	0x19, 0x0f, 0xa1, 0xf2, 0x88, 0x84, 0x17, 0xa0, 0xf2, 0x19, 0xc1, 0xb5, 0x79, 0x84, 0x15, 0x76,
	0xf7, 0x00, 0x76, 0x66, 0xd9, 0xd8, 0xa3, 0x17, 0xf4, 0x94, 0x04, 0x13, 0x59, 0x57, 0x61, 0x3d,
	0x1c, 0xdb, 0x52, 0x4c, 0x6a, 0x18, 0xdf, 0x10, 0xdc, 0x50, 0xa4, 0xad, 0x4e, 0x4b, 0xf3, 0xf7,
	0x3a, 0x94, 0xdb, 0x51, 0x57, 0x9a, 0x1d, 0xc2, 0x63, 0xdf, 0x23, 0x38, 0x86, 0xf2, 0xc2, 0xae,
	0xc4, 0x96, 0x02, 0x5b, 0xf5, 0x37, 0xd0, 0xf7, 0x96, 0x4f, 0x90, 0x1d, 0x88, 0xa1, 0xbc, 0xb0,
	0x10, 0x94, 0x75, 0x55, 0xbb, 0x53, 0x59, 0x57, 0xbd, 0x6b, 0x22, 0xd8, 0x9e, 0x9f, 0x5f, 0x6c,
	0x2a, 0x50, 0x14, 0x0b, 0x42, 0xb7, 0x96, 0x8e, 0x4f, 0x8b, 0xee, 0xa1, 0xb1, 0xdc, 0x85, 0x11,
	0x53, 0xca, 0x55, 0x0d, 0xb4, 0x52, 0xae, 0x7a, 0x7a, 0x07, 0xb0, 0x39, 0x7b, 0x13, 0xf1, 0x5d,
	0x05, 0x46, 0xee, 0xdc, 0xea, 0xf7, 0x96, 0x8c, 0x96, 0xe5, 0x3e, 0xa1, 0xf9, 0x05, 0x20, 0x6f,
	0x3e, 0xde, 0x5f, 0x0a, 0x68, 0x76, 0xbc, 0xf4, 0x83, 0xb3, 0x25, 0xa5, 0x24, 0xec, 0x11, 0x54,
	0x3d, 0x3a, 0xc8, 0x4f, 0xb5, 0xb7, 0x33, 0x89, 0xed, 0xf1, 0x33, 0xa7, 0x8d, 0xde, 0x6c, 0xcd,
	0x84, 0xb0, 0xee, 0xf7, 0x42, 0xb1, 0x65, 0x3f, 0x7f, 0x7d, 0x68, 0xff, 0x28, 0x54, 0x5a, 0x47,
	0xce, 0x33, 0x9b, 0x0e, 0xb3, 0x63, 0xf5, 0xaa, 0xf1, 0xab, 0x50, 0x69, 0x11, 0x1e, 0xd8, 0x74,
	0xf8, 0x76, 0xe6, 0x7b, 0xb7, 0x98, 0xbc, 0x9c, 0xf6, 0xff, 0x06, 0x00, 0x00, 0xff, 0xff, 0xb5,
	0xa5, 0x9b, 0x7c, 0xdc, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PubicShareServiceClient is the client API for PubicShareService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PubicShareServiceClient interface {
	CreatePublicShare(ctx context.Context, in *CreatePublicShareRequest, opts ...grpc.CallOption) (*CreatePublicShareResponse, error)
	UpdatePublicShare(ctx context.Context, in *UpdatePublicShareRequest, opts ...grpc.CallOption) (*UpdatePublicShareResponse, error)
	ListPublicShares(ctx context.Context, in *ListPublicSharesRequest, opts ...grpc.CallOption) (PubicShareService_ListPublicSharesClient, error)
	RevokePublicShare(ctx context.Context, in *RevokePublicShareRequest, opts ...grpc.CallOption) (*RevokePublicShareResponse, error)
	GetPublicShare(ctx context.Context, in *GetPublicShareRequest, opts ...grpc.CallOption) (*GetPublicShareResponse, error)
	GetPublicShareByToken(ctx context.Context, in *GetPublicShareByTokenRequest, opts ...grpc.CallOption) (*GetPublicShareByTokenResponse, error)
}

type pubicShareServiceClient struct {
	cc *grpc.ClientConn
}

func NewPubicShareServiceClient(cc *grpc.ClientConn) PubicShareServiceClient {
	return &pubicShareServiceClient{cc}
}

func (c *pubicShareServiceClient) CreatePublicShare(ctx context.Context, in *CreatePublicShareRequest, opts ...grpc.CallOption) (*CreatePublicShareResponse, error) {
	out := new(CreatePublicShareResponse)
	err := c.cc.Invoke(ctx, "/cernbox.publicsharev1.PubicShareService/CreatePublicShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubicShareServiceClient) UpdatePublicShare(ctx context.Context, in *UpdatePublicShareRequest, opts ...grpc.CallOption) (*UpdatePublicShareResponse, error) {
	out := new(UpdatePublicShareResponse)
	err := c.cc.Invoke(ctx, "/cernbox.publicsharev1.PubicShareService/UpdatePublicShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubicShareServiceClient) ListPublicShares(ctx context.Context, in *ListPublicSharesRequest, opts ...grpc.CallOption) (PubicShareService_ListPublicSharesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PubicShareService_serviceDesc.Streams[0], "/cernbox.publicsharev1.PubicShareService/ListPublicShares", opts...)
	if err != nil {
		return nil, err
	}
	x := &pubicShareServiceListPublicSharesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PubicShareService_ListPublicSharesClient interface {
	Recv() (*ListPublicSharesResponse, error)
	grpc.ClientStream
}

type pubicShareServiceListPublicSharesClient struct {
	grpc.ClientStream
}

func (x *pubicShareServiceListPublicSharesClient) Recv() (*ListPublicSharesResponse, error) {
	m := new(ListPublicSharesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pubicShareServiceClient) RevokePublicShare(ctx context.Context, in *RevokePublicShareRequest, opts ...grpc.CallOption) (*RevokePublicShareResponse, error) {
	out := new(RevokePublicShareResponse)
	err := c.cc.Invoke(ctx, "/cernbox.publicsharev1.PubicShareService/RevokePublicShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubicShareServiceClient) GetPublicShare(ctx context.Context, in *GetPublicShareRequest, opts ...grpc.CallOption) (*GetPublicShareResponse, error) {
	out := new(GetPublicShareResponse)
	err := c.cc.Invoke(ctx, "/cernbox.publicsharev1.PubicShareService/GetPublicShare", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pubicShareServiceClient) GetPublicShareByToken(ctx context.Context, in *GetPublicShareByTokenRequest, opts ...grpc.CallOption) (*GetPublicShareByTokenResponse, error) {
	out := new(GetPublicShareByTokenResponse)
	err := c.cc.Invoke(ctx, "/cernbox.publicsharev1.PubicShareService/GetPublicShareByToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PubicShareServiceServer is the server API for PubicShareService service.
type PubicShareServiceServer interface {
	CreatePublicShare(context.Context, *CreatePublicShareRequest) (*CreatePublicShareResponse, error)
	UpdatePublicShare(context.Context, *UpdatePublicShareRequest) (*UpdatePublicShareResponse, error)
	ListPublicShares(*ListPublicSharesRequest, PubicShareService_ListPublicSharesServer) error
	RevokePublicShare(context.Context, *RevokePublicShareRequest) (*RevokePublicShareResponse, error)
	GetPublicShare(context.Context, *GetPublicShareRequest) (*GetPublicShareResponse, error)
	GetPublicShareByToken(context.Context, *GetPublicShareByTokenRequest) (*GetPublicShareByTokenResponse, error)
}

func RegisterPubicShareServiceServer(s *grpc.Server, srv PubicShareServiceServer) {
	s.RegisterService(&_PubicShareService_serviceDesc, srv)
}

func _PubicShareService_CreatePublicShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePublicShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubicShareServiceServer).CreatePublicShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cernbox.publicsharev1.PubicShareService/CreatePublicShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubicShareServiceServer).CreatePublicShare(ctx, req.(*CreatePublicShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubicShareService_UpdatePublicShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePublicShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubicShareServiceServer).UpdatePublicShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cernbox.publicsharev1.PubicShareService/UpdatePublicShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubicShareServiceServer).UpdatePublicShare(ctx, req.(*UpdatePublicShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubicShareService_ListPublicShares_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListPublicSharesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PubicShareServiceServer).ListPublicShares(m, &pubicShareServiceListPublicSharesServer{stream})
}

type PubicShareService_ListPublicSharesServer interface {
	Send(*ListPublicSharesResponse) error
	grpc.ServerStream
}

type pubicShareServiceListPublicSharesServer struct {
	grpc.ServerStream
}

func (x *pubicShareServiceListPublicSharesServer) Send(m *ListPublicSharesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PubicShareService_RevokePublicShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokePublicShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubicShareServiceServer).RevokePublicShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cernbox.publicsharev1.PubicShareService/RevokePublicShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubicShareServiceServer).RevokePublicShare(ctx, req.(*RevokePublicShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubicShareService_GetPublicShare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicShareRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubicShareServiceServer).GetPublicShare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cernbox.publicsharev1.PubicShareService/GetPublicShare",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubicShareServiceServer).GetPublicShare(ctx, req.(*GetPublicShareRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PubicShareService_GetPublicShareByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublicShareByTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PubicShareServiceServer).GetPublicShareByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cernbox.publicsharev1.PubicShareService/GetPublicShareByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PubicShareServiceServer).GetPublicShareByToken(ctx, req.(*GetPublicShareByTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PubicShareService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cernbox.publicsharev1.PubicShareService",
	HandlerType: (*PubicShareServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePublicShare",
			Handler:    _PubicShareService_CreatePublicShare_Handler,
		},
		{
			MethodName: "UpdatePublicShare",
			Handler:    _PubicShareService_UpdatePublicShare_Handler,
		},
		{
			MethodName: "RevokePublicShare",
			Handler:    _PubicShareService_RevokePublicShare_Handler,
		},
		{
			MethodName: "GetPublicShare",
			Handler:    _PubicShareService_GetPublicShare_Handler,
		},
		{
			MethodName: "GetPublicShareByToken",
			Handler:    _PubicShareService_GetPublicShareByToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListPublicShares",
			Handler:       _PubicShareService_ListPublicShares_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cernbox/publicshare/v1/publicshare.proto",
}
