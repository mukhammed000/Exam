// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: progress.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ProgressService_GetProgressByLanguage_FullMethodName    = "/progress.ProgressService/GetProgressByLanguage"
	ProgressService_GetLeaderboardByLanguage_FullMethodName = "/progress.ProgressService/GetLeaderboardByLanguage"
	ProgressService_GetSkillsByLanguage_FullMethodName      = "/progress.ProgressService/GetSkillsByLanguage"
	ProgressService_GetDailyProgress_FullMethodName         = "/progress.ProgressService/GetDailyProgress"
	ProgressService_GetWeeklyProgress_FullMethodName        = "/progress.ProgressService/GetWeeklyProgress"
	ProgressService_GetMonthlyProgress_FullMethodName       = "/progress.ProgressService/GetMonthlyProgress"
	ProgressService_GetAchievements_FullMethodName          = "/progress.ProgressService/GetAchievements"
)

// ProgressServiceClient is the client API for ProgressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProgressServiceClient interface {
	GetProgressByLanguage(ctx context.Context, in *ProgressByLanguageReq, opts ...grpc.CallOption) (*ProgressByLanguageResponse, error)
	GetLeaderboardByLanguage(ctx context.Context, in *ProgressByLanguageRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error)
	GetSkillsByLanguage(ctx context.Context, in *ProgressByLanguageReq, opts ...grpc.CallOption) (*SkillsResponse, error)
	GetDailyProgress(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*DailyProgressResponse, error)
	GetWeeklyProgress(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*WeeklyProgressResponse, error)
	GetMonthlyProgress(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*MonthlyProgressResponse, error)
	GetAchievements(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*AchievementsResponse, error)
}

type progressServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProgressServiceClient(cc grpc.ClientConnInterface) ProgressServiceClient {
	return &progressServiceClient{cc}
}

func (c *progressServiceClient) GetProgressByLanguage(ctx context.Context, in *ProgressByLanguageReq, opts ...grpc.CallOption) (*ProgressByLanguageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProgressByLanguageResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetProgressByLanguage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetLeaderboardByLanguage(ctx context.Context, in *ProgressByLanguageRequest, opts ...grpc.CallOption) (*LeaderboardResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaderboardResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetLeaderboardByLanguage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetSkillsByLanguage(ctx context.Context, in *ProgressByLanguageReq, opts ...grpc.CallOption) (*SkillsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SkillsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetSkillsByLanguage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetDailyProgress(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*DailyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DailyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetDailyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetWeeklyProgress(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*WeeklyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WeeklyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetWeeklyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetMonthlyProgress(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*MonthlyProgressResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MonthlyProgressResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetMonthlyProgress_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *progressServiceClient) GetAchievements(ctx context.Context, in *UserId, opts ...grpc.CallOption) (*AchievementsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AchievementsResponse)
	err := c.cc.Invoke(ctx, ProgressService_GetAchievements_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProgressServiceServer is the server API for ProgressService service.
// All implementations must embed UnimplementedProgressServiceServer
// for forward compatibility
type ProgressServiceServer interface {
	GetProgressByLanguage(context.Context, *ProgressByLanguageReq) (*ProgressByLanguageResponse, error)
	GetLeaderboardByLanguage(context.Context, *ProgressByLanguageRequest) (*LeaderboardResponse, error)
	GetSkillsByLanguage(context.Context, *ProgressByLanguageReq) (*SkillsResponse, error)
	GetDailyProgress(context.Context, *UserId) (*DailyProgressResponse, error)
	GetWeeklyProgress(context.Context, *UserId) (*WeeklyProgressResponse, error)
	GetMonthlyProgress(context.Context, *UserId) (*MonthlyProgressResponse, error)
	GetAchievements(context.Context, *UserId) (*AchievementsResponse, error)
	mustEmbedUnimplementedProgressServiceServer()
}

// UnimplementedProgressServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProgressServiceServer struct {
}

func (UnimplementedProgressServiceServer) GetProgressByLanguage(context.Context, *ProgressByLanguageReq) (*ProgressByLanguageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProgressByLanguage not implemented")
}
func (UnimplementedProgressServiceServer) GetLeaderboardByLanguage(context.Context, *ProgressByLanguageRequest) (*LeaderboardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLeaderboardByLanguage not implemented")
}
func (UnimplementedProgressServiceServer) GetSkillsByLanguage(context.Context, *ProgressByLanguageReq) (*SkillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSkillsByLanguage not implemented")
}
func (UnimplementedProgressServiceServer) GetDailyProgress(context.Context, *UserId) (*DailyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDailyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetWeeklyProgress(context.Context, *UserId) (*WeeklyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWeeklyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetMonthlyProgress(context.Context, *UserId) (*MonthlyProgressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMonthlyProgress not implemented")
}
func (UnimplementedProgressServiceServer) GetAchievements(context.Context, *UserId) (*AchievementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAchievements not implemented")
}
func (UnimplementedProgressServiceServer) mustEmbedUnimplementedProgressServiceServer() {}

// UnsafeProgressServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProgressServiceServer will
// result in compilation errors.
type UnsafeProgressServiceServer interface {
	mustEmbedUnimplementedProgressServiceServer()
}

func RegisterProgressServiceServer(s grpc.ServiceRegistrar, srv ProgressServiceServer) {
	s.RegisterService(&ProgressService_ServiceDesc, srv)
}

func _ProgressService_GetProgressByLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgressByLanguageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetProgressByLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetProgressByLanguage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetProgressByLanguage(ctx, req.(*ProgressByLanguageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetLeaderboardByLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgressByLanguageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetLeaderboardByLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetLeaderboardByLanguage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetLeaderboardByLanguage(ctx, req.(*ProgressByLanguageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetSkillsByLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProgressByLanguageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetSkillsByLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetSkillsByLanguage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetSkillsByLanguage(ctx, req.(*ProgressByLanguageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetDailyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetDailyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetDailyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetDailyProgress(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetWeeklyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetWeeklyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetWeeklyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetWeeklyProgress(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetMonthlyProgress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetMonthlyProgress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetMonthlyProgress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetMonthlyProgress(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProgressService_GetAchievements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProgressServiceServer).GetAchievements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProgressService_GetAchievements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProgressServiceServer).GetAchievements(ctx, req.(*UserId))
	}
	return interceptor(ctx, in, info, handler)
}

// ProgressService_ServiceDesc is the grpc.ServiceDesc for ProgressService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProgressService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "progress.ProgressService",
	HandlerType: (*ProgressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProgressByLanguage",
			Handler:    _ProgressService_GetProgressByLanguage_Handler,
		},
		{
			MethodName: "GetLeaderboardByLanguage",
			Handler:    _ProgressService_GetLeaderboardByLanguage_Handler,
		},
		{
			MethodName: "GetSkillsByLanguage",
			Handler:    _ProgressService_GetSkillsByLanguage_Handler,
		},
		{
			MethodName: "GetDailyProgress",
			Handler:    _ProgressService_GetDailyProgress_Handler,
		},
		{
			MethodName: "GetWeeklyProgress",
			Handler:    _ProgressService_GetWeeklyProgress_Handler,
		},
		{
			MethodName: "GetMonthlyProgress",
			Handler:    _ProgressService_GetMonthlyProgress_Handler,
		},
		{
			MethodName: "GetAchievements",
			Handler:    _ProgressService_GetAchievements_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "progress.proto",
}
