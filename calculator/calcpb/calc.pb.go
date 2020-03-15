// Code generated by protoc-gen-go. DO NOT EDIT.
// source: calculator/calcpb/calc.proto

package calcpb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SumNumbers struct {
	Num_1                int64    `protobuf:"varint,1,opt,name=num_1,json=num1,proto3" json:"num_1,omitempty"`
	Num_2                int64    `protobuf:"varint,2,opt,name=num_2,json=num2,proto3" json:"num_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumNumbers) Reset()         { *m = SumNumbers{} }
func (m *SumNumbers) String() string { return proto.CompactTextString(m) }
func (*SumNumbers) ProtoMessage()    {}
func (*SumNumbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{0}
}

func (m *SumNumbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumNumbers.Unmarshal(m, b)
}
func (m *SumNumbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumNumbers.Marshal(b, m, deterministic)
}
func (m *SumNumbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumNumbers.Merge(m, src)
}
func (m *SumNumbers) XXX_Size() int {
	return xxx_messageInfo_SumNumbers.Size(m)
}
func (m *SumNumbers) XXX_DiscardUnknown() {
	xxx_messageInfo_SumNumbers.DiscardUnknown(m)
}

var xxx_messageInfo_SumNumbers proto.InternalMessageInfo

func (m *SumNumbers) GetNum_1() int64 {
	if m != nil {
		return m.Num_1
	}
	return 0
}

func (m *SumNumbers) GetNum_2() int64 {
	if m != nil {
		return m.Num_2
	}
	return 0
}

type SumRequest struct {
	SumResult            *SumNumbers `protobuf:"bytes,1,opt,name=sumResult,proto3" json:"sumResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{1}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetSumResult() *SumNumbers {
	if m != nil {
		return m.SumResult
	}
	return nil
}

type SumResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{2}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type PrimeRequest struct {
	Num                  int64    `protobuf:"varint,1,opt,name=num,proto3" json:"num,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeRequest) Reset()         { *m = PrimeRequest{} }
func (m *PrimeRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeRequest) ProtoMessage()    {}
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{3}
}

func (m *PrimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeRequest.Unmarshal(m, b)
}
func (m *PrimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeRequest.Marshal(b, m, deterministic)
}
func (m *PrimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeRequest.Merge(m, src)
}
func (m *PrimeRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeRequest.Size(m)
}
func (m *PrimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeRequest proto.InternalMessageInfo

func (m *PrimeRequest) GetNum() int64 {
	if m != nil {
		return m.Num
	}
	return 0
}

type PrimeResponse struct {
	Result               int64    `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeResponse) Reset()         { *m = PrimeResponse{} }
func (m *PrimeResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeResponse) ProtoMessage()    {}
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{4}
}

