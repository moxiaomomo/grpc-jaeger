package wrapper

import (
	"fmt"
	"net"
	"os"
	"time"

	pb "github.com/moxiaomomo/grpc-jaeger/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"testing"
)

func Test_Tracing(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:8001")
	if err != nil {
		os.Exit(-1)
	}

	var servOpts []grpc.ServerOption
	tracer, _, err := NewJaegerTracer("testSrv", "127.0.0.1:6831")
	if err != nil {
		fmt.Printf("new tracer err: %+v\n", err)
		os.Exit(-1)
	}
	if tracer != nil {
		servOpts = append(servOpts, ServerOption(tracer))
	}
	svr := grpc.NewServer(servOpts...)
	pb.RegisterGreeterSrv(svr)

	go func() {
		time.Sleep(time.Second)

		dialOpts := []grpc.DialOption{grpc.WithInsecure()}
		tracer, _, err := NewJaegerTracer("testCli", "127.0.0.1:6831")
		if err != nil {
			fmt.Printf("new tracer err: %+v\n", err)
			os.Exit(-1)
		}

		if tracer != nil {
			dialOpts = append(dialOpts, DialOption(tracer))
		}

		conn, err := grpc.Dial("127.0.0.1:8001", dialOpts...)
		if err != nil {
			fmt.Printf("grpc connect failed, err:%+v\n", err)
			os.Exit(-1)
		}
		defer conn.Close()

		client := pb.NewGreeterClient(conn)
		reqbody := pb.HelloRequest{
			Name:    "im tester",
			Message: "just4test",
		}
		resp, err := client.SayHello(context.Background(), &reqbody)
		if err != nil {
			fmt.Printf("call sayhello failed, err:%+v\n", err)
			os.Exit(-1)
		} else {
			fmt.Printf("call sayhello suc, res:%+v\n", resp)
		}
	}()

	go func() {
		timeout := time.After(time.Second * 3)
		for {
			select {
			case <-timeout:
				os.Exit(0)
			}
		}
	}()

	svr.Serve(ln)
}
