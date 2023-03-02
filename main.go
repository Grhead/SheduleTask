package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
	"timesheet/GetInfo"
	"timesheet/ProtoApi"
	"timesheet/Server"
	srv "timesheet/Service"
	"timesheet/SetInfo"
)

func main() {
	srv.InitDB()
	fmt.Println(GetInfo.GetSubjectFromDb())
	//LaunchServer()
	var q []SetInfo.ObjectPattern
	var t SetInfo.ObjectPattern
	t.Auditorium = 1
	t.Tutor = 2
	t.Type = 3
	t.Subject = 2
	t.Number = 5
	t.Group = 1
	t.Dates = time.Date(2023, time.March, 22, 0, 0, 0, 0, time.UTC)
	q = append(q, t)
	fmt.Println(SetInfo.InsertionToDb(q))
}
func LaunchServer() {
	s := grpc.NewServer()
	srv := &Server.GRPCServer{}
	ProtoApi.RegisterSheduleServiceServer(s, srv)
	l, _ := net.Listen("tcp", ":8080")
	s.Serve(l)
}
