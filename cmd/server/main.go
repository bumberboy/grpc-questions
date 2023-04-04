package main

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"grpc-questions/service/question"
	"net"
	"net/http"
)

func main() {
	s := grpc.NewServer()
	mux := runtime.NewServeMux()

	question.RegisterService(s, "localhost:8090", mux)
	go startGrpcServer(s)
	go startGateway(mux)

	select {}
}

func startGrpcServer(svr *grpc.Server) {
	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		panic(err)
	}
	if err := svr.Serve(l); err != nil {
		panic(err)
	}
}

func startGateway(mux *runtime.ServeMux) {

	if err := http.ListenAndServe(":8091", mux); err != nil {
		panic(err)
	}
}
