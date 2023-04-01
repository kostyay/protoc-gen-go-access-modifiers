package example

import (
	"context"
)

// ExampleService is a service that implements the ExampleServiceServer interface
type ExampleService struct {
	ExampleServiceServer
}

// NewExampleService returns a new ExampleService
func NewExampleService() *ExampleService {
	return &ExampleService{}
}

func (s *ExampleService) PublicHello(ctx context.Context, r *PublicHelloRequest) (*PublicHelloResponse, error) {
	return &PublicHelloResponse{
		Name: "public-hello-response",
	}, nil
}

func (s *ExampleService) PrivateHello(context.Context, *PrivateHelloRequest) (*PrivateHelloResponse, error) {
	return &PrivateHelloResponse{
		Name: "private-hello-response",
	}, nil
}
