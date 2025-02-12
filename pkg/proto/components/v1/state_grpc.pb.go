// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: dapr/proto/components/v1/state.proto

package components

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

// QueriableStateStoreClient is the client API for QueriableStateStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueriableStateStoreClient interface {
	// Query performs a query request on the statestore.
	Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error)
}

type queriableStateStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewQueriableStateStoreClient(cc grpc.ClientConnInterface) QueriableStateStoreClient {
	return &queriableStateStoreClient{cc}
}

func (c *queriableStateStoreClient) Query(ctx context.Context, in *QueryRequest, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.QueriableStateStore/Query", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueriableStateStoreServer is the server API for QueriableStateStore service.
// All implementations should embed UnimplementedQueriableStateStoreServer
// for forward compatibility
type QueriableStateStoreServer interface {
	// Query performs a query request on the statestore.
	Query(context.Context, *QueryRequest) (*QueryResponse, error)
}

// UnimplementedQueriableStateStoreServer should be embedded to have forward compatible implementations.
type UnimplementedQueriableStateStoreServer struct {
}

func (UnimplementedQueriableStateStoreServer) Query(context.Context, *QueryRequest) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Query not implemented")
}

// UnsafeQueriableStateStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueriableStateStoreServer will
// result in compilation errors.
type UnsafeQueriableStateStoreServer interface {
	mustEmbedUnimplementedQueriableStateStoreServer()
}

func RegisterQueriableStateStoreServer(s grpc.ServiceRegistrar, srv QueriableStateStoreServer) {
	s.RegisterService(&QueriableStateStore_ServiceDesc, srv)
}

