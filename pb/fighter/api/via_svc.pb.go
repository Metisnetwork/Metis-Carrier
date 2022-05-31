// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fighter/api/via_svc.proto

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

type ServiceType int32

const (
	ServiceType_DATA_SVC     ServiceType = 0
	ServiceType_COMPUTE_SVC  ServiceType = 1
	ServiceType_SCHEDULE_SVC ServiceType = 2
	ServiceType_NET_COMM_SVC ServiceType = 3
)

var ServiceType_name = map[int32]string{
	0: "DATA_SVC",
	1: "COMPUTE_SVC",
	2: "SCHEDULE_SVC",
	3: "NET_COMM_SVC",
}

var ServiceType_value = map[string]int32{
	"DATA_SVC":     0,
	"COMPUTE_SVC":  1,
	"SCHEDULE_SVC": 2,
	"NET_COMM_SVC": 3,
}

func (x ServiceType) String() string {
	return proto.EnumName(ServiceType_name, int32(x))
}

func (ServiceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b884649ae566ccd8, []int{0}
}

type ExposeReq struct {
	TaskId               string      `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	SvcType              ServiceType `protobuf:"varint,2,opt,name=svc_type,json=svcType,proto3,enum=fighter.api.ServiceType" json:"svc_type,omitempty"`
	PartyId              string      `protobuf:"bytes,3,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Ip                   string      `protobuf:"bytes,4,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32       `protobuf:"varint,5,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ExposeReq) Reset()         { *m = ExposeReq{} }
func (m *ExposeReq) String() string { return proto.CompactTextString(m) }
func (*ExposeReq) ProtoMessage()    {}
func (*ExposeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_b884649ae566ccd8, []int{0}
}
func (m *ExposeReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExposeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExposeReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExposeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExposeReq.Merge(m, src)
}
func (m *ExposeReq) XXX_Size() int {
	return m.Size()
}
func (m *ExposeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ExposeReq.DiscardUnknown(m)
}

var xxx_messageInfo_ExposeReq proto.InternalMessageInfo

func (m *ExposeReq) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *ExposeReq) GetSvcType() ServiceType {
	if m != nil {
		return m.SvcType
	}
	return ServiceType_DATA_SVC
}

func (m *ExposeReq) GetPartyId() string {
	if m != nil {
		return m.PartyId
	}
	return ""
}

func (m *ExposeReq) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ExposeReq) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ExposeAns struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	Ip                   string   `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Port                 int32    `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExposeAns) Reset()         { *m = ExposeAns{} }
func (m *ExposeAns) String() string { return proto.CompactTextString(m) }
func (*ExposeAns) ProtoMessage()    {}
func (*ExposeAns) Descriptor() ([]byte, []int) {
	return fileDescriptor_b884649ae566ccd8, []int{1}
}
func (m *ExposeAns) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ExposeAns) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ExposeAns.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ExposeAns) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExposeAns.Merge(m, src)
}
func (m *ExposeAns) XXX_Size() int {
	return m.Size()
}
func (m *ExposeAns) XXX_DiscardUnknown() {
	xxx_messageInfo_ExposeAns.DiscardUnknown(m)
}

var xxx_messageInfo_ExposeAns proto.InternalMessageInfo

func (m *ExposeAns) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func (m *ExposeAns) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ExposeAns) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterEnum("fighter.api.ServiceType", ServiceType_name, ServiceType_value)
	proto.RegisterType((*ExposeReq)(nil), "fighter.api.ExposeReq")
	proto.RegisterType((*ExposeAns)(nil), "fighter.api.ExposeAns")
}

func init() { proto.RegisterFile("fighter/api/via_svc.proto", fileDescriptor_b884649ae566ccd8) }

var fileDescriptor_b884649ae566ccd8 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0xaa, 0xda, 0x40,
	0x14, 0x86, 0xef, 0x24, 0xf7, 0x6a, 0xee, 0x44, 0xac, 0xcc, 0xa2, 0x8d, 0x5d, 0x88, 0xb8, 0x92,
	0x82, 0x09, 0xe8, 0xa2, 0xe0, 0xa6, 0xa4, 0x31, 0xa0, 0x50, 0xab, 0xc4, 0xe8, 0xa2, 0x9b, 0x30,
	0x26, 0xa3, 0x0e, 0x41, 0x67, 0x3a, 0x33, 0xa6, 0x75, 0xdb, 0xa7, 0xe8, 0x23, 0x75, 0xd9, 0x47,
	0x28, 0x3e, 0x49, 0xc9, 0x68, 0x5b, 0x4b, 0xbb, 0xb9, 0xbb, 0x7f, 0xbe, 0xe1, 0x7c, 0xfc, 0x1c,
	0x0e, 0x6c, 0x6e, 0xe8, 0x76, 0xa7, 0x88, 0xf0, 0x30, 0xa7, 0x5e, 0x41, 0x71, 0x22, 0x8b, 0xd4,
	0xe5, 0x82, 0x29, 0x86, 0xec, 0xeb, 0x97, 0x8b, 0x39, 0xed, 0x7c, 0x05, 0xf0, 0x31, 0xfc, 0xcc,
	0x99, 0x24, 0x11, 0xf9, 0x88, 0x5e, 0xc0, 0xaa, 0xc2, 0x32, 0x4f, 0x68, 0xe6, 0x80, 0x36, 0xe8,
	0x3e, 0x46, 0x95, 0xf2, 0x39, 0xc9, 0xd0, 0x00, 0x5a, 0xb2, 0x48, 0x13, 0x75, 0xe2, 0xc4, 0x31,
	0xda, 0xa0, 0x5b, 0xef, 0x3b, 0xee, 0x8d, 0xc6, 0x5d, 0x10, 0x51, 0xd0, 0x94, 0xc4, 0x27, 0x4e,
	0xa2, 0xaa, 0x2c, 0xd2, 0x32, 0xa0, 0x26, 0xb4, 0x38, 0x16, 0xea, 0x54, 0xea, 0x4c, 0xad, 0xab,
	0xea, 0xf7, 0x24, 0x43, 0x75, 0x68, 0x50, 0xee, 0xdc, 0x6b, 0x68, 0x50, 0x8e, 0x10, 0xbc, 0xe7,
	0x4c, 0x28, 0xe7, 0xa1, 0x0d, 0xba, 0x0f, 0x91, 0xce, 0x9d, 0x37, 0xbf, 0x9a, 0xf9, 0x07, 0x59,
	0x0e, 0xb0, 0x5c, 0x97, 0xb2, 0x22, 0x83, 0xe5, 0x57, 0x81, 0xf1, 0x8f, 0xc0, 0xfc, 0x23, 0x78,
	0x35, 0x87, 0xf6, 0x4d, 0x2f, 0x54, 0x83, 0xd6, 0xc8, 0x8f, 0xfd, 0x64, 0xb1, 0x0a, 0x1a, 0x77,
	0xe8, 0x19, 0xb4, 0x83, 0xd9, 0x74, 0xbe, 0x8c, 0x43, 0x0d, 0x00, 0x6a, 0xc0, 0xda, 0x22, 0x18,
	0x87, 0xa3, 0xe5, 0xbb, 0x0b, 0x31, 0x4a, 0xf2, 0x3e, 0x8c, 0x93, 0x60, 0x36, 0x9d, 0x6a, 0x62,
	0xf6, 0xbf, 0x00, 0x68, 0xaf, 0x28, 0x9e, 0x0b, 0x56, 0xd0, 0x8c, 0x08, 0x34, 0x84, 0x95, 0x4b,
	0x45, 0xf4, 0xfc, 0xaf, 0x75, 0xfc, 0xde, 0xe8, 0xcb, 0xff, 0x71, 0xff, 0x20, 0x3b, 0x77, 0xe8,
	0x35, 0x34, 0x67, 0x9b, 0xcd, 0xd3, 0x07, 0xdf, 0x8e, 0xbf, 0x9d, 0x5b, 0xe0, 0xfb, 0xb9, 0x05,
	0x7e, 0x9c, 0x5b, 0xe0, 0xc3, 0x70, 0x4b, 0xd5, 0xee, 0xb8, 0x76, 0x53, 0xb6, 0xf7, 0x32, 0xac,
	0x8e, 0x7b, 0x45, 0xd2, 0x9d, 0xbc, 0xc4, 0xde, 0x81, 0xa8, 0x4f, 0x4c, 0xe4, 0xbd, 0x14, 0x0b,
	0x41, 0x89, 0xf0, 0xf8, 0xda, 0xbb, 0xb9, 0x8b, 0x75, 0x45, 0x1f, 0xc4, 0xe0, 0x67, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x64, 0xf2, 0x87, 0x92, 0x2d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ViaProviderClient is the client API for ViaProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ViaProviderClient interface {
	Expose(ctx context.Context, in *ExposeReq, opts ...grpc.CallOption) (*ExposeAns, error)
	Off(ctx context.Context, in *ExposeReq, opts ...grpc.CallOption) (*ExposeAns, error)
}

type viaProviderClient struct {
	cc *grpc.ClientConn
}

func NewViaProviderClient(cc *grpc.ClientConn) ViaProviderClient {
	return &viaProviderClient{cc}
}

func (c *viaProviderClient) Expose(ctx context.Context, in *ExposeReq, opts ...grpc.CallOption) (*ExposeAns, error) {
	out := new(ExposeAns)
	err := c.cc.Invoke(ctx, "/fighter.api.ViaProvider/Expose", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *viaProviderClient) Off(ctx context.Context, in *ExposeReq, opts ...grpc.CallOption) (*ExposeAns, error) {
	out := new(ExposeAns)
	err := c.cc.Invoke(ctx, "/fighter.api.ViaProvider/Off", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ViaProviderServer is the server API for ViaProvider service.
type ViaProviderServer interface {
	Expose(context.Context, *ExposeReq) (*ExposeAns, error)
	Off(context.Context, *ExposeReq) (*ExposeAns, error)
}

// UnimplementedViaProviderServer can be embedded to have forward compatible implementations.
type UnimplementedViaProviderServer struct {
}

func (*UnimplementedViaProviderServer) Expose(ctx context.Context, req *ExposeReq) (*ExposeAns, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Expose not implemented")
}
func (*UnimplementedViaProviderServer) Off(ctx context.Context, req *ExposeReq) (*ExposeAns, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Off not implemented")
}

func RegisterViaProviderServer(s *grpc.Server, srv ViaProviderServer) {
	s.RegisterService(&_ViaProvider_serviceDesc, srv)
}

func _ViaProvider_Expose_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExposeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViaProviderServer).Expose(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fighter.api.ViaProvider/Expose",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViaProviderServer).Expose(ctx, req.(*ExposeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ViaProvider_Off_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExposeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ViaProviderServer).Off(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fighter.api.ViaProvider/Off",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ViaProviderServer).Off(ctx, req.(*ExposeReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _ViaProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fighter.api.ViaProvider",
	HandlerType: (*ViaProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Expose",
			Handler:    _ViaProvider_Expose_Handler,
		},
		{
			MethodName: "Off",
			Handler:    _ViaProvider_Off_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fighter/api/via_svc.proto",
}

func (m *ExposeReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExposeReq) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExposeReq) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Port != 0 {
		i = encodeVarintViaSvc(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x28
	}
	if len(m.Ip) > 0 {
		i -= len(m.Ip)
		copy(dAtA[i:], m.Ip)
		i = encodeVarintViaSvc(dAtA, i, uint64(len(m.Ip)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.PartyId) > 0 {
		i -= len(m.PartyId)
		copy(dAtA[i:], m.PartyId)
		i = encodeVarintViaSvc(dAtA, i, uint64(len(m.PartyId)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SvcType != 0 {
		i = encodeVarintViaSvc(dAtA, i, uint64(m.SvcType))
		i--
		dAtA[i] = 0x10
	}
	if len(m.TaskId) > 0 {
		i -= len(m.TaskId)
		copy(dAtA[i:], m.TaskId)
		i = encodeVarintViaSvc(dAtA, i, uint64(len(m.TaskId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ExposeAns) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ExposeAns) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ExposeAns) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Port != 0 {
		i = encodeVarintViaSvc(dAtA, i, uint64(m.Port))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Ip) > 0 {
		i -= len(m.Ip)
		copy(dAtA[i:], m.Ip)
		i = encodeVarintViaSvc(dAtA, i, uint64(len(m.Ip)))
		i--
		dAtA[i] = 0x12
	}
	if m.Ok {
		i--
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintViaSvc(dAtA []byte, offset int, v uint64) int {
	offset -= sovViaSvc(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ExposeReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TaskId)
	if l > 0 {
		n += 1 + l + sovViaSvc(uint64(l))
	}
	if m.SvcType != 0 {
		n += 1 + sovViaSvc(uint64(m.SvcType))
	}
	l = len(m.PartyId)
	if l > 0 {
		n += 1 + l + sovViaSvc(uint64(l))
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovViaSvc(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovViaSvc(uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ExposeAns) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ok {
		n += 2
	}
	l = len(m.Ip)
	if l > 0 {
		n += 1 + l + sovViaSvc(uint64(l))
	}
	if m.Port != 0 {
		n += 1 + sovViaSvc(uint64(m.Port))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovViaSvc(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozViaSvc(x uint64) (n int) {
	return sovViaSvc(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ExposeReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowViaSvc
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
			return fmt.Errorf("proto: ExposeReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExposeReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TaskId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowViaSvc
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
				return ErrInvalidLengthViaSvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthViaSvc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TaskId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SvcType", wireType)
			}
			m.SvcType = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowViaSvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SvcType |= ServiceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PartyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowViaSvc
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
				return ErrInvalidLengthViaSvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthViaSvc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PartyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowViaSvc
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
				return ErrInvalidLengthViaSvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthViaSvc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowViaSvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipViaSvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthViaSvc
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
func (m *ExposeAns) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowViaSvc
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
			return fmt.Errorf("proto: ExposeAns: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ExposeAns: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowViaSvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Ok = bool(v != 0)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ip", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowViaSvc
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
				return ErrInvalidLengthViaSvc
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthViaSvc
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Ip = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Port", wireType)
			}
			m.Port = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowViaSvc
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Port |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipViaSvc(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthViaSvc
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
func skipViaSvc(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowViaSvc
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
					return 0, ErrIntOverflowViaSvc
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
					return 0, ErrIntOverflowViaSvc
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
				return 0, ErrInvalidLengthViaSvc
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupViaSvc
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthViaSvc
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthViaSvc        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowViaSvc          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupViaSvc = fmt.Errorf("proto: unexpected end of group")
)
