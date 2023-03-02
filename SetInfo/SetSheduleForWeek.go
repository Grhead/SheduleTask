package SetInfo

import (
	"time"
	srv "timesheet/Service"
)

type ObjectPattern struct {
	Subject    int       `db:"Subject_item"`
	Tutor      int       `db:"Tutor"`
	Auditorium int       `db:"Classroom"`
	Type       int       `db:"Type"`
	Group      int       `db:"Group_name"`
	Dates      time.Time `db:"Dates"`
	Number     int       `db:"Number"`
}

func CreateObjectOfShedule(NewObjectStruct ObjectPattern) error {
	sqlStr := "INSERT INTO Timesheet (\"Subject_item\", \"Classroom\", \"Tutor\", \"Type\", \"Group_name\", \"Dates\", \"Number\") VALUES (?, ?, ?, ?, ?, ?, ?)"
	_, err := srv.Db.Exec(sqlStr, NewObjectStruct.Subject, NewObjectStruct.Auditorium, NewObjectStruct.Tutor, NewObjectStruct.Type, NewObjectStruct.Group, NewObjectStruct.Dates, NewObjectStruct.Number)
	return err
}
