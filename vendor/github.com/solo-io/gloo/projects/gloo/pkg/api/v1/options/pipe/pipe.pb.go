// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/pipe/pipe.proto

package pipe

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	options "github.com/solo-io/gloo/projects/gloo/pkg/api/v1/options"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Pipe upstreams are used to route request to services listening at a Unix Domain Socket.
// Pipe upstreams can be used to proxy any kind of service, and therefore contain a ServiceSpec
// for additional service-specific configuration.
// Unlike upstreams created by service discovery, Pipe Upstreams must be created manually by users
type UpstreamSpec struct {
	// The Unix Domain Socket path.
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	// An optional Service Spec describing the service listening at this address
	ServiceSpec          *options.ServiceSpec `protobuf:"bytes,2,opt,name=service_spec,json=serviceSpec,proto3" json:"service_spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpstreamSpec) Reset()         { *m = UpstreamSpec{} }
func (m *UpstreamSpec) String() string { return proto.CompactTextString(m) }
func (*UpstreamSpec) ProtoMessage()    {}
func (*UpstreamSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c6840b9e9c2b2b5, []int{0}
}
func (m *UpstreamSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamSpec.Unmarshal(m, b)
}
func (m *UpstreamSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamSpec.Marshal(b, m, deterministic)
}
func (m *UpstreamSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamSpec.Merge(m, src)
}
func (m *UpstreamSpec) XXX_Size() int {
	return xxx_messageInfo_UpstreamSpec.Size(m)
}
func (m *UpstreamSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamSpec.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamSpec proto.InternalMessageInfo

func (m *UpstreamSpec) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *UpstreamSpec) GetServiceSpec() *options.ServiceSpec {
	if m != nil {
		return m.ServiceSpec
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamSpec)(nil), "pipe.options.gloo.solo.io.UpstreamSpec")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/pipe/pipe.proto", fileDescriptor_2c6840b9e9c2b2b5)
}

var fileDescriptor_2c6840b9e9c2b2b5 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x72, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0x4f,
	0xcf, 0xc9, 0xcf, 0xd7, 0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0x86, 0xf0, 0x12, 0x0b,
	0x32, 0xf5, 0xcb, 0x0c, 0xf5, 0xf3, 0x0b, 0x4a, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x0b, 0x32, 0x0b,
	0x52, 0xc1, 0x84, 0x5e, 0x41, 0x51, 0x7e, 0x49, 0xbe, 0x90, 0x24, 0x98, 0x0d, 0x95, 0xd5, 0x03,
	0xe9, 0xd0, 0x03, 0x19, 0xa6, 0x97, 0x99, 0x2f, 0x25, 0x92, 0x9e, 0x9f, 0x9e, 0x0f, 0x56, 0xa5,
	0x0f, 0x62, 0x41, 0x34, 0x48, 0x09, 0xa5, 0x56, 0x94, 0x40, 0x04, 0x53, 0x2b, 0x4a, 0xa0, 0x62,
	0x26, 0x84, 0xed, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x8d, 0x2f, 0x2e, 0x48, 0x4d, 0x86,
	0xe8, 0x52, 0xca, 0xe0, 0xe2, 0x09, 0x2d, 0x28, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x0d, 0x2e, 0x48,
	0x4d, 0x16, 0x12, 0xe2, 0x62, 0x29, 0x48, 0x2c, 0xc9, 0x90, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c,
	0x02, 0xb3, 0x85, 0x5c, 0xb8, 0x78, 0x90, 0x75, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b, 0x29,
	0x62, 0x75, 0xb0, 0x5e, 0x30, 0x44, 0x25, 0xc8, 0xb0, 0x20, 0xee, 0x62, 0x04, 0xc7, 0xc9, 0x7d,
	0xc7, 0x57, 0x16, 0xc6, 0x15, 0x8f, 0xe4, 0x18, 0xa3, 0x6c, 0x89, 0x0b, 0xb4, 0x82, 0xec, 0x74,
	0x6c, 0x01, 0x97, 0xc4, 0x06, 0x76, 0xb9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x90, 0x43,
	0xe8, 0x7c, 0x01, 0x00, 0x00,
}

func (this *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
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
	if this.Path != that1.Path {
		return false
	}
	if !this.ServiceSpec.Equal(that1.ServiceSpec) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}