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
	var arrayOfAuditoriums []Classrooms
	srv.Db.Select(&arrayOfAuditoriums, "select * from Classrooms")
	return arrayOfAuditoriums
}

func GetTypeFromDb() []Types {
	var arrayOfTypes []Types
	srv.Db.Select(&arrayOfTypes, "select * from Types")
	return arrayOfTypes
}

func GetGroupFromDb() []Groups {
	var arrayOfGroups []Groups
	srv.Db.Select(&arrayOfGroups, "select * from Groups_name")
	return arrayOfGroups
}
