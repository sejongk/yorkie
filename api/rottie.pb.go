// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/rottie.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ActivateRequest struct {
	ClientKey            string   `protobuf:"bytes,1,opt,name=client_key,json=clientKey,proto3" json:"client_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateRequest) Reset()         { *m = ActivateRequest{} }
func (m *ActivateRequest) String() string { return proto.CompactTextString(m) }
func (*ActivateRequest) ProtoMessage()    {}
func (*ActivateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_da93f86f306eca03, []int{0}
}
func (m *ActivateRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivateRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateRequest.Merge(m, src)
}
func (m *ActivateRequest) XXX_Size() int {
	return m.Size()
}
func (m *ActivateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateRequest proto.InternalMessageInfo

func (m *ActivateRequest) GetClientKey() string {
	if m != nil {
		return m.ClientKey
	}
	return ""
}

type ActivateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ActivateResponse) Reset()         { *m = ActivateResponse{} }
func (m *ActivateResponse) String() string { return proto.CompactTextString(m) }
func (*ActivateResponse) ProtoMessage()    {}
func (*ActivateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_da93f86f306eca03, []int{1}
}
func (m *ActivateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ActivateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ActivateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ActivateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActivateResponse.Merge(m, src)
}
func (m *ActivateResponse) XXX_Size() int {
	return m.Size()
}
func (m *ActivateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ActivateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ActivateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ActivateRequest)(nil), "api.ActivateRequest")
	proto.RegisterType((*ActivateResponse)(nil), "api.ActivateResponse")
}

func init() { proto.RegisterFile("api/rottie.proto", fileDescriptor_da93f86f306eca03) }

var fileDescriptor_da93f86f306eca03 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0xc8, 0xd4,
	0x2f, 0xca, 0x2f, 0x29, 0xc9, 0x4c, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c,
	0xc8, 0x54, 0x32, 0xe0, 0xe2, 0x77, 0x4c, 0x2e, 0xc9, 0x2c, 0x4b, 0x2c, 0x49, 0x0d, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe5, 0xe2, 0x4a, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89, 0xcf,
	0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x84, 0x88, 0x78, 0xa7, 0x56, 0x2a,
	0x09, 0x71, 0x09, 0x20, 0x74, 0x14, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x1a, 0x39, 0x73, 0xb1, 0x05,
	0x81, 0x8d, 0x16, 0xb2, 0xe4, 0xe2, 0x80, 0xc9, 0x0a, 0x89, 0xe8, 0x25, 0x16, 0x64, 0xea, 0xa1,
	0x19, 0x2f, 0x25, 0x8a, 0x26, 0x0a, 0x31, 0x42, 0x89, 0xc1, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6,
	0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec,
	0x50, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x75, 0x3c, 0xc5, 0x83, 0xbc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RottieClient is the client API for Rottie service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RottieClient interface {
	Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error)
}

type rottieClient struct {
	cc *grpc.ClientConn
}

func NewRottieClient(cc *grpc.ClientConn) RottieClient {
	return &rottieClient{cc}
}

func (c *rottieClient) Activate(ctx context.Context, in *ActivateRequest, opts ...grpc.CallOption) (*ActivateResponse, error) {
	out := new(ActivateResponse)
	err := c.cc.Invoke(ctx, "/api.Rottie/Activate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RottieServer is the server API for Rottie service.
type RottieServer interface {
	Activate(context.Context, *ActivateRequest) (*ActivateResponse, error)
}

// UnimplementedRottieServer can be embedded to have forward compatible implementations.
type UnimplementedRottieServer struct {
}

func (*UnimplementedRottieServer) Activate(ctx context.Context, req *ActivateRequest) (*ActivateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Activate not implemented")
}

func RegisterRottieServer(s *grpc.Server, srv RottieServer) {
	s.RegisterService(&_Rottie_serviceDesc, srv)
}

func _Rottie_Activate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RottieServer).Activate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Rottie/Activate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RottieServer).Activate(ctx, req.(*ActivateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rottie_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Rottie",
	HandlerType: (*RottieServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Activate",
			Handler:    _Rottie_Activate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/rottie.proto",
}

func (m *ActivateRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivateRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivateRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.ClientKey) > 0 {
		i -= len(m.ClientKey)
		copy(dAtA[i:], m.ClientKey)
		i = encodeVarintRottie(dAtA, i, uint64(len(m.ClientKey)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ActivateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ActivateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ActivateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	return len(dAtA) - i, nil
}

func encodeVarintRottie(dAtA []byte, offset int, v uint64) int {
	offset -= sovRottie(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ActivateRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientKey)
	if l > 0 {
		n += 1 + l + sovRottie(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ActivateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovRottie(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRottie(x uint64) (n int) {
	return sovRottie(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ActivateRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRottie
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
			return fmt.Errorf("proto: ActivateRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivateRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRottie
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
				return ErrInvalidLengthRottie
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRottie
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRottie(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRottie
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRottie
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
func (m *ActivateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRottie
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
			return fmt.Errorf("proto: ActivateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ActivateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipRottie(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthRottie
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthRottie
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
func skipRottie(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRottie
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
					return 0, ErrIntOverflowRottie
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
					return 0, ErrIntOverflowRottie
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
				return 0, ErrInvalidLengthRottie
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRottie
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRottie
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRottie        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRottie          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRottie = fmt.Errorf("proto: unexpected end of group")
)