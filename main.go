package main

import (
	"net"
	"os"

	"github.com/joaquinto/Todo-List-gRPC/model"
	"github.com/joaquinto/Todo-List-gRPC/server/handler"
	"google.golang.org/grpc"
	glog "google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

var grpcLog glog.LoggerV2

func init() {
	grpcLog = glog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		grpcLog.Warningf("Unable to create tcp listener: %v", err)
	}
	serv := &handler.TodoServiceServer{}
	server := grpc.NewServer()
	model.RegisterTodoServiceServer(server, serv)
	reflection.Register(server)

	grpcLog.Info("Starting server at port :5000")
	if err := server.Serve(listener); err != nil {
		grpcLog.Warningf("Unable to create server: %v", err)
	}
}
