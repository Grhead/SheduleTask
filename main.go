package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"timesheet/FilterPack"
	"timesheet/ProtoApi"
	"timesheet/Server"
	srv "timesheet/Service"
)

func main() {
	srv.InitDB()
	LaunchServer()
	//FilterPack.FilterFunction("0", 0)
	//var q = FilterPack.GetDaysOfWeek(FilterPack.FilterFunction("0", 0))
	var q [7]FilterPack.ArrayStruct
	q = FilterPack.GetDaysOfWeek(FilterPack.FilterFunction("0", 0))
	var t = FilterPack.AdditionSubjectToEachDay(q)
	fmt.Println(t[0].SubjectsOfThisDay[0].Dates.Day())

	//fmt.Println(q[0].SubjectsOfThisDay[2])
}
func LaunchServer() {
	s := grpc.NewServer()
	srvi := &Server.GRPCServer{}
	ProtoApi.RegisterSheduleServiceServer(s, srvi)
	l, _ := net.Listen("tcp", ":8787")
	s.Serve(l)
}
