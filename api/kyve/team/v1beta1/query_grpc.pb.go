// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package teamv1beta1

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

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// TeamInfo queries all important information from the team module
	TeamInfo(ctx context.Context, in *QueryTeamInfoRequest, opts ...grpc.CallOption) (*QueryTeamInfoResponse, error)
	// TeamVestingAccounts queries all team vesting accounts of the module.
	TeamVestingAccounts(ctx context.Context, in *QueryTeamVestingAccountsRequest, opts ...grpc.CallOption) (*QueryTeamVestingAccountsResponse, error)
	// TeamVestingAccount queries the team vesting accounts of the module.
	TeamVestingAccount(ctx context.Context, in *QueryTeamVestingAccountRequest, opts ...grpc.CallOption) (*QueryTeamVestingAccountResponse, error)
	// TeamCurrentVestingStatus queries the current vesting progress of a team vesting account
	TeamVestingStatus(ctx context.Context, in *QueryTeamVestingStatusRequest, opts ...grpc.CallOption) (*QueryTeamVestingStatusResponse, error)
	// TeamCurrentVestingStatus queries the current vesting progress of a team vesting account
	TeamVestingStatusByTime(ctx context.Context, in *QueryTeamVestingStatusByTimeRequest, opts ...grpc.CallOption) (*QueryTeamVestingStatusByTimeResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) TeamInfo(ctx context.Context, in *QueryTeamInfoRequest, opts ...grpc.CallOption) (*QueryTeamInfoResponse, error) {
	out := new(QueryTeamInfoResponse)
	err := c.cc.Invoke(ctx, "/kyve.team.v1beta1.Query/TeamInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TeamVestingAccounts(ctx context.Context, in *QueryTeamVestingAccountsRequest, opts ...grpc.CallOption) (*QueryTeamVestingAccountsResponse, error) {
	out := new(QueryTeamVestingAccountsResponse)
	err := c.cc.Invoke(ctx, "/kyve.team.v1beta1.Query/TeamVestingAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TeamVestingAccount(ctx context.Context, in *QueryTeamVestingAccountRequest, opts ...grpc.CallOption) (*QueryTeamVestingAccountResponse, error) {
	out := new(QueryTeamVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/kyve.team.v1beta1.Query/TeamVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TeamVestingStatus(ctx context.Context, in *QueryTeamVestingStatusRequest, opts ...grpc.CallOption) (*QueryTeamVestingStatusResponse, error) {
	out := new(QueryTeamVestingStatusResponse)
	err := c.cc.Invoke(ctx, "/kyve.team.v1beta1.Query/TeamVestingStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TeamVestingStatusByTime(ctx context.Context, in *QueryTeamVestingStatusByTimeRequest, opts ...grpc.CallOption) (*QueryTeamVestingStatusByTimeResponse, error) {
	out := new(QueryTeamVestingStatusByTimeResponse)
	err := c.cc.Invoke(ctx, "/kyve.team.v1beta1.Query/TeamVestingStatusByTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// TeamInfo queries all important information from the team module
	TeamInfo(context.Context, *QueryTeamInfoRequest) (*QueryTeamInfoResponse, error)
	// TeamVestingAccounts queries all team vesting accounts of the module.
	TeamVestingAccounts(context.Context, *QueryTeamVestingAccountsRequest) (*QueryTeamVestingAccountsResponse, error)
	// TeamVestingAccount queries the team vesting accounts of the module.
	TeamVestingAccount(context.Context, *QueryTeamVestingAccountRequest) (*QueryTeamVestingAccountResponse, error)
	// TeamCurrentVestingStatus queries the current vesting progress of a team vesting account
	TeamVestingStatus(context.Context, *QueryTeamVestingStatusRequest) (*QueryTeamVestingStatusResponse, error)
	// TeamCurrentVestingStatus queries the current vesting progress of a team vesting account
	TeamVestingStatusByTime(context.Context, *QueryTeamVestingStatusByTimeRequest) (*QueryTeamVestingStatusByTimeResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) TeamInfo(context.Context, *QueryTeamInfoRequest) (*QueryTeamInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamInfo not implemented")
}
func (UnimplementedQueryServer) TeamVestingAccounts(context.Context, *QueryTeamVestingAccountsRequest) (*QueryTeamVestingAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamVestingAccounts not implemented")
}
func (UnimplementedQueryServer) TeamVestingAccount(context.Context, *QueryTeamVestingAccountRequest) (*QueryTeamVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamVestingAccount not implemented")
}
func (UnimplementedQueryServer) TeamVestingStatus(context.Context, *QueryTeamVestingStatusRequest) (*QueryTeamVestingStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamVestingStatus not implemented")
}
func (UnimplementedQueryServer) TeamVestingStatusByTime(context.Context, *QueryTeamVestingStatusByTimeRequest) (*QueryTeamVestingStatusByTimeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TeamVestingStatusByTime not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_TeamInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTeamInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TeamInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.team.v1beta1.Query/TeamInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TeamInfo(ctx, req.(*QueryTeamInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TeamVestingAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTeamVestingAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TeamVestingAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.team.v1beta1.Query/TeamVestingAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TeamVestingAccounts(ctx, req.(*QueryTeamVestingAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TeamVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTeamVestingAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TeamVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.team.v1beta1.Query/TeamVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TeamVestingAccount(ctx, req.(*QueryTeamVestingAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TeamVestingStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTeamVestingStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TeamVestingStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.team.v1beta1.Query/TeamVestingStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TeamVestingStatus(ctx, req.(*QueryTeamVestingStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TeamVestingStatusByTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryTeamVestingStatusByTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TeamVestingStatusByTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kyve.team.v1beta1.Query/TeamVestingStatusByTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TeamVestingStatusByTime(ctx, req.(*QueryTeamVestingStatusByTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "kyve.team.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TeamInfo",
			Handler:    _Query_TeamInfo_Handler,
		},
		{
			MethodName: "TeamVestingAccounts",
			Handler:    _Query_TeamVestingAccounts_Handler,
		},
		{
			MethodName: "TeamVestingAccount",
			Handler:    _Query_TeamVestingAccount_Handler,
		},
		{
			MethodName: "TeamVestingStatus",
			Handler:    _Query_TeamVestingStatus_Handler,
		},
		{
			MethodName: "TeamVestingStatusByTime",
			Handler:    _Query_TeamVestingStatusByTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kyve/team/v1beta1/query.proto",
}