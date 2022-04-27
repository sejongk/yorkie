// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: admin.proto

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

type ListDocumentsRequest struct {
	PreviousId           string   `protobuf:"bytes,1,opt,name=previous_id,json=previousId,proto3" json:"previous_id,omitempty"`
	PageSize             int32    `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	IsForward            bool     `protobuf:"varint,3,opt,name=is_forward,json=isForward,proto3" json:"is_forward,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListDocumentsRequest) Reset()         { *m = ListDocumentsRequest{} }
func (m *ListDocumentsRequest) String() string { return proto.CompactTextString(m) }
func (*ListDocumentsRequest) ProtoMessage()    {}
func (*ListDocumentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{0}
}
func (m *ListDocumentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListDocumentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListDocumentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListDocumentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDocumentsRequest.Merge(m, src)
}
func (m *ListDocumentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *ListDocumentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDocumentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListDocumentsRequest proto.InternalMessageInfo

func (m *ListDocumentsRequest) GetPreviousId() string {
	if m != nil {
		return m.PreviousId
	}
	return ""
}

func (m *ListDocumentsRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListDocumentsRequest) GetIsForward() bool {
	if m != nil {
		return m.IsForward
	}
	return false
}

type ListDocumentsResponse struct {
	Documents            []*DocumentSummary `protobuf:"bytes,1,rep,name=documents,proto3" json:"documents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ListDocumentsResponse) Reset()         { *m = ListDocumentsResponse{} }
func (m *ListDocumentsResponse) String() string { return proto.CompactTextString(m) }
func (*ListDocumentsResponse) ProtoMessage()    {}
func (*ListDocumentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_73a7fc70dcc2027c, []int{1}
}
func (m *ListDocumentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListDocumentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListDocumentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListDocumentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListDocumentsResponse.Merge(m, src)
}
func (m *ListDocumentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *ListDocumentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListDocumentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListDocumentsResponse proto.InternalMessageInfo

func (m *ListDocumentsResponse) GetDocuments() []*DocumentSummary {
	if m != nil {
		return m.Documents
	}
	return nil
}

func init() {
	proto.RegisterType((*ListDocumentsRequest)(nil), "api.ListDocumentsRequest")
	proto.RegisterType((*ListDocumentsResponse)(nil), "api.ListDocumentsResponse")
}

func init() { proto.RegisterFile("admin.proto", fileDescriptor_73a7fc70dcc2027c) }

var fileDescriptor_73a7fc70dcc2027c = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0x87, 0xbb, 0x86, 0x4a, 0x33, 0x41, 0x94, 0xa5, 0x42, 0x8c, 0x18, 0x43, 0x4e, 0x39, 0xe5,
	0x10, 0x9f, 0x40, 0x11, 0x51, 0xf4, 0x62, 0xfa, 0x00, 0x61, 0x6d, 0x46, 0x99, 0x43, 0xb2, 0xeb,
	0x4e, 0x56, 0xb1, 0x4f, 0xe2, 0x23, 0x79, 0xf4, 0x11, 0x24, 0xbe, 0x88, 0xb4, 0x21, 0x88, 0xd2,
	0xeb, 0x37, 0xbf, 0xf9, 0xf3, 0x0d, 0x04, 0xaa, 0x6e, 0xa8, 0xcd, 0x8d, 0xd5, 0x9d, 0x96, 0x9e,
	0x32, 0x14, 0xed, 0x5b, 0x64, 0xed, 0xec, 0x12, 0x79, 0xa0, 0x29, 0xc3, 0xfc, 0x8e, 0xb8, 0xbb,
	0xd4, 0x4b, 0xd7, 0x60, 0xdb, 0x71, 0x89, 0xcf, 0x0e, 0xb9, 0x93, 0xa7, 0x10, 0x18, 0x8b, 0x2f,
	0xa4, 0x1d, 0x57, 0x54, 0x87, 0x22, 0x11, 0x99, 0x5f, 0xc2, 0x88, 0x6e, 0x6a, 0x79, 0x0c, 0xbe,
	0x51, 0x4f, 0x58, 0x31, 0xad, 0x30, 0xdc, 0x49, 0x44, 0x36, 0x2d, 0x67, 0x6b, 0xb0, 0xa0, 0x15,
	0xca, 0x13, 0x00, 0xe2, 0xea, 0x51, 0xdb, 0x57, 0x65, 0xeb, 0xd0, 0x4b, 0x44, 0x36, 0x2b, 0x7d,
	0xe2, 0xab, 0x01, 0xa4, 0xb7, 0x70, 0xf8, 0x6f, 0x29, 0x1b, 0xdd, 0x32, 0xca, 0x02, 0xfc, 0x7a,
	0x84, 0xa1, 0x48, 0xbc, 0x2c, 0x28, 0xe6, 0xb9, 0x32, 0x94, 0x8f, 0xd1, 0x85, 0x6b, 0x1a, 0x65,
	0xdf, 0xca, 0xdf, 0x58, 0x71, 0x0f, 0xd3, 0xf3, 0xb5, 0xa6, 0xbc, 0x86, 0xbd, 0x3f, 0x53, 0xe5,
	0xd1, 0xa6, 0x75, 0x9b, 0x5e, 0x14, 0x6d, 0x2b, 0x0d, 0x47, 0xa4, 0x93, 0x8b, 0x83, 0x8f, 0x3e,
	0x16, 0x9f, 0x7d, 0x2c, 0xbe, 0xfa, 0x58, 0xbc, 0x7f, 0xc7, 0x93, 0x87, 0xdd, 0xcd, 0xb7, 0xce,
	0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x34, 0x37, 0x57, 0x21, 0x52, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AdminClient is the client API for Admin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminClient interface {
	ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) ListDocuments(ctx context.Context, in *ListDocumentsRequest, opts ...grpc.CallOption) (*ListDocumentsResponse, error) {
	out := new(ListDocumentsResponse)
	err := c.cc.Invoke(ctx, "/api.Admin/ListDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AdminServer is the server API for Admin service.
type AdminServer interface {
	ListDocuments(context.Context, *ListDocumentsRequest) (*ListDocumentsResponse, error)
}

// UnimplementedAdminServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServer struct {
}

func (*UnimplementedAdminServer) ListDocuments(ctx context.Context, req *ListDocumentsRequest) (*ListDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDocuments not implemented")
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_ListDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).ListDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Admin/ListDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).ListDocuments(ctx, req.(*ListDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDocuments",
			Handler:    _Admin_ListDocuments_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "admin.proto",
}

func (m *ListDocumentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListDocumentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListDocumentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.IsForward {
		i--
		if m.IsForward {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.PageSize != 0 {
		i = encodeVarintAdmin(dAtA, i, uint64(m.PageSize))
		i--
		dAtA[i] = 0x10
	}
	if len(m.PreviousId) > 0 {
		i -= len(m.PreviousId)
		copy(dAtA[i:], m.PreviousId)
		i = encodeVarintAdmin(dAtA, i, uint64(len(m.PreviousId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ListDocumentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListDocumentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListDocumentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Documents) > 0 {
		for iNdEx := len(m.Documents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Documents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAdmin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAdmin(dAtA []byte, offset int, v uint64) int {
	offset -= sovAdmin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListDocumentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PreviousId)
	if l > 0 {
		n += 1 + l + sovAdmin(uint64(l))
	}
	if m.PageSize != 0 {
		n += 1 + sovAdmin(uint64(m.PageSize))
	}
	if m.IsForward {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ListDocumentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Documents) > 0 {
		for _, e := range m.Documents {
			l = e.Size()
			n += 1 + l + sovAdmin(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovAdmin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAdmin(x uint64) (n int) {
	return sovAdmin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListDocumentsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: ListDocumentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListDocumentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PageSize", wireType)
			}
			m.PageSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PageSize |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsForward", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
			m.IsForward = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func (m *ListDocumentsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAdmin
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
			return fmt.Errorf("proto: ListDocumentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListDocumentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Documents", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAdmin
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
				return ErrInvalidLengthAdmin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAdmin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Documents = append(m.Documents, &DocumentSummary{})
			if err := m.Documents[len(m.Documents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAdmin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAdmin
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
func skipAdmin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
					return 0, ErrIntOverflowAdmin
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
				return 0, ErrInvalidLengthAdmin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAdmin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAdmin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAdmin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAdmin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAdmin = fmt.Errorf("proto: unexpected end of group")
)
