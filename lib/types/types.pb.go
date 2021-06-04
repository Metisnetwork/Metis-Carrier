// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/types/types.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
	io "io"
	math "math"
	math_bits "math/bits"
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

// MetaData body data struct.
type BlockData struct {
	Header       Header         `protobuf:"bytes,1,opt,name=header,proto3" json:"header" xml:"header"`
	Metadata     []MetaData     `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata" xml:"metadata"`
	Resourcedata []ResourceData `protobuf:"bytes,3,rep,name=resourcedata,proto3" json:"resourcedata" xml:"resourcedata"`
	Identitydata []IdentityData `protobuf:"bytes,4,rep,name=identitydata,proto3" json:"identitydata" xml:"identitydata"`
	Taskdata     []TaskData     `protobuf:"bytes,5,rep,name=taskdata,proto3" json:"taskdata" xml:"taskdata"`
	//
	ReceivedAt   uint64 `protobuf:"varint,6,opt,name=receivedAt,proto3" json:"receivedat" xml:"receivedat"`
	ReceivedFrom string `protobuf:"bytes,7,opt,name=receivedFrom,proto3" json:"receivedfrom" xml:"receivedfrom"`
}

func (m *BlockData) Reset()         { *m = BlockData{} }
func (m *BlockData) String() string { return proto.CompactTextString(m) }
func (*BlockData) ProtoMessage()    {}
func (*BlockData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc2b82598ae74bd, []int{0}
}
func (m *BlockData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockData.Merge(m, src)
}
func (m *BlockData) XXX_Size() int {
	return m.ProtoSize()
}
func (m *BlockData) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockData.DiscardUnknown(m)
}

var xxx_messageInfo_BlockData proto.InternalMessageInfo

type BodyData struct {
	Metadata     []MetaData     `protobuf:"bytes,1,rep,name=metadata,proto3" json:"metadata" xml:"metadata"`
	Resourcedata []ResourceData `protobuf:"bytes,2,rep,name=resourcedata,proto3" json:"resourcedata" xml:"resourcedata"`
	Identitydata []IdentityData `protobuf:"bytes,3,rep,name=identitydata,proto3" json:"identitydata" xml:"identitydata"`
	Taskdata     []TaskData     `protobuf:"bytes,4,rep,name=taskdata,proto3" json:"taskdata" xml:"taskdata"`
	ExtraData    []byte         `protobuf:"bytes,5,opt,name=extraData,proto3" json:"extradata" xml:"extradata"`
}

func (m *BodyData) Reset()         { *m = BodyData{} }
func (m *BodyData) String() string { return proto.CompactTextString(m) }
func (*BodyData) ProtoMessage()    {}
func (*BodyData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc2b82598ae74bd, []int{1}
}
func (m *BodyData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BodyData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BodyData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BodyData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BodyData.Merge(m, src)
}
func (m *BodyData) XXX_Size() int {
	return m.ProtoSize()
}
func (m *BodyData) XXX_DiscardUnknown() {
	xxx_messageInfo_BodyData.DiscardUnknown(m)
}

var xxx_messageInfo_BodyData proto.InternalMessageInfo

// MetadataLookupEntry is a positional metadata to help looking up the data content of
// a metadata given only its dataId.
type DataLookupEntry struct {
	BlockHash  []byte `protobuf:"bytes,1,opt,name=blockHash,proto3" json:"blockhash" xml:"blockhash"`
	BlockIndex uint64 `protobuf:"varint,2,opt,name=blockIndex,proto3" json:"blockindex" xml:"blockindex"`
	Index      uint64 `protobuf:"varint,3,opt,name=index,proto3" json:"index" xml:"index"`
	NodeId     string `protobuf:"bytes,4,opt,name=nodeId,proto3" json:"nodeid" xml:"nodeid"`
	Type       string `protobuf:"bytes,5,opt,name=type,proto3" json:"type" xml:"type"`
}

func (m *DataLookupEntry) Reset()         { *m = DataLookupEntry{} }
func (m *DataLookupEntry) String() string { return proto.CompactTextString(m) }
func (*DataLookupEntry) ProtoMessage()    {}
func (*DataLookupEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_7dc2b82598ae74bd, []int{2}
}
func (m *DataLookupEntry) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DataLookupEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DataLookupEntry.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DataLookupEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataLookupEntry.Merge(m, src)
}
func (m *DataLookupEntry) XXX_Size() int {
	return m.ProtoSize()
}
func (m *DataLookupEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_DataLookupEntry.DiscardUnknown(m)
}

var xxx_messageInfo_DataLookupEntry proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BlockData)(nil), "types.BlockData")
	proto.RegisterType((*BodyData)(nil), "types.BodyData")
	proto.RegisterType((*DataLookupEntry)(nil), "types.DataLookupEntry")
}

