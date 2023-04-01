package mw

import (
	"context"
	"github.com/kostyay/protoc-gen-go-access-modifiers/example"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"net"
	"testing"
	"time"
)

func Test_MethodsMW(t *testing.T) {

	mw := NewPublicMethods()
	testService := example.NewExampleService()
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(mw.UnaryServerInterceptor()))

	example.RegisterExampleServiceServer(grpcServer, testService)

	netListener, err := net.Listen("tcp", "127.0.0.1:0")
	require.Nil(t, err)

	go grpcServer.Serve(netListener)

	time.Sleep(2 * time.Second)

	clientConn, err := grpc.Dial(
		netListener.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	require.NoError(t, err)
	defer clientConn.Close()

	testClient := example.NewExampleServiceClient(clientConn)

	// call public method
	_, err = testClient.PublicHello(context.Background(), &example.PublicHelloRequest{})
	require.NoError(t, err)

	// call private method
	_, err = testClient.PrivateHello(context.Background(), &example.PrivateHelloRequest{})
	require.Error(t, err)
	require.Equal(t, status.Code(err), codes.Unknown)
}
