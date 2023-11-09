// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: kyve/pool/v1beta1/tx.proto

package poolv1beta1

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
	Msg_CreatePool_FullMethodName             = "/kyve.pool.v1beta1.Msg/CreatePool"
	Msg_UpdatePool_FullMethodName             = "/kyve.pool.v1beta1.Msg/UpdatePool"
	Msg_DisablePool_FullMethodName            = "/kyve.pool.v1beta1.Msg/DisablePool"
	Msg_EnablePool_FullMethodName             = "/kyve.pool.v1beta1.Msg/EnablePool"
	Msg_ScheduleRuntimeUpgrade_FullMethodName = "/kyve.pool.v1beta1.Msg/ScheduleRuntimeUpgrade"
	Msg_CancelRuntimeUpgrade_FullMethodName   = "/kyve.pool.v1beta1.Msg/CancelRuntimeUpgrade"
	Msg_UpdateParams_FullMethodName           = "/kyve.pool.v1beta1.Msg/UpdateParams"
)

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// CreatePool defines a governance operation for creating a new pool.
	// The authority is hard-coded to the x/gov module account.
	CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error)
	// UpdatePool defines a governance operation for updating an existing pool.
	// The authority is hard-coded to the x/gov module account.
	UpdatePool(ctx context.Context, in *MsgUpdatePool, opts ...grpc.CallOption) (*MsgUpdatePoolResponse, error)
	// DisablePool defines a governance operation for disabling an existing pool.
	// The authority is hard-coded to the x/gov module account.
	DisablePool(ctx context.Context, in *MsgDisablePool, opts ...grpc.CallOption) (*MsgDisablePoolResponse, error)
	// EnablePool defines a governance operation for enabling an existing pool.
	// The authority is hard-coded to the x/gov module account.
	EnablePool(ctx context.Context, in *MsgEnablePool, opts ...grpc.CallOption) (*MsgEnablePoolResponse, error)
	// ScheduleRuntimeUpgrade defines a governance operation for scheduling a runtime upgrade.
	// The authority is hard-coded to the x/gov module account.
	ScheduleRuntimeUpgrade(ctx context.Context, in *MsgScheduleRuntimeUpgrade, opts ...grpc.CallOption) (*MsgScheduleRuntimeUpgradeResponse, error)
	// CancelRuntimeUpgrade defines a governance operation for cancelling a runtime upgrade.
	// The authority is hard-coded to the x/gov module account.
	CancelRuntimeUpgrade(ctx context.Context, in *MsgCancelRuntimeUpgrade, opts ...grpc.CallOption) (*MsgCancelRuntimeUpgradeResponse, error)
	// UpdateParams defines a governance operation for updating the x/pool module
	// parameters. The authority is hard-coded to the x/gov module account.
	UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePool(ctx context.Context, in *MsgCreatePool, opts ...grpc.CallOption) (*MsgCreatePoolResponse, error) {
	out := new(MsgCreatePoolResponse)
	err := c.cc.Invoke(ctx, Msg_CreatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePool(ctx context.Context, in *MsgUpdatePool, opts ...grpc.CallOption) (*MsgUpdatePoolResponse, error) {
	out := new(MsgUpdatePoolResponse)
	err := c.cc.Invoke(ctx, Msg_UpdatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DisablePool(ctx context.Context, in *MsgDisablePool, opts ...grpc.CallOption) (*MsgDisablePoolResponse, error) {
	out := new(MsgDisablePoolResponse)
	err := c.cc.Invoke(ctx, Msg_DisablePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) EnablePool(ctx context.Context, in *MsgEnablePool, opts ...grpc.CallOption) (*MsgEnablePoolResponse, error) {
	out := new(MsgEnablePoolResponse)
	err := c.cc.Invoke(ctx, Msg_EnablePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ScheduleRuntimeUpgrade(ctx context.Context, in *MsgScheduleRuntimeUpgrade, opts ...grpc.CallOption) (*MsgScheduleRuntimeUpgradeResponse, error) {
	out := new(MsgScheduleRuntimeUpgradeResponse)
	err := c.cc.Invoke(ctx, Msg_ScheduleRuntimeUpgrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CancelRuntimeUpgrade(ctx context.Context, in *MsgCancelRuntimeUpgrade, opts ...grpc.CallOption) (*MsgCancelRuntimeUpgradeResponse, error) {
	out := new(MsgCancelRuntimeUpgradeResponse)
	err := c.cc.Invoke(ctx, Msg_CancelRuntimeUpgrade_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateParams(ctx context.Context, in *MsgUpdateParams, opts ...grpc.CallOption) (*MsgUpdateParamsResponse, error) {
	out := new(MsgUpdateParamsResponse)
	err := c.cc.Invoke(ctx, Msg_UpdateParams_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	// CreatePool defines a governance operation for creating a new pool.
	// The authority is hard-coded to the x/gov module account.
	CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error)
	// UpdatePool defines a governance operation for updating an existing pool.
	// The authority is hard-coded to the x/gov module account.
	UpdatePool(context.Context, *MsgUpdatePool) (*MsgUpdatePoolResponse, error)
	// DisablePool defines a governance operation for disabling an existing pool.
	// The authority is hard-coded to the x/gov module account.
	DisablePool(context.Context, *MsgDisablePool) (*MsgDisablePoolResponse, error)
	// EnablePool defines a governance operation for enabling an existing pool.
	// The authority is hard-coded to the x/gov module account.
	EnablePool(context.Context, *MsgEnablePool) (*MsgEnablePoolResponse, error)
	// ScheduleRuntimeUpgrade defines a governance operation for scheduling a runtime upgrade.
	// The authority is hard-coded to the x/gov module account.
	ScheduleRuntimeUpgrade(context.Context, *MsgScheduleRuntimeUpgrade) (*MsgScheduleRuntimeUpgradeResponse, error)
	// CancelRuntimeUpgrade defines a governance operation for cancelling a runtime upgrade.
	// The authority is hard-coded to the x/gov module account.
	CancelRuntimeUpgrade(context.Context, *MsgCancelRuntimeUpgrade) (*MsgCancelRuntimeUpgradeResponse, error)
	// UpdateParams defines a governance operation for updating the x/pool module
	// parameters. The authority is hard-coded to the x/gov module account.
	UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreatePool(context.Context, *MsgCreatePool) (*MsgCreatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedMsgServer) UpdatePool(context.Context, *MsgUpdatePool) (*MsgUpdatePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePool not implemented")
}
func (UnimplementedMsgServer) DisablePool(context.Context, *MsgDisablePool) (*MsgDisablePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisablePool not implemented")
}
func (UnimplementedMsgServer) EnablePool(context.Context, *MsgEnablePool) (*MsgEnablePoolResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnablePool not implemented")
}
func (UnimplementedMsgServer) ScheduleRuntimeUpgrade(context.Context, *MsgScheduleRuntimeUpgrade) (*MsgScheduleRuntimeUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleRuntimeUpgrade not implemented")
}
func (UnimplementedMsgServer) CancelRuntimeUpgrade(context.Context, *MsgCancelRuntimeUpgrade) (*MsgCancelRuntimeUpgradeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelRuntimeUpgrade not implemented")
}
func (UnimplementedMsgServer) UpdateParams(context.Context, *MsgUpdateParams) (*MsgUpdateParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateParams not implemented")
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

func _Msg_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePool(ctx, req.(*MsgCreatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_UpdatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePool(ctx, req.(*MsgUpdatePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DisablePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDisablePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DisablePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_DisablePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DisablePool(ctx, req.(*MsgDisablePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_EnablePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgEnablePool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).EnablePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_EnablePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).EnablePool(ctx, req.(*MsgEnablePool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ScheduleRuntimeUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgScheduleRuntimeUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ScheduleRuntimeUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_ScheduleRuntimeUpgrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ScheduleRuntimeUpgrade(ctx, req.(*MsgScheduleRuntimeUpgrade))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CancelRuntimeUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCancelRuntimeUpgrade)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CancelRuntimeUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Msg_CancelRuntimeUpgrade_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CancelRuntimeUpgrade(ctx, req.(*MsgCancelRuntimeUpgrade))
	}
	return interceptor(ctx, in, info, handler)
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

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kyve.pool.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _Msg_CreatePool_Handler,
		},
		{
			MethodName: "UpdatePool",
			Handler:    _Msg_UpdatePool_Handler,
		},
		{
			MethodName: "DisablePool",
			Handler:    _Msg_DisablePool_Handler,
		},
		{
			MethodName: "EnablePool",
			Handler:    _Msg_EnablePool_Handler,
		},
		{
			MethodName: "ScheduleRuntimeUpgrade",
			Handler:    _Msg_ScheduleRuntimeUpgrade_Handler,
		},
		{
			MethodName: "CancelRuntimeUpgrade",
			Handler:    _Msg_CancelRuntimeUpgrade_Handler,
		},
		{
			MethodName: "UpdateParams",
			Handler:    _Msg_UpdateParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kyve/pool/v1beta1/tx.proto",
}
