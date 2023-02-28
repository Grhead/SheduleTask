package main

import "fmt"
import getting "timesheet/GetInfo"

func main() {
	q := getting.GetSubjectFromDb()
	for _, i := range q {
		fmt.Println(i)
	}
}
