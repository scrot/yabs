// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: blogpost-service/blogposts.proto

package blogpost

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Blogpost struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title     string               `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content   string               `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Author    string               `protobuf:"bytes,4,opt,name=author,proto3" json:"author,omitempty"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Blogpost) Reset() {
	*x = Blogpost{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogpost_service_blogposts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Blogpost) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Blogpost) ProtoMessage() {}

func (x *Blogpost) ProtoReflect() protoreflect.Message {
	mi := &file_blogpost_service_blogposts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Blogpost.ProtoReflect.Descriptor instead.
func (*Blogpost) Descriptor() ([]byte, []int) {
	return file_blogpost_service_blogposts_proto_rawDescGZIP(), []int{0}
}

func (x *Blogpost) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Blogpost) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Blogpost) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Blogpost) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Blogpost) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Blogpost) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Created  bool      `protobuf:"varint,1,opt,name=created,proto3" json:"created,omitempty"`
	Blogpost *Blogpost `protobuf:"bytes,2,opt,name=blogpost,proto3" json:"blogpost,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_blogpost_service_blogposts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_blogpost_service_blogposts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_blogpost_service_blogposts_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetCreated() bool {
	if x != nil {
		return x.Created
	}
	return false
}

func (x *Response) GetBlogpost() *Blogpost {
	if x != nil {
		return x.Blogpost
	}
	return nil
}

var File_blogpost_service_blogposts_proto protoreflect.FileDescriptor

var file_blogpost_service_blogposts_proto_rawDesc = []byte{
	0x0a, 0x20, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd8, 0x01,
	0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x54, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x2e,
	0x0a, 0x08, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x42, 0x6c, 0x6f, 0x67,
	0x70, 0x6f, 0x73, 0x74, 0x52, 0x08, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x32, 0x4d,
	0x0a, 0x0f, 0x42, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x67, 0x70,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x2e, 0x42,
	0x6c, 0x6f, 0x67, 0x70, 0x6f, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x70, 0x6f,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_blogpost_service_blogposts_proto_rawDescOnce sync.Once
	file_blogpost_service_blogposts_proto_rawDescData = file_blogpost_service_blogposts_proto_rawDesc
)

func file_blogpost_service_blogposts_proto_rawDescGZIP() []byte {
	file_blogpost_service_blogposts_proto_rawDescOnce.Do(func() {
		file_blogpost_service_blogposts_proto_rawDescData = protoimpl.X.CompressGZIP(file_blogpost_service_blogposts_proto_rawDescData)
	})
	return file_blogpost_service_blogposts_proto_rawDescData
}

var file_blogpost_service_blogposts_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_blogpost_service_blogposts_proto_goTypes = []interface{}{
	(*Blogpost)(nil),            // 0: blogpost.Blogpost
	(*Response)(nil),            // 1: blogpost.Response
	(*timestamp.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_blogpost_service_blogposts_proto_depIdxs = []int32{
	2, // 0: blogpost.Blogpost.created_at:type_name -> google.protobuf.Timestamp
	2, // 1: blogpost.Blogpost.updated_at:type_name -> google.protobuf.Timestamp
	0, // 2: blogpost.Response.blogpost:type_name -> blogpost.Blogpost
	0, // 3: blogpost.BlogpostService.CreateBlogpost:input_type -> blogpost.Blogpost
	1, // 4: blogpost.BlogpostService.CreateBlogpost:output_type -> blogpost.Response
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_blogpost_service_blogposts_proto_init() }
func file_blogpost_service_blogposts_proto_init() {
	if File_blogpost_service_blogposts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_blogpost_service_blogposts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Blogpost); i {
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
		file_blogpost_service_blogposts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_blogpost_service_blogposts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_blogpost_service_blogposts_proto_goTypes,
		DependencyIndexes: file_blogpost_service_blogposts_proto_depIdxs,
		MessageInfos:      file_blogpost_service_blogposts_proto_msgTypes,
	}.Build()
	File_blogpost_service_blogposts_proto = out.File
	file_blogpost_service_blogposts_proto_rawDesc = nil
	file_blogpost_service_blogposts_proto_goTypes = nil
	file_blogpost_service_blogposts_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BlogpostServiceClient is the client API for BlogpostService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BlogpostServiceClient interface {
	CreateBlogpost(ctx context.Context, in *Blogpost, opts ...grpc.CallOption) (*Response, error)
}

type blogpostServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlogpostServiceClient(cc grpc.ClientConnInterface) BlogpostServiceClient {
	return &blogpostServiceClient{cc}
}

func (c *blogpostServiceClient) CreateBlogpost(ctx context.Context, in *Blogpost, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/blogpost.BlogpostService/CreateBlogpost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlogpostServiceServer is the server API for BlogpostService service.
type BlogpostServiceServer interface {
	CreateBlogpost(context.Context, *Blogpost) (*Response, error)
}

// UnimplementedBlogpostServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBlogpostServiceServer struct {
}

func (*UnimplementedBlogpostServiceServer) CreateBlogpost(context.Context, *Blogpost) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlogpost not implemented")
}

func RegisterBlogpostServiceServer(s *grpc.Server, srv BlogpostServiceServer) {
	s.RegisterService(&_BlogpostService_serviceDesc, srv)
}

func _BlogpostService_CreateBlogpost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Blogpost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlogpostServiceServer).CreateBlogpost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blogpost.BlogpostService/CreateBlogpost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlogpostServiceServer).CreateBlogpost(ctx, req.(*Blogpost))
	}
	return interceptor(ctx, in, info, handler)
}

var _BlogpostService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "blogpost.BlogpostService",
	HandlerType: (*BlogpostServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBlogpost",
			Handler:    _BlogpostService_CreateBlogpost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blogpost-service/blogposts.proto",
}
