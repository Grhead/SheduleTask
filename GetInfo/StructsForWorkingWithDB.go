package GetInfo

import "time"

type ObjectPattern struct {
	Subject    int
	Tutor      int
	Auditorium int
	Type       int
	Group      int
	Date       time.Time
}
type Subjects struct {
	Id          int    `db:"Id"`
	SubjectItem string `db:"Subject_item"`
}
type Classrooms struct {
	Id        int    `db:"Id"`
	Classroom string `db:"Classroom"`
}
type GroupsName struct {
	Id        int    `db:"Id"`
	GroupName string `db:"Group_name"`
}
type Types struct {
	Id   int    `db:"Id"`
	Type string `db:"Types"`
}
type Tutors struct {
	Id         int    `db:"Id"`
	SecondName string `db:"SecondName"`
	FirstName  string `db:"FirstName"`
	LastName   string `db:"LastName"`
}
