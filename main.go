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
	ProtoApi.RegisterSheduleServiceServer(s, srvi)
	l, _ := net.Listen("tcp", ":8383")
	s.Serve(l)
}
