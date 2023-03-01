package GetInfo

import (
	"database/sql"
	"fmt"
	srv "timesheet/Service"
)

func GetSubjectFromDb() []Subjects {
	var arrayOfSubjects []Subjects
	err := srv.Db.Select(&arrayOfSubjects, "select * from Subjects")
	if err != sql.ErrNoRows {
		fmt.Println(err)
	}
	defer srv.Db.Close()
	return arrayOfSubjects
}

func GetTutorsFromDb() []Tutors {
	var arrayOfTutors []Tutors
	err := srv.Db.Select(&arrayOfTutors, "select * from Tutors")
	if err != sql.ErrNoRows {
		fmt.Println(err)
	}
	defer srv.Db.Close()
	return arrayOfTutors
}

func GetAuditoriumFromDb() []Classrooms {
	var arrayOfTutors []Classrooms
	err := srv.Db.Select(&arrayOfTutors, "select * from Classrooms")
	if err != sql.ErrNoRows {
		fmt.Println(err)
	}
	defer srv.Db.Close()
	return arrayOfTutors
}

func GetTypeFromDb() []Types {
	var arrayOfTutors []Types
	err := srv.Db.Select(&arrayOfTutors, "select * from Types")
	if err != sql.ErrNoRows {
		fmt.Println(err)
	}
	defer srv.Db.Close()
	return arrayOfTutors
}

func GetGroupFromDb() []GroupsName {
	var arrayOfGroups []GroupsName
	err := srv.Db.Select(&arrayOfGroups, "select * from Groups_name")
	if err != sql.ErrNoRows {
		fmt.Println(err)
	}
	defer srv.Db.Close()
	return arrayOfGroups
}
