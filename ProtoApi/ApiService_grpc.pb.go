// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.10
// source: ApiService.proto

package ProtoApi

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
	SheduleService_GetSubjectFromDbRPC_FullMethodName    = "/ProtoApi.SheduleService/GetSubjectFromDbRPC"
	SheduleService_GetTutorsFromDbRPC_FullMethodName     = "/ProtoApi.SheduleService/GetTutorsFromDbRPC"
	SheduleService_GetAuditoriumFromDbRPC_FullMethodName = "/ProtoApi.SheduleService/GetAuditoriumFromDbRPC"
	SheduleService_GetTypeFromDbRPC_FullMethodName       = "/ProtoApi.SheduleService/GetTypeFromDbRPC"
	SheduleService_GetGroupFromDbRPC_FullMethodName      = "/ProtoApi.SheduleService/GetGroupFromDbRPC"
	SheduleService_GetSheduleFromDb_FullMethodName       = "/ProtoApi.SheduleService/GetSheduleFromDb"
	SheduleService_AddShedule_FullMethodName             = "/ProtoApi.SheduleService/AddShedule"
)

// SheduleServiceClient is the client API for SheduleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SheduleServiceClient interface {
	GetSubjectFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*SubjectsFromDb, error)
	GetTutorsFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*TutorsFromDb, error)
	GetAuditoriumFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*AuditoriumsFromDb, error)
	GetTypeFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*TypesFromDb, error)
	GetGroupFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*GroupsFromDb, error)
	GetSheduleFromDb(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*SheduleArray, error)
	AddShedule(ctx context.Context, in *SheduleArray, opts ...grpc.CallOption) (*Wrap, error)
}

type sheduleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSheduleServiceClient(cc grpc.ClientConnInterface) SheduleServiceClient {
	return &sheduleServiceClient{cc}
}

func (c *sheduleServiceClient) GetSubjectFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*SubjectsFromDb, error) {
	out := new(SubjectsFromDb)
	err := c.cc.Invoke(ctx, SheduleService_GetSubjectFromDbRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheduleServiceClient) GetTutorsFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*TutorsFromDb, error) {
	out := new(TutorsFromDb)
	err := c.cc.Invoke(ctx, SheduleService_GetTutorsFromDbRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheduleServiceClient) GetAuditoriumFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*AuditoriumsFromDb, error) {
	out := new(AuditoriumsFromDb)
	err := c.cc.Invoke(ctx, SheduleService_GetAuditoriumFromDbRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheduleServiceClient) GetTypeFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*TypesFromDb, error) {
	out := new(TypesFromDb)
	err := c.cc.Invoke(ctx, SheduleService_GetTypeFromDbRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheduleServiceClient) GetGroupFromDbRPC(ctx context.Context, in *Wrap, opts ...grpc.CallOption) (*GroupsFromDb, error) {
	out := new(GroupsFromDb)
	err := c.cc.Invoke(ctx, SheduleService_GetGroupFromDbRPC_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheduleServiceClient) GetSheduleFromDb(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*SheduleArray, error) {
	out := new(SheduleArray)
	err := c.cc.Invoke(ctx, SheduleService_GetSheduleFromDb_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sheduleServiceClient) AddShedule(ctx context.Context, in *SheduleArray, opts ...grpc.CallOption) (*Wrap, error) {
	out := new(Wrap)
	err := c.cc.Invoke(ctx, SheduleService_AddShedule_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SheduleServiceServer is the server API for SheduleService service.
// All implementations must embed UnimplementedSheduleServiceServer
// for forward compatibility
type SheduleServiceServer interface {
	GetSubjectFromDbRPC(context.Context, *Wrap) (*SubjectsFromDb, error)
	GetTutorsFromDbRPC(context.Context, *Wrap) (*TutorsFromDb, error)
	GetAuditoriumFromDbRPC(context.Context, *Wrap) (*AuditoriumsFromDb, error)
	GetTypeFromDbRPC(context.Context, *Wrap) (*TypesFromDb, error)
	GetGroupFromDbRPC(context.Context, *Wrap) (*GroupsFromDb, error)
	GetSheduleFromDb(context.Context, *Filter) (*SheduleArray, error)
	AddShedule(context.Context, *SheduleArray) (*Wrap, error)
	mustEmbedUnimplementedSheduleServiceServer()
}

// UnimplementedSheduleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSheduleServiceServer struct {
}

func (UnimplementedSheduleServiceServer) GetSubjectFromDbRPC(context.Context, *Wrap) (*SubjectsFromDb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubjectFromDbRPC not implemented")
}
func (UnimplementedSheduleServiceServer) GetTutorsFromDbRPC(context.Context, *Wrap) (*TutorsFromDb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTutorsFromDbRPC not implemented")
}
func (UnimplementedSheduleServiceServer) GetAuditoriumFromDbRPC(context.Context, *Wrap) (*AuditoriumsFromDb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuditoriumFromDbRPC not implemented")
}
func (UnimplementedSheduleServiceServer) GetTypeFromDbRPC(context.Context, *Wrap) (*TypesFromDb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTypeFromDbRPC not implemented")
}
func (UnimplementedSheduleServiceServer) GetGroupFromDbRPC(context.Context, *Wrap) (*GroupsFromDb, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupFromDbRPC not implemented")
}
func (UnimplementedSheduleServiceServer) GetSheduleFromDb(context.Context, *Filter) (*SheduleArray, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSheduleFromDb not implemented")
}
func (UnimplementedSheduleServiceServer) AddShedule(context.Context, *SheduleArray) (*Wrap, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddShedule not implemented")
}
func (UnimplementedSheduleServiceServer) mustEmbedUnimplementedSheduleServiceServer() {}

// UnsafeSheduleServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SheduleServiceServer will
// result in compilation errors.
type UnsafeSheduleServiceServer interface {
	mustEmbedUnimplementedSheduleServiceServer()
}

func RegisterSheduleServiceServer(s grpc.ServiceRegistrar, srv SheduleServiceServer) {
	s.RegisterService(&SheduleService_ServiceDesc, srv)
}

func _SheduleService_GetSubjectFromDbRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Wrap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheduleServiceServer).GetSubjectFromDbRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SheduleService_GetSubjectFromDbRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheduleServiceServer).GetSubjectFromDbRPC(ctx, req.(*Wrap))
	}
	return interceptor(ctx, in, info, handler)
}