func init() { proto.RegisterFile("lib/types/types.proto", fileDescriptor_7dc2b82598ae74bd) }

var fileDescriptor_7dc2b82598ae74bd = []byte{
	// 652 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0x4f, 0x6b, 0x13, 0x41,
	0x18, 0xc6, 0xb3, 0xf9, 0x67, 0x33, 0xad, 0x56, 0x56, 0x2c, 0x63, 0x91, 0xdd, 0x30, 0x82, 0x84,
	0x8a, 0x09, 0xd4, 0x9b, 0x07, 0xc1, 0xd5, 0xd6, 0x56, 0x2c, 0xc2, 0xe0, 0xc9, 0xdb, 0x24, 0x3b,
	0x4d, 0x96, 0xfc, 0x99, 0x30, 0x3b, 0xd1, 0xe6, 0x5b, 0xf8, 0x11, 0xbc, 0xf8, 0x15, 0xfc, 0x0c,
	0x3d, 0x49, 0x8f, 0x9e, 0x16, 0xda, 0xdc, 0x72, 0xcc, 0x27, 0x90, 0x7d, 0x67, 0x76, 0x76, 0x17,
	0x3c, 0x5a, 0x2f, 0x61, 0xdf, 0xdf, 0xfb, 0x3e, 0xcf, 0x30, 0x4f, 0x5e, 0x06, 0x3d, 0x9c, 0x44,
	0xfd, 0x9e, 0x5a, 0xce, 0x79, 0xac, 0x7f, 0xbb, 0x73, 0x29, 0x94, 0x70, 0x1b, 0x50, 0xec, 0xef,
	0xe5, 0xdd, 0x11, 0x67, 0x21, 0x97, 0xba, 0xbd, 0x8f, 0x73, 0x3e, 0xe5, 0x8a, 0x85, 0x4c, 0x31,
	0xd3, 0x79, 0x9c, 0x77, 0x24, 0x8f, 0xc5, 0x42, 0x0e, 0x78, 0xa1, 0x5b, 0xd0, 0x29, 0x16, 0x8f,
	0xff, 0xae, 0x8b, 0x42, 0x3e, 0x53, 0x91, 0x5a, 0x16, 0xba, 0x4f, 0x24, 0x9f, 0x8b, 0xb8, 0x07,
	0x45, 0x7f, 0x71, 0xde, 0x1b, 0x8a, 0xa1, 0x80, 0x02, 0xbe, 0xcc, 0xd0, 0xa3, 0xa1, 0x10, 0xc3,
	0x09, 0xcf, 0xa7, 0xd8, 0x6c, 0xa9, 0x5b, 0xe4, 0x57, 0x1d, 0xb5, 0x82, 0x89, 0x18, 0x8c, 0xdf,
	0x32, 0xc5, 0xdc, 0x23, 0xd4, 0xd4, 0xb7, 0xc1, 0x4e, 0xdb, 0xe9, 0x6c, 0x1f, 0xde, 0xed, 0xea,
	0xab, 0x9f, 0x00, 0x0c, 0xbc, 0xcb, 0xc4, 0xaf, 0xac, 0x13, 0xdf, 0x0c, 0x6d, 0x12, 0x7f, 0xe7,
	0x62, 0x3a, 0x79, 0x49, 0x74, 0x49, 0xa8, 0xe1, 0xee, 0x47, 0xb4, 0x95, 0x5d, 0x1e, 0x57, 0xdb,
	0xb5, 0xce, 0xf6, 0xe1, 0xae, 0x31, 0x3a, 0xe3, 0x8a, 0xa5, 0x27, 0x05, 0xc4, 0x58, 0xd9, 0xc1,
	0x4d, 0xe2, 0xdf, 0x03, 0xb3, 0x0c, 0x10, 0x6a, 0x7b, 0xee, 0x00, 0xed, 0x14, 0x33, 0xc3, 0x35,
	0x30, 0x7d, 0x60, 0x4c, 0xa9, 0x69, 0x81, 0xf1, 0x81, 0x31, 0x2e, 0x09, 0x36, 0x89, 0xef, 0x82,
	0x79, 0x11, 0x12, 0x5a, 0x9a, 0x49, 0x0f, 0x29, 0x06, 0x8c, 0xeb, 0xa5, 0x43, 0x4e, 0x4d, 0xab,
	0x7c, 0x48, 0x51, 0x60, 0x0f, 0x29, 0x42, 0x42, 0x4b, 0x33, 0x69, 0x34, 0xd9, 0xff, 0x8b, 0x1b,
	0xa5, 0x68, 0x3e, 0xb1, 0x78, 0x5c, 0x8e, 0x26, 0x1b, 0xb4, 0xd1, 0x64, 0x80, 0x50, 0xdb, 0x73,
	0x03, 0x84, 0x24, 0x1f, 0xf0, 0xe8, 0x0b, 0x0f, 0x5f, 0x2b, 0xdc, 0x6c, 0x3b, 0x9d, 0x7a, 0x40,
	0xd6, 0x89, 0x6f, 0x29, 0x53, 0x9b, 0xc4, 0xbf, 0x6f, 0x6e, 0x9f, 0x21, 0x42, 0x0b, 0x2a, 0xf7,
	0x7d, 0x1a, 0xaf, 0xae, 0x8e, 0xa5, 0x98, 0xe2, 0x3b, 0x6d, 0xa7, 0xd3, 0x0a, 0x9e, 0xea, 0x14,
	0x35, 0x3f, 0x97, 0x62, 0x5a, 0x48, 0x31, 0x87, 0x90, 0x62, 0xae, 0x25, 0x3f, 0x6b, 0x68, 0x2b,
	0x10, 0x21, 0xe4, 0x54, 0x5a, 0x04, 0xe7, 0x36, 0x16, 0xa1, 0xfa, 0x3f, 0x16, 0xa1, 0x76, 0xdb,
	0x8b, 0x50, 0xff, 0x17, 0x8b, 0xf0, 0x0a, 0xb5, 0xf8, 0x85, 0x92, 0x90, 0x2a, 0x6e, 0xb4, 0x9d,
	0xce, 0x4e, 0xd0, 0x5e, 0x27, 0xbe, 0x86, 0x46, 0xbd, 0x0b, 0x6a, 0x4b, 0x08, 0xcd, 0x25, 0xe4,
	0x47, 0x15, 0xed, 0xa6, 0x1f, 0x1f, 0x84, 0x18, 0x2f, 0xe6, 0x47, 0x33, 0x25, 0x97, 0xa9, 0x67,
	0x3f, 0x7d, 0x1c, 0x4e, 0x58, 0x3c, 0x82, 0x27, 0xc1, 0x78, 0x02, 0x1c, 0xb1, 0x78, 0x64, 0x3d,
	0x2d, 0x21, 0x34, 0x97, 0xa4, 0xcb, 0x09, 0xc5, 0xe9, 0x2c, 0xe4, 0x17, 0xb8, 0x9a, 0x2f, 0x27,
	0xd0, 0x28, 0xa5, 0x76, 0x39, 0x73, 0x44, 0x68, 0x41, 0xe5, 0x76, 0x51, 0x03, 0x28, 0xae, 0x81,
	0x1c, 0xaf, 0x13, 0xbf, 0x91, 0x29, 0xb7, 0x75, 0xcc, 0x5a, 0xa4, 0xa9, 0x7b, 0x88, 0x9a, 0x33,
	0x11, 0xf2, 0xd3, 0x10, 0xd7, 0x61, 0x8d, 0xf7, 0xd3, 0x07, 0x2b, 0x25, 0x51, 0x68, 0x1f, 0x2c,
	0x5d, 0x12, 0x6a, 0x26, 0xdd, 0x03, 0x54, 0x4f, 0xb3, 0x87, 0xd8, 0x5a, 0xc1, 0xde, 0x3a, 0xf1,
	0xa1, 0xde, 0x24, 0x3e, 0xd2, 0x79, 0x2f, 0xe7, 0x9c, 0x50, 0x60, 0xc1, 0xd9, 0xe5, 0xb5, 0x57,
	0xb9, 0xba, 0xf6, 0x2a, 0x97, 0x37, 0x9e, 0x73, 0x75, 0xe3, 0x39, 0xdf, 0x56, 0x5e, 0xe5, 0xfb,
	0xca, 0x73, 0xae, 0x56, 0x5e, 0xe5, 0xf7, 0xca, 0xab, 0x7c, 0x7e, 0x36, 0x8c, 0xd4, 0x68, 0xd1,
	0xef, 0x0e, 0xc4, 0xb4, 0x47, 0x45, 0xcc, 0x95, 0x62, 0xc7, 0x13, 0xf1, 0xb5, 0xf7, 0x86, 0x49,
	0x19, 0x71, 0xf9, 0xfc, 0x9d, 0xe8, 0xd9, 0x37, 0xbd, 0xdf, 0x84, 0x77, 0xf8, 0xc5, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xbb, 0xa8, 0xad, 0xf9, 0x6f, 0x06, 0x00, 0x00,
}

