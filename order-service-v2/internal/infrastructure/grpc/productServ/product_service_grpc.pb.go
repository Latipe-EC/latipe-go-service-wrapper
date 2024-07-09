// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.2
// source: internal/infrastructure/grpc/productServ/product_service.proto

package productgrpc

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

// ProductServiceClient is the client API for ProductService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductServiceClient interface {
	CheckInStock(ctx context.Context, in *GetPurchaseProductRequest, opts ...grpc.CallOption) (*GetPurchaseItemResponse, error)
	UpdateQuantity(ctx context.Context, in *UpdateProductQuantityRequest, opts ...grpc.CallOption) (*UpdateProductQuantityResponse, error)
}

type productServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProductServiceClient(cc grpc.ClientConnInterface) ProductServiceClient {
	return &productServiceClient{cc}
}

func (c *productServiceClient) CheckInStock(ctx context.Context, in *GetPurchaseProductRequest, opts ...grpc.CallOption) (*GetPurchaseItemResponse, error) {
	out := new(GetPurchaseItemResponse)
	err := c.cc.Invoke(ctx, "/protobuf.ProductService/CheckInStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productServiceClient) UpdateQuantity(ctx context.Context, in *UpdateProductQuantityRequest, opts ...grpc.CallOption) (*UpdateProductQuantityResponse, error) {
	out := new(UpdateProductQuantityResponse)
	err := c.cc.Invoke(ctx, "/protobuf.ProductService/UpdateQuantity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductServiceServer is the server API for ProductService service.
// All implementations must embed UnimplementedProductServiceServer
// for forward compatibility
type ProductServiceServer interface {
	CheckInStock(context.Context, *GetPurchaseProductRequest) (*GetPurchaseItemResponse, error)
	UpdateQuantity(context.Context, *UpdateProductQuantityRequest) (*UpdateProductQuantityResponse, error)
	mustEmbedUnimplementedProductServiceServer()
}

// UnimplementedProductServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProductServiceServer struct {
}

func (UnimplementedProductServiceServer) CheckInStock(context.Context, *GetPurchaseProductRequest) (*GetPurchaseItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckInStock not implemented")
}
func (UnimplementedProductServiceServer) UpdateQuantity(context.Context, *UpdateProductQuantityRequest) (*UpdateProductQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQuantity not implemented")
}
func (UnimplementedProductServiceServer) mustEmbedUnimplementedProductServiceServer() {}

// UnsafeProductServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductServiceServer will
// result in compilation errors.
type UnsafeProductServiceServer interface {
	mustEmbedUnimplementedProductServiceServer()
}

func RegisterProductServiceServer(s grpc.ServiceRegistrar, srv ProductServiceServer) {
	s.RegisterService(&ProductService_ServiceDesc, srv)
}

func _ProductService_CheckInStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPurchaseProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).CheckInStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ProductService/CheckInStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).CheckInStock(ctx, req.(*GetPurchaseProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductService_UpdateQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProductQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductServiceServer).UpdateQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.ProductService/UpdateQuantity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductServiceServer).UpdateQuantity(ctx, req.(*UpdateProductQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductService_ServiceDesc is the grpc.ServiceDesc for ProductService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.ProductService",
	HandlerType: (*ProductServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckInStock",
			Handler:    _ProductService_CheckInStock_Handler,
		},
		{
			MethodName: "UpdateQuantity",
			Handler:    _ProductService_UpdateQuantity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/infrastructure/grpc/productServ/product_service.proto",
}