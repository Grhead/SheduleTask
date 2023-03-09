package main

import (
	"google.golang.org/grpc"
	"net"
	"timesheet/ProtoApi"
	"timesheet/Server"
	srv "timesheet/Service"
)

func main() {
	srv.InitDB()
	LaunchServer()
}
func LaunchServer() {
	s := grpc.NewServer()
	srvi := &Server.GRPCServer{}
	l, _ := net.Listen("tcp", ":8787")
	ProtoApi.RegisterSheduleServiceServer(s, srvi)
	s.Serve(l)
}
