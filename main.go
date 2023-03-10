package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"timesheet/AuthPack"
	"timesheet/ProtoApi"
	"timesheet/Server"
	srv "timesheet/Service"
)

func main() {
	srv.InitDB()
	fmt.Println(AuthPack.Authorization("user", "user"))
	LaunchServer()
}
func LaunchServer() {
	s := grpc.NewServer()
	srvi := &Server.GRPCServer{}
	ProtoApi.RegisterSheduleServiceServer(s, srvi)
	l, _ := net.Listen("tcp", ":8383")
	s.Serve(l)
}
