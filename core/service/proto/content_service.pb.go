// Code generated by protoc-gen-go. DO NOT EDIT.
// source: content_service.proto

package proto // import "."

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PagingArticleRequest struct {
	Cat                  string   `protobuf:"bytes,1,opt,name=cat,proto3" json:"cat,omitempty"`
	Begin                int32    `protobuf:"zigzag32,2,opt,name=begin,proto3" json:"begin,omitempty"`
	Size                 int32    `protobuf:"zigzag32,3,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PagingArticleRequest) Reset()         { *m = PagingArticleRequest{} }
func (m *PagingArticleRequest) String() string { return proto.CompactTextString(m) }
func (*PagingArticleRequest) ProtoMessage()    {}
func (*PagingArticleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_dd70072697eddb23, []int{0}
}
func (m *PagingArticleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PagingArticleRequest.Unmarshal(m, b)
}
func (m *PagingArticleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PagingArticleRequest.Marshal(b, m, deterministic)
}
func (dst *PagingArticleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PagingArticleRequest.Merge(dst, src)
}
func (m *PagingArticleRequest) XXX_Size() int {
	return xxx_messageInfo_PagingArticleRequest.Size(m)
}
func (m *PagingArticleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PagingArticleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PagingArticleRequest proto.InternalMessageInfo

func (m *PagingArticleRequest) GetCat() string {
	if m != nil {
		return m.Cat
	}
	return ""
}

func (m *PagingArticleRequest) GetBegin() int32 {
	if m != nil {
		return m.Begin
	}
	return 0
}

func (m *PagingArticleRequest) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type ArticlesResponse struct {
	List                 []*SArticle `protobuf:"bytes,1,rep,name=List,proto3" json:"List,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ArticlesResponse) Reset()         { *m = ArticlesResponse{} }
func (m *ArticlesResponse) String() string { return proto.CompactTextString(m) }
func (*ArticlesResponse) ProtoMessage()    {}
func (*ArticlesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_dd70072697eddb23, []int{1}
}
func (m *ArticlesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ArticlesResponse.Unmarshal(m, b)
}
func (m *ArticlesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ArticlesResponse.Marshal(b, m, deterministic)
}
func (dst *ArticlesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ArticlesResponse.Merge(dst, src)
}
func (m *ArticlesResponse) XXX_Size() int {
	return xxx_messageInfo_ArticlesResponse.Size(m)
}
func (m *ArticlesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ArticlesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ArticlesResponse proto.InternalMessageInfo

func (m *ArticlesResponse) GetList() []*SArticle {
	if m != nil {
		return m.List
	}
	return nil
}

// * 文章
type SArticle struct {
	// * 编号
	Id int32 `protobuf:"zigzag32,1,opt,name=Id,proto3" json:"Id,omitempty"`
	// * 栏目编号
	CatId int32 `protobuf:"zigzag32,2,opt,name=CatId,proto3" json:"CatId,omitempty"`
	// * 标题
	Title string `protobuf:"bytes,3,opt,name=Title,proto3" json:"Title,omitempty"`
	// * 小标题
	SmallTitle string `protobuf:"bytes,4,opt,name=SmallTitle,proto3" json:"SmallTitle,omitempty"`
	// * 文章附图
	Thumbnail string `protobuf:"bytes,5,opt,name=Thumbnail,proto3" json:"Thumbnail,omitempty"`
	// * 重定向URL
	PublisherId int32 `protobuf:"zigzag32,6,opt,name=PublisherId,proto3" json:"PublisherId,omitempty"`
	// * 重定向URL
	Location string `protobuf:"bytes,7,opt,name=Location,proto3" json:"Location,omitempty"`
	// * 优先级,优先级越高，则置顶
	Priority int32 `protobuf:"zigzag32,8,opt,name=Priority,proto3" json:"Priority,omitempty"`
	// * 浏览钥匙
	AccessKey string `protobuf:"bytes,9,opt,name=AccessKey,proto3" json:"AccessKey,omitempty"`
	// * 文档内容
	Content string `protobuf:"bytes,10,opt,name=Content,proto3" json:"Content,omitempty"`
	// * 标签（关键词）
	Tags string `protobuf:"bytes,11,opt,name=Tags,proto3" json:"Tags,omitempty"`
	// * 显示次数
	ViewCount int32 `protobuf:"zigzag32,12,opt,name=ViewCount,proto3" json:"ViewCount,omitempty"`
	// * 排序序号
	SortNum int32 `protobuf:"zigzag32,13,opt,name=SortNum,proto3" json:"SortNum,omitempty"`
	// * 创建时间
	CreateTime int32 `protobuf:"zigzag32,14,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"`
	// * 最后修改时间
	UpdateTime           int32    `protobuf:"zigzag32,15,opt,name=UpdateTime,proto3" json:"UpdateTime,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SArticle) Reset()         { *m = SArticle{} }
func (m *SArticle) String() string { return proto.CompactTextString(m) }
func (*SArticle) ProtoMessage()    {}
func (*SArticle) Descriptor() ([]byte, []int) {
	return fileDescriptor_content_service_dd70072697eddb23, []int{2}
}
func (m *SArticle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SArticle.Unmarshal(m, b)
}
func (m *SArticle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SArticle.Marshal(b, m, deterministic)
}
func (dst *SArticle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SArticle.Merge(dst, src)
}
func (m *SArticle) XXX_Size() int {
	return xxx_messageInfo_SArticle.Size(m)
}
func (m *SArticle) XXX_DiscardUnknown() {
	xxx_messageInfo_SArticle.DiscardUnknown(m)
}

var xxx_messageInfo_SArticle proto.InternalMessageInfo

func (m *SArticle) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SArticle) GetCatId() int32 {
	if m != nil {
		return m.CatId
	}
	return 0
}

func (m *SArticle) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *SArticle) GetSmallTitle() string {
	if m != nil {
		return m.SmallTitle
	}
	return ""
}

func (m *SArticle) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *SArticle) GetPublisherId() int32 {
	if m != nil {
		return m.PublisherId
	}
	return 0
}

func (m *SArticle) GetLocation() string {
	if m != nil {
		return m.Location
	}
	return ""
}

func (m *SArticle) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func (m *SArticle) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *SArticle) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *SArticle) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *SArticle) GetViewCount() int32 {
	if m != nil {
		return m.ViewCount
	}
	return 0
}

func (m *SArticle) GetSortNum() int32 {
	if m != nil {
		return m.SortNum
	}
	return 0
}

func (m *SArticle) GetCreateTime() int32 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *SArticle) GetUpdateTime() int32 {
	if m != nil {
		return m.UpdateTime
	}
	return 0
}

func init() {
	proto.RegisterType((*PagingArticleRequest)(nil), "PagingArticleRequest")
	proto.RegisterType((*ArticlesResponse)(nil), "ArticlesResponse")
	proto.RegisterType((*SArticle)(nil), "SArticle")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContentServiceClient is the client API for ContentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContentServiceClient interface {
	// * 获取置顶的文章,cat
	QueryTopArticles(ctx context.Context, in *String, opts ...grpc.CallOption) (*ArticlesResponse, error)
	// * 获取分页文章
	QueryPagingArticles(ctx context.Context, in *PagingArticleRequest, opts ...grpc.CallOption) (*SPagingResult, error)
}

type contentServiceClient struct {
	cc *grpc.ClientConn
}

func NewContentServiceClient(cc *grpc.ClientConn) ContentServiceClient {
	return &contentServiceClient{cc}
}

func (c *contentServiceClient) QueryTopArticles(ctx context.Context, in *String, opts ...grpc.CallOption) (*ArticlesResponse, error) {
	out := new(ArticlesResponse)
	err := c.cc.Invoke(ctx, "/ContentService/QueryTopArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contentServiceClient) QueryPagingArticles(ctx context.Context, in *PagingArticleRequest, opts ...grpc.CallOption) (*SPagingResult, error) {
	out := new(SPagingResult)
	err := c.cc.Invoke(ctx, "/ContentService/QueryPagingArticles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContentServiceServer is the server API for ContentService service.
type ContentServiceServer interface {
	// * 获取置顶的文章,cat
	QueryTopArticles(context.Context, *String) (*ArticlesResponse, error)
	// * 获取分页文章
	QueryPagingArticles(context.Context, *PagingArticleRequest) (*SPagingResult, error)
}

func RegisterContentServiceServer(s *grpc.Server, srv ContentServiceServer) {
	s.RegisterService(&_ContentService_serviceDesc, srv)
}

func _ContentService_QueryTopArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(String)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).QueryTopArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/QueryTopArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).QueryTopArticles(ctx, req.(*String))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContentService_QueryPagingArticles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PagingArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContentServiceServer).QueryPagingArticles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ContentService/QueryPagingArticles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContentServiceServer).QueryPagingArticles(ctx, req.(*PagingArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ContentService",
	HandlerType: (*ContentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryTopArticles",
			Handler:    _ContentService_QueryTopArticles_Handler,
		},
		{
			MethodName: "QueryPagingArticles",
			Handler:    _ContentService_QueryPagingArticles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "content_service.proto",
}

func init() {
	proto.RegisterFile("content_service.proto", fileDescriptor_content_service_dd70072697eddb23)
}

var fileDescriptor_content_service_dd70072697eddb23 = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x8f, 0xd3, 0x30,
	0x10, 0xc5, 0xb7, 0x7f, 0x76, 0xdb, 0x4c, 0xa1, 0xb4, 0x66, 0x57, 0xb2, 0x2a, 0x40, 0x55, 0x4e,
	0x3d, 0x45, 0xb0, 0x1c, 0x91, 0x90, 0x96, 0x9e, 0x2a, 0x56, 0xa8, 0x24, 0x81, 0x03, 0x17, 0x94,
	0x26, 0xa3, 0xac, 0xa5, 0xd4, 0x0e, 0xb6, 0x03, 0x0a, 0x47, 0x3e, 0x0a, 0x9f, 0x14, 0x79, 0xdc,
	0xb4, 0x05, 0x71, 0x8a, 0xdf, 0xef, 0x69, 0xe6, 0x45, 0xf6, 0x83, 0x9b, 0x5c, 0x49, 0x8b, 0xd2,
	0x7e, 0x35, 0xa8, 0xbf, 0x8b, 0x1c, 0xa3, 0x5a, 0x2b, 0xab, 0x16, 0x13, 0x6b, 0xdb, 0xfa, 0x20,
	0xc2, 0x18, 0xae, 0xb7, 0x59, 0x29, 0x64, 0x79, 0xa7, 0xad, 0xc8, 0x2b, 0x8c, 0xf1, 0x5b, 0x83,
	0xc6, 0xb2, 0x19, 0x0c, 0xf2, 0xcc, 0xf2, 0xde, 0xb2, 0xb7, 0x0a, 0x62, 0x77, 0x64, 0xd7, 0x70,
	0xb9, 0xc3, 0x52, 0x48, 0xde, 0x5f, 0xf6, 0x56, 0xf3, 0xd8, 0x0b, 0xc6, 0x60, 0x68, 0xc4, 0x4f,
	0xe4, 0x03, 0x82, 0x74, 0x0e, 0x5f, 0xc1, 0xec, 0xb0, 0xcd, 0xc4, 0x68, 0x6a, 0x25, 0x0d, 0xb2,
	0xe7, 0x30, 0xbc, 0x17, 0xc6, 0x2d, 0x1c, 0xac, 0x26, 0xb7, 0x41, 0x94, 0x74, 0x79, 0x84, 0xc3,
	0xdf, 0x03, 0x18, 0x77, 0x88, 0x4d, 0xa1, 0xbf, 0x29, 0x28, 0x7a, 0x1e, 0xf7, 0x37, 0x85, 0x4b,
	0x5e, 0x67, 0x76, 0x53, 0x74, 0xc9, 0x24, 0x1c, 0x4d, 0x85, 0xad, 0x7c, 0x74, 0x10, 0x7b, 0xc1,
	0x5e, 0x00, 0x24, 0xfb, 0xac, 0xaa, 0xbc, 0x35, 0x24, 0xeb, 0x8c, 0xb0, 0x67, 0x10, 0xa4, 0x0f,
	0xcd, 0x7e, 0x27, 0x33, 0x51, 0xf1, 0x4b, 0xb2, 0x4f, 0x80, 0x2d, 0x61, 0xb2, 0x6d, 0x76, 0x95,
	0x30, 0x0f, 0xa8, 0x37, 0x05, 0xbf, 0xa2, 0xbc, 0x73, 0xc4, 0x16, 0x30, 0xbe, 0x57, 0x79, 0x66,
	0x85, 0x92, 0x7c, 0x44, 0xe3, 0x47, 0xed, 0xbc, 0xad, 0x16, 0x4a, 0x0b, 0xdb, 0xf2, 0x31, 0x8d,
	0x1e, 0xb5, 0xcb, 0xbd, 0xcb, 0x73, 0x34, 0xe6, 0x3d, 0xb6, 0x3c, 0xf0, 0xb9, 0x47, 0xc0, 0x38,
	0x8c, 0xd6, 0xfe, 0xad, 0x38, 0x90, 0xd7, 0x49, 0x77, 0xbf, 0x69, 0x56, 0x1a, 0x3e, 0x21, 0x4c,
	0x67, 0xb7, 0xeb, 0xb3, 0xc0, 0x1f, 0x6b, 0xd5, 0x48, 0xcb, 0x1f, 0x51, 0xd0, 0x09, 0xb8, 0x5d,
	0x89, 0xd2, 0xf6, 0x43, 0xb3, 0xe7, 0x8f, 0xc9, 0xeb, 0xa4, 0xbb, 0x9b, 0xb5, 0xc6, 0xcc, 0x62,
	0x2a, 0xf6, 0xc8, 0xa7, 0x64, 0x9e, 0x11, 0xe7, 0x7f, 0xaa, 0x8b, 0xce, 0x7f, 0xe2, 0xfd, 0x13,
	0xb9, 0xfd, 0xd5, 0x83, 0xe9, 0xe1, 0xbf, 0x12, 0xdf, 0x28, 0xf6, 0x12, 0x66, 0x1f, 0x1b, 0xd4,
	0x6d, 0xaa, 0xea, 0xee, 0xc9, 0xd9, 0x28, 0x4a, 0xac, 0x16, 0xb2, 0x5c, 0xcc, 0xa3, 0x7f, 0x6b,
	0x10, 0x5e, 0xb0, 0xb7, 0xf0, 0x94, 0x26, 0xfe, 0x6a, 0x9d, 0x61, 0x37, 0xd1, 0xff, 0x6a, 0xb8,
	0x98, 0x46, 0x89, 0xe7, 0x31, 0x9a, 0xa6, 0xb2, 0xe1, 0xc5, 0xbb, 0xe0, 0xcb, 0x28, 0x7a, 0x43,
	0xdd, 0xdd, 0x5d, 0xd1, 0xe7, 0xf5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xd1, 0xba, 0x39,
	0xe8, 0x02, 0x00, 0x00,
}
