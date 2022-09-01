// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: api_docs/catalogs/read_service/protobuf/products/service_clients/products_service_client.proto

package products_service

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

// ProductsServiceClient is the client API for ProductsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductsServiceClient interface {
	CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductRes, error)
	UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductRes, error)
	GetProductById(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdRes, error)
}

type productsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductsServiceClient(cc grpc.ClientConnInterface) ProductsServiceClient {
	return &productsServiceClient{cc}
}

func (c *productsServiceClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductRes, error) {
	out := new(CreateProductRes)
	err := c.cc.Invoke(ctx, "/products_service.ProductsService/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductRes, error) {
	out := new(UpdateProductRes)
	err := c.cc.Invoke(ctx, "/products_service.ProductsService/UpdateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productsServiceClient) GetProductById(ctx context.Context, in *GetProductByIdReq, opts ...grpc.CallOption) (*GetProductByIdRes, error) {
	out := new(GetProductByIdRes)
	err := c.cc.Invoke(ctx, "/products_service.ProductsService/GetProductById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductsServiceServer is the server API for ProductsService service.
// All implementations should embed UnimplementedProductsServiceServer
// for forward compatibility
type ProductsServiceServer interface {
	CreateProduct(context.Context, *CreateProductReq) (*CreateProductRes, error)
	UpdateProduct(context.Context, *UpdateProductReq) (*UpdateProductRes, error)
	GetProductById(context.Context, *GetProductByIdReq) (*GetProductByIdRes, error)
}

// UnimplementedProductsServiceServer should be embedded to have forward compatible implementations.
type UnimplementedProductsServiceServer struct {
}

func (UnimplementedProductsServiceServer) CreateProduct(context.Context, *CreateProductReq) (*CreateProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductsServiceServer) UpdateProduct(context.Context, *UpdateProductReq) (*UpdateProductRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProduct not implemented")
}
func (UnimplementedProductsServiceServer) GetProductById(context.Context, *GetProductByIdReq) (*GetProductByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductById not implemented")
}

// UnsafeProductsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductsServiceServer will
// result in compilation http_errors.
type UnsafeProductsServiceServer interface {
	mustEmbedUnimplementedProductsServiceServer()
}

func RegisterProductsServiceServer(s grpc.ServiceRegistrar, srv ProductsServiceServer) {
	s.RegisterService(&ProductsService_ServiceDesc, srv)
}

func _ProductsService_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products_service.ProductsService/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).CreateProduct(ctx, req.(*CreateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_UpdateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).UpdateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products_service.ProductsService/UpdateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).UpdateProduct(ctx, req.(*UpdateProductReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductsService_GetProductById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProductByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductsServiceServer).GetProductById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/products_service.ProductsService/GetProductById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductsServiceServer).GetProductById(ctx, req.(*GetProductByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductsService_ServiceDesc is the grpc.ServiceDesc for ProductsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "products_service.ProductsService",
	HandlerType: (*ProductsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductsService_CreateProduct_Handler,
		},
		{
			MethodName: "UpdateProduct",
			Handler:    _ProductsService_UpdateProduct_Handler,
		},
		{
			MethodName: "GetProductById",
			Handler:    _ProductsService_GetProductById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api_docs/catalogs/read_service/protobuf/products/service_clients/products_service_client.proto",
}
