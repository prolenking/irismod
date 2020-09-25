// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: random/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// QueryRandomRequest is request type for the Query/Random RPC method
type QueryRandomRequest struct {
	ReqId string `protobuf:"bytes,1,opt,name=req_id,json=reqId,proto3" json:"req_id,omitempty"`
}

func (m *QueryRandomRequest) Reset()         { *m = QueryRandomRequest{} }
func (m *QueryRandomRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRandomRequest) ProtoMessage()    {}
func (*QueryRandomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7e1fe88061ff84, []int{0}
}
func (m *QueryRandomRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRandomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRandomRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRandomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRandomRequest.Merge(m, src)
}
func (m *QueryRandomRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRandomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRandomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRandomRequest proto.InternalMessageInfo

func (m *QueryRandomRequest) GetReqId() string {
	if m != nil {
		return m.ReqId
	}
	return ""
}

// QueryParametersResponse is response type for the Query/Random RPC method
type QueryRandomResponse struct {
	Random *Random `protobuf:"bytes,1,opt,name=random,proto3" json:"random,omitempty"`
}

func (m *QueryRandomResponse) Reset()         { *m = QueryRandomResponse{} }
func (m *QueryRandomResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRandomResponse) ProtoMessage()    {}
func (*QueryRandomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7e1fe88061ff84, []int{1}
}
func (m *QueryRandomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRandomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRandomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRandomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRandomResponse.Merge(m, src)
}
func (m *QueryRandomResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRandomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRandomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRandomResponse proto.InternalMessageInfo

func (m *QueryRandomResponse) GetRandom() *Random {
	if m != nil {
		return m.Random
	}
	return nil
}

// QueryRandomRequestQueueRequest is request type for the Query/RandomRequestQueue RPC method
type QueryRandomRequestQueueRequest struct {
	Height int64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
}

func (m *QueryRandomRequestQueueRequest) Reset()         { *m = QueryRandomRequestQueueRequest{} }
func (m *QueryRandomRequestQueueRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRandomRequestQueueRequest) ProtoMessage()    {}
func (*QueryRandomRequestQueueRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7e1fe88061ff84, []int{2}
}
func (m *QueryRandomRequestQueueRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRandomRequestQueueRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRandomRequestQueueRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRandomRequestQueueRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRandomRequestQueueRequest.Merge(m, src)
}
func (m *QueryRandomRequestQueueRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryRandomRequestQueueRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRandomRequestQueueRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRandomRequestQueueRequest proto.InternalMessageInfo

func (m *QueryRandomRequestQueueRequest) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

// QueryRandomRequestQueueResponse is response type for the Query/RandomRequestQueue RPC method
type QueryRandomRequestQueueResponse struct {
	Requests []Request `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests"`
}

func (m *QueryRandomRequestQueueResponse) Reset()         { *m = QueryRandomRequestQueueResponse{} }
func (m *QueryRandomRequestQueueResponse) String() string { return proto.CompactTextString(m) }
func (*QueryRandomRequestQueueResponse) ProtoMessage()    {}
func (*QueryRandomRequestQueueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e7e1fe88061ff84, []int{3}
}
func (m *QueryRandomRequestQueueResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryRandomRequestQueueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryRandomRequestQueueResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryRandomRequestQueueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRandomRequestQueueResponse.Merge(m, src)
}
func (m *QueryRandomRequestQueueResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryRandomRequestQueueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRandomRequestQueueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRandomRequestQueueResponse proto.InternalMessageInfo

func (m *QueryRandomRequestQueueResponse) GetRequests() []Request {
	if m != nil {
		return m.Requests
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryRandomRequest)(nil), "irismod.random.QueryRandomRequest")
	proto.RegisterType((*QueryRandomResponse)(nil), "irismod.random.QueryRandomResponse")
	proto.RegisterType((*QueryRandomRequestQueueRequest)(nil), "irismod.random.QueryRandomRequestQueueRequest")
	proto.RegisterType((*QueryRandomRequestQueueResponse)(nil), "irismod.random.QueryRandomRequestQueueResponse")
}

func init() { proto.RegisterFile("random/query.proto", fileDescriptor_0e7e1fe88061ff84) }

var fileDescriptor_0e7e1fe88061ff84 = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0xdb, 0xed, 0xb7, 0xf2, 0x33, 0x03, 0x0f, 0x99, 0xfb, 0x43, 0xd1, 0x6e, 0xd4, 0xcb,
	0x40, 0x68, 0x60, 0x5e, 0xf4, 0x3a, 0xf0, 0xb0, 0xe3, 0x7a, 0x14, 0x41, 0x3a, 0x1b, 0xda, 0xc2,
	0xda, 0xb4, 0x4d, 0x72, 0x18, 0xc3, 0x8b, 0xaf, 0x40, 0xd0, 0x9b, 0x6f, 0x68, 0xc7, 0x81, 0x17,
	0x4f, 0x22, 0x9b, 0x2f, 0x44, 0x9a, 0x64, 0xe2, 0x56, 0xfc, 0x73, 0x6a, 0x9a, 0xe7, 0xf3, 0x3c,
	0xdf, 0x4f, 0xd3, 0x00, 0x98, 0x7b, 0x89, 0x4f, 0x62, 0x94, 0x71, 0x9c, 0xcf, 0x9c, 0x34, 0x27,
	0x8c, 0xc0, 0xfd, 0x28, 0x8f, 0x68, 0x4c, 0x7c, 0x47, 0xd6, 0xcc, 0x86, 0x62, 0xe4, 0x43, 0x42,
	0xe6, 0x41, 0x40, 0x02, 0x22, 0x96, 0xa8, 0x58, 0xa9, 0xdd, 0xc3, 0x80, 0x90, 0x60, 0x8a, 0x91,
	0x97, 0x46, 0xc8, 0x4b, 0x12, 0xc2, 0x3c, 0x16, 0x91, 0x84, 0xca, 0xaa, 0x7d, 0x02, 0xe0, 0xb8,
	0xc8, 0x71, 0xc5, 0x20, 0x17, 0x67, 0x1c, 0x53, 0x06, 0x9b, 0xc0, 0xc8, 0x71, 0x76, 0x1d, 0xf9,
	0x1d, 0xbd, 0xa7, 0xf7, 0xf7, 0xdc, 0x5a, 0x8e, 0xb3, 0x91, 0x6f, 0x5f, 0x80, 0xc6, 0x16, 0x4c,
	0x53, 0x92, 0x50, 0x0c, 0x1d, 0x60, 0x48, 0x0f, 0x41, 0xd7, 0x07, 0x2d, 0x67, 0xdb, 0xd6, 0x51,
	0xbc, 0xa2, 0xec, 0x33, 0x60, 0x95, 0x33, 0xc7, 0x1c, 0x73, 0xbc, 0xc9, 0x6f, 0x01, 0x23, 0xc4,
	0x51, 0x10, 0x32, 0x31, 0xb1, 0xea, 0xaa, 0x37, 0xfb, 0x0a, 0x74, 0xbf, 0xed, 0x54, 0x32, 0xe7,
	0xe0, 0x7f, 0x2e, 0xf7, 0x69, 0x47, 0xef, 0x55, 0xfb, 0xf5, 0x41, 0xbb, 0xa4, 0x23, 0xeb, 0xc3,
	0x7f, 0x8b, 0xd7, 0xae, 0xe6, 0x7e, 0xe2, 0x83, 0xa7, 0x0a, 0xa8, 0x89, 0xf1, 0x70, 0x0e, 0x0c,
	0x19, 0x01, 0xed, 0xdd, 0xe6, 0x72, 0xbe, 0x79, 0xfc, 0x23, 0x23, 0xbd, 0xec, 0xfe, 0xdd, 0xf3,
	0xfb, 0x43, 0xc5, 0x86, 0x3d, 0x54, 0xc0, 0x21, 0x9f, 0xa0, 0xad, 0x5f, 0x48, 0xd1, 0x5c, 0x9e,
	0xf8, 0x2d, 0x7c, 0xd4, 0x01, 0x2c, 0x7f, 0x20, 0x74, 0x7e, 0x37, 0xf9, 0x7a, 0x86, 0x26, 0xfa,
	0x33, 0xaf, 0x0c, 0x8f, 0x84, 0x61, 0x1b, 0x36, 0x77, 0x0d, 0xb3, 0x02, 0x1b, 0x8e, 0x16, 0x2b,
	0x4b, 0x5f, 0xae, 0x2c, 0xfd, 0x6d, 0x65, 0xe9, 0xf7, 0x6b, 0x4b, 0x5b, 0xae, 0x2d, 0xed, 0x65,
	0x6d, 0x69, 0x97, 0x28, 0x88, 0x58, 0xc8, 0x27, 0xce, 0x0d, 0x89, 0x45, 0x6b, 0x82, 0x19, 0x52,
	0xd9, 0x28, 0x26, 0x3e, 0x9f, 0x62, 0xba, 0x19, 0xc5, 0x66, 0x29, 0xa6, 0x13, 0x43, 0xdc, 0xbd,
	0xd3, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x89, 0xbe, 0x54, 0xf5, 0xea, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Random queries the random result
	Random(ctx context.Context, in *QueryRandomRequest, opts ...grpc.CallOption) (*QueryRandomResponse, error)
	// RandomRequestQueue queries the random request queue
	RandomRequestQueue(ctx context.Context, in *QueryRandomRequestQueueRequest, opts ...grpc.CallOption) (*QueryRandomRequestQueueResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Random(ctx context.Context, in *QueryRandomRequest, opts ...grpc.CallOption) (*QueryRandomResponse, error) {
	out := new(QueryRandomResponse)
	err := c.cc.Invoke(ctx, "/irismod.random.Query/Random", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) RandomRequestQueue(ctx context.Context, in *QueryRandomRequestQueueRequest, opts ...grpc.CallOption) (*QueryRandomRequestQueueResponse, error) {
	out := new(QueryRandomRequestQueueResponse)
	err := c.cc.Invoke(ctx, "/irismod.random.Query/RandomRequestQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Random queries the random result
	Random(context.Context, *QueryRandomRequest) (*QueryRandomResponse, error)
	// RandomRequestQueue queries the random request queue
	RandomRequestQueue(context.Context, *QueryRandomRequestQueueRequest) (*QueryRandomRequestQueueResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Random(ctx context.Context, req *QueryRandomRequest) (*QueryRandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Random not implemented")
}
func (*UnimplementedQueryServer) RandomRequestQueue(ctx context.Context, req *QueryRandomRequestQueueRequest) (*QueryRandomRequestQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RandomRequestQueue not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Random_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRandomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Random(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.random.Query/Random",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Random(ctx, req.(*QueryRandomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_RandomRequestQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRandomRequestQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).RandomRequestQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/irismod.random.Query/RandomRequestQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).RandomRequestQueue(ctx, req.(*QueryRandomRequestQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "irismod.random.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Random",
			Handler:    _Query_Random_Handler,
		},
		{
			MethodName: "RandomRequestQueue",
			Handler:    _Query_RandomRequestQueue_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "random/query.proto",
}

func (m *QueryRandomRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRandomRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRandomRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReqId) > 0 {
		i -= len(m.ReqId)
		copy(dAtA[i:], m.ReqId)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.ReqId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRandomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRandomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRandomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Random != nil {
		{
			size, err := m.Random.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryRandomRequestQueueRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRandomRequestQueueRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRandomRequestQueueRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Height != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryRandomRequestQueueResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryRandomRequestQueueResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryRandomRequestQueueResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for iNdEx := len(m.Requests) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Requests[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryRandomRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ReqId)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRandomResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Random != nil {
		l = m.Random.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryRandomRequestQueueRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovQuery(uint64(m.Height))
	}
	return n
}

func (m *QueryRandomRequestQueueResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Requests) > 0 {
		for _, e := range m.Requests {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryRandomRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRandomRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRandomRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReqId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReqId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRandomResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRandomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRandomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Random", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Random == nil {
				m.Random = &Random{}
			}
			if err := m.Random.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRandomRequestQueueRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRandomRequestQueueRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRandomRequestQueueRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryRandomRequestQueueResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryRandomRequestQueueResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryRandomRequestQueueResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requests", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requests = append(m.Requests, Request{})
			if err := m.Requests[len(m.Requests)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
