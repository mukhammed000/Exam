// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: learn.proto

package learning

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
	LearningService_GetAllLanguages_FullMethodName       = "/learning.LearningService/GetAllLanguages"
	LearningService_StartLearningLanguage_FullMethodName = "/learning.LearningService/StartLearningLanguage"
	LearningService_GetLessonInfo_FullMethodName         = "/learning.LearningService/GetLessonInfo"
	LearningService_CompleteLesson_FullMethodName        = "/learning.LearningService/CompleteLesson"
	LearningService_StartExercise_FullMethodName         = "/learning.LearningService/StartExercise"
	LearningService_AnswerExercise_FullMethodName        = "/learning.LearningService/AnswerExercise"
	LearningService_GetAllVocabulary_FullMethodName      = "/learning.LearningService/GetAllVocabulary"
	LearningService_AddNewWord_FullMethodName            = "/learning.LearningService/AddNewWord"
)

// LearningServiceClient is the client API for LearningService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LearningServiceClient interface {
	GetAllLanguages(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*GetAllLanguagesResp, error)
	StartLearningLanguage(ctx context.Context, in *StartLearningReq, opts ...grpc.CallOption) (*StartLearningResp, error)
	GetLessonInfo(ctx context.Context, in *ParticipateLessonReq, opts ...grpc.CallOption) (*ParticipateLessonResp, error)
	CompleteLesson(ctx context.Context, in *CompleteLessonReq, opts ...grpc.CallOption) (*CompleteLessonResp, error)
	StartExercise(ctx context.Context, in *StartExerciceReq, opts ...grpc.CallOption) (*StartExerciceResp, error)
	AnswerExercise(ctx context.Context, in *AnswerExerciseReq, opts ...grpc.CallOption) (*AnswerExerciseResp, error)
	GetAllVocabulary(ctx context.Context, in *GetAllVocabularyReq, opts ...grpc.CallOption) (*GetAllVocabularyResp, error)
	AddNewWord(ctx context.Context, in *AddNewWordReq, opts ...grpc.CallOption) (*AddNewWordResp, error)
}

type learningServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLearningServiceClient(cc grpc.ClientConnInterface) LearningServiceClient {
	return &learningServiceClient{cc}
}

