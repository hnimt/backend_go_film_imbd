// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: film.proto

package pb_film

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

type Film struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID         string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name       string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Year       string  `protobuf:"bytes,3,opt,name=Year,proto3" json:"Year,omitempty"`
	Rating     float64 `protobuf:"fixed64,4,opt,name=Rating,proto3" json:"Rating,omitempty"`
	Image      string  `protobuf:"bytes,5,opt,name=Image,proto3" json:"Image,omitempty"`
	Bookmarked string  `protobuf:"bytes,6,opt,name=Bookmarked,proto3" json:"Bookmarked,omitempty"`
}

func (x *Film) Reset() {
	*x = Film{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Film) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Film) ProtoMessage() {}

func (x *Film) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Film.ProtoReflect.Descriptor instead.
func (*Film) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{0}
}

func (x *Film) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Film) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Film) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Film) GetRating() float64 {
	if x != nil {
		return x.Rating
	}
	return 0
}

func (x *Film) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Film) GetBookmarked() string {
	if x != nil {
		return x.Bookmarked
	}
	return ""
}

type ReqAllFilms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReqAllFilms) Reset() {
	*x = ReqAllFilms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqAllFilms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqAllFilms) ProtoMessage() {}

func (x *ReqAllFilms) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqAllFilms.ProtoReflect.Descriptor instead.
func (*ReqAllFilms) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{1}
}

type ResAllFilms struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Films []*Film `protobuf:"bytes,1,rep,name=Films,proto3" json:"Films,omitempty"`
}

func (x *ResAllFilms) Reset() {
	*x = ResAllFilms{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResAllFilms) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResAllFilms) ProtoMessage() {}

func (x *ResAllFilms) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResAllFilms.ProtoReflect.Descriptor instead.
func (*ResAllFilms) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{2}
}

func (x *ResAllFilms) GetFilms() []*Film {
	if x != nil {
		return x.Films
	}
	return nil
}

type ReqFindFilmsByUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
}

func (x *ReqFindFilmsByUser) Reset() {
	*x = ReqFindFilmsByUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqFindFilmsByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqFindFilmsByUser) ProtoMessage() {}

func (x *ReqFindFilmsByUser) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqFindFilmsByUser.ProtoReflect.Descriptor instead.
func (*ReqFindFilmsByUser) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{3}
}

func (x *ReqFindFilmsByUser) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

type ResFindFilmsByUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Films []*Film `protobuf:"bytes,1,rep,name=Films,proto3" json:"Films,omitempty"`
}

func (x *ResFindFilmsByUser) Reset() {
	*x = ResFindFilmsByUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_film_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResFindFilmsByUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResFindFilmsByUser) ProtoMessage() {}

func (x *ResFindFilmsByUser) ProtoReflect() protoreflect.Message {
	mi := &file_film_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResFindFilmsByUser.ProtoReflect.Descriptor instead.
func (*ResFindFilmsByUser) Descriptor() ([]byte, []int) {
	return file_film_proto_rawDescGZIP(), []int{4}
}

func (x *ResFindFilmsByUser) GetFilms() []*Film {
	if x != nil {
		return x.Films
	}
	return nil
}

var File_film_proto protoreflect.FileDescriptor

var file_film_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x62,
	0x5f, 0x66, 0x69, 0x6c, 0x6d, 0x22, 0x8c, 0x01, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x59, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x59, 0x65, 0x61, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x14,
	0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49,
	0x6d, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x72, 0x6b,
	0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x42, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x64, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x41, 0x6c, 0x6c, 0x46, 0x69,
	0x6c, 0x6d, 0x73, 0x22, 0x32, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c,
	0x6d, 0x73, 0x12, 0x23, 0x0a, 0x05, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x5f, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x6d,
	0x52, 0x05, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x22, 0x2c, 0x0a, 0x12, 0x52, 0x65, 0x71, 0x46, 0x69,
	0x6e, 0x64, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x39, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x46, 0x69, 0x6e, 0x64,
	0x46, 0x69, 0x6c, 0x6d, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x05, 0x46,
	0x69, 0x6c, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x70, 0x62, 0x5f,
	0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x46, 0x69, 0x6c, 0x6d, 0x52, 0x05, 0x46, 0x69, 0x6c, 0x6d, 0x73,
	0x32, 0x95, 0x01, 0x0a, 0x0a, 0x42, 0x69, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x38, 0x0a, 0x08, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x12, 0x14, 0x2e, 0x70, 0x62,
	0x5f, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x6d,
	0x73, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x5f, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x41,
	0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x46, 0x69, 0x6e,
	0x64, 0x46, 0x69, 0x6c, 0x6d, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70,
	0x62, 0x5f, 0x66, 0x69, 0x6c, 0x6d, 0x2e, 0x52, 0x65, 0x71, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x69,
	0x6c, 0x6d, 0x73, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x5f, 0x66,
	0x69, 0x6c, 0x6d, 0x2e, 0x52, 0x65, 0x73, 0x46, 0x69, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x6d, 0x73,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x3b, 0x70,
	0x62, 0x5f, 0x66, 0x69, 0x6c, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_film_proto_rawDescOnce sync.Once
	file_film_proto_rawDescData = file_film_proto_rawDesc
)

