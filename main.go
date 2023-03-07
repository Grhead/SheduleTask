package main

import (
	"fmt"
	"github.com/jinzhu/now"
	"google.golang.org/grpc"
	"net"
	"strconv"
	"timesheet/FilterPack"
	"timesheet/ProtoApi"
	"timesheet/Server"
	srv "timesheet/Service"
)

func main() {
	srv.InitDB()
	//LaunchServer()
	//FilterPack.LowFilterFunction("0", 0)
	//var q = FilterPack.GetDaysOfWeek(FilterPack.LowFilterFunction("0", 0))
	//var q [7]FilterPack.ArrayStructString
	fmt.Println(strconv.Itoa(now.BeginningOfWeek().Year()) + strconv.Itoa(int(now.BeginningOfWeek().Month())) + strconv.Itoa(now.BeginningOfWeek().Day()))
	q := FilterPack.FilterFunction("0", 0)
	//var t = FilterPack.AdditionSubjectToEachDayForStrings(q)
	for _, i := range q {
		fmt.Println(i)
	}

	//fmt.Println(q[0].SubjectsOfThisDay[2])
}
func LaunchServer() {
	s := grpc.NewServer()
	srvi := &Server.GRPCServer{}
	ProtoApi.RegisterSheduleServiceServer(s, srvi)
	l, _ := net.Listen("tcp", ":8787")
	s.Serve(l)
}
