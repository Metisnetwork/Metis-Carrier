// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fighter/api/io_channel.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SendRequest struct {
	Nodeid               string   `protobuf:"bytes,1,opt,name=nodeid,proto3" json:"nodeid,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Data                 []byte   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendRequest) Reset()         { *m = SendRequest{} }
func (m *SendRequest) String() string { return proto.CompactTextString(m) }
func (*SendRequest) ProtoMessage()    {}
func (*SendRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_585c4a2c48f31701, []int{0}
}
func (m *SendRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendRequest.Merge(m, src)
}
func (m *SendRequest) XXX_Size() int {
	return m.Size()
}
func (m *SendRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SendRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SendRequest proto.InternalMessageInfo

func (m *SendRequest) GetNodeid() string {
	if m != nil {
		return m.Nodeid
	}
	return ""
}

func (m *SendRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SendRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RetCode struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RetCode) Reset()         { *m = RetCode{} }
func (m *RetCode) String() string { return proto.CompactTextString(m) }
func (*RetCode) ProtoMessage()    {}
func (*RetCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_585c4a2c48f31701, []int{1}
}
func (m *RetCode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RetCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RetCode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RetCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RetCode.Merge(m, src)
}
func (m *RetCode) XXX_Size() int {
	return m.Size()
}
func (m *RetCode) XXX_DiscardUnknown() {
	xxx_messageInfo_RetCode.DiscardUnknown(m)
}

var xxx_messageInfo_RetCode proto.InternalMessageInfo

func (m *RetCode) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func init() {
	proto.RegisterType((*SendRequest)(nil), "fighter.api.SendRequest")
	proto.RegisterType((*RetCode)(nil), "fighter.api.RetCode")
}

func init() { proto.RegisterFile("fighter/api/io_channel.proto", fileDescriptor_585c4a2c48f31701) }

var fileDescriptor_585c4a2c48f31701 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x14, 0x45, 0x71, 0x28, 0x45, 0x75, 0x11, 0x83, 0x85, 0x50, 0x84, 0x20, 0xaa, 0x32, 0x75, 0xa9,
	0x2d, 0xc1, 0x82, 0x18, 0xa9, 0x18, 0x2a, 0x36, 0xb3, 0xb1, 0x20, 0xc7, 0x7e, 0x34, 0x16, 0xd4,
	0x2f, 0xb8, 0x2f, 0xe2, 0x17, 0x19, 0xf9, 0x04, 0x94, 0xaf, 0x60, 0x44, 0x75, 0x33, 0x84, 0xed,
	0xda, 0x47, 0xba, 0xef, 0xe8, 0xf2, 0xcb, 0x57, 0xbf, 0xae, 0x09, 0xa2, 0x32, 0x8d, 0x57, 0x1e,
	0x5f, 0x6c, 0x6d, 0x42, 0x80, 0x77, 0xd9, 0x44, 0x24, 0x14, 0xd3, 0x9e, 0x4a, 0xd3, 0xf8, 0x72,
	0xc5, 0xa7, 0x4f, 0x10, 0x9c, 0x86, 0x8f, 0x16, 0xb6, 0x24, 0xce, 0xf9, 0x38, 0xa0, 0x03, 0xef,
	0x72, 0x36, 0x63, 0xf3, 0x89, 0xee, 0x5f, 0xe2, 0x94, 0x67, 0xde, 0xe5, 0x59, 0xfa, 0xcb, 0xbc,
	0x13, 0x82, 0x8f, 0x9c, 0x21, 0x93, 0x1f, 0xce, 0xd8, 0xfc, 0x44, 0xa7, 0x5c, 0x5e, 0xf1, 0x63,
	0x0d, 0xb4, 0x44, 0x07, 0x3b, 0x6c, 0xd1, 0x41, 0x2a, 0x39, 0xd2, 0x29, 0x5f, 0x3f, 0xf0, 0xc9,
	0x0a, 0x97, 0x7b, 0x13, 0x71, 0xcb, 0x47, 0xbb, 0xb3, 0x22, 0x97, 0x03, 0x19, 0x39, 0x30, 0xb9,
	0x38, 0xfb, 0x47, 0xfa, 0xe2, 0xf2, 0xe0, 0xfe, 0xf1, 0xab, 0x2b, 0xd8, 0x77, 0x57, 0xb0, 0x9f,
	0xae, 0x60, 0xcf, 0x77, 0x6b, 0x4f, 0x75, 0x5b, 0x49, 0x8b, 0x1b, 0xe5, 0x0c, 0xb5, 0x1b, 0x02,
	0x5b, 0x6f, 0xf7, 0x71, 0x11, 0x80, 0x3e, 0x31, 0xbe, 0x2d, 0xac, 0x89, 0xd1, 0x43, 0x54, 0x4d,
	0xa5, 0x06, 0xc3, 0xfc, 0x32, 0x56, 0x8d, 0xd3, 0x22, 0x37, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x57, 0x18, 0x2a, 0x73, 0x31, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// IoChannelClient is the client API for IoChannel service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type IoChannelClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*RetCode, error)
}

type ioChannelClient struct {
	cc *grpc.ClientConn
}

func NewIoChannelClient(cc *grpc.ClientConn) IoChannelClient {
	return &ioChannelClient{cc}
}

func (c *ioChannelClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*RetCode, error) {
	out := new(RetCode)
	err := c.cc.Invoke(ctx, "/fighter.api.IoChannel/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IoChannelServer is the server API for IoChannel service.
type IoChannelServer interface {
	Send(context.Context, *SendRequest) (*RetCode, error)
}

// UnimplementedIoChannelServer can be embedded to have forward compatible implementations.
type UnimplementedIoChannelServer struct {
}

func (*UnimplementedIoChannelServer) Send(ctx context.Context, req *SendRequest) (*RetCode, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}

func RegisterIoChannelServer(s *grpc.Server, srv IoChannelServer) {
	s.RegisterService(&_IoChannel_serviceDesc, srv)
}

func _IoChannel_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IoChannelServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fighter.api.IoChannel/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IoChannelServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IoChannel_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fighter.api.IoChannel",
	HandlerType: (*IoChannelServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _IoChannel_Send_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fighter/api/io_channel.proto",
}

func (m *SendRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintIoChannel(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintIoChannel(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Nodeid) > 0 {
		i -= len(m.Nodeid)
		copy(dAtA[i:], m.Nodeid)
		i = encodeVarintIoChannel(dAtA, i, uint64(len(m.Nodeid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RetCode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RetCode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RetCode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Code != 0 {
		i = encodeVarintIoChannel(dAtA, i, uint64(m.Code))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintIoChannel(dAtA []byte, offset int, v uint64) int {
	offset -= sovIoChannel(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SendRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Nodeid)
	if l > 0 {
		n += 1 + l + sovIoChannel(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovIoChannel(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovIoChannel(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *RetCode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Code != 0 {
		n += 1 + sovIoChannel(uint64(m.Code))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovIoChannel(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIoChannel(x uint64) (n int) {
	return sovIoChannel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SendRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIoChannel
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
			return fmt.Errorf("proto: SendRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodeid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoChannel
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
				return ErrInvalidLengthIoChannel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIoChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodeid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoChannel
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
				return ErrInvalidLengthIoChannel
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIoChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoChannel
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
				return ErrInvalidLengthIoChannel
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIoChannel
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data[:0], dAtA[iNdEx:postIndex]...)
			if m.Data == nil {
				m.Data = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIoChannel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIoChannel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *RetCode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIoChannel
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
			return fmt.Errorf("proto: RetCode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RetCode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Code", wireType)
			}
			m.Code = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIoChannel
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Code |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipIoChannel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIoChannel
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipIoChannel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIoChannel
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
					return 0, ErrIntOverflowIoChannel
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
					return 0, ErrIntOverflowIoChannel
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
				return 0, ErrInvalidLengthIoChannel
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIoChannel
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIoChannel
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIoChannel        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIoChannel          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIoChannel = fmt.Errorf("proto: unexpected end of group")
)