func (c *learningServiceClient) GetAllLanguages(ctx context.Context, in *EmptyReq, opts ...grpc.CallOption) (*GetAllLanguagesResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllLanguagesResp)
	err := c.cc.Invoke(ctx, LearningService_GetAllLanguages_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) StartLearningLanguage(ctx context.Context, in *StartLearningReq, opts ...grpc.CallOption) (*StartLearningResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartLearningResp)
	err := c.cc.Invoke(ctx, LearningService_StartLearningLanguage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetLessonInfo(ctx context.Context, in *ParticipateLessonReq, opts ...grpc.CallOption) (*ParticipateLessonResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ParticipateLessonResp)
	err := c.cc.Invoke(ctx, LearningService_GetLessonInfo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) CompleteLesson(ctx context.Context, in *CompleteLessonReq, opts ...grpc.CallOption) (*CompleteLessonResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CompleteLessonResp)
	err := c.cc.Invoke(ctx, LearningService_CompleteLesson_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) StartExercise(ctx context.Context, in *StartExerciceReq, opts ...grpc.CallOption) (*StartExerciceResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(StartExerciceResp)
	err := c.cc.Invoke(ctx, LearningService_StartExercise_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) AnswerExercise(ctx context.Context, in *AnswerExerciseReq, opts ...grpc.CallOption) (*AnswerExerciseResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnswerExerciseResp)
	err := c.cc.Invoke(ctx, LearningService_AnswerExercise_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) GetAllVocabulary(ctx context.Context, in *GetAllVocabularyReq, opts ...grpc.CallOption) (*GetAllVocabularyResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetAllVocabularyResp)
	err := c.cc.Invoke(ctx, LearningService_GetAllVocabulary_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *learningServiceClient) AddNewWord(ctx context.Context, in *AddNewWordReq, opts ...grpc.CallOption) (*AddNewWordResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddNewWordResp)
	err := c.cc.Invoke(ctx, LearningService_AddNewWord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LearningServiceServer is the server API for LearningService service.
// All implementations must embed UnimplementedLearningServiceServer
// for forward compatibility
type LearningServiceServer interface {
	GetAllLanguages(context.Context, *EmptyReq) (*GetAllLanguagesResp, error)
	StartLearningLanguage(context.Context, *StartLearningReq) (*StartLearningResp, error)
	GetLessonInfo(context.Context, *ParticipateLessonReq) (*ParticipateLessonResp, error)
	CompleteLesson(context.Context, *CompleteLessonReq) (*CompleteLessonResp, error)
	StartExercise(context.Context, *StartExerciceReq) (*StartExerciceResp, error)
	AnswerExercise(context.Context, *AnswerExerciseReq) (*AnswerExerciseResp, error)
	GetAllVocabulary(context.Context, *GetAllVocabularyReq) (*GetAllVocabularyResp, error)
	AddNewWord(context.Context, *AddNewWordReq) (*AddNewWordResp, error)
	mustEmbedUnimplementedLearningServiceServer()
}

// UnimplementedLearningServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLearningServiceServer struct {
}

func (UnimplementedLearningServiceServer) GetAllLanguages(context.Context, *EmptyReq) (*GetAllLanguagesResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllLanguages not implemented")
}
func (UnimplementedLearningServiceServer) StartLearningLanguage(context.Context, *StartLearningReq) (*StartLearningResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartLearningLanguage not implemented")
}
func (UnimplementedLearningServiceServer) GetLessonInfo(context.Context, *ParticipateLessonReq) (*ParticipateLessonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLessonInfo not implemented")
}
func (UnimplementedLearningServiceServer) CompleteLesson(context.Context, *CompleteLessonReq) (*CompleteLessonResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteLesson not implemented")
}
func (UnimplementedLearningServiceServer) StartExercise(context.Context, *StartExerciceReq) (*StartExerciceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartExercise not implemented")
}
func (UnimplementedLearningServiceServer) AnswerExercise(context.Context, *AnswerExerciseReq) (*AnswerExerciseResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerExercise not implemented")
}
func (UnimplementedLearningServiceServer) GetAllVocabulary(context.Context, *GetAllVocabularyReq) (*GetAllVocabularyResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllVocabulary not implemented")
}
func (UnimplementedLearningServiceServer) AddNewWord(context.Context, *AddNewWordReq) (*AddNewWordResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewWord not implemented")
}
func (UnimplementedLearningServiceServer) mustEmbedUnimplementedLearningServiceServer() {}

// UnsafeLearningServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LearningServiceServer will
// result in compilation errors.
type UnsafeLearningServiceServer interface {
	mustEmbedUnimplementedLearningServiceServer()
}

func RegisterLearningServiceServer(s grpc.ServiceRegistrar, srv LearningServiceServer) {
	s.RegisterService(&LearningService_ServiceDesc, srv)
}

func _LearningService_GetAllLanguages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetAllLanguages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningService_GetAllLanguages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetAllLanguages(ctx, req.(*EmptyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_StartLearningLanguage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartLearningReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).StartLearningLanguage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningService_StartLearningLanguage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).StartLearningLanguage(ctx, req.(*StartLearningReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetLessonInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParticipateLessonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetLessonInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningService_GetLessonInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetLessonInfo(ctx, req.(*ParticipateLessonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_CompleteLesson_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteLessonReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).CompleteLesson(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningService_CompleteLesson_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).CompleteLesson(ctx, req.(*CompleteLessonReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_StartExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartExerciceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).StartExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningService_StartExercise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).StartExercise(ctx, req.(*StartExerciceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_AnswerExercise_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerExerciseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).AnswerExercise(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningService_AnswerExercise_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).AnswerExercise(ctx, req.(*AnswerExerciseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_GetAllVocabulary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllVocabularyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).GetAllVocabulary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningService_GetAllVocabulary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).GetAllVocabulary(ctx, req.(*GetAllVocabularyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LearningService_AddNewWord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewWordReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LearningServiceServer).AddNewWord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LearningService_AddNewWord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LearningServiceServer).AddNewWord(ctx, req.(*AddNewWordReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LearningService_ServiceDesc is the grpc.ServiceDesc for LearningService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LearningService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "learning.LearningService",
	HandlerType: (*LearningServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAllLanguages",
			Handler:    _LearningService_GetAllLanguages_Handler,
		},
		{
			MethodName: "StartLearningLanguage",
			Handler:    _LearningService_StartLearningLanguage_Handler,
		},
		{
			MethodName: "GetLessonInfo",
			Handler:    _LearningService_GetLessonInfo_Handler,
		},
		{
			MethodName: "CompleteLesson",
			Handler:    _LearningService_CompleteLesson_Handler,
		},
		{
			MethodName: "StartExercise",
			Handler:    _LearningService_StartExercise_Handler,
		},
		{
			MethodName: "AnswerExercise",
			Handler:    _LearningService_AnswerExercise_Handler,
		},
		{
			MethodName: "GetAllVocabulary",
			Handler:    _LearningService_GetAllVocabulary_Handler,
		},
		{
			MethodName: "AddNewWord",
			Handler:    _LearningService_AddNewWord_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "learn.proto",
}
