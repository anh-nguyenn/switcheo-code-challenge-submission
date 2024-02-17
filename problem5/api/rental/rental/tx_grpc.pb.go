// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: rental/rental/tx.proto

package rental

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

const (
	Msg_UpdateParams_FullMethodName    = "/rental.rental.Msg/UpdateParams"
	Msg_RequestRental_FullMethodName   = "/rental.rental.Msg/RequestRental"
	Msg_ApproveRental_FullMethodName   = "/rental.rental.Msg/ApproveRental"
	Msg_RepayRental_FullMethodName     = "/rental.rental.Msg/RepayRental"
	Msg_LiquidateRental_FullMethodName = "/rental.rental.Msg/LiquidateRental"
	Msg_CancelRental_FullMethodName    = "/rental.rental.Msg/CancelRental"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
	RequestRental(ctx context.Context, in *MsgRequestRental, opts ...grpc.CallOption) (*MsgRequestRentalResponse, error)
	ApproveRental(ctx context.Context, in *MsgApproveRental, opts ...grpc.CallOption) (*MsgApproveRentalResponse, error)
	RepayRental(ctx context.Context, in *MsgRepayRental, opts ...grpc.CallOption) (*MsgRepayRentalResponse, error)
	LiquidateRental(ctx context.Context, in *MsgLiquidateRental, opts ...grpc.CallOption) (*MsgLiquidateRentalResponse, error)
	CancelRental(ctx context.Context, in *MsgCancelRental, opts ...grpc.CallOption) (*MsgCancelRentalResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestRental(ctx context.Context, in *MsgRequestRental, opts ...grpc.CallOption) (*MsgRequestRentalResponse, error) {
	out := new(MsgRequestRentalResponse)
	err := c.cc.Invoke(ctx, Msg_RequestRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveRental(ctx context.Context, in *MsgApproveRental, opts ...grpc.CallOption) (*MsgApproveRentalResponse, error) {
	out := new(MsgApproveRentalResponse)
	err := c.cc.Invoke(ctx, Msg_ApproveRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RepayRental(ctx context.Context, in *MsgRepayRental, opts ...grpc.CallOption) (*MsgRepayRentalResponse, error) {
	out := new(MsgRepayRentalResponse)
	err := c.cc.Invoke(ctx, Msg_RepayRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) LiquidateRental(ctx context.Context, in *MsgLiquidateRental, opts ...grpc.CallOption) (*MsgLiquidateRentalResponse, error) {
	out := new(MsgLiquidateRentalResponse)
	err := c.cc.Invoke(ctx, Msg_LiquidateRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelRental(ctx context.Context, in *MsgCancelRental, opts ...grpc.CallOption) (*MsgCancelRentalResponse, error) {
	out := new(MsgCancelRentalResponse)
	err := c.cc.Invoke(ctx, Msg_CancelRental_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// UpdateParams defines a (governance) operation for updating the module
	// parameters. The authority defaults to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	RequestRental(context.Context, *MsgRequestRental) (*MsgRequestRentalResponse, error)
	ApproveRental(context.Context, *MsgApproveRental) (*MsgApproveRentalResponse, error)
	RepayRental(context.Context, *MsgRepayRental) (*MsgRepayRentalResponse, error)
	LiquidateRental(context.Context, *MsgLiquidateRental) (*MsgLiquidateRentalResponse, error)
	CancelRental(context.Context, *MsgCancelRental) (*MsgCancelRentalResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
}
func (UnimplementedMsgServer) RequestRental(context.Context, *MsgRequestRental) (*MsgRequestRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRental not implemented")
}
func (UnimplementedMsgServer) ApproveRental(context.Context, *MsgApproveRental) (*MsgApproveRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveRental not implemented")
}
func (UnimplementedMsgServer) RepayRental(context.Context, *MsgRepayRental) (*MsgRepayRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepayRental not implemented")
}
func (UnimplementedMsgServer) LiquidateRental(context.Context, *MsgLiquidateRental) (*MsgLiquidateRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LiquidateRental not implemented")
}
func (UnimplementedMsgServer) CancelRental(context.Context, *MsgCancelRental) (*MsgCancelRentalResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelRental not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_UpdateParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdateParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateParams(ctx, req.(*MsgUpdateParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestRental)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RequestRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestRental(ctx, req.(*MsgRequestRental))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveRental)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ApproveRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveRental(ctx, req.(*MsgApproveRental))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RepayRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRepayRental)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RepayRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_RepayRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RepayRental(ctx, req.(*MsgRepayRental))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_LiquidateRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgLiquidateRental)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).LiquidateRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_LiquidateRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).LiquidateRental(ctx, req.(*MsgLiquidateRental))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelRental_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelRental)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelRental(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelRental_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelRental(ctx, req.(*MsgCancelRental))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rental.rental.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
		{
			MethodName: "RequestRental",
			Handler:    _Msg_RequestRental_Handler,
		},
		{
			MethodName: "ApproveRental",
			Handler:    _Msg_ApproveRental_Handler,
		},
		{
			MethodName: "RepayRental",
			Handler:    _Msg_RepayRental_Handler,
		},
		{
			MethodName: "LiquidateRental",
			Handler:    _Msg_LiquidateRental_Handler,
		},
		{
			MethodName: "CancelRental",
			Handler:    _Msg_CancelRental_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rental/rental/tx.proto",
}
