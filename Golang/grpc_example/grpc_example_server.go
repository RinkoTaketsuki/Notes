package grpc_example

import (
	"context"
	"fmt"
	"net"

	pd "example/grpc_example/protoes"

	"google.golang.org/grpc"
)

type Server struct{}

func (this *Server) SayHello(ctx context.Context, in *pd.HelloReq) (out *pd.HelloRsp, err error) {
	return &pd.HelloRsp{Msg: "hello"}, nil
}

func (this *Server) SayName(ctx context.Context, in *pd.NameReq) (out *pd.NameRsp, err error) {
	return &pd.NameRsp{Msg: in.Name + "it is name"}, nil
}

func main() {
	ln, err := net.Listen("tcp", ":10088")
	if err != nil {
		fmt.Println("network error", err)
	}

	//创建grpc服务
	srv := grpc.NewServer()
	//注册服务
	pd.RegisterHelloServerServer(srv, &server{})
	err = srv.Serve(ln)
	if err != nil {
		fmt.Println("Serve error", err)
	}
}
