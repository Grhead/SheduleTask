package SetInfo

import (
	"time"
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
	Tutor      string    `db:"Tutor"`
	Auditorium string    `db:"Classroom"`
	Type       string    `db:"Type"`
	Group      string    `db:"Group_name"`
	Dates      time.Time `db:"Dates"`
	Number     int32     `db:"Number"`
}

func сreateObjectOfShedule(NewObjectStruct *ObjectPattern) error {
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
