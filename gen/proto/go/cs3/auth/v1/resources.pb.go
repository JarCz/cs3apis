// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cs3/auth/v1/resources.proto

package authv1pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type User struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Mail                 string   `protobuf:"bytes,2,opt,name=mail,proto3" json:"mail,omitempty"`
	DisplayName          string   `protobuf:"bytes,3,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Groups               []string `protobuf:"bytes,4,rep,name=groups,proto3" json:"groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_905048cb25412094, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetMail() string {
	if m != nil {
		return m.Mail
	}
	return ""
}

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetGroups() []string {
	if m != nil {
		return m.Groups
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "cs3.authv1.User")
}

func init() { proto.RegisterFile("cs3/auth/v1/resources.proto", fileDescriptor_905048cb25412094) }

var fileDescriptor_905048cb25412094 = []byte{
	// 206 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0x2e, 0x36, 0xd6,
	0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0x4a, 0x2d, 0xce, 0x2f, 0x2d, 0x4a, 0x4e,
	0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4a, 0x2e, 0x36, 0xd6, 0x03, 0x49, 0x96,
	0x19, 0x2a, 0x15, 0x72, 0xb1, 0x84, 0x16, 0xa7, 0x16, 0x09, 0x49, 0x71, 0x71, 0x94, 0x16, 0xa7,
	0x16, 0xe5, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x42, 0x42,
	0x5c, 0x2c, 0xb9, 0x89, 0x99, 0x39, 0x12, 0x4c, 0x60, 0x71, 0x30, 0x5b, 0x48, 0x91, 0x8b, 0x27,
	0x25, 0xb3, 0xb8, 0x20, 0x27, 0xb1, 0x32, 0x1e, 0xac, 0x87, 0x19, 0x2c, 0xc7, 0x0d, 0x15, 0xf3,
	0x03, 0x69, 0x13, 0xe3, 0x62, 0x4b, 0x2f, 0xca, 0x2f, 0x2d, 0x28, 0x96, 0x60, 0x51, 0x60, 0xd6,
	0xe0, 0x0c, 0x82, 0xf2, 0x9c, 0xfc, 0xb9, 0xf8, 0x92, 0xf3, 0x73, 0xf5, 0x10, 0x8e, 0x70, 0xe2,
	0x0b, 0x82, 0xb9, 0x30, 0x00, 0xe4, 0xc0, 0x00, 0xc6, 0x28, 0x0e, 0x88, 0x4c, 0x41, 0xd2, 0x22,
	0x26, 0x36, 0x67, 0x27, 0xff, 0x08, 0x47, 0xa7, 0x55, 0x4c, 0x5c, 0xce, 0xc1, 0xc6, 0x7a, 0x8e,
	0xa5, 0x25, 0x19, 0x61, 0x86, 0xa7, 0xc0, 0x9c, 0x18, 0x08, 0x27, 0x89, 0x0d, 0xec, 0x2d, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x52, 0x4e, 0xd9, 0xc0, 0xf5, 0x00, 0x00, 0x00,
}
