package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"timesheet/GetInfo"
	"timesheet/ProtoApi"
	"timesheet/Server"
	srv "timesheet/Service"
)

func main() {
	srv.InitDB()
	fmt.Println(GetInfo.GetSubjectFromDb())
	LaunchServer()
}
func LaunchServer() {
	s := grpc.NewServer()
	srv := &Server.GRPCServer{}
	ProtoApi.RegisterSheduleServiceServer(s, srv)
	l, _ := net.Listen("tcp", ":8080")
	s.Serve(l)
}
