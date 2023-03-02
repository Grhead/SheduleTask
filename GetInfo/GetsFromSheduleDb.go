package GetInfo

import (
	srv "timesheet/Service"
)

func GetSubjectFromDb() []Subjects {
	var arrayOfSubjects []Subjects
	srv.Db.Select(&arrayOfSubjects, "select * from Subjects")
	return arrayOfSubjects
}

func GetTutorsFromDb() []Tutors {
	var arrayOfTutors []Tutors
	srv.Db.Select(&arrayOfTutors, "select * from Tutors")
	return arrayOfTutors
}

func GetAuditoriumFromDb() []Classrooms {
	var arrayOfTutors []Classrooms
	srv.Db.Select(&arrayOfTutors, "select * from Classrooms")
	return arrayOfTutors
}

func GetTypeFromDb() []Types {
	var arrayOfTutors []Types
	srv.Db.Select(&arrayOfTutors, "select * from Types")
	return arrayOfTutors
}

func GetGroupFromDb() []Groups {
	var arrayOfGroups []Groups
	srv.Db.Select(&arrayOfGroups, "select * from Groups_name")
	return arrayOfGroups
}
