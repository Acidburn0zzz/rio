// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/parameters.proto

package transformation

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
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

type Parameters struct {
	// headers that will be used to extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         headers:
	//           x-user-id: '{userId}'
	Headers map[string]string `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// part of the (or the entire) path that will be used extract data for processing output templates
	// Gloo will search for parameters by their name in header value strings, enclosed in single
	// curly braces
	// Example:
	//   extensions:
	//     parameters:
	//         path: /users/{ userId }
	Path                 *types.StringValue `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Parameters) Reset()         { *m = Parameters{} }
func (m *Parameters) String() string { return proto.CompactTextString(m) }
func (*Parameters) ProtoMessage()    {}
func (*Parameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_bfd93bc8f6e984ec, []int{0}
}
func (m *Parameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Parameters.Unmarshal(m, b)
}
func (m *Parameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Parameters.Marshal(b, m, deterministic)
}
func (m *Parameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Parameters.Merge(m, src)
}
func (m *Parameters) XXX_Size() int {
	return xxx_messageInfo_Parameters.Size(m)
}
func (m *Parameters) XXX_DiscardUnknown() {
	xxx_messageInfo_Parameters.DiscardUnknown(m)
}

var xxx_messageInfo_Parameters proto.InternalMessageInfo

func (m *Parameters) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Parameters) GetPath() *types.StringValue {
	if m != nil {
		return m.Path
	}
	return nil
}

func init() {
	proto.RegisterType((*Parameters)(nil), "transformation.options.gloo.solo.io.Parameters")
	proto.RegisterMapType((map[string]string)(nil), "transformation.options.gloo.solo.io.Parameters.HeadersEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/options/transformation/parameters.proto", fileDescriptor_bfd93bc8f6e984ec)
}

var fileDescriptor_bfd93bc8f6e984ec = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xb1, 0x4e, 0x42, 0x31,
	0x18, 0x85, 0x53, 0x40, 0x0d, 0xc5, 0xc1, 0x34, 0x0c, 0x84, 0x18, 0x42, 0x74, 0x61, 0xf1, 0xaf,
	0xe2, 0x62, 0x88, 0x93, 0x89, 0xd1, 0x51, 0xaf, 0x86, 0xc1, 0xad, 0x60, 0x29, 0x95, 0x7b, 0xef,
	0xdf, 0xb4, 0xbd, 0x08, 0x6f, 0xe4, 0x23, 0xf8, 0x1c, 0x3e, 0x82, 0xef, 0xe0, 0x6e, 0x7a, 0x0b,
	0xa8, 0x8b, 0x61, 0x3b, 0xe7, 0xa4, 0xe7, 0xeb, 0x49, 0x7e, 0xfa, 0xa8, 0xb4, 0x9f, 0x16, 0x23,
	0x18, 0x63, 0xc6, 0x1d, 0xa6, 0x78, 0xa2, 0x91, 0xab, 0x14, 0x91, 0x1b, 0x8b, 0x2f, 0x72, 0xec,
	0x5d, 0x74, 0xc2, 0x68, 0x3e, 0x3f, 0xe3, 0x68, 0xbc, 0xc6, 0xdc, 0x71, 0x6f, 0x45, 0xee, 0x26,
	0x68, 0x33, 0x11, 0x3c, 0x37, 0xc2, 0x8a, 0x4c, 0x7a, 0x69, 0x1d, 0x18, 0x8b, 0x1e, 0xd9, 0xf1,
	0xdf, 0x07, 0xb0, 0xea, 0x41, 0x60, 0x41, 0xf8, 0x06, 0x34, 0xb6, 0x3b, 0x0a, 0x51, 0xa5, 0x92,
	0x97, 0x95, 0x51, 0x31, 0xe1, 0xaf, 0x56, 0x18, 0xb3, 0x81, 0xb4, 0x9b, 0x0a, 0x15, 0x96, 0x92,
	0x07, 0xb5, 0x4a, 0x99, 0x5c, 0xf8, 0x18, 0xca, 0x85, 0x8f, 0xd9, 0xd1, 0x07, 0xa1, 0xf4, 0x6e,
	0xb3, 0x81, 0x0d, 0xe9, 0xde, 0x54, 0x8a, 0x67, 0x69, 0x5d, 0x8b, 0x74, 0xab, 0xbd, 0x46, 0xff,
	0x12, 0xb6, 0xd8, 0x03, 0x3f, 0x04, 0xb8, 0x8d, 0xf5, 0xeb, 0xdc, 0xdb, 0x65, 0xb2, 0x86, 0xb1,
	0x53, 0x5a, 0x33, 0xc2, 0x4f, 0x5b, 0x95, 0x2e, 0xe9, 0x35, 0xfa, 0x87, 0x10, 0xf7, 0xc3, 0x7a,
	0x3f, 0x3c, 0x78, 0xab, 0x73, 0x35, 0x14, 0x69, 0x21, 0x93, 0xf2, 0x65, 0x7b, 0x40, 0xf7, 0x7f,
	0xa3, 0xd8, 0x01, 0xad, 0xce, 0xe4, 0xb2, 0x45, 0xba, 0xa4, 0x57, 0x4f, 0x82, 0x64, 0x4d, 0xba,
	0x33, 0x0f, 0x85, 0x12, 0x5a, 0x4f, 0xa2, 0x19, 0x54, 0x2e, 0xc8, 0xd5, 0xfd, 0xfb, 0x57, 0x8d,
	0xbc, 0x7d, 0x76, 0xc8, 0xd3, 0xcd, 0x76, 0x37, 0x32, 0x33, 0xf5, 0xff, 0x9d, 0x46, 0xbb, 0xe5,
	0xd4, 0xf3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0xd1, 0x0b, 0xac, 0xf5, 0x01, 0x00, 0x00,
}

func (this *Parameters) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Parameters)
	if !ok {
		that2, ok := that.(Parameters)
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
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if this.Headers[i] != that1.Headers[i] {
			return false
		}
	}
	if !this.Path.Equal(that1.Path) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}