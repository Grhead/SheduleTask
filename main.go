package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"timesheet/FilterPack"
	"timesheet/ProtoApi"
	"timesheet/Server"
	srv "timesheet/Service"
	"timesheet/SetInfo"
)

func main() {
	srv.InitDB()
	//LaunchServer()
	//FilterPack.FilterFunction("0", 0)
	//var q = FilterPack.GetDayOfWeek(FilterPack.FilterFunction("0", 0))
	var q [7]FilterPack.ArrayStruct
	q = FilterPack.GetDayOfWeek(FilterPack.FilterFunction("0", 0))
	temp := SetInfo.ObjectPattern{}
	for j := 0; j < 8; j++ {
		for _, i := range q {
			if i.Arr[j] != temp {
				fmt.Println(i.Arr[j])
			} else {
				i.Arr = append(i.Arr, temp)
			}
		}
	}
	//fmt.Println(q[0].Arr[2])
}
func LaunchServer() {
	s := grpc.NewServer()
	srvi := &Server.GRPCServer{}
	ProtoApi.RegisterSheduleServiceServer(s, srvi)
	l, _ := net.Listen("tcp", ":8787")
	s.Serve(l)
}
