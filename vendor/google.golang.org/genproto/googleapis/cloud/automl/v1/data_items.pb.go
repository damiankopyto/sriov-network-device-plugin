// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1/data_items.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// A representation of a text snippet.
type TextSnippet struct {
	// Required. The content of the text snippet as a string. Up to 250000
	// characters long.
	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	// Optional. The format of [content][google.cloud.automl.v1.TextSnippet.content]. Currently the only two allowed
	// values are "text/html" and "text/plain". If left blank, the format is
	// automatically determined from the type of the uploaded [content][google.cloud.automl.v1.TextSnippet.content].
	MimeType string `protobuf:"bytes,2,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	// Output only. HTTP URI where you can download the content.
	ContentUri           string   `protobuf:"bytes,4,opt,name=content_uri,json=contentUri,proto3" json:"content_uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TextSnippet) Reset()         { *m = TextSnippet{} }
func (m *TextSnippet) String() string { return proto.CompactTextString(m) }
func (*TextSnippet) ProtoMessage()    {}
func (*TextSnippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_9aa6620a22bddfe1, []int{0}
}

func (m *TextSnippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TextSnippet.Unmarshal(m, b)
}
func (m *TextSnippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TextSnippet.Marshal(b, m, deterministic)
}
func (m *TextSnippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TextSnippet.Merge(m, src)
}
func (m *TextSnippet) XXX_Size() int {
	return xxx_messageInfo_TextSnippet.Size(m)
}
func (m *TextSnippet) XXX_DiscardUnknown() {
	xxx_messageInfo_TextSnippet.DiscardUnknown(m)
}

var xxx_messageInfo_TextSnippet proto.InternalMessageInfo

func (m *TextSnippet) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *TextSnippet) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

func (m *TextSnippet) GetContentUri() string {
	if m != nil {
		return m.ContentUri
	}
	return ""
}

// Example data used for training or prediction.
type ExamplePayload struct {
	// Required. Input only. The example data.
	//
	// Types that are valid to be assigned to Payload:
	//	*ExamplePayload_TextSnippet
	Payload              isExamplePayload_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ExamplePayload) Reset()         { *m = ExamplePayload{} }
func (m *ExamplePayload) String() string { return proto.CompactTextString(m) }
func (*ExamplePayload) ProtoMessage()    {}
func (*ExamplePayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_9aa6620a22bddfe1, []int{1}
}

func (m *ExamplePayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExamplePayload.Unmarshal(m, b)
}
func (m *ExamplePayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExamplePayload.Marshal(b, m, deterministic)
}
func (m *ExamplePayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExamplePayload.Merge(m, src)
}
func (m *ExamplePayload) XXX_Size() int {
	return xxx_messageInfo_ExamplePayload.Size(m)
}
func (m *ExamplePayload) XXX_DiscardUnknown() {
	xxx_messageInfo_ExamplePayload.DiscardUnknown(m)
}

var xxx_messageInfo_ExamplePayload proto.InternalMessageInfo

type isExamplePayload_Payload interface {
	isExamplePayload_Payload()
}

type ExamplePayload_TextSnippet struct {
	TextSnippet *TextSnippet `protobuf:"bytes,2,opt,name=text_snippet,json=textSnippet,proto3,oneof"`
}

func (*ExamplePayload_TextSnippet) isExamplePayload_Payload() {}

func (m *ExamplePayload) GetPayload() isExamplePayload_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ExamplePayload) GetTextSnippet() *TextSnippet {
	if x, ok := m.GetPayload().(*ExamplePayload_TextSnippet); ok {
		return x.TextSnippet
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExamplePayload) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExamplePayload_TextSnippet)(nil),
	}
}

func init() {
	proto.RegisterType((*TextSnippet)(nil), "google.cloud.automl.v1.TextSnippet")
	proto.RegisterType((*ExamplePayload)(nil), "google.cloud.automl.v1.ExamplePayload")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1/data_items.proto", fileDescriptor_9aa6620a22bddfe1)
}

