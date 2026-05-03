// 由protoc-gen-go生成的代码,请勿编辑。
// source: implant.proto

/*
Package grpcapi是一个生成的协议缓冲区包。

它是从以下文件生成的:

	implant.proto

它有以下顶级消息:

	Command
	Empty
*/
package grpcapi

import (
	fmt "fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	context "golang.org/x/net/context"

	grpc "google.golang.org/grpc"
)

// 引用导入以抑制未使用时的错误。
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 这是编译时断言,用于确保生成的此文件与编译时使用的proto包兼容。
// 如果此行出现编译错误,可能意味着您的proto包副本需要更新。
const _ = proto.ProtoPackageIsVersion2 // 请升级proto包

// Command定义具有输入和输出字段的消息
type Command struct {
	In  string `protobuf:"bytes,1,opt,name=In" json:"In,omitempty"`
	Out string `protobuf:"bytes,2,opt,name=Out" json:"Out,omitempty"`
}

func (m *Command) Reset()                    { *m = Command{} }
func (m *Command) String() string            { return proto.CompactTextString(m) }
func (*Command) ProtoMessage()               {}
func (*Command) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Command) GetIn() string {
	if m != nil {
		return m.In
	}
	return ""
}

func (m *Command) GetOut() string {
	if m != nil {
		return m.Out
	}
	return ""
}

// Empty定义一个用于替代null的空消息
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Command)(nil), "grpcapi.Command")
	proto.RegisterType((*Empty)(nil), "grpcapi.Empty")
}

// 引用导入以抑制未使用时的错误。
var _ context.Context
var _ grpc.ClientConn

// 这是编译时断言,用于确保生成的此文件与grpc包兼容。
const _ = grpc.SupportPackageIsVersion4

// Implant服务的客户端API

type ImplantClient interface {
	FetchCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Command, error)
	SendOutput(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Empty, error)
}

type implantClient struct {
	cc *grpc.ClientConn
}

func NewImplantClient(cc *grpc.ClientConn) ImplantClient {
	return &implantClient{cc}
}

func (c *implantClient) FetchCommand(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Command, error) {
	out := new(Command)
	err := grpc.Invoke(ctx, "/grpcapi.Implant/FetchCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *implantClient) SendOutput(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/grpcapi.Implant/SendOutput", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Implant服务的服务端API

type ImplantServer interface {
	FetchCommand(context.Context, *Empty) (*Command, error)
	SendOutput(context.Context, *Command) (*Empty, error)
}

func RegisterImplantServer(s *grpc.Server, srv ImplantServer) {
	s.RegisterService(&_Implant_serviceDesc, srv)
}

func _Implant_FetchCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).FetchCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.Implant/FetchCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).FetchCommand(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Implant_SendOutput_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImplantServer).SendOutput(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.Implant/SendOutput",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImplantServer).SendOutput(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Implant_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.Implant",
	HandlerType: (*ImplantServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchCommand",
			Handler:    _Implant_FetchCommand_Handler,
		},
		{
			MethodName: "SendOutput",
			Handler:    _Implant_SendOutput_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "implant.proto",
}

// Admin服务的客户端API

type AdminClient interface {
	RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Command, error)
}

type adminClient struct {
	cc *grpc.ClientConn
}

func NewAdminClient(cc *grpc.ClientConn) AdminClient {
	return &adminClient{cc}
}

func (c *adminClient) RunCommand(ctx context.Context, in *Command, opts ...grpc.CallOption) (*Command, error) {
	out := new(Command)
	err := grpc.Invoke(ctx, "/grpcapi.Admin/RunCommand", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Admin服务的服务端API

type AdminServer interface {
	RunCommand(context.Context, *Command) (*Command, error)
}

func RegisterAdminServer(s *grpc.Server, srv AdminServer) {
	s.RegisterService(&_Admin_serviceDesc, srv)
}

func _Admin_RunCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Command)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServer).RunCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcapi.Admin/RunCommand",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServer).RunCommand(ctx, req.(*Command))
	}
	return interceptor(ctx, in, info, handler)
}

var _Admin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpcapi.Admin",
	HandlerType: (*AdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunCommand",
			Handler:    _Admin_RunCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "implant.proto",
}

func init() { proto.RegisterFile("implant.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcd, 0xcc, 0x2d, 0xc8,
	0x49, 0xcc, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x2f, 0x2a, 0x48, 0x4e,
	0x2c, 0xc8, 0x54, 0xd2, 0xe6, 0x62, 0x77, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x11, 0xe2, 0xe3,
	0x62, 0xf2, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xf2, 0xcc, 0x13, 0x12, 0xe0,
	0x62, 0xf6, 0x2f, 0x2d, 0x91, 0x60, 0x02, 0x0b, 0x80, 0x98, 0x4a, 0xec, 0x5c, 0xac, 0xae, 0xb9,
	0x05, 0x25, 0x95, 0x46, 0xd9, 0x5c, 0xec, 0x9e, 0x10, 0xf3, 0x84, 0x0c, 0xb8, 0x78, 0xdc, 0x52,
	0x4b, 0x92, 0x33, 0xe0, 0xa6, 0xe8, 0x41, 0x8d, 0xd6, 0x03, 0x2b, 0x95, 0x12, 0x80, 0xf3, 0x61,
	0x2a, 0xf4, 0xb8, 0xb8, 0x82, 0x53, 0xf3, 0x52, 0xfc, 0x4b, 0x4b, 0x0a, 0x4a, 0x4b, 0x84, 0x30,
	0xe4, 0xa5, 0xd0, 0x4c, 0x30, 0xb2, 0xe4, 0x62, 0x75, 0x4c, 0xc9, 0xcd, 0xcc, 0x13, 0x32, 0xe0,
	0xe2, 0x0a, 0x2a, 0xcd, 0x83, 0x19, 0x83, 0xa9, 0x11, 0x43, 0x24, 0x89, 0x0d, 0xec, 0x5b, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x25, 0xf4, 0x80, 0xfe, 0x00, 0x00, 0x00,
}
