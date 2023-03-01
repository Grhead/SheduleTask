package main

import "fmt"
import getting "timesheet/GetInfo"
import srv "timesheet/Service"

func main() {
	srv.InitDB()
	//q := getting.GetSubjectFromDb()
	q := getting.GetTutorsFromDb()
	for _, i := range q {
		fmt.Println(i)
	}
}