var fileDescriptor_9aa6620a22bddfe1 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0x5f, 0xca, 0xe3, 0xf5, 0x75, 0xf3, 0x78, 0x87, 0x1c, 0x4a, 0x5a, 0xc5, 0x4a, 0x3d,
	0xe8, 0x29, 0x4b, 0xf4, 0x16, 0xbd, 0x58, 0x11, 0x7b, 0x50, 0x28, 0xb5, 0xf6, 0x20, 0x85, 0xb0,
	0x4d, 0xd6, 0xb0, 0x90, 0xec, 0x2c, 0xc9, 0xa4, 0x34, 0x1f, 0xc3, 0xaf, 0xe5, 0x47, 0xf1, 0x53,
	0x48, 0x77, 0x57, 0x94, 0x60, 0x6f, 0x99, 0xf9, 0xfd, 0xe6, 0x9f, 0x61, 0x96, 0x9c, 0x66, 0x00,
	0x59, 0xce, 0x69, 0x92, 0x43, 0x9d, 0x52, 0x56, 0x23, 0x14, 0x39, 0xdd, 0x84, 0x34, 0x65, 0xc8,
	0x62, 0x81, 0xbc, 0xa8, 0x02, 0x55, 0x02, 0x82, 0xd7, 0x37, 0x62, 0xa0, 0xc5, 0xc0, 0x88, 0xc1,
	0x26, 0x1c, 0x8e, 0xf6, 0x04, 0x08, 0x30, 0x83, 0xc3, 0x81, 0x15, 0x74, 0xb5, 0xae, 0x5f, 0x28,
	0x93, 0x8d, 0x45, 0x47, 0x6d, 0x94, 0xd6, 0x25, 0x43, 0x01, 0xd2, 0xf2, 0xc3, 0x36, 0xaf, 0xb0,
	0xac, 0x13, 0x6c, 0x51, 0xa6, 0x04, 0x65, 0x52, 0x02, 0xea, 0x51, 0xbb, 0xef, 0x98, 0x13, 0x77,
	0xc1, 0xb7, 0xf8, 0x28, 0x85, 0x52, 0x1c, 0x3d, 0x9f, 0x74, 0x13, 0x90, 0xc8, 0x25, 0xfa, 0xce,
	0xb1, 0x73, 0xd6, 0x9b, 0x7f, 0x96, 0xde, 0x01, 0xe9, 0x15, 0xa2, 0xe0, 0x31, 0x36, 0x8a, 0xfb,
	0x1d, 0xcd, 0xfe, 0xee, 0x1a, 0x8b, 0x46, 0x71, 0x6f, 0x44, 0x5c, 0xeb, 0xc5, 0x75, 0x29, 0xfc,
	0xdf, 0x1a, 0x13, 0xdb, 0x7a, 0x2a, 0xc5, 0x98, 0x93, 0xff, 0xb7, 0x5b, 0x56, 0xa8, 0x9c, 0xcf,
	0x58, 0x93, 0x03, 0x4b, 0xbd, 0x29, 0xf9, 0x87, 0x7c, 0x8b, 0x71, 0x65, 0xfe, 0xac, 0x23, 0xdd,
	0xf3, 0x93, 0xe0, 0xe7, 0xfb, 0x05, 0xdf, 0x96, 0x9c, 0xfe, 0x9a, 0xbb, 0xf8, 0x55, 0x4e, 0x7a,
	0xa4, 0xab, 0x4c, 0xe8, 0xe4, 0xd5, 0x21, 0xc3, 0x04, 0x8a, 0x3d, 0x21, 0x33, 0xe7, 0xf9, 0xca,
	0x92, 0x0c, 0x72, 0x26, 0xb3, 0x00, 0xca, 0x8c, 0x66, 0x5c, 0xea, 0x53, 0x50, 0x83, 0x98, 0x12,
	0x55, 0xfb, 0x95, 0x2e, 0xcd, 0xd7, 0x5b, 0xa7, 0x7f, 0xa7, 0x9d, 0xd5, 0xcd, 0x8e, 0xaf, 0xae,
	0x6b, 0x84, 0x87, 0x7c, 0xb5, 0x0c, 0xdf, 0x3b, 0x03, 0x03, 0xa2, 0x48, 0x93, 0x28, 0xd2, 0xe8,
	0x3e, 0x8a, 0x96, 0xe1, 0xfa, 0x8f, 0x4e, 0xbf, 0xf8, 0x08, 0x00, 0x00, 0xff, 0xff, 0xc5, 0xd3,
	0x75, 0x0d, 0x43, 0x02, 0x00, 0x00,
}