func _QueriableStateStore_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueriableStateStoreServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.QueriableStateStore/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueriableStateStoreServer).Query(ctx, req.(*QueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueriableStateStore_ServiceDesc is the grpc.ServiceDesc for QueriableStateStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueriableStateStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.components.v1.QueriableStateStore",
	HandlerType: (*QueriableStateStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _QueriableStateStore_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapr/proto/components/v1/state.proto",
}

// TransactionalStateStoreClient is the client API for TransactionalStateStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TransactionalStateStoreClient interface {
	// Transact executes multiples operation in a transactional environment.
	Transact(ctx context.Context, in *TransactionalStateRequest, opts ...grpc.CallOption) (*TransactionalStateResponse, error)
}

type transactionalStateStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionalStateStoreClient(cc grpc.ClientConnInterface) TransactionalStateStoreClient {
	return &transactionalStateStoreClient{cc}
}

func (c *transactionalStateStoreClient) Transact(ctx context.Context, in *TransactionalStateRequest, opts ...grpc.CallOption) (*TransactionalStateResponse, error) {
	out := new(TransactionalStateResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.TransactionalStateStore/Transact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionalStateStoreServer is the server API for TransactionalStateStore service.
// All implementations should embed UnimplementedTransactionalStateStoreServer
// for forward compatibility
type TransactionalStateStoreServer interface {
	// Transact executes multiples operation in a transactional environment.
	Transact(context.Context, *TransactionalStateRequest) (*TransactionalStateResponse, error)
}

// UnimplementedTransactionalStateStoreServer should be embedded to have forward compatible implementations.
type UnimplementedTransactionalStateStoreServer struct {
}

func (UnimplementedTransactionalStateStoreServer) Transact(context.Context, *TransactionalStateRequest) (*TransactionalStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transact not implemented")
}

// UnsafeTransactionalStateStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TransactionalStateStoreServer will
// result in compilation errors.
type UnsafeTransactionalStateStoreServer interface {
	mustEmbedUnimplementedTransactionalStateStoreServer()
}

func RegisterTransactionalStateStoreServer(s grpc.ServiceRegistrar, srv TransactionalStateStoreServer) {
	s.RegisterService(&TransactionalStateStore_ServiceDesc, srv)
}

func _TransactionalStateStore_Transact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransactionalStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionalStateStoreServer).Transact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.TransactionalStateStore/Transact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionalStateStoreServer).Transact(ctx, req.(*TransactionalStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TransactionalStateStore_ServiceDesc is the grpc.ServiceDesc for TransactionalStateStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TransactionalStateStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.components.v1.TransactionalStateStore",
	HandlerType: (*TransactionalStateStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transact",
			Handler:    _TransactionalStateStore_Transact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapr/proto/components/v1/state.proto",
}

// StateStoreClient is the client API for StateStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StateStoreClient interface {
	// Initializes the state store component with the given metadata.
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	// Returns a list of implemented state store features.
	Features(ctx context.Context, in *FeaturesRequest, opts ...grpc.CallOption) (*FeaturesResponse, error)
	// Deletes the specified key from the state store.
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// Get data from the given key.
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	// Sets the value of the specified key.
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	// Ping the state store. Used for liveness porpuses.
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// Deletes many keys at once.
	BulkDelete(ctx context.Context, in *BulkDeleteRequest, opts ...grpc.CallOption) (*BulkDeleteResponse, error)
	// Retrieves many keys at once.
	BulkGet(ctx context.Context, in *BulkGetRequest, opts ...grpc.CallOption) (*BulkGetResponse, error)
	// Set the value of many keys at once.
	BulkSet(ctx context.Context, in *BulkSetRequest, opts ...grpc.CallOption) (*BulkSetResponse, error)
}

type stateStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewStateStoreClient(cc grpc.ClientConnInterface) StateStoreClient {
	return &stateStoreClient{cc}
}

func (c *stateStoreClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) Features(ctx context.Context, in *FeaturesRequest, opts ...grpc.CallOption) (*FeaturesResponse, error) {
	out := new(FeaturesResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/Features", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) BulkDelete(ctx context.Context, in *BulkDeleteRequest, opts ...grpc.CallOption) (*BulkDeleteResponse, error) {
	out := new(BulkDeleteResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/BulkDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) BulkGet(ctx context.Context, in *BulkGetRequest, opts ...grpc.CallOption) (*BulkGetResponse, error) {
	out := new(BulkGetResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/BulkGet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stateStoreClient) BulkSet(ctx context.Context, in *BulkSetRequest, opts ...grpc.CallOption) (*BulkSetResponse, error) {
	out := new(BulkSetResponse)
	err := c.cc.Invoke(ctx, "/dapr.proto.components.v1.StateStore/BulkSet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StateStoreServer is the server API for StateStore service.
// All implementations should embed UnimplementedStateStoreServer
// for forward compatibility
type StateStoreServer interface {
	// Initializes the state store component with the given metadata.
	Init(context.Context, *InitRequest) (*InitResponse, error)
	// Returns a list of implemented state store features.
	Features(context.Context, *FeaturesRequest) (*FeaturesResponse, error)
	// Deletes the specified key from the state store.
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// Get data from the given key.
	Get(context.Context, *GetRequest) (*GetResponse, error)
	// Sets the value of the specified key.
	Set(context.Context, *SetRequest) (*SetResponse, error)
	// Ping the state store. Used for liveness porpuses.
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// Deletes many keys at once.
	BulkDelete(context.Context, *BulkDeleteRequest) (*BulkDeleteResponse, error)
	// Retrieves many keys at once.
	BulkGet(context.Context, *BulkGetRequest) (*BulkGetResponse, error)
	// Set the value of many keys at once.
	BulkSet(context.Context, *BulkSetRequest) (*BulkSetResponse, error)
}

// UnimplementedStateStoreServer should be embedded to have forward compatible implementations.
type UnimplementedStateStoreServer struct {
}

func (UnimplementedStateStoreServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedStateStoreServer) Features(context.Context, *FeaturesRequest) (*FeaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Features not implemented")
}
func (UnimplementedStateStoreServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedStateStoreServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedStateStoreServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedStateStoreServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedStateStoreServer) BulkDelete(context.Context, *BulkDeleteRequest) (*BulkDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkDelete not implemented")
}
func (UnimplementedStateStoreServer) BulkGet(context.Context, *BulkGetRequest) (*BulkGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkGet not implemented")
}
func (UnimplementedStateStoreServer) BulkSet(context.Context, *BulkSetRequest) (*BulkSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BulkSet not implemented")
}

// UnsafeStateStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StateStoreServer will
// result in compilation errors.
type UnsafeStateStoreServer interface {
	mustEmbedUnimplementedStateStoreServer()
}

func RegisterStateStoreServer(s grpc.ServiceRegistrar, srv StateStoreServer) {
	s.RegisterService(&StateStore_ServiceDesc, srv)
}

func _StateStore_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_Features_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).Features(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/Features",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).Features(ctx, req.(*FeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_BulkDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).BulkDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/BulkDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).BulkDelete(ctx, req.(*BulkDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_BulkGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkGetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).BulkGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/BulkGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).BulkGet(ctx, req.(*BulkGetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StateStore_BulkSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BulkSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StateStoreServer).BulkSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dapr.proto.components.v1.StateStore/BulkSet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StateStoreServer).BulkSet(ctx, req.(*BulkSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StateStore_ServiceDesc is the grpc.ServiceDesc for StateStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StateStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dapr.proto.components.v1.StateStore",
	HandlerType: (*StateStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Init",
			Handler:    _StateStore_Init_Handler,
		},
		{
			MethodName: "Features",
			Handler:    _StateStore_Features_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _StateStore_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _StateStore_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _StateStore_Set_Handler,
		},
		{
			MethodName: "Ping",
			Handler:    _StateStore_Ping_Handler,
		},
		{
			MethodName: "BulkDelete",
			Handler:    _StateStore_BulkDelete_Handler,
		},
		{
			MethodName: "BulkGet",
			Handler:    _StateStore_BulkGet_Handler,
		},
		{
			MethodName: "BulkSet",
			Handler:    _StateStore_BulkSet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dapr/proto/components/v1/state.proto",
}
