package mw

import (
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	pb "github.com/kostyay/protoc-gen-go-access-modifiers/pkg/access/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"strings"
)

type publicMethods struct {
}

// NewPublicMethods returns a new instance of public methods interceptor
// It checks if the method is private and returns an error if it is thus protecting external calls to private methods
func NewPublicMethods() *publicMethods {
	return &publicMethods{}
}

func (i *publicMethods) interceptor(ctx context.Context, server interface{}, fullMethod string) (context.Context, error) {
	if i.isPrivateMethod(fullMethod) {
		return nil, status.Error(codes.Unknown, "unknown method")
	}

	return ctx, nil
}

func (i *publicMethods) isPrivateMethod(fullMethod string) bool {
	methodOpt, err := getMethodOptionsByName(formatFullMethod(fullMethod))
	if err != nil {
		return false
	}

	return methodOpt.Private
}

func formatFullMethod(fullMethod string) string {
	fullMethod = strings.TrimLeft(fullMethod, "/")
	fullMethod = strings.Replace(fullMethod, "/", ".", -1)
	return fullMethod
}

func getMethodOptionsByName(fullMethod string) (*pb.MethodOption, error) {
	desc, err := protoregistry.GlobalFiles.FindDescriptorByName(protoreflect.FullName(fullMethod))
	if err != nil {
		return nil, fmt.Errorf("FindDescriptorByName: %w", err)
	}

	opts, ok := desc.Options().(*descriptorpb.MethodOptions)
	if !ok {
		return nil, fmt.Errorf("cast Options() to MethodOptions")
	}

	if opts == nil {
		return nil, fmt.Errorf("method options are nil")
	}

	ext := proto.GetExtension(opts, pb.E_Mo)

	methodOpt, ok := ext.(*pb.MethodOption)
	if !ok {
		return nil, fmt.Errorf("cast Extension to MethodOption")
	}

	return methodOpt, nil
}

func (i *publicMethods) UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		newCtx, err := i.interceptor(ctx, info.Server, info.FullMethod)
		if err != nil {
			return nil, err
		}

		return handler(newCtx, req)
	}
}

func (i *publicMethods) StreamServerInterceptor() grpc.StreamServerInterceptor {
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, err := i.interceptor(stream.Context(), srv, info.FullMethod)
		if err != nil {
			return err
		}

		wrapped := grpc_middleware.WrapServerStream(stream)
		wrapped.WrappedContext = ctx
		return handler(srv, wrapped)
	}
}
