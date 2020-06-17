// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto

package v1

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

func init() { proto.RegisterFile("proto", fileDescriptor_00212fb1f9d3bf1c) }

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x5d, 0x0f, 0xa5, 0x8d, 0x20, 0x32, 0x08, 0x6a, 0x85, 0x3d, 0x2a, 0x82, 0xa4, 0x54,
	0x4f, 0x1e, 0x5b, 0xab, 0x27, 0x2f, 0xb6, 0x78, 0x29, 0x78, 0xd8, 0x6e, 0x07, 0x0c, 0x2c, 0xcd,
	0x34, 0xc9, 0x2e, 0x6c, 0x5f, 0xc3, 0x8b, 0x8f, 0xe4, 0xd1, 0x47, 0x90, 0x0a, 0x3e, 0x87, 0x24,
	0xc1, 0xa5, 0x0a, 0x66, 0x5b, 0x6f, 0xb3, 0xfc, 0xdf, 0x9f, 0x7f, 0x99, 0xfc, 0x61, 0xad, 0x84,
	0x04, 0x27, 0x25, 0x8d, 0x84, 0xa3, 0x49, 0xae, 0xc5, 0x0c, 0xb5, 0xe6, 0x1a, 0x55, 0x21, 0x52,
	0xe4, 0xda, 0x48, 0x85, 0xbc, 0xe8, 0xb6, 0x0f, 0xdc, 0xd4, 0x49, 0x48, 0x74, 0x14, 0xce, 0x73,
	0xd4, 0xc6, 0x7b, 0xda, 0x87, 0xab, 0x82, 0x26, 0x39, 0xd3, 0xe8, 0x95, 0x8b, 0xcf, 0x06, 0x6b,
	0x8e, 0xac, 0x38, 0x52, 0x05, 0x3c, 0xb2, 0x96, 0x9b, 0xef, 0x84, 0x36, 0x70, 0xca, 0xff, 0x0c,
	0xe2, 0x15, 0x35, 0xc4, 0x79, 0x7b, 0x4d, 0x90, 0x20, 0x65, 0x3b, 0xee, 0x7b, 0x80, 0x26, 0x11,
	0x19, 0x9c, 0xd5, 0xf9, 0x3c, 0x67, 0x23, 0xd6, 0x46, 0x09, 0x14, 0xdb, 0xbb, 0xcd, 0x17, 0x8b,
	0x72, 0x84, 0x89, 0x4a, 0x9f, 0x9c, 0x08, 0x3c, 0x60, 0xff, 0x0d, 0xdb, 0xb8, 0xcd, 0x78, 0x82,
	0x31, 0x6b, 0xf6, 0xa6, 0x53, 0x9f, 0x75, 0x12, 0xf0, 0x7e, 0x43, 0x36, 0x63, 0x3d, 0x8e, 0xec,
	0x9d, 0xdc, 0x4c, 0x85, 0xf1, 0x87, 0x87, 0x56, 0x5d, 0x51, 0x75, 0x77, 0xb2, 0x02, 0x12, 0x64,
	0x6c, 0xd7, 0xcd, 0xbd, 0xdc, 0xc8, 0x21, 0x52, 0x56, 0xc2, 0x79, 0xdd, 0xae, 0x2b, 0xd4, 0x06,
	0x6d, 0x42, 0x13, 0x3c, 0x47, 0xec, 0xf8, 0x3e, 0x47, 0x55, 0x56, 0xbd, 0xe8, 0x97, 0xd7, 0x72,
	0x66, 0x92, 0xd4, 0x3c, 0x68, 0x54, 0x70, 0x15, 0x38, 0x2d, 0xe0, 0xb3, 0x3f, 0xf2, 0x6f, 0xab,
	0xeb, 0xe5, 0x00, 0x33, 0x34, 0xe8, 0x97, 0x1c, 0x2a, 0xdb, 0x0a, 0x57, 0xd7, 0xcb, 0x1f, 0x28,
	0xf5, 0xf7, 0x5f, 0x97, 0x71, 0xf4, 0xb6, 0x8c, 0xa3, 0xf7, 0x65, 0x1c, 0xbd, 0x7c, 0xc4, 0x5b,
	0xe3, 0xed, 0xa2, 0x3b, 0x69, 0xb8, 0x57, 0x78, 0xf9, 0x15, 0x00, 0x00, 0xff, 0xff, 0x30, 0xe6,
	0x69, 0xdc, 0xe0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StoreSrvClient is the client API for StoreSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StoreSrvClient interface {
	// 线下门店列表
	StoreList(ctx context.Context, in *StoreListReq, opts ...grpc.CallOption) (*StoreListRep, error)
	// 线下门店详情
	StoreDetail(ctx context.Context, in *StoreDetailReq, opts ...grpc.CallOption) (*StoreDetailRep, error)
	// 全文检索门店
	FuzzySearchStore(ctx context.Context, in *FuzzySearchStoreReq, opts ...grpc.CallOption) (*FuzzySearchStoreRep, error)
	// 增加门店
	AddStore(ctx context.Context, in *AddStoreReq, opts ...grpc.CallOption) (*AddStoreRep, error)
	// 编辑门店
	EditStore(ctx context.Context, in *EditStoreReq, opts ...grpc.CallOption) (*EditStoreRep, error)
	// 门店联系人自动回复
	StoreAutoReply(ctx context.Context, in *StoreAutoReplyReq, opts ...grpc.CallOption) (*StoreAutoReplyRep, error)
	// 根据门店联系人获取门店列表
	QueryStoreListByContactUser(ctx context.Context, in *QueryStoreListByContactUserReq, opts ...grpc.CallOption) (*QueryStoreListByContactUserRep, error)
	// 删除门店
	DeleteStore(ctx context.Context, in *DeleteStoreReq, opts ...grpc.CallOption) (*DeleteStoreRep, error)
}

type storeSrvClient struct {
	cc *grpc.ClientConn
}

func NewStoreSrvClient(cc *grpc.ClientConn) StoreSrvClient {
	return &storeSrvClient{cc}
}

func (c *storeSrvClient) StoreList(ctx context.Context, in *StoreListReq, opts ...grpc.CallOption) (*StoreListRep, error) {
	out := new(StoreListRep)
	err := c.cc.Invoke(ctx, "/business.service.store.v1.StoreSrv/StoreList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeSrvClient) StoreDetail(ctx context.Context, in *StoreDetailReq, opts ...grpc.CallOption) (*StoreDetailRep, error) {
	out := new(StoreDetailRep)
	err := c.cc.Invoke(ctx, "/business.service.store.v1.StoreSrv/StoreDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeSrvClient) FuzzySearchStore(ctx context.Context, in *FuzzySearchStoreReq, opts ...grpc.CallOption) (*FuzzySearchStoreRep, error) {
	out := new(FuzzySearchStoreRep)
	err := c.cc.Invoke(ctx, "/business.service.store.v1.StoreSrv/FuzzySearchStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeSrvClient) AddStore(ctx context.Context, in *AddStoreReq, opts ...grpc.CallOption) (*AddStoreRep, error) {
	out := new(AddStoreRep)
	err := c.cc.Invoke(ctx, "/business.service.store.v1.StoreSrv/AddStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeSrvClient) EditStore(ctx context.Context, in *EditStoreReq, opts ...grpc.CallOption) (*EditStoreRep, error) {
	out := new(EditStoreRep)
	err := c.cc.Invoke(ctx, "/business.service.store.v1.StoreSrv/EditStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeSrvClient) StoreAutoReply(ctx context.Context, in *StoreAutoReplyReq, opts ...grpc.CallOption) (*StoreAutoReplyRep, error) {
	out := new(StoreAutoReplyRep)
	err := c.cc.Invoke(ctx, "/business.service.store.v1.StoreSrv/StoreAutoReply", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeSrvClient) QueryStoreListByContactUser(ctx context.Context, in *QueryStoreListByContactUserReq, opts ...grpc.CallOption) (*QueryStoreListByContactUserRep, error) {
	out := new(QueryStoreListByContactUserRep)
	err := c.cc.Invoke(ctx, "/business.service.store.v1.StoreSrv/QueryStoreListByContactUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeSrvClient) DeleteStore(ctx context.Context, in *DeleteStoreReq, opts ...grpc.CallOption) (*DeleteStoreRep, error) {
	out := new(DeleteStoreRep)
	err := c.cc.Invoke(ctx, "/business.service.store.v1.StoreSrv/DeleteStore", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreSrvServer is the server API for StoreSrv service.
type StoreSrvServer interface {
	// 线下门店列表
	StoreList(context.Context, *StoreListReq) (*StoreListRep, error)
	// 线下门店详情
	StoreDetail(context.Context, *StoreDetailReq) (*StoreDetailRep, error)
	// 全文检索门店
	FuzzySearchStore(context.Context, *FuzzySearchStoreReq) (*FuzzySearchStoreRep, error)
	// 增加门店
	AddStore(context.Context, *AddStoreReq) (*AddStoreRep, error)
	// 编辑门店
	EditStore(context.Context, *EditStoreReq) (*EditStoreRep, error)
	// 门店联系人自动回复
	StoreAutoReply(context.Context, *StoreAutoReplyReq) (*StoreAutoReplyRep, error)
	// 根据门店联系人获取门店列表
	QueryStoreListByContactUser(context.Context, *QueryStoreListByContactUserReq) (*QueryStoreListByContactUserRep, error)
	// 删除门店
	DeleteStore(context.Context, *DeleteStoreReq) (*DeleteStoreRep, error)
}

// UnimplementedStoreSrvServer can be embedded to have forward compatible implementations.
type UnimplementedStoreSrvServer struct {
}

func (*UnimplementedStoreSrvServer) StoreList(ctx context.Context, req *StoreListReq) (*StoreListRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreList not implemented")
}
func (*UnimplementedStoreSrvServer) StoreDetail(ctx context.Context, req *StoreDetailReq) (*StoreDetailRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreDetail not implemented")
}
func (*UnimplementedStoreSrvServer) FuzzySearchStore(ctx context.Context, req *FuzzySearchStoreReq) (*FuzzySearchStoreRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FuzzySearchStore not implemented")
}
func (*UnimplementedStoreSrvServer) AddStore(ctx context.Context, req *AddStoreReq) (*AddStoreRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStore not implemented")
}
func (*UnimplementedStoreSrvServer) EditStore(ctx context.Context, req *EditStoreReq) (*EditStoreRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EditStore not implemented")
}
func (*UnimplementedStoreSrvServer) StoreAutoReply(ctx context.Context, req *StoreAutoReplyReq) (*StoreAutoReplyRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoreAutoReply not implemented")
}
func (*UnimplementedStoreSrvServer) QueryStoreListByContactUser(ctx context.Context, req *QueryStoreListByContactUserReq) (*QueryStoreListByContactUserRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryStoreListByContactUser not implemented")
}
func (*UnimplementedStoreSrvServer) DeleteStore(ctx context.Context, req *DeleteStoreReq) (*DeleteStoreRep, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteStore not implemented")
}

func RegisterStoreSrvServer(s *grpc.Server, srv StoreSrvServer) {
	s.RegisterService(&_StoreSrv_serviceDesc, srv)
}

func _StoreSrv_StoreList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSrvServer).StoreList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.service.store.v1.StoreSrv/StoreList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSrvServer).StoreList(ctx, req.(*StoreListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreSrv_StoreDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreDetailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSrvServer).StoreDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.service.store.v1.StoreSrv/StoreDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSrvServer).StoreDetail(ctx, req.(*StoreDetailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreSrv_FuzzySearchStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FuzzySearchStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSrvServer).FuzzySearchStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.service.store.v1.StoreSrv/FuzzySearchStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSrvServer).FuzzySearchStore(ctx, req.(*FuzzySearchStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreSrv_AddStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSrvServer).AddStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.service.store.v1.StoreSrv/AddStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSrvServer).AddStore(ctx, req.(*AddStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreSrv_EditStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSrvServer).EditStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.service.store.v1.StoreSrv/EditStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSrvServer).EditStore(ctx, req.(*EditStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreSrv_StoreAutoReply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreAutoReplyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSrvServer).StoreAutoReply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.service.store.v1.StoreSrv/StoreAutoReply",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSrvServer).StoreAutoReply(ctx, req.(*StoreAutoReplyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreSrv_QueryStoreListByContactUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryStoreListByContactUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSrvServer).QueryStoreListByContactUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.service.store.v1.StoreSrv/QueryStoreListByContactUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSrvServer).QueryStoreListByContactUser(ctx, req.(*QueryStoreListByContactUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreSrv_DeleteStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteStoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreSrvServer).DeleteStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/business.service.store.v1.StoreSrv/DeleteStore",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreSrvServer).DeleteStore(ctx, req.(*DeleteStoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _StoreSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "business.service.store.v1.StoreSrv",
	HandlerType: (*StoreSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreList",
			Handler:    _StoreSrv_StoreList_Handler,
		},
		{
			MethodName: "StoreDetail",
			Handler:    _StoreSrv_StoreDetail_Handler,
		},
		{
			MethodName: "FuzzySearchStore",
			Handler:    _StoreSrv_FuzzySearchStore_Handler,
		},
		{
			MethodName: "AddStore",
			Handler:    _StoreSrv_AddStore_Handler,
		},
		{
			MethodName: "EditStore",
			Handler:    _StoreSrv_EditStore_Handler,
		},
		{
			MethodName: "StoreAutoReply",
			Handler:    _StoreSrv_StoreAutoReply_Handler,
		},
		{
			MethodName: "QueryStoreListByContactUser",
			Handler:    _StoreSrv_QueryStoreListByContactUser_Handler,
		},
		{
			MethodName: "DeleteStore",
			Handler:    _StoreSrv_DeleteStore_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto",
}
