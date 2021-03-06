// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/crypto-zero/micro/debug/log/proto/log.proto

package go_micro_debug_log

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Record struct {
	// timestamp of log record
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// record metadata
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// record value
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1ac715ceb139de, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Record) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Record) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReadRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1ac715ceb139de, []int{1}
}

func (m *ReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadRequest.Unmarshal(m, b)
}
func (m *ReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadRequest.Marshal(b, m, deterministic)
}
func (m *ReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadRequest.Merge(m, src)
}
func (m *ReadRequest) XXX_Size() int {
	return xxx_messageInfo_ReadRequest.Size(m)
}
func (m *ReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReadRequest proto.InternalMessageInfo

func (m *ReadRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *ReadRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type ReadResponse struct {
	Records              []*Record `protobuf:"bytes,1,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ReadResponse) Reset()         { *m = ReadResponse{} }
func (m *ReadResponse) String() string { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()    {}
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e1ac715ceb139de, []int{2}
}

func (m *ReadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadResponse.Unmarshal(m, b)
}
func (m *ReadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadResponse.Marshal(b, m, deterministic)
}
func (m *ReadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadResponse.Merge(m, src)
}
func (m *ReadResponse) XXX_Size() int {
	return xxx_messageInfo_ReadResponse.Size(m)
}
func (m *ReadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReadResponse proto.InternalMessageInfo

func (m *ReadResponse) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "go.micro.debug.log.Record")
	proto.RegisterMapType((map[string]string)(nil), "go.micro.debug.log.Record.MetadataEntry")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.debug.log.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.debug.log.ReadResponse")
}

func init() {
	proto.RegisterFile("github.com/crypto-zero/micro/debug/log/proto/log.proto", fileDescriptor_1e1ac715ceb139de)
}

var fileDescriptor_1e1ac715ceb139de = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x51, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x35, 0xab, 0x6e, 0xee, 0x9b, 0x82, 0x04, 0x0f, 0x61, 0x08, 0x96, 0x9e, 0x7a, 0xca, 0x64,
	0x7a, 0x10, 0x3d, 0x09, 0xf3, 0xa4, 0x5e, 0xf2, 0x0f, 0xb2, 0xf6, 0x23, 0x16, 0x9b, 0xa6, 0x26,
	0x69, 0x61, 0xbf, 0xcf, 0x3f, 0x26, 0x4d, 0xdb, 0x89, 0xe8, 0x2e, 0xe1, 0x7b, 0x5f, 0xde, 0x7b,
	0x79, 0x8f, 0xc0, 0x8d, 0x2a, 0xfc, 0x7b, 0xb3, 0xe5, 0x99, 0xd1, 0x2b, 0x5d, 0x64, 0xd6, 0x0c,
	0x67, 0x8e, 0xdb, 0x46, 0xad, 0x4a, 0xa3, 0x56, 0xb5, 0x35, 0xde, 0x74, 0x13, 0x0f, 0x13, 0xa5,
	0xca, 0xf0, 0xc0, 0xe1, 0x81, 0xc3, 0x4b, 0xa3, 0x92, 0x2f, 0x02, 0x53, 0x81, 0x99, 0xb1, 0x39,
	0xbd, 0x82, 0xb9, 0x2f, 0x34, 0x3a, 0x2f, 0x75, 0xcd, 0x48, 0x4c, 0xd2, 0x48, 0xfc, 0x2c, 0xe8,
	0x06, 0x4e, 0x35, 0x7a, 0x99, 0x4b, 0x2f, 0xd9, 0x24, 0x8e, 0xd2, 0xc5, 0x3a, 0xe5, 0x7f, 0xfd,
	0x78, 0xef, 0xc5, 0xdf, 0x06, 0xea, 0x73, 0xe5, 0xed, 0x4e, 0xec, 0x95, 0x94, 0xc1, 0x4c, 0xa3,
	0x73, 0x52, 0x21, 0x8b, 0x62, 0x92, 0xce, 0xc5, 0x08, 0x97, 0x8f, 0x70, 0xfe, 0x4b, 0x44, 0x2f,
	0x20, 0xfa, 0xc0, 0x5d, 0x08, 0x32, 0x17, 0xdd, 0x48, 0x2f, 0xe1, 0xa4, 0x95, 0x65, 0x83, 0x6c,
	0x12, 0x76, 0x3d, 0x78, 0x98, 0xdc, 0x93, 0xe4, 0x09, 0x16, 0x02, 0x65, 0x2e, 0xf0, 0xb3, 0x41,
	0xe7, 0xbb, 0x57, 0x1c, 0xda, 0xb6, 0xc8, 0x70, 0x90, 0x8f, 0xb0, 0xbb, 0x69, 0xd1, 0xba, 0xc2,
	0x54, 0x83, 0xc9, 0x08, 0x93, 0x0d, 0x9c, 0xf5, 0x16, 0xae, 0x36, 0x95, 0x43, 0x7a, 0x07, 0x33,
	0x1b, 0xba, 0x38, 0x46, 0x42, 0xdd, 0xe5, 0xe1, 0xba, 0x62, 0xa4, 0xae, 0x05, 0x44, 0xaf, 0x46,
	0xd1, 0x17, 0x38, 0xee, 0xcc, 0xe8, 0xf5, 0xff, 0x9a, 0x7d, 0xd2, 0x65, 0x7c, 0x98, 0xd0, 0xe7,
	0x48, 0x8e, 0xb6, 0xd3, 0xf0, 0x7b, 0xb7, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x33, 0xdc, 0x44,
	0xa4, 0xf1, 0x01, 0x00, 0x00,
}
