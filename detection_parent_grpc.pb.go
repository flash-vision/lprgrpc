// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: detection_parent.proto

package lpr_grpc

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

// EncounterServiceClient is the client API for EncounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncounterServiceClient interface {
	// A server-to-client streaming RPC.
	SubscribeToDetections(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (EncounterService_SubscribeToDetectionsClient, error)
}

type encounterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncounterServiceClient(cc grpc.ClientConnInterface) EncounterServiceClient {
	return &encounterServiceClient{cc}
}

func (c *encounterServiceClient) SubscribeToDetections(ctx context.Context, in *SubscriptionRequest, opts ...grpc.CallOption) (EncounterService_SubscribeToDetectionsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EncounterService_ServiceDesc.Streams[0], "/lprgrpc.EncounterService/SubscribeToDetections", opts...)
	if err != nil {
		return nil, err
	}
	x := &encounterServiceSubscribeToDetectionsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EncounterService_SubscribeToDetectionsClient interface {
	Recv() (*DetectionEncounter, error)
	grpc.ClientStream
}

type encounterServiceSubscribeToDetectionsClient struct {
	grpc.ClientStream
}

func (x *encounterServiceSubscribeToDetectionsClient) Recv() (*DetectionEncounter, error) {
	m := new(DetectionEncounter)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EncounterServiceServer is the server API for EncounterService service.
// All implementations must embed UnimplementedEncounterServiceServer
// for forward compatibility
type EncounterServiceServer interface {
	// A server-to-client streaming RPC.
	SubscribeToDetections(*SubscriptionRequest, EncounterService_SubscribeToDetectionsServer) error
	mustEmbedUnimplementedEncounterServiceServer()
}

// UnimplementedEncounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEncounterServiceServer struct {
}

func (UnimplementedEncounterServiceServer) SubscribeToDetections(*SubscriptionRequest, EncounterService_SubscribeToDetectionsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeToDetections not implemented")
}
func (UnimplementedEncounterServiceServer) mustEmbedUnimplementedEncounterServiceServer() {}

// UnsafeEncounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncounterServiceServer will
// result in compilation errors.
type UnsafeEncounterServiceServer interface {
	mustEmbedUnimplementedEncounterServiceServer()
}

func RegisterEncounterServiceServer(s grpc.ServiceRegistrar, srv EncounterServiceServer) {
	s.RegisterService(&EncounterService_ServiceDesc, srv)
}

func _EncounterService_SubscribeToDetections_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscriptionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EncounterServiceServer).SubscribeToDetections(m, &encounterServiceSubscribeToDetectionsServer{stream})
}

type EncounterService_SubscribeToDetectionsServer interface {
	Send(*DetectionEncounter) error
	grpc.ServerStream
}

type encounterServiceSubscribeToDetectionsServer struct {
	grpc.ServerStream
}

func (x *encounterServiceSubscribeToDetectionsServer) Send(m *DetectionEncounter) error {
	return x.ServerStream.SendMsg(m)
}

// EncounterService_ServiceDesc is the grpc.ServiceDesc for EncounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "lprgrpc.EncounterService",
	HandlerType: (*EncounterServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeToDetections",
			Handler:       _EncounterService_SubscribeToDetections_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "detection_parent.proto",
}
