package server_http

import (
	"net/http"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"

	gw "wps_store/proto"
)

var (
    ServerPort string
	ServerHttpPort string
	EndPoint string
	HttpEndPoint string
)

func Run() (err error) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// grpc服务地址
	EndPoint = ":" + ServerPort
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	// HTTP转grpc
	err = gw.RegisterStoreApiServiceHandlerFromEndpoint(ctx, mux, EndPoint, opts)
	if err != nil {
		grpclog.Fatalf("Register handler err:%v\n", err)
	}
	HttpEndPoint = ":" + ServerHttpPort
	log.Println("HTTP Listen success:", HttpEndPoint)
	http.ListenAndServe(HttpEndPoint, mux)
	return err
}