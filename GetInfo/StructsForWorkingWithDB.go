package GetInfo

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
