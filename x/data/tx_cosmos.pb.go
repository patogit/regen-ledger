// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package data

import (
	context "context"
	types "github.com/regen-network/regen-ledger/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	// AnchorData "anchors" a piece of data to the blockchain based on its secure
	// hash, effectively providing a tamper resistant timestamp.
	//
	// The sender in AnchorData is not attesting to the veracity of the underlying
	// data. They can simply be a intermediary providing timestamp services.
	// SignData should be used to create a digital signature attesting to the
	// veracity of some piece of data.
	AnchorData(ctx context.Context, in *MsgAnchorDataRequest, opts ...grpc.CallOption) (*MsgAnchorDataResponse, error)
	// SignData allows for signing of an arbitrary piece of data on the
	// blockchain. By "signing" data the signers are making a statement about the
	// veracity of the data itself. It is like signing a legal document, meaning
	// that I agree to all conditions and to the best of my knowledge everything
	// is true. When anchoring data, the sender is not attesting to the veracity
	// of the data, they are simply communicating that it exists.
	//
	// On-chain signatures have the following benefits:
	// - on-chain identities can be managed using different cryptographic keys
	//   that change over time through key rotation practices
	// - an on-chain identity may represent an organization and through delegation
	//   individual members may sign on behalf of the group
	// - the blockchain transaction envelope provides built-in replay protection
	//   and timestamping
	//
	// SignData implicitly calls AnchorData if the data was not already anchored.
	//
	// SignData can be called multiple times for the same CID with different
	// signers and those signers will be appended to the list of signers.
	SignData(ctx context.Context, in *MsgSignDataRequest, opts ...grpc.CallOption) (*MsgSignDataResponse, error)
	// StoreData stores a piece of data corresponding to a CID on the blockchain.
	//
	// Currently only data for CID's using sha2-256 and blake2b-256 hash
	// algorithms can be stored on-chain.
	//
	// StoreData implicitly calls AnchorData if the data was not already anchored.
	//
	// The sender in StoreData is not attesting to the veracity of the underlying
	// data. They can simply be a intermediary providing storage services.
	// SignData should be used to create a digital signature attesting to the
	// veracity of some piece of data.
	StoreData(ctx context.Context, in *MsgStoreDataRequest, opts ...grpc.CallOption) (*MsgStoreDataResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AnchorData(ctx context.Context, in *MsgAnchorDataRequest, opts ...grpc.CallOption) (*MsgAnchorDataResponse, error) {
	out := new(MsgAnchorDataResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Msg/AnchorData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SignData(ctx context.Context, in *MsgSignDataRequest, opts ...grpc.CallOption) (*MsgSignDataResponse, error) {
	out := new(MsgSignDataResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Msg/SignData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) StoreData(ctx context.Context, in *MsgStoreDataRequest, opts ...grpc.CallOption) (*MsgStoreDataResponse, error) {
	out := new(MsgStoreDataResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha1.Msg/StoreData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// AnchorData "anchors" a piece of data to the blockchain based on its secure
	// hash, effectively providing a tamper resistant timestamp.
	//
	// The sender in AnchorData is not attesting to the veracity of the underlying
	// data. They can simply be a intermediary providing timestamp services.
	// SignData should be used to create a digital signature attesting to the
	// veracity of some piece of data.
	AnchorData(types.Context, *MsgAnchorDataRequest) (*MsgAnchorDataResponse, error)
	// SignData allows for signing of an arbitrary piece of data on the
	// blockchain. By "signing" data the signers are making a statement about the
	// veracity of the data itself. It is like signing a legal document, meaning
	// that I agree to all conditions and to the best of my knowledge everything
	// is true. When anchoring data, the sender is not attesting to the veracity
	// of the data, they are simply communicating that it exists.
	//
	// On-chain signatures have the following benefits:
	// - on-chain identities can be managed using different cryptographic keys
	//   that change over time through key rotation practices
	// - an on-chain identity may represent an organization and through delegation
	//   individual members may sign on behalf of the group
	// - the blockchain transaction envelope provides built-in replay protection
	//   and timestamping
	//
	// SignData implicitly calls AnchorData if the data was not already anchored.
	//
	// SignData can be called multiple times for the same CID with different
	// signers and those signers will be appended to the list of signers.
	SignData(types.Context, *MsgSignDataRequest) (*MsgSignDataResponse, error)
	// StoreData stores a piece of data corresponding to a CID on the blockchain.
	//
	// Currently only data for CID's using sha2-256 and blake2b-256 hash
	// algorithms can be stored on-chain.
	//
	// StoreData implicitly calls AnchorData if the data was not already anchored.
	//
	// The sender in StoreData is not attesting to the veracity of the underlying
	// data. They can simply be a intermediary providing storage services.
	// SignData should be used to create a digital signature attesting to the
	// veracity of some piece of data.
	StoreData(types.Context, *MsgStoreDataRequest) (*MsgStoreDataResponse, error)
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_AnchorData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAnchorDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AnchorData(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Msg/AnchorData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AnchorData(types.UnwrapSDKContext(ctx), req.(*MsgAnchorDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SignData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSignDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SignData(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Msg/SignData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SignData(types.UnwrapSDKContext(ctx), req.(*MsgSignDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_StoreData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgStoreDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).StoreData(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha1.Msg/StoreData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).StoreData(types.UnwrapSDKContext(ctx), req.(*MsgStoreDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.data.v1alpha1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnchorData",
			Handler:    _Msg_AnchorData_Handler,
		},
		{
			MethodName: "SignData",
			Handler:    _Msg_SignData_Handler,
		},
		{
			MethodName: "StoreData",
			Handler:    _Msg_StoreData_Handler,
		},
	},
	Metadata: "regen/data/v1alpha1/tx.proto",
}
