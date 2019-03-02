// Code generated by protoc-gen-go. DO NOT EDIT.
// source: signature.proto

package legacy_pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Signature_Version int32

const (
	Signature_UNKNOWN_VERSION Signature_Version = 0
	Signature__0_0_1          Signature_Version = 1
)

var Signature_Version_name = map[int32]string{
	0: "UNKNOWN_VERSION",
	1: "_0_0_1",
}
var Signature_Version_value = map[string]int32{
	"UNKNOWN_VERSION": 0,
	"_0_0_1":          1,
}

func (x Signature_Version) Enum() *Signature_Version {
	p := new(Signature_Version)
	*p = x
	return p
}
func (x Signature_Version) String() string {
	return proto.EnumName(Signature_Version_name, int32(x))
}
func (x *Signature_Version) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Signature_Version_value, data, "Signature_Version")
	if err != nil {
		return err
	}
	*x = Signature_Version(value)
	return nil
}
func (Signature_Version) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_signature_db49c854424da426, []int{0, 0}
}

type Signature struct {
	Version              *Signature_Version `protobuf:"varint,1,req,name=version,enum=legacy_pb.Signature_Version" json:"version,omitempty"`
	SignatureType        *KeyType           `protobuf:"varint,2,req,name=signatureType,enum=legacy_pb.KeyType" json:"signatureType,omitempty"`
	Signature            []byte             `protobuf:"bytes,3,req,name=signature" json:"signature,omitempty"`
	CertificateId        []byte             `protobuf:"bytes,4,req,name=certificateId" json:"certificateId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Signature) Reset()         { *m = Signature{} }
func (m *Signature) String() string { return proto.CompactTextString(m) }
func (*Signature) ProtoMessage()    {}
func (*Signature) Descriptor() ([]byte, []int) {
	return fileDescriptor_signature_db49c854424da426, []int{0}
}
func (m *Signature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Signature.Unmarshal(m, b)
}
func (m *Signature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Signature.Marshal(b, m, deterministic)
}
func (dst *Signature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Signature.Merge(dst, src)
}
func (m *Signature) XXX_Size() int {
	return xxx_messageInfo_Signature.Size(m)
}
func (m *Signature) XXX_DiscardUnknown() {
	xxx_messageInfo_Signature.DiscardUnknown(m)
}

var xxx_messageInfo_Signature proto.InternalMessageInfo

func (m *Signature) GetVersion() Signature_Version {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return Signature_UNKNOWN_VERSION
}

func (m *Signature) GetSignatureType() KeyType {
	if m != nil && m.SignatureType != nil {
		return *m.SignatureType
	}
	return KeyType_UNKNOWN_PUBLIC_KEY_TYPE
}

func (m *Signature) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Signature) GetCertificateId() []byte {
	if m != nil {
		return m.CertificateId
	}
	return nil
}

func init() {
	proto.RegisterType((*Signature)(nil), "legacy_pb.Signature")
	proto.RegisterEnum("legacy_pb.Signature_Version", Signature_Version_name, Signature_Version_value)
}

func init() { proto.RegisterFile("signature.proto", fileDescriptor_signature_db49c854424da426) }

var fileDescriptor_signature_db49c854424da426 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xce, 0x4c, 0xcf,
	0x4b, 0x2c, 0x29, 0x2d, 0x4a, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x49, 0x4d,
	0x4f, 0x4c, 0xae, 0x8c, 0x2f, 0x48, 0x92, 0x12, 0x4c, 0x4e, 0x2d, 0x2a, 0xc9, 0x4c, 0xcb, 0x4c,
	0x4e, 0x2c, 0x81, 0xca, 0x2a, 0xbd, 0x65, 0xe4, 0xe2, 0x0c, 0x86, 0xe9, 0x10, 0x32, 0xe3, 0x62,
	0x2f, 0x4b, 0x2d, 0x2a, 0xce, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd2, 0xe0, 0x33, 0x92, 0xd1,
	0x83, 0xeb, 0xd6, 0x83, 0x2b, 0xd3, 0x0b, 0x83, 0xa8, 0x09, 0x82, 0x29, 0x16, 0xb2, 0xe0, 0xe2,
	0x85, 0x5b, 0x1b, 0x52, 0x59, 0x90, 0x2a, 0xc1, 0x04, 0xd6, 0x2d, 0x84, 0xa4, 0xdb, 0x3b, 0xb5,
	0x12, 0x24, 0x13, 0x84, 0xaa, 0x50, 0x48, 0x86, 0x8b, 0x13, 0x2e, 0x20, 0xc1, 0xac, 0xc0, 0xa4,
	0xc1, 0x13, 0x84, 0x10, 0x10, 0x52, 0xe1, 0xe2, 0x45, 0x72, 0xb2, 0x67, 0x8a, 0x04, 0x0b, 0x58,
	0x05, 0xaa, 0xa0, 0x92, 0x16, 0x17, 0x3b, 0xd4, 0x45, 0x42, 0xc2, 0x5c, 0xfc, 0xa1, 0x7e, 0xde,
	0x7e, 0xfe, 0xe1, 0x7e, 0xf1, 0x61, 0xae, 0x41, 0xc1, 0x9e, 0xfe, 0x7e, 0x02, 0x0c, 0x42, 0x5c,
	0x5c, 0x6c, 0xf1, 0x06, 0xf1, 0x06, 0xf1, 0x86, 0x02, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x48, 0xf0, 0x60, 0xdb, 0x1f, 0x01, 0x00, 0x00,
}
