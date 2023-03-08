package SetInfo

import (
	"fmt"
	"github.com/jinzhu/now"
	"time"
	"timesheet/ProtoApi"
	srv "timesheet/Service"
)

type ObjectPattern struct {
	Id         int32     `db:"Id"`
	Subject    int32     `db:"Subject_item"`
	Tutor      int32     `db:"Tutor"`
	Auditorium int32     `db:"Classroom"`
	Type       int32     `db:"Type"`
	Group      int32     `db:"Group_name"`
	Dates      time.Time `db:"Dates"`
	Number     int32     `db:"Number"`
}

type StringObjectPattern struct {
	Id         int32     `db:"Id"`
	Subject    string    `db:"Subject_item"`
	Tid        string    `db:"T.Id"`
	SecondName string    `db:"SecondName"`
	FirstName  string    `db:"FirstName"`
	LastName   string    `db:"LastName"`
	Auditorium string    `db:"Classroom"`
	Type       string    `db:"Type"`
	Group      string    `db:"Group_name"`
	Dates      time.Time `db:"Dates"`
	Number     int32     `db:"Number"`
}

func сreateObjectOfShedule(NewObjectStruct *ObjectPattern) error {
	fmt.Println(NewObjectStruct)
	sqlStr1 := "DELETE FROM Timesheet WHERE Group_name = ? AND Number = ? AND Dates = ?"
	srv.Db.Exec(sqlStr1, NewObjectStruct.Group, NewObjectStruct.Number, NewObjectStruct.Dates)
	sqlStr := "INSERT INTO Timesheet (Subject_item, Classroom, Tutor, Type, Group_name, Dates, Number) VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := srv.Db.Exec(sqlStr, NewObjectStruct.Subject, NewObjectStruct.Auditorium, NewObjectStruct.Tutor, NewObjectStruct.Type, NewObjectStruct.Group, NewObjectStruct.Dates, NewObjectStruct.Number)

	return err
}
func InsertionToDb(SheduleTable []*ObjectPattern) error {
	var err error
	for _, i := range SheduleTable {
		err = сreateObjectOfShedule(i)
	}
	return err
}

func DatesForDaysOfWeek() []*ProtoApi.DateTime {
	ArrayOfDates := []*ProtoApi.DateTime{}
	for i := 0; i < 7; i++ {
		ArrayOfDates = append(ArrayOfDates, &ProtoApi.DateTime{Year: int32(now.BeginningOfWeek().Year()), Month: int32(now.BeginningOfWeek().Month()), Day: int32(now.BeginningOfWeek().Day()) + 1})
	}
	return ArrayOfDates
}
