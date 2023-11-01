// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: injective_auction_rpc.proto

/*
Package injective_auction_rpcpb is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package injective_auction_rpcpb

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_InjectiveAuctionRPC_AuctionEndpoint_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveAuctionRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AuctionEndpointRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.AuctionEndpoint(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InjectiveAuctionRPC_AuctionEndpoint_0(ctx context.Context, marshaler runtime.Marshaler, server InjectiveAuctionRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AuctionEndpointRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.AuctionEndpoint(ctx, &protoReq)
	return msg, metadata, err

}

func request_InjectiveAuctionRPC_Auctions_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveAuctionRPCClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AuctionsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Auctions(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_InjectiveAuctionRPC_Auctions_0(ctx context.Context, marshaler runtime.Marshaler, server InjectiveAuctionRPCServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq AuctionsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.Auctions(ctx, &protoReq)
	return msg, metadata, err

}

func request_InjectiveAuctionRPC_StreamBids_0(ctx context.Context, marshaler runtime.Marshaler, client InjectiveAuctionRPCClient, req *http.Request, pathParams map[string]string) (InjectiveAuctionRPC_StreamBidsClient, runtime.ServerMetadata, error) {
	var protoReq StreamBidsRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	stream, err := client.StreamBids(ctx, &protoReq)
	if err != nil {
		return nil, metadata, err
	}
	header, err := stream.Header()
	if err != nil {
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil

}

// RegisterInjectiveAuctionRPCHandlerServer registers the http handlers for service InjectiveAuctionRPC to "mux".
// UnaryRPC     :call InjectiveAuctionRPCServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterInjectiveAuctionRPCHandlerFromEndpoint instead.
func RegisterInjectiveAuctionRPCHandlerServer(ctx context.Context, mux *runtime.ServeMux, server InjectiveAuctionRPCServer) error {

	mux.Handle("POST", pattern_InjectiveAuctionRPC_AuctionEndpoint_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/injective_auction_rpc.InjectiveAuctionRPC/AuctionEndpoint", runtime.WithHTTPPathPattern("/injective_auction_rpc.InjectiveAuctionRPC/AuctionEndpoint"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InjectiveAuctionRPC_AuctionEndpoint_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveAuctionRPC_AuctionEndpoint_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveAuctionRPC_Auctions_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/injective_auction_rpc.InjectiveAuctionRPC/Auctions", runtime.WithHTTPPathPattern("/injective_auction_rpc.InjectiveAuctionRPC/Auctions"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_InjectiveAuctionRPC_Auctions_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveAuctionRPC_Auctions_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveAuctionRPC_StreamBids_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterInjectiveAuctionRPCHandlerFromEndpoint is same as RegisterInjectiveAuctionRPCHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterInjectiveAuctionRPCHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterInjectiveAuctionRPCHandler(ctx, mux, conn)
}

// RegisterInjectiveAuctionRPCHandler registers the http handlers for service InjectiveAuctionRPC to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterInjectiveAuctionRPCHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterInjectiveAuctionRPCHandlerClient(ctx, mux, NewInjectiveAuctionRPCClient(conn))
}

// RegisterInjectiveAuctionRPCHandlerClient registers the http handlers for service InjectiveAuctionRPC
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "InjectiveAuctionRPCClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "InjectiveAuctionRPCClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "InjectiveAuctionRPCClient" to call the correct interceptors.
func RegisterInjectiveAuctionRPCHandlerClient(ctx context.Context, mux *runtime.ServeMux, client InjectiveAuctionRPCClient) error {

	mux.Handle("POST", pattern_InjectiveAuctionRPC_AuctionEndpoint_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_auction_rpc.InjectiveAuctionRPC/AuctionEndpoint", runtime.WithHTTPPathPattern("/injective_auction_rpc.InjectiveAuctionRPC/AuctionEndpoint"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveAuctionRPC_AuctionEndpoint_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveAuctionRPC_AuctionEndpoint_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveAuctionRPC_Auctions_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_auction_rpc.InjectiveAuctionRPC/Auctions", runtime.WithHTTPPathPattern("/injective_auction_rpc.InjectiveAuctionRPC/Auctions"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveAuctionRPC_Auctions_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveAuctionRPC_Auctions_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("POST", pattern_InjectiveAuctionRPC_StreamBids_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/injective_auction_rpc.InjectiveAuctionRPC/StreamBids", runtime.WithHTTPPathPattern("/injective_auction_rpc.InjectiveAuctionRPC/StreamBids"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_InjectiveAuctionRPC_StreamBids_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_InjectiveAuctionRPC_StreamBids_0(annotatedContext, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_InjectiveAuctionRPC_AuctionEndpoint_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_auction_rpc.InjectiveAuctionRPC", "AuctionEndpoint"}, ""))

	pattern_InjectiveAuctionRPC_Auctions_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_auction_rpc.InjectiveAuctionRPC", "Auctions"}, ""))

	pattern_InjectiveAuctionRPC_StreamBids_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"injective_auction_rpc.InjectiveAuctionRPC", "StreamBids"}, ""))
)

var (
	forward_InjectiveAuctionRPC_AuctionEndpoint_0 = runtime.ForwardResponseMessage

	forward_InjectiveAuctionRPC_Auctions_0 = runtime.ForwardResponseMessage

	forward_InjectiveAuctionRPC_StreamBids_0 = runtime.ForwardResponseStream
)
