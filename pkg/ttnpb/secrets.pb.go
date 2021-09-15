// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/secrets.proto

package ttnpb

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
	reflect "reflect"
	strings "strings"
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

// Secret contains a secret value. It also contains the ID of the Encryption key used to encrypt it.
type Secret struct {
	// ID of the Key used to encrypt the secret.
	KeyId                string   `protobuf:"bytes,1,opt,name=key_id,json=keyId,proto3" json:"key_id,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Secret) Reset()      { *m = Secret{} }
func (*Secret) ProtoMessage() {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c9d796b7f7ca235, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetKeyId() string {
	if m != nil {
		return m.KeyId
	}
	return ""
}

func (m *Secret) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*Secret)(nil), "ttn.lorawan.v3.Secret")
	golang_proto.RegisterType((*Secret)(nil), "ttn.lorawan.v3.Secret")
}

func init() { proto.RegisterFile("lorawan-stack/api/secrets.proto", fileDescriptor_8c9d796b7f7ca235) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/secrets.proto", fileDescriptor_8c9d796b7f7ca235)
}

var fileDescriptor_8c9d796b7f7ca235 = []byte{
	// 266 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0xc9, 0x2f, 0x4a,
	0x2c, 0x4f, 0xcc, 0xd3, 0x2d, 0x2e, 0x49, 0x4c, 0xce, 0xd6, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0x4e,
	0x4d, 0x2e, 0x4a, 0x2d, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0x29, 0xc9,
	0xd3, 0x83, 0x2a, 0xd2, 0x2b, 0x33, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b,
	0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4b, 0x2a, 0x4d, 0x03, 0xf3, 0xc0,
	0x1c, 0x30, 0x0b, 0xa2, 0x5d, 0xca, 0x11, 0x49, 0x79, 0x6a, 0x5e, 0x59, 0x7e, 0x65, 0x41, 0x51,
	0x7e, 0x45, 0x25, 0x44, 0x53, 0xb2, 0x6e, 0x7a, 0x6a, 0x9e, 0x6e, 0x59, 0x62, 0x4e, 0x66, 0x4a,
	0x62, 0x49, 0xaa, 0x3e, 0x06, 0x03, 0x62, 0x84, 0x92, 0x3d, 0x17, 0x5b, 0x30, 0xd8, 0x49, 0x42,
	0xa2, 0x5c, 0x6c, 0xd9, 0xa9, 0x95, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41,
	0xac, 0xd9, 0xa9, 0x95, 0x9e, 0x29, 0x42, 0x72, 0x5c, 0xac, 0x65, 0x89, 0x39, 0xa5, 0xa9, 0x12,
	0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x4e, 0x1c, 0xbf, 0x9c, 0x58, 0xab, 0x98, 0x25, 0x1a, 0x04, 0x82,
	0x20, 0xc2, 0x4e, 0xbe, 0x37, 0x1e, 0xca, 0x31, 0x34, 0x3c, 0x92, 0x63, 0x5c, 0xf1, 0x48, 0x8e,
	0xf1, 0xc5, 0x23, 0x39, 0x86, 0x0f, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0x38, 0xf0, 0x58,
	0x8e, 0x31, 0x4a, 0x3f, 0x3d, 0x5f, 0xaf, 0x24, 0x23, 0xb5, 0x24, 0x23, 0x33, 0x2f, 0xbd, 0x58,
	0x2f, 0x2f, 0xb5, 0xa4, 0x3c, 0xbf, 0x28, 0x5b, 0x1f, 0x35, 0x58, 0xca, 0x8c, 0xf5, 0x0b, 0xb2,
	0xd3, 0xf5, 0x4b, 0x4a, 0xf2, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0xce, 0x32, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x3f, 0x5f, 0x1c, 0x10, 0x3b, 0x01, 0x00, 0x00,
}

func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.KeyId != that1.KeyId {
		return false
	}
	if !bytes.Equal(this.Value, that1.Value) {
		return false
	}
	return true
}
func (this *Secret) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Secret{`,
		`KeyId:` + fmt.Sprintf("%v", this.KeyId) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSecrets(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
