// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: splitty/v1/splitty.proto

package v1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "example/gen/splitty/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// SplittyServiceName is the fully-qualified name of the SplittyService service.
	SplittyServiceName = "splitty.v1.SplittyService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// SplittyServiceSignupProcedure is the fully-qualified name of the SplittyService's Signup RPC.
	SplittyServiceSignupProcedure = "/splitty.v1.SplittyService/Signup"
	// SplittyServiceLoginProcedure is the fully-qualified name of the SplittyService's Login RPC.
	SplittyServiceLoginProcedure = "/splitty.v1.SplittyService/Login"
	// SplittyServiceReceiptAnalyzeProcedure is the fully-qualified name of the SplittyService's
	// ReceiptAnalyze RPC.
	SplittyServiceReceiptAnalyzeProcedure = "/splitty.v1.SplittyService/ReceiptAnalyze"
)

// SplittyServiceClient is a client for the splitty.v1.SplittyService service.
type SplittyServiceClient interface {
	Signup(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	ReceiptAnalyze(context.Context, *connect.Request[v1.ReceiptAnalyzeRequest]) (*connect.Response[v1.ReceiptAnalyzeResponse], error)
}

// NewSplittyServiceClient constructs a client for the splitty.v1.SplittyService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewSplittyServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) SplittyServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	splittyServiceMethods := v1.File_splitty_v1_splitty_proto.Services().ByName("SplittyService").Methods()
	return &splittyServiceClient{
		signup: connect.NewClient[v1.SignupRequest, v1.SignupResponse](
			httpClient,
			baseURL+SplittyServiceSignupProcedure,
			connect.WithSchema(splittyServiceMethods.ByName("Signup")),
			connect.WithClientOptions(opts...),
		),
		login: connect.NewClient[v1.LoginRequest, v1.LoginResponse](
			httpClient,
			baseURL+SplittyServiceLoginProcedure,
			connect.WithSchema(splittyServiceMethods.ByName("Login")),
			connect.WithClientOptions(opts...),
		),
		receiptAnalyze: connect.NewClient[v1.ReceiptAnalyzeRequest, v1.ReceiptAnalyzeResponse](
			httpClient,
			baseURL+SplittyServiceReceiptAnalyzeProcedure,
			connect.WithSchema(splittyServiceMethods.ByName("ReceiptAnalyze")),
			connect.WithClientOptions(opts...),
		),
	}
}

// splittyServiceClient implements SplittyServiceClient.
type splittyServiceClient struct {
	signup         *connect.Client[v1.SignupRequest, v1.SignupResponse]
	login          *connect.Client[v1.LoginRequest, v1.LoginResponse]
	receiptAnalyze *connect.Client[v1.ReceiptAnalyzeRequest, v1.ReceiptAnalyzeResponse]
}

// Signup calls splitty.v1.SplittyService.Signup.
func (c *splittyServiceClient) Signup(ctx context.Context, req *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	return c.signup.CallUnary(ctx, req)
}

// Login calls splitty.v1.SplittyService.Login.
func (c *splittyServiceClient) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return c.login.CallUnary(ctx, req)
}

// ReceiptAnalyze calls splitty.v1.SplittyService.ReceiptAnalyze.
func (c *splittyServiceClient) ReceiptAnalyze(ctx context.Context, req *connect.Request[v1.ReceiptAnalyzeRequest]) (*connect.Response[v1.ReceiptAnalyzeResponse], error) {
	return c.receiptAnalyze.CallUnary(ctx, req)
}

// SplittyServiceHandler is an implementation of the splitty.v1.SplittyService service.
type SplittyServiceHandler interface {
	Signup(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error)
	Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error)
	ReceiptAnalyze(context.Context, *connect.Request[v1.ReceiptAnalyzeRequest]) (*connect.Response[v1.ReceiptAnalyzeResponse], error)
}

// NewSplittyServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewSplittyServiceHandler(svc SplittyServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	splittyServiceMethods := v1.File_splitty_v1_splitty_proto.Services().ByName("SplittyService").Methods()
	splittyServiceSignupHandler := connect.NewUnaryHandler(
		SplittyServiceSignupProcedure,
		svc.Signup,
		connect.WithSchema(splittyServiceMethods.ByName("Signup")),
		connect.WithHandlerOptions(opts...),
	)
	splittyServiceLoginHandler := connect.NewUnaryHandler(
		SplittyServiceLoginProcedure,
		svc.Login,
		connect.WithSchema(splittyServiceMethods.ByName("Login")),
		connect.WithHandlerOptions(opts...),
	)
	splittyServiceReceiptAnalyzeHandler := connect.NewUnaryHandler(
		SplittyServiceReceiptAnalyzeProcedure,
		svc.ReceiptAnalyze,
		connect.WithSchema(splittyServiceMethods.ByName("ReceiptAnalyze")),
		connect.WithHandlerOptions(opts...),
	)
	return "/splitty.v1.SplittyService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case SplittyServiceSignupProcedure:
			splittyServiceSignupHandler.ServeHTTP(w, r)
		case SplittyServiceLoginProcedure:
			splittyServiceLoginHandler.ServeHTTP(w, r)
		case SplittyServiceReceiptAnalyzeProcedure:
			splittyServiceReceiptAnalyzeHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedSplittyServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedSplittyServiceHandler struct{}

func (UnimplementedSplittyServiceHandler) Signup(context.Context, *connect.Request[v1.SignupRequest]) (*connect.Response[v1.SignupResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("splitty.v1.SplittyService.Signup is not implemented"))
}

func (UnimplementedSplittyServiceHandler) Login(context.Context, *connect.Request[v1.LoginRequest]) (*connect.Response[v1.LoginResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("splitty.v1.SplittyService.Login is not implemented"))
}

func (UnimplementedSplittyServiceHandler) ReceiptAnalyze(context.Context, *connect.Request[v1.ReceiptAnalyzeRequest]) (*connect.Response[v1.ReceiptAnalyzeResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("splitty.v1.SplittyService.ReceiptAnalyze is not implemented"))
}
