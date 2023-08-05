// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: omdb.proto

package pb

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

// OMDBServiceClient is the client API for OMDBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OMDBServiceClient interface {
	GetMovieByID(ctx context.Context, in *GetMovieByIDRequest, opts ...grpc.CallOption) (*GetMovieByIDResponse, error)
	SearchMovies(ctx context.Context, in *SearchMoviesRequest, opts ...grpc.CallOption) (*SearchMoviesResponse, error)
}

type oMDBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOMDBServiceClient(cc grpc.ClientConnInterface) OMDBServiceClient {
	return &oMDBServiceClient{cc}
}

func (c *oMDBServiceClient) GetMovieByID(ctx context.Context, in *GetMovieByIDRequest, opts ...grpc.CallOption) (*GetMovieByIDResponse, error) {
	out := new(GetMovieByIDResponse)
	err := c.cc.Invoke(ctx, "/omdb.OMDBService/GetMovieByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *oMDBServiceClient) SearchMovies(ctx context.Context, in *SearchMoviesRequest, opts ...grpc.CallOption) (*SearchMoviesResponse, error) {
	out := new(SearchMoviesResponse)
	err := c.cc.Invoke(ctx, "/omdb.OMDBService/SearchMovies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OMDBServiceServer is the server API for OMDBService service.
// All implementations should embed UnimplementedOMDBServiceServer
// for forward compatibility
type OMDBServiceServer interface {
	GetMovieByID(context.Context, *GetMovieByIDRequest) (*GetMovieByIDResponse, error)
	SearchMovies(context.Context, *SearchMoviesRequest) (*SearchMoviesResponse, error)
}

// UnimplementedOMDBServiceServer should be embedded to have forward compatible implementations.
type UnimplementedOMDBServiceServer struct {
}

func (UnimplementedOMDBServiceServer) GetMovieByID(context.Context, *GetMovieByIDRequest) (*GetMovieByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMovieByID not implemented")
}
func (UnimplementedOMDBServiceServer) SearchMovies(context.Context, *SearchMoviesRequest) (*SearchMoviesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchMovies not implemented")
}

// UnsafeOMDBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OMDBServiceServer will
// result in compilation errors.
type UnsafeOMDBServiceServer interface {
	mustEmbedUnimplementedOMDBServiceServer()
}

func RegisterOMDBServiceServer(s grpc.ServiceRegistrar, srv OMDBServiceServer) {
	s.RegisterService(&OMDBService_ServiceDesc, srv)
}

func _OMDBService_GetMovieByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMovieByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMDBServiceServer).GetMovieByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omdb.OMDBService/GetMovieByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMDBServiceServer).GetMovieByID(ctx, req.(*GetMovieByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OMDBService_SearchMovies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchMoviesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OMDBServiceServer).SearchMovies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/omdb.OMDBService/SearchMovies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OMDBServiceServer).SearchMovies(ctx, req.(*SearchMoviesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OMDBService_ServiceDesc is the grpc.ServiceDesc for OMDBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OMDBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "omdb.OMDBService",
	HandlerType: (*OMDBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMovieByID",
			Handler:    _OMDBService_GetMovieByID_Handler,
		},
		{
			MethodName: "SearchMovies",
			Handler:    _OMDBService_SearchMovies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "omdb.proto",
}