// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/srv_marketplace.proto

package service_marketplace

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MarketplaceClient is the client API for Marketplace service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MarketplaceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error)
	GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error)
	AddReview(ctx context.Context, in *AddReviewRequest, opts ...grpc.CallOption) (*AddReviewResponse, error)
	GetReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error)
	UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*UpdateCartResponse, error)
	GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error)
}

type marketplaceClient struct {
	cc grpc.ClientConnInterface
}

func NewMarketplaceClient(cc grpc.ClientConnInterface) MarketplaceClient {
	return &marketplaceClient{cc}
}

func (c *marketplaceClient) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*CreateProductResponse, error) {
	out := new(CreateProductResponse)
	err := c.cc.Invoke(ctx, "/api_service_marketplace.Marketplace/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) GetProduct(ctx context.Context, in *GetProductRequest, opts ...grpc.CallOption) (*GetProductResponse, error) {
	out := new(GetProductResponse)
	err := c.cc.Invoke(ctx, "/api_service_marketplace.Marketplace/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) AddReview(ctx context.Context, in *AddReviewRequest, opts ...grpc.CallOption) (*AddReviewResponse, error) {
	out := new(AddReviewResponse)
	err := c.cc.Invoke(ctx, "/api_service_marketplace.Marketplace/AddReview", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) GetReviews(ctx context.Context, in *GetReviewsRequest, opts ...grpc.CallOption) (*GetReviewsResponse, error) {
	out := new(GetReviewsResponse)
	err := c.cc.Invoke(ctx, "/api_service_marketplace.Marketplace/GetReviews", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) UpdateCart(ctx context.Context, in *UpdateCartRequest, opts ...grpc.CallOption) (*UpdateCartResponse, error) {
	out := new(UpdateCartResponse)
	err := c.cc.Invoke(ctx, "/api_service_marketplace.Marketplace/UpdateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *marketplaceClient) GetCart(ctx context.Context, in *GetCartRequest, opts ...grpc.CallOption) (*GetCartResponse, error) {
	out := new(GetCartResponse)
	err := c.cc.Invoke(ctx, "/api_service_marketplace.Marketplace/GetCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MarketplaceServer is the server API for Marketplace service.
// All implementations must embed UnimplementedMarketplaceServer
// for forward compatibility
type MarketplaceServer interface {
	CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error)
	GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error)
	AddReview(context.Context, *AddReviewRequest) (*AddReviewResponse, error)
	GetReviews(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error)
	UpdateCart(context.Context, *UpdateCartRequest) (*UpdateCartResponse, error)
	GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error)
	mustEmbedUnimplementedMarketplaceServer()
}

// UnimplementedMarketplaceServer must be embedded to have forward compatible implementations.
type UnimplementedMarketplaceServer struct {
}

func (UnimplementedMarketplaceServer) CreateProduct(context.Context, *CreateProductRequest) (*CreateProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedMarketplaceServer) GetProduct(context.Context, *GetProductRequest) (*GetProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedMarketplaceServer) AddReview(context.Context, *AddReviewRequest) (*AddReviewResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddReview not implemented")
}
func (UnimplementedMarketplaceServer) GetReviews(context.Context, *GetReviewsRequest) (*GetReviewsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReviews not implemented")
}
func (UnimplementedMarketplaceServer) UpdateCart(context.Context, *UpdateCartRequest) (*UpdateCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCart not implemented")
}
func (UnimplementedMarketplaceServer) GetCart(context.Context, *GetCartRequest) (*GetCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCart not implemented")
}
func (UnimplementedMarketplaceServer) mustEmbedUnimplementedMarketplaceServer() {}

// UnsafeMarketplaceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MarketplaceServer will
// result in compilation errors.
type UnsafeMarketplaceServer interface {
	mustEmbedUnimplementedMarketplaceServer()
}

func RegisterMarketplaceServer(s grpc.ServiceRegistrar, srv MarketplaceServer) {
	s.RegisterService(&Marketplace_ServiceDesc, srv)
}

func _Marketplace_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_service_marketplace.Marketplace/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).CreateProduct(ctx, req.(*CreateProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_service_marketplace.Marketplace/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).GetProduct(ctx, req.(*GetProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_AddReview_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddReviewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).AddReview(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_service_marketplace.Marketplace/AddReview",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).AddReview(ctx, req.(*AddReviewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_GetReviews_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReviewsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).GetReviews(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_service_marketplace.Marketplace/GetReviews",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).GetReviews(ctx, req.(*GetReviewsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_UpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).UpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_service_marketplace.Marketplace/UpdateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).UpdateCart(ctx, req.(*UpdateCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Marketplace_GetCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MarketplaceServer).GetCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api_service_marketplace.Marketplace/GetCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MarketplaceServer).GetCart(ctx, req.(*GetCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Marketplace_ServiceDesc is the grpc.ServiceDesc for Marketplace service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Marketplace_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api_service_marketplace.Marketplace",
	HandlerType: (*MarketplaceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _Marketplace_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _Marketplace_GetProduct_Handler,
		},
		{
			MethodName: "AddReview",
			Handler:    _Marketplace_AddReview_Handler,
		},
		{
			MethodName: "GetReviews",
			Handler:    _Marketplace_GetReviews_Handler,
		},
		{
			MethodName: "UpdateCart",
			Handler:    _Marketplace_UpdateCart_Handler,
		},
		{
			MethodName: "GetCart",
			Handler:    _Marketplace_GetCart_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/srv_marketplace.proto",
}