func (m *PrimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeResponse.Unmarshal(m, b)
}
func (m *PrimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeResponse.Marshal(b, m, deterministic)
}
func (m *PrimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeResponse.Merge(m, src)
}
func (m *PrimeResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeResponse.Size(m)
}
func (m *PrimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeResponse proto.InternalMessageInfo

func (m *PrimeResponse) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

type ComputeAvgRequest struct {
	Number               int64    `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAvgRequest) Reset()         { *m = ComputeAvgRequest{} }
func (m *ComputeAvgRequest) String() string { return proto.CompactTextString(m) }
func (*ComputeAvgRequest) ProtoMessage()    {}
func (*ComputeAvgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{5}
}

func (m *ComputeAvgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAvgRequest.Unmarshal(m, b)
}
func (m *ComputeAvgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAvgRequest.Marshal(b, m, deterministic)
}
func (m *ComputeAvgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAvgRequest.Merge(m, src)
}
func (m *ComputeAvgRequest) XXX_Size() int {
	return xxx_messageInfo_ComputeAvgRequest.Size(m)
}
func (m *ComputeAvgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAvgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAvgRequest proto.InternalMessageInfo

func (m *ComputeAvgRequest) GetNumber() int64 {
	if m != nil {
		return m.Number
	}
	return 0
}

type ComputeAvgResponse struct {
	AvgResult            float64  `protobuf:"fixed64,1,opt,name=avgResult,proto3" json:"avgResult,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputeAvgResponse) Reset()         { *m = ComputeAvgResponse{} }
func (m *ComputeAvgResponse) String() string { return proto.CompactTextString(m) }
func (*ComputeAvgResponse) ProtoMessage()    {}
func (*ComputeAvgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cfc74e58bc0fa04b, []int{6}
}

func (m *ComputeAvgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputeAvgResponse.Unmarshal(m, b)
}
func (m *ComputeAvgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputeAvgResponse.Marshal(b, m, deterministic)
}
func (m *ComputeAvgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputeAvgResponse.Merge(m, src)
}
func (m *ComputeAvgResponse) XXX_Size() int {
	return xxx_messageInfo_ComputeAvgResponse.Size(m)
}
func (m *ComputeAvgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputeAvgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ComputeAvgResponse proto.InternalMessageInfo

func (m *ComputeAvgResponse) GetAvgResult() float64 {
	if m != nil {
		return m.AvgResult
	}
	return 0
}

func init() {
	proto.RegisterType((*SumNumbers)(nil), "calculator.SumNumbers")
	proto.RegisterType((*SumRequest)(nil), "calculator.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "calculator.SumResponse")
	proto.RegisterType((*PrimeRequest)(nil), "calculator.PrimeRequest")
	proto.RegisterType((*PrimeResponse)(nil), "calculator.PrimeResponse")
	proto.RegisterType((*ComputeAvgRequest)(nil), "calculator.ComputeAvgRequest")
	proto.RegisterType((*ComputeAvgResponse)(nil), "calculator.ComputeAvgResponse")
}

func init() {
	proto.RegisterFile("calculator/calcpb/calc.proto", fileDescriptor_cfc74e58bc0fa04b)
}

var fileDescriptor_cfc74e58bc0fa04b = []byte{
	// 322 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0xad, 0xd5, 0xe2, 0xee, 0x54, 0xf4, 0x0a, 0x73, 0x8e, 0x29, 0x23, 0x20, 0x0e, 0x84,
	0xe9, 0xaa, 0x88, 0xaf, 0x6e, 0xbe, 0x2a, 0xd2, 0xbe, 0xf9, 0x22, 0x5d, 0x09, 0xa3, 0xd0, 0x24,
	0x35, 0x7f, 0xf6, 0xa5, 0xfd, 0x12, 0xb2, 0xac, 0x59, 0x8a, 0x53, 0x7c, 0x6a, 0xee, 0xbd, 0xe7,
	0xfc, 0x4e, 0x73, 0x09, 0xf4, 0xf3, 0xac, 0xcc, 0x4d, 0x99, 0x69, 0x21, 0x6f, 0x96, 0xc7, 0x6a,
	0x66, 0x3f, 0xa3, 0x4a, 0x0a, 0x2d, 0x10, 0xfc, 0x94, 0x3c, 0x00, 0xa4, 0x86, 0xbd, 0x1a, 0x36,
	0xa3, 0x52, 0xe1, 0x09, 0xec, 0x72, 0xc3, 0x3e, 0xc6, 0xdd, 0x60, 0x10, 0x0c, 0xc3, 0x64, 0x87,
	0x1b, 0x36, 0x76, 0xcd, 0xb8, 0xbb, 0xbd, 0x6e, 0xc6, 0x64, 0x62, 0x7d, 0x09, 0xfd, 0x34, 0x54,
	0x69, 0xbc, 0x87, 0x96, 0x5a, 0x56, 0xca, 0x94, 0xda, 0x7a, 0xdb, 0x71, 0x67, 0xe4, 0x53, 0x46,
	0x3e, 0x22, 0xf1, 0x42, 0x72, 0x09, 0x6d, 0xcb, 0x50, 0x95, 0xe0, 0x8a, 0x62, 0x07, 0x22, 0xe9,
	0x09, 0x61, 0x52, 0x57, 0x64, 0x00, 0xfb, 0x6f, 0xb2, 0x60, 0xd4, 0x85, 0x1d, 0x41, 0xc8, 0x0d,
	0xab, 0x45, 0xcb, 0x23, 0xb9, 0x82, 0x83, 0x5a, 0xf1, 0x0f, 0xea, 0x1a, 0x8e, 0xa7, 0x82, 0x55,
	0x46, 0xd3, 0xa7, 0xc5, 0xdc, 0xf1, 0x3a, 0x10, 0x71, 0xfb, 0x73, 0x4e, 0xbc, 0xaa, 0x48, 0x0c,
	0xd8, 0x14, 0xd7, 0xe8, 0x3e, 0xb4, 0x32, 0x5b, 0x3a, 0x7a, 0x90, 0xf8, 0x46, 0xfc, 0x15, 0x40,
	0x7b, 0x9a, 0x95, 0x79, 0x4a, 0xe5, 0xa2, 0xc8, 0x29, 0x3e, 0x42, 0x98, 0x1a, 0x86, 0x3f, 0x97,
	0x51, 0x47, 0xf7, 0x4e, 0x37, 0xfa, 0xab, 0x14, 0xb2, 0x85, 0x2f, 0x80, 0xf6, 0x4e, 0xcf, 0x34,
	0x17, 0xac, 0x12, 0xaa, 0xd0, 0x85, 0xe0, 0xd8, 0x6d, 0x1a, 0x9a, 0x5b, 0xe9, 0x9d, 0xfd, 0x32,
	0x71, 0xb0, 0xdb, 0x00, 0x53, 0x38, 0x5c, 0x5f, 0x86, 0xca, 0x6c, 0x4e, 0xf1, 0xbc, 0x69, 0xd8,
	0xd8, 0x4a, 0xef, 0xe2, 0xaf, 0xb1, 0x83, 0x0e, 0x83, 0xc9, 0xde, 0x7b, 0xb4, 0x7a, 0x5d, 0xb3,
	0xc8, 0xbe, 0xac, 0xbb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xee, 0x8c, 0xd4, 0xce, 0x79, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CalcServiceClient is the client API for CalcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CalcServiceClient interface {
	// Unary
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	// Server Streaming
	PrimeDecomposition(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalcService_PrimeDecompositionClient, error)
	// Client Streaming
	ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalcService_ComputeAverageClient, error)
}

type calcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalcServiceClient(cc grpc.ClientConnInterface) CalcServiceClient {
	return &calcServiceClient{cc}
}

func (c *calcServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalcService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calcServiceClient) PrimeDecomposition(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalcService_PrimeDecompositionClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[0], "/calculator.CalcService/PrimeDecomposition", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServicePrimeDecompositionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalcService_PrimeDecompositionClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type calcServicePrimeDecompositionClient struct {
	grpc.ClientStream
}

func (x *calcServicePrimeDecompositionClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calcServiceClient) ComputeAverage(ctx context.Context, opts ...grpc.CallOption) (CalcService_ComputeAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CalcService_serviceDesc.Streams[1], "/calculator.CalcService/ComputeAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calcServiceComputeAverageClient{stream}
	return x, nil
}

type CalcService_ComputeAverageClient interface {
	Send(*ComputeAvgRequest) error
	CloseAndRecv() (*ComputeAvgResponse, error)
	grpc.ClientStream
}

type calcServiceComputeAverageClient struct {
	grpc.ClientStream
}

func (x *calcServiceComputeAverageClient) Send(m *ComputeAvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calcServiceComputeAverageClient) CloseAndRecv() (*ComputeAvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ComputeAvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CalcServiceServer is the server API for CalcService service.
type CalcServiceServer interface {
	// Unary
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	// Server Streaming
	PrimeDecomposition(*PrimeRequest, CalcService_PrimeDecompositionServer) error
	// Client Streaming
	ComputeAverage(CalcService_ComputeAverageServer) error
}

// UnimplementedCalcServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCalcServiceServer struct {
}

func (*UnimplementedCalcServiceServer) Sum(ctx context.Context, req *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (*UnimplementedCalcServiceServer) PrimeDecomposition(req *PrimeRequest, srv CalcService_PrimeDecompositionServer) error {
	return status.Errorf(codes.Unimplemented, "method PrimeDecomposition not implemented")
}
func (*UnimplementedCalcServiceServer) ComputeAverage(srv CalcService_ComputeAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAverage not implemented")
}

func RegisterCalcServiceServer(s *grpc.Server, srv CalcServiceServer) {
	s.RegisterService(&_CalcService_serviceDesc, srv)
}

func _CalcService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalcServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalcService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalcServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalcService_PrimeDecomposition_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalcServiceServer).PrimeDecomposition(m, &calcServicePrimeDecompositionServer{stream})
}

type CalcService_PrimeDecompositionServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type calcServicePrimeDecompositionServer struct {
	grpc.ServerStream
}

func (x *calcServicePrimeDecompositionServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalcService_ComputeAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalcServiceServer).ComputeAverage(&calcServiceComputeAverageServer{stream})
}

type CalcService_ComputeAverageServer interface {
	SendAndClose(*ComputeAvgResponse) error
	Recv() (*ComputeAvgRequest, error)
	grpc.ServerStream
}

type calcServiceComputeAverageServer struct {
	grpc.ServerStream
}

func (x *calcServiceComputeAverageServer) SendAndClose(m *ComputeAvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calcServiceComputeAverageServer) Recv() (*ComputeAvgRequest, error) {
	m := new(ComputeAvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _CalcService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalcService",
	HandlerType: (*CalcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalcService_Sum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrimeDecomposition",
			Handler:       _CalcService_PrimeDecomposition_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAverage",
			Handler:       _CalcService_ComputeAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "calculator/calcpb/calc.proto",
}
