package greeter

import (
	"fmt"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloResponse, error) {
	fmt.Println("SayHello Called.")
	return &HelloResponse{Message: "Hi " + in.Name + "\n"}, nil
}

// RegisterGreeterSrv register service into grpc
func RegisterGreeterSrv(gsvr *grpc.Server) {
	RegisterGreeterServer(gsvr, &server{})
}
