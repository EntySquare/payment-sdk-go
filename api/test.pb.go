// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.15.8
// source: test.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Req struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	JsonStr string `protobuf:"bytes,1,opt,name=jsonStr,proto3" json:"jsonStr,omitempty"`
}

func (x *Req) Reset() {
	*x = Req{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Req) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Req) ProtoMessage() {}

func (x *Req) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Req.ProtoReflect.Descriptor instead.
func (*Req) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{0}
}

func (x *Req) GetJsonStr() string {
	if x != nil {
		return x.JsonStr
	}
	return ""
}

type Res struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BackJson string `protobuf:"bytes,1,opt,name=backJson,proto3" json:"backJson,omitempty"`
}

func (x *Res) Reset() {
	*x = Res{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Res) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Res) ProtoMessage() {}

func (x *Res) ProtoReflect() protoreflect.Message {
	mi := &file_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Res.ProtoReflect.Descriptor instead.
func (*Res) Descriptor() ([]byte, []int) {
	return file_test_proto_rawDescGZIP(), []int{1}
}

func (x *Res) GetBackJson() string {
	if x != nil {
		return x.BackJson
	}
	return ""
}

var File_test_proto protoreflect.FileDescriptor

var file_test_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70,
	0x69, 0x22, 0x1f, 0x0a, 0x03, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6a, 0x73, 0x6f, 0x6e,
	0x53, 0x74, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6a, 0x73, 0x6f, 0x6e, 0x53,
	0x74, 0x72, 0x22, 0x21, 0x0a, 0x03, 0x52, 0x65, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x61, 0x63,
	0x6b, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x61, 0x63,
	0x6b, 0x4a, 0x73, 0x6f, 0x6e, 0x32, 0x27, 0x0a, 0x06, 0x57, 0x61, 0x69, 0x74, 0x65, 0x72, 0x12,
	0x1d, 0x0a, 0x05, 0x44, 0x6f, 0x4d, 0x44, 0x35, 0x12, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52,
	0x65, 0x71, 0x1a, 0x08, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x04,
	0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_proto_rawDescOnce sync.Once
	file_test_proto_rawDescData = file_test_proto_rawDesc
)

func file_test_proto_rawDescGZIP() []byte {
	file_test_proto_rawDescOnce.Do(func() {
		file_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_proto_rawDescData)
	})
	return file_test_proto_rawDescData
}

var file_test_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_test_proto_goTypes = []interface{}{
	(*Req)(nil), // 0: api.Req
	(*Res)(nil), // 1: api.Res
}
var file_test_proto_depIdxs = []int32{
	0, // 0: api.Waiter.DoMD5:input_type -> api.Req
	1, // 1: api.Waiter.DoMD5:output_type -> api.Res
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_proto_init() }
func file_test_proto_init() {
	if File_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Req); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Res); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_test_proto_goTypes,
		DependencyIndexes: file_test_proto_depIdxs,
		MessageInfos:      file_test_proto_msgTypes,
	}.Build()
	File_test_proto = out.File
	file_test_proto_rawDesc = nil
	file_test_proto_goTypes = nil
	file_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WaiterClient is the client API for Waiter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WaiterClient interface {
	DoMD5(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error)
}

type waiterClient struct {
	cc grpc.ClientConnInterface
}

func NewWaiterClient(cc grpc.ClientConnInterface) WaiterClient {
	return &waiterClient{cc}
}

func (c *waiterClient) DoMD5(ctx context.Context, in *Req, opts ...grpc.CallOption) (*Res, error) {
	out := new(Res)
	err := c.cc.Invoke(ctx, "/api.Waiter/DoMD5", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaiterServer is the server API for Waiter service.
type WaiterServer interface {
	DoMD5(context.Context, *Req) (*Res, error)
}

// UnimplementedWaiterServer can be embedded to have forward compatible implementations.
type UnimplementedWaiterServer struct {
}

func (*UnimplementedWaiterServer) DoMD5(context.Context, *Req) (*Res, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoMD5 not implemented")
}

func RegisterWaiterServer(s *grpc.Server, srv WaiterServer) {
	s.RegisterService(&_Waiter_serviceDesc, srv)
}

func _Waiter_DoMD5_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Req)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaiterServer).DoMD5(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Waiter/DoMD5",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaiterServer).DoMD5(ctx, req.(*Req))
	}
	return interceptor(ctx, in, info, handler)
}

var _Waiter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Waiter",
	HandlerType: (*WaiterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoMD5",
			Handler:    _Waiter_DoMD5_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "test.proto",
}