func _SheduleService_GetTutorsFromDbRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Wrap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheduleServiceServer).GetTutorsFromDbRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SheduleService_GetTutorsFromDbRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheduleServiceServer).GetTutorsFromDbRPC(ctx, req.(*Wrap))
	}
	return interceptor(ctx, in, info, handler)
}

func _SheduleService_GetAuditoriumFromDbRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Wrap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheduleServiceServer).GetAuditoriumFromDbRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SheduleService_GetAuditoriumFromDbRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheduleServiceServer).GetAuditoriumFromDbRPC(ctx, req.(*Wrap))
	}
	return interceptor(ctx, in, info, handler)
}

func _SheduleService_GetTypeFromDbRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Wrap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheduleServiceServer).GetTypeFromDbRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SheduleService_GetTypeFromDbRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheduleServiceServer).GetTypeFromDbRPC(ctx, req.(*Wrap))
	}
	return interceptor(ctx, in, info, handler)
}

func _SheduleService_GetGroupFromDbRPC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Wrap)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheduleServiceServer).GetGroupFromDbRPC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SheduleService_GetGroupFromDbRPC_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheduleServiceServer).GetGroupFromDbRPC(ctx, req.(*Wrap))
	}
	return interceptor(ctx, in, info, handler)
}

func _SheduleService_GetSheduleFromDb_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheduleServiceServer).GetSheduleFromDb(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SheduleService_GetSheduleFromDb_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheduleServiceServer).GetSheduleFromDb(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

func _SheduleService_AddShedule_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SheduleArray)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SheduleServiceServer).AddShedule(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SheduleService_AddShedule_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SheduleServiceServer).AddShedule(ctx, req.(*SheduleArray))
	}
	return interceptor(ctx, in, info, handler)
}

// SheduleService_ServiceDesc is the grpc.ServiceDesc for SheduleService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SheduleService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ProtoApi.SheduleService",
	HandlerType: (*SheduleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSubjectFromDbRPC",
			Handler:    _SheduleService_GetSubjectFromDbRPC_Handler,
		},
		{
			MethodName: "GetTutorsFromDbRPC",
			Handler:    _SheduleService_GetTutorsFromDbRPC_Handler,
		},
		{
			MethodName: "GetAuditoriumFromDbRPC",
			Handler:    _SheduleService_GetAuditoriumFromDbRPC_Handler,
		},
		{
			MethodName: "GetTypeFromDbRPC",
			Handler:    _SheduleService_GetTypeFromDbRPC_Handler,
		},
		{
			MethodName: "GetGroupFromDbRPC",
			Handler:    _SheduleService_GetGroupFromDbRPC_Handler,
		},
		{
			MethodName: "GetSheduleFromDb",
			Handler:    _SheduleService_GetSheduleFromDb_Handler,
		},
		{
			MethodName: "AddShedule",
			Handler:    _SheduleService_AddShedule_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ApiService.proto",
}