func (m *BlockData) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockData) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReceivedFrom) > 0 {
		i -= len(m.ReceivedFrom)
		copy(dAtA[i:], m.ReceivedFrom)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ReceivedFrom)))
		i--
		dAtA[i] = 0x3a
	}
	if m.ReceivedAt != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.ReceivedAt))
		i--
		dAtA[i] = 0x30
	}
	if len(m.Taskdata) > 0 {
		for iNdEx := len(m.Taskdata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Taskdata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Identitydata) > 0 {
		for iNdEx := len(m.Identitydata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Identitydata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Resourcedata) > 0 {
		for iNdEx := len(m.Resourcedata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Resourcedata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Metadata) > 0 {
		for iNdEx := len(m.Metadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Metadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *BodyData) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BodyData) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BodyData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExtraData) > 0 {
		i -= len(m.ExtraData)
		copy(dAtA[i:], m.ExtraData)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ExtraData)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Taskdata) > 0 {
		for iNdEx := len(m.Taskdata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Taskdata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Identitydata) > 0 {
		for iNdEx := len(m.Identitydata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Identitydata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Resourcedata) > 0 {
		for iNdEx := len(m.Resourcedata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Resourcedata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Metadata) > 0 {
		for iNdEx := len(m.Metadata) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Metadata[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *DataLookupEntry) Marshal() (dAtA []byte, err error) {
	size := m.ProtoSize()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DataLookupEntry) MarshalTo(dAtA []byte) (int, error) {
	size := m.ProtoSize()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DataLookupEntry) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.NodeId) > 0 {
		i -= len(m.NodeId)
		copy(dAtA[i:], m.NodeId)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NodeId)))
		i--
		dAtA[i] = 0x22
	}
	if m.Index != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x18
	}
	if m.BlockIndex != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.BlockIndex))
		i--
		dAtA[i] = 0x10
	}
	if len(m.BlockHash) > 0 {
		i -= len(m.BlockHash)
		copy(dAtA[i:], m.BlockHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.BlockHash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockData) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Header.ProtoSize()
	n += 1 + l + sovTypes(uint64(l))
	if len(m.Metadata) > 0 {
		for _, e := range m.Metadata {
			l = e.ProtoSize()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Resourcedata) > 0 {
		for _, e := range m.Resourcedata {
			l = e.ProtoSize()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Identitydata) > 0 {
		for _, e := range m.Identitydata {
			l = e.ProtoSize()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Taskdata) > 0 {
		for _, e := range m.Taskdata {
			l = e.ProtoSize()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if m.ReceivedAt != 0 {
		n += 1 + sovTypes(uint64(m.ReceivedAt))
	}
	l = len(m.ReceivedFrom)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *BodyData) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Metadata) > 0 {
		for _, e := range m.Metadata {
			l = e.ProtoSize()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Resourcedata) > 0 {
		for _, e := range m.Resourcedata {
			l = e.ProtoSize()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Identitydata) > 0 {
		for _, e := range m.Identitydata {
			l = e.ProtoSize()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	if len(m.Taskdata) > 0 {
		for _, e := range m.Taskdata {
			l = e.ProtoSize()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	l = len(m.ExtraData)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *DataLookupEntry) ProtoSize() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BlockHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.BlockIndex != 0 {
		n += 1 + sovTypes(uint64(m.BlockIndex))
	}
	if m.Index != 0 {
		n += 1 + sovTypes(uint64(m.Index))
	}
	l = len(m.NodeId)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BlockData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = append(m.Metadata, MetaData{})
			if err := m.Metadata[len(m.Metadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resourcedata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resourcedata = append(m.Resourcedata, ResourceData{})
			if err := m.Resourcedata[len(m.Resourcedata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identitydata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identitydata = append(m.Identitydata, IdentityData{})
			if err := m.Identitydata[len(m.Identitydata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Taskdata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Taskdata = append(m.Taskdata, TaskData{})
			if err := m.Taskdata[len(m.Taskdata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedAt", wireType)
			}
			m.ReceivedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReceivedAt |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReceivedFrom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReceivedFrom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BodyData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BodyData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BodyData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = append(m.Metadata, MetaData{})
			if err := m.Metadata[len(m.Metadata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resourcedata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resourcedata = append(m.Resourcedata, ResourceData{})
			if err := m.Resourcedata[len(m.Resourcedata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identitydata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identitydata = append(m.Identitydata, IdentityData{})
			if err := m.Identitydata[len(m.Identitydata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Taskdata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Taskdata = append(m.Taskdata, TaskData{})
			if err := m.Taskdata[len(m.Taskdata)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtraData = append(m.ExtraData[:0], dAtA[iNdEx:postIndex]...)
			if m.ExtraData == nil {
				m.ExtraData = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DataLookupEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DataLookupEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DataLookupEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BlockHash = append(m.BlockHash[:0], dAtA[iNdEx:postIndex]...)
			if m.BlockHash == nil {
				m.BlockHash = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockIndex", wireType)
			}
			m.BlockIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
