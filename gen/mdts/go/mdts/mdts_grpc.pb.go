// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: mdts/mdts.proto

package mdtsv1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_GetById_FullMethodName       = "/mdts.UserService/GetById"
	UserService_GetByUsername_FullMethodName = "/mdts.UserService/GetByUsername"
	UserService_GetAll_FullMethodName        = "/mdts.UserService/GetAll"
	UserService_Update_FullMethodName        = "/mdts.UserService/Update"
	UserService_Delete_FullMethodName        = "/mdts.UserService/Delete"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetAll(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetByUsername(ctx context.Context, in *GetUserByUsernameRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, UserService_GetByUsername_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetAll(ctx context.Context, in *GetAllUsersRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UsersResponse)
	err := c.cc.Invoke(ctx, UserService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*UpdateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateUserResponse)
	err := c.cc.Invoke(ctx, UserService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*DeleteUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteUserResponse)
	err := c.cc.Invoke(ctx, UserService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	GetById(context.Context, *GetUserByIdRequest) (*UserResponse, error)
	GetByUsername(context.Context, *GetUserByUsernameRequest) (*UserResponse, error)
	GetAll(context.Context, *GetAllUsersRequest) (*UsersResponse, error)
	Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error)
	Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetById(context.Context, *GetUserByIdRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedUserServiceServer) GetByUsername(context.Context, *GetUserByUsernameRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByUsername not implemented")
}
func (UnimplementedUserServiceServer) GetAll(context.Context, *GetAllUsersRequest) (*UsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedUserServiceServer) Update(context.Context, *UpdateUserRequest) (*UpdateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedUserServiceServer) Delete(context.Context, *DeleteUserRequest) (*DeleteUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetByUsername_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByUsernameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetByUsername(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetByUsername_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetByUsername(ctx, req.(*GetUserByUsernameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetAll(ctx, req.(*GetAllUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Update(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Delete(ctx, req.(*DeleteUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mdts.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetById",
			Handler:    _UserService_GetById_Handler,
		},
		{
			MethodName: "GetByUsername",
			Handler:    _UserService_GetByUsername_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _UserService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _UserService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _UserService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mdts/mdts.proto",
}

const (
	AlertService_Create_FullMethodName      = "/mdts.AlertService/Create"
	AlertService_GetById_FullMethodName     = "/mdts.AlertService/GetById"
	AlertService_GetByPeriod_FullMethodName = "/mdts.AlertService/GetByPeriod"
	AlertService_GetAll_FullMethodName      = "/mdts.AlertService/GetAll"
	AlertService_Update_FullMethodName      = "/mdts.AlertService/Update"
	AlertService_Delete_FullMethodName      = "/mdts.AlertService/Delete"
)

// AlertServiceClient is the client API for AlertService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AlertServiceClient interface {
	Create(ctx context.Context, in *CreateAlertRequest, opts ...grpc.CallOption) (*AlertResponse, error)
	GetById(ctx context.Context, in *GetAlertByIdRequest, opts ...grpc.CallOption) (*AlertResponse, error)
	GetByPeriod(ctx context.Context, in *GetAlertsByPeriodRequest, opts ...grpc.CallOption) (*AlertsResponse, error)
	GetAll(ctx context.Context, in *GetAllAlertsRequest, opts ...grpc.CallOption) (*AlertsResponse, error)
	Update(ctx context.Context, in *UpdateAlertRequest, opts ...grpc.CallOption) (*UpdateAlertResponse, error)
	Delete(ctx context.Context, in *DeleteAlertRequest, opts ...grpc.CallOption) (*DeleteAlertResponse, error)
}

type alertServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAlertServiceClient(cc grpc.ClientConnInterface) AlertServiceClient {
	return &alertServiceClient{cc}
}

func (c *alertServiceClient) Create(ctx context.Context, in *CreateAlertRequest, opts ...grpc.CallOption) (*AlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertResponse)
	err := c.cc.Invoke(ctx, AlertService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) GetById(ctx context.Context, in *GetAlertByIdRequest, opts ...grpc.CallOption) (*AlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertResponse)
	err := c.cc.Invoke(ctx, AlertService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) GetByPeriod(ctx context.Context, in *GetAlertsByPeriodRequest, opts ...grpc.CallOption) (*AlertsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertsResponse)
	err := c.cc.Invoke(ctx, AlertService_GetByPeriod_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) GetAll(ctx context.Context, in *GetAllAlertsRequest, opts ...grpc.CallOption) (*AlertsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AlertsResponse)
	err := c.cc.Invoke(ctx, AlertService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) Update(ctx context.Context, in *UpdateAlertRequest, opts ...grpc.CallOption) (*UpdateAlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateAlertResponse)
	err := c.cc.Invoke(ctx, AlertService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *alertServiceClient) Delete(ctx context.Context, in *DeleteAlertRequest, opts ...grpc.CallOption) (*DeleteAlertResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteAlertResponse)
	err := c.cc.Invoke(ctx, AlertService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AlertServiceServer is the server API for AlertService service.
// All implementations must embed UnimplementedAlertServiceServer
// for forward compatibility.
type AlertServiceServer interface {
	Create(context.Context, *CreateAlertRequest) (*AlertResponse, error)
	GetById(context.Context, *GetAlertByIdRequest) (*AlertResponse, error)
	GetByPeriod(context.Context, *GetAlertsByPeriodRequest) (*AlertsResponse, error)
	GetAll(context.Context, *GetAllAlertsRequest) (*AlertsResponse, error)
	Update(context.Context, *UpdateAlertRequest) (*UpdateAlertResponse, error)
	Delete(context.Context, *DeleteAlertRequest) (*DeleteAlertResponse, error)
	mustEmbedUnimplementedAlertServiceServer()
}

// UnimplementedAlertServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAlertServiceServer struct{}

func (UnimplementedAlertServiceServer) Create(context.Context, *CreateAlertRequest) (*AlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedAlertServiceServer) GetById(context.Context, *GetAlertByIdRequest) (*AlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedAlertServiceServer) GetByPeriod(context.Context, *GetAlertsByPeriodRequest) (*AlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPeriod not implemented")
}
func (UnimplementedAlertServiceServer) GetAll(context.Context, *GetAllAlertsRequest) (*AlertsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAlertServiceServer) Update(context.Context, *UpdateAlertRequest) (*UpdateAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedAlertServiceServer) Delete(context.Context, *DeleteAlertRequest) (*DeleteAlertResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedAlertServiceServer) mustEmbedUnimplementedAlertServiceServer() {}
func (UnimplementedAlertServiceServer) testEmbeddedByValue()                      {}

// UnsafeAlertServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AlertServiceServer will
// result in compilation errors.
type UnsafeAlertServiceServer interface {
	mustEmbedUnimplementedAlertServiceServer()
}

func RegisterAlertServiceServer(s grpc.ServiceRegistrar, srv AlertServiceServer) {
	// If the following call pancis, it indicates UnimplementedAlertServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&AlertService_ServiceDesc, srv)
}

func _AlertService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).Create(ctx, req.(*CreateAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).GetById(ctx, req.(*GetAlertByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_GetByPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAlertsByPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).GetByPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_GetByPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).GetByPeriod(ctx, req.(*GetAlertsByPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllAlertsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).GetAll(ctx, req.(*GetAllAlertsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).Update(ctx, req.(*UpdateAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AlertService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAlertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlertServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AlertService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlertServiceServer).Delete(ctx, req.(*DeleteAlertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AlertService_ServiceDesc is the grpc.ServiceDesc for AlertService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AlertService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mdts.AlertService",
	HandlerType: (*AlertServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _AlertService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _AlertService_GetById_Handler,
		},
		{
			MethodName: "GetByPeriod",
			Handler:    _AlertService_GetByPeriod_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AlertService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _AlertService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _AlertService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mdts/mdts.proto",
}

const (
	BplaService_Create_FullMethodName              = "/mdts.BplaService/Create"
	BplaService_GetById_FullMethodName             = "/mdts.BplaService/GetById"
	BplaService_GetByEntryDateRange_FullMethodName = "/mdts.BplaService/GetByEntryDateRange"
	BplaService_GetByInUsage_FullMethodName        = "/mdts.BplaService/GetByInUsage"
	BplaService_GetByModelAndBrand_FullMethodName  = "/mdts.BplaService/GetByModelAndBrand"
	BplaService_GetAll_FullMethodName              = "/mdts.BplaService/GetAll"
	BplaService_Update_FullMethodName              = "/mdts.BplaService/Update"
	BplaService_Delete_FullMethodName              = "/mdts.BplaService/Delete"
)

// BplaServiceClient is the client API for BplaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BplaServiceClient interface {
	Create(ctx context.Context, in *CreateBplaRequest, opts ...grpc.CallOption) (*BplaResponse, error)
	GetById(ctx context.Context, in *GetBplaByIdRequest, opts ...grpc.CallOption) (*BplaResponse, error)
	GetByEntryDateRange(ctx context.Context, in *GetBplaByEntryDateRangeRequest, opts ...grpc.CallOption) (*BplasResponse, error)
	GetByInUsage(ctx context.Context, in *GetBplaByInUsageRequest, opts ...grpc.CallOption) (*BplasResponse, error)
	GetByModelAndBrand(ctx context.Context, in *GetBplaByModelAndBrandRequest, opts ...grpc.CallOption) (*BplasResponse, error)
	GetAll(ctx context.Context, in *GetAllBplasRequest, opts ...grpc.CallOption) (*BplasResponse, error)
	Update(ctx context.Context, in *UpdateBplaRequest, opts ...grpc.CallOption) (*UpdateBplaResponse, error)
	Delete(ctx context.Context, in *DeleteBplaRequest, opts ...grpc.CallOption) (*DeleteBplaResponse, error)
}

type bplaServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBplaServiceClient(cc grpc.ClientConnInterface) BplaServiceClient {
	return &bplaServiceClient{cc}
}

func (c *bplaServiceClient) Create(ctx context.Context, in *CreateBplaRequest, opts ...grpc.CallOption) (*BplaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BplaResponse)
	err := c.cc.Invoke(ctx, BplaService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bplaServiceClient) GetById(ctx context.Context, in *GetBplaByIdRequest, opts ...grpc.CallOption) (*BplaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BplaResponse)
	err := c.cc.Invoke(ctx, BplaService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bplaServiceClient) GetByEntryDateRange(ctx context.Context, in *GetBplaByEntryDateRangeRequest, opts ...grpc.CallOption) (*BplasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BplasResponse)
	err := c.cc.Invoke(ctx, BplaService_GetByEntryDateRange_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bplaServiceClient) GetByInUsage(ctx context.Context, in *GetBplaByInUsageRequest, opts ...grpc.CallOption) (*BplasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BplasResponse)
	err := c.cc.Invoke(ctx, BplaService_GetByInUsage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bplaServiceClient) GetByModelAndBrand(ctx context.Context, in *GetBplaByModelAndBrandRequest, opts ...grpc.CallOption) (*BplasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BplasResponse)
	err := c.cc.Invoke(ctx, BplaService_GetByModelAndBrand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bplaServiceClient) GetAll(ctx context.Context, in *GetAllBplasRequest, opts ...grpc.CallOption) (*BplasResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BplasResponse)
	err := c.cc.Invoke(ctx, BplaService_GetAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bplaServiceClient) Update(ctx context.Context, in *UpdateBplaRequest, opts ...grpc.CallOption) (*UpdateBplaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateBplaResponse)
	err := c.cc.Invoke(ctx, BplaService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bplaServiceClient) Delete(ctx context.Context, in *DeleteBplaRequest, opts ...grpc.CallOption) (*DeleteBplaResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteBplaResponse)
	err := c.cc.Invoke(ctx, BplaService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BplaServiceServer is the server API for BplaService service.
// All implementations must embed UnimplementedBplaServiceServer
// for forward compatibility.
type BplaServiceServer interface {
	Create(context.Context, *CreateBplaRequest) (*BplaResponse, error)
	GetById(context.Context, *GetBplaByIdRequest) (*BplaResponse, error)
	GetByEntryDateRange(context.Context, *GetBplaByEntryDateRangeRequest) (*BplasResponse, error)
	GetByInUsage(context.Context, *GetBplaByInUsageRequest) (*BplasResponse, error)
	GetByModelAndBrand(context.Context, *GetBplaByModelAndBrandRequest) (*BplasResponse, error)
	GetAll(context.Context, *GetAllBplasRequest) (*BplasResponse, error)
	Update(context.Context, *UpdateBplaRequest) (*UpdateBplaResponse, error)
	Delete(context.Context, *DeleteBplaRequest) (*DeleteBplaResponse, error)
	mustEmbedUnimplementedBplaServiceServer()
}

// UnimplementedBplaServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedBplaServiceServer struct{}

func (UnimplementedBplaServiceServer) Create(context.Context, *CreateBplaRequest) (*BplaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedBplaServiceServer) GetById(context.Context, *GetBplaByIdRequest) (*BplaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedBplaServiceServer) GetByEntryDateRange(context.Context, *GetBplaByEntryDateRangeRequest) (*BplasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByEntryDateRange not implemented")
}
func (UnimplementedBplaServiceServer) GetByInUsage(context.Context, *GetBplaByInUsageRequest) (*BplasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByInUsage not implemented")
}
func (UnimplementedBplaServiceServer) GetByModelAndBrand(context.Context, *GetBplaByModelAndBrandRequest) (*BplasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByModelAndBrand not implemented")
}
func (UnimplementedBplaServiceServer) GetAll(context.Context, *GetAllBplasRequest) (*BplasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedBplaServiceServer) Update(context.Context, *UpdateBplaRequest) (*UpdateBplaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedBplaServiceServer) Delete(context.Context, *DeleteBplaRequest) (*DeleteBplaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedBplaServiceServer) mustEmbedUnimplementedBplaServiceServer() {}
func (UnimplementedBplaServiceServer) testEmbeddedByValue()                     {}

// UnsafeBplaServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BplaServiceServer will
// result in compilation errors.
type UnsafeBplaServiceServer interface {
	mustEmbedUnimplementedBplaServiceServer()
}

func RegisterBplaServiceServer(s grpc.ServiceRegistrar, srv BplaServiceServer) {
	// If the following call pancis, it indicates UnimplementedBplaServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&BplaService_ServiceDesc, srv)
}

func _BplaService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBplaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BplaServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BplaService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BplaServiceServer).Create(ctx, req.(*CreateBplaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BplaService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBplaByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BplaServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BplaService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BplaServiceServer).GetById(ctx, req.(*GetBplaByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BplaService_GetByEntryDateRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBplaByEntryDateRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BplaServiceServer).GetByEntryDateRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BplaService_GetByEntryDateRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BplaServiceServer).GetByEntryDateRange(ctx, req.(*GetBplaByEntryDateRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BplaService_GetByInUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBplaByInUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BplaServiceServer).GetByInUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BplaService_GetByInUsage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BplaServiceServer).GetByInUsage(ctx, req.(*GetBplaByInUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BplaService_GetByModelAndBrand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBplaByModelAndBrandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BplaServiceServer).GetByModelAndBrand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BplaService_GetByModelAndBrand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BplaServiceServer).GetByModelAndBrand(ctx, req.(*GetBplaByModelAndBrandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BplaService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllBplasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BplaServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BplaService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BplaServiceServer).GetAll(ctx, req.(*GetAllBplasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BplaService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBplaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BplaServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BplaService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BplaServiceServer).Update(ctx, req.(*UpdateBplaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BplaService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBplaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BplaServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BplaService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BplaServiceServer).Delete(ctx, req.(*DeleteBplaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BplaService_ServiceDesc is the grpc.ServiceDesc for BplaService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BplaService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mdts.BplaService",
	HandlerType: (*BplaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _BplaService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _BplaService_GetById_Handler,
		},
		{
			MethodName: "GetByEntryDateRange",
			Handler:    _BplaService_GetByEntryDateRange_Handler,
		},
		{
			MethodName: "GetByInUsage",
			Handler:    _BplaService_GetByInUsage_Handler,
		},
		{
			MethodName: "GetByModelAndBrand",
			Handler:    _BplaService_GetByModelAndBrand_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _BplaService_GetAll_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _BplaService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _BplaService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "mdts/mdts.proto",
}
