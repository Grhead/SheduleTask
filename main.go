package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"timesheet/Filter"
	"timesheet/ProtoApi"
	"timesheet/Server"
	srv "timesheet/Service"
)

func main() {
	srv.InitDB()
	//LaunchServer()
	Filter.FilterFunction("0", 0)
	var q = Filter.GetDayOfWeek(Filter.FilterFunction("0", 0))
	for _, i := range q {
		fmt.Println(i)
	}
	//for _, i := range *q {
	//	fmt.Println(i.SheduleTableItem)
	//}
}
func LaunchServer() {
	s := grpc.NewServer()
	srv := &Server.GRPCServer{}
	ProtoApi.RegisterSheduleServiceServer(s, srv)
	l, _ := net.Listen("tcp", ":8787")
	s.Serve(l)
}
