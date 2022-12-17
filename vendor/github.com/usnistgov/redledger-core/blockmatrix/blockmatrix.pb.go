// Code generated by protoc-gen-go. DO NOT EDIT.
// source: blockmatrix.proto

package blockmatrix

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

// Contains information about the block matrix ledger such as size, block count, and row/column hashes.
type Info struct {
	Size                 uint64   `protobuf:"varint,1,opt,name=size,proto3" json:"size,omitempty"`
	BlockCount           uint64   `protobuf:"varint,2,opt,name=blockCount,proto3" json:"blockCount,omitempty"`
	RowHashes            [][]byte `protobuf:"bytes,3,rep,name=rowHashes,proto3" json:"rowHashes,omitempty"`
	ColumnHashes         [][]byte `protobuf:"bytes,4,rep,name=columnHashes,proto3" json:"columnHashes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Info) Reset()         { *m = Info{} }
func (m *Info) String() string { return proto.CompactTextString(m) }
func (*Info) ProtoMessage()    {}
func (*Info) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05697675af98205, []int{0}
}

func (m *Info) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Info.Unmarshal(m, b)
}
func (m *Info) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Info.Marshal(b, m, deterministic)
}
func (m *Info) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Info.Merge(m, src)
}
func (m *Info) XXX_Size() int {
	return xxx_messageInfo_Info.Size(m)
}
func (m *Info) XXX_DiscardUnknown() {
	xxx_messageInfo_Info.DiscardUnknown(m)
}

var xxx_messageInfo_Info proto.InternalMessageInfo

func (m *Info) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *Info) GetBlockCount() uint64 {
	if m != nil {
		return m.BlockCount
	}
	return 0
}

func (m *Info) GetRowHashes() [][]byte {
	if m != nil {
		return m.RowHashes
	}
	return nil
}

func (m *Info) GetColumnHashes() [][]byte {
	if m != nil {
		return m.ColumnHashes
	}
	return nil
}

type LedgerType struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LedgerType) Reset()         { *m = LedgerType{} }
func (m *LedgerType) String() string { return proto.CompactTextString(m) }
func (*LedgerType) ProtoMessage()    {}
func (*LedgerType) Descriptor() ([]byte, []int) {
	return fileDescriptor_f05697675af98205, []int{1}
}

func (m *LedgerType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LedgerType.Unmarshal(m, b)
}
func (m *LedgerType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LedgerType.Marshal(b, m, deterministic)
}
func (m *LedgerType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LedgerType.Merge(m, src)
}
func (m *LedgerType) XXX_Size() int {
	return xxx_messageInfo_LedgerType.Size(m)
}
func (m *LedgerType) XXX_DiscardUnknown() {
	xxx_messageInfo_LedgerType.DiscardUnknown(m)
}

var xxx_messageInfo_LedgerType proto.InternalMessageInfo

func (m *LedgerType) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func init() {
	proto.RegisterType((*Info)(nil), "blockmatrix.Info")
	proto.RegisterType((*LedgerType)(nil), "blockmatrix.LedgerType")
}

func init() { proto.RegisterFile("blockmatrix.proto", fileDescriptor_f05697675af98205) }

var fileDescriptor_f05697675af98205 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xca, 0xc9, 0x4f,
	0xce, 0xce, 0x4d, 0x2c, 0x29, 0xca, 0xac, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x46,
	0x12, 0x52, 0xaa, 0xe1, 0x62, 0xf1, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xe2, 0x62, 0x29, 0xce, 0xac,
	0x4a, 0x95, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x09, 0x02, 0xb3, 0x85, 0xe4, 0xb8, 0xb8, 0xc0, 0x4a,
	0x9d, 0xf3, 0x4b, 0xf3, 0x4a, 0x24, 0x98, 0xc0, 0x32, 0x48, 0x22, 0x42, 0x32, 0x5c, 0x9c, 0x45,
	0xf9, 0xe5, 0x1e, 0x89, 0xc5, 0x19, 0xa9, 0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0x3c, 0x41, 0x08,
	0x01, 0x21, 0x25, 0x2e, 0x9e, 0xe4, 0xfc, 0x9c, 0xd2, 0xdc, 0x3c, 0xa8, 0x02, 0x16, 0xb0, 0x02,
	0x14, 0x31, 0x25, 0x05, 0x2e, 0x2e, 0x9f, 0xd4, 0x94, 0xf4, 0xd4, 0xa2, 0x90, 0xca, 0x82, 0x54,
	0x90, 0x1b, 0x4a, 0x2a, 0x0b, 0x20, 0x6e, 0xe0, 0x0c, 0x02, 0xb3, 0x9d, 0x2c, 0xa3, 0xcc, 0xd3,
	0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0x33, 0x2a, 0x0b, 0x52, 0x8b, 0x72,
	0xc0, 0x3a, 0xf4, 0xd3, 0x12, 0x93, 0x8a, 0x32, 0x93, 0xf5, 0x93, 0xf3, 0x73, 0x73, 0xf3, 0xf3,
	0xf4, 0xa1, 0x82, 0x48, 0x5e, 0x4b, 0x62, 0x03, 0x7b, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0x65, 0x49, 0x41, 0x2c, 0x03, 0x01, 0x00, 0x00,
}