func file_film_proto_rawDescGZIP() []byte {
	file_film_proto_rawDescOnce.Do(func() {
		file_film_proto_rawDescData = protoimpl.X.CompressGZIP(file_film_proto_rawDescData)
	})
	return file_film_proto_rawDescData
}

var file_film_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_film_proto_goTypes = []interface{}{
	(*Film)(nil),               // 0: pb_film.Film
	(*ReqAllFilms)(nil),        // 1: pb_film.ReqAllFilms
	(*ResAllFilms)(nil),        // 2: pb_film.ResAllFilms
	(*ReqFindFilmsByUser)(nil), // 3: pb_film.ReqFindFilmsByUser
	(*ResFindFilmsByUser)(nil), // 4: pb_film.ResFindFilmsByUser
}
var file_film_proto_depIdxs = []int32{
	0, // 0: pb_film.ResAllFilms.Films:type_name -> pb_film.Film
	0, // 1: pb_film.ResFindFilmsByUser.Films:type_name -> pb_film.Film
	1, // 2: pb_film.BizService.AllFilms:input_type -> pb_film.ReqAllFilms
	3, // 3: pb_film.BizService.FindFilmsByUser:input_type -> pb_film.ReqFindFilmsByUser
	2, // 4: pb_film.BizService.AllFilms:output_type -> pb_film.ResAllFilms
	4, // 5: pb_film.BizService.FindFilmsByUser:output_type -> pb_film.ResFindFilmsByUser
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_film_proto_init() }
func file_film_proto_init() {
	if File_film_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_film_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Film); i {
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
		file_film_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqAllFilms); i {
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
		file_film_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResAllFilms); i {
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
		file_film_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqFindFilmsByUser); i {
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
		file_film_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResFindFilmsByUser); i {
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
			RawDescriptor: file_film_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_film_proto_goTypes,
		DependencyIndexes: file_film_proto_depIdxs,
		MessageInfos:      file_film_proto_msgTypes,
	}.Build()
	File_film_proto = out.File
	file_film_proto_rawDesc = nil
	file_film_proto_goTypes = nil
	file_film_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BizServiceClient is the client API for BizService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BizServiceClient interface {
	AllFilms(ctx context.Context, in *ReqAllFilms, opts ...grpc.CallOption) (*ResAllFilms, error)
	FindFilmsByUser(ctx context.Context, in *ReqFindFilmsByUser, opts ...grpc.CallOption) (*ResFindFilmsByUser, error)
}

type bizServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBizServiceClient(cc grpc.ClientConnInterface) BizServiceClient {
	return &bizServiceClient{cc}
}

func (c *bizServiceClient) AllFilms(ctx context.Context, in *ReqAllFilms, opts ...grpc.CallOption) (*ResAllFilms, error) {
	out := new(ResAllFilms)
	err := c.cc.Invoke(ctx, "/pb_film.BizService/AllFilms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bizServiceClient) FindFilmsByUser(ctx context.Context, in *ReqFindFilmsByUser, opts ...grpc.CallOption) (*ResFindFilmsByUser, error) {
	out := new(ResFindFilmsByUser)
	err := c.cc.Invoke(ctx, "/pb_film.BizService/FindFilmsByUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BizServiceServer is the server API for BizService service.
type BizServiceServer interface {
	AllFilms(context.Context, *ReqAllFilms) (*ResAllFilms, error)
	FindFilmsByUser(context.Context, *ReqFindFilmsByUser) (*ResFindFilmsByUser, error)
}

// UnimplementedBizServiceServer can be embedded to have forward compatible implementations.
type UnimplementedBizServiceServer struct {
}

func (*UnimplementedBizServiceServer) AllFilms(context.Context, *ReqAllFilms) (*ResAllFilms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllFilms not implemented")
}
func (*UnimplementedBizServiceServer) FindFilmsByUser(context.Context, *ReqFindFilmsByUser) (*ResFindFilmsByUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindFilmsByUser not implemented")
}

func RegisterBizServiceServer(s *grpc.Server, srv BizServiceServer) {
	s.RegisterService(&_BizService_serviceDesc, srv)
}

func _BizService_AllFilms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqAllFilms)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).AllFilms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_film.BizService/AllFilms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).AllFilms(ctx, req.(*ReqAllFilms))
	}
	return interceptor(ctx, in, info, handler)
}

func _BizService_FindFilmsByUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReqFindFilmsByUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BizServiceServer).FindFilmsByUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb_film.BizService/FindFilmsByUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BizServiceServer).FindFilmsByUser(ctx, req.(*ReqFindFilmsByUser))
	}
	return interceptor(ctx, in, info, handler)
}

var _BizService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb_film.BizService",
	HandlerType: (*BizServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AllFilms",
			Handler:    _BizService_AllFilms_Handler,
		},
		{
			MethodName: "FindFilmsByUser",
			Handler:    _BizService_FindFilmsByUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "film.proto",
}
