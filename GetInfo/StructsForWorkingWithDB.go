package GetInfo

type Subjects struct {
	Id          int32  `db:"Id"`
	SubjectItem string `db:"Subject_item"`
}
type Classrooms struct {
	Id        int32  `db:"Id"`
	Classroom string `db:"Classroom"`
}
type Groups struct {
	Id        int32  `db:"Id"`
	GroupName string `db:"Group_name"`
}
type Types struct {
	Id   int32  `db:"Id"`
	Type string `db:"Type"`
}
type Tutors struct {
	Id         int32  `db:"Id"`
	SecondName string `db:"SecondName"`
	FirstName  string `db:"FirstName"`
	LastName   string `db:"LastName"`
}
