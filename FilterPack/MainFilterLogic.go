package FilterPack

import (
	"fmt"
	"github.com/jinzhu/now"
	"sort"
	"time"
	srv "timesheet/Service"
	"timesheet/SetInfo"
)

type ArrayStruct struct {
	SubjectsOfThisDay []SetInfo.ObjectPattern
}

type ArrayStructString struct {
	SubjectsOfThisDay []SetInfo.StringObjectPattern
}

func LowFilterFunction(Filter string, value int32) []*SetInfo.ObjectPattern {
	var AllObjectsFromDb []*SetInfo.ObjectPattern
	BeginningOfThisWeek := now.BeginningOfWeek()
	EndingOfThisWeek := now.EndOfWeek()
	if Filter == "0" && value == 0 {
		query := "SELECT * FROM Timesheet WHERE (Dates < ? AND Dates > ?)"
		err := srv.Db.Select(&AllObjectsFromDb, query, EndingOfThisWeek, BeginningOfThisWeek)
		fmt.Println(err)
	} else {
		query := "SELECT * FROM Timesheet WHERE (? = ?)"
		srv.Db.Select(&AllObjectsFromDb, query)
	}
	return AllObjectsFromDb
}

func FilterFunction(Filter string, value int32) []*SetInfo.StringObjectPattern {
	var AllObjectsFromDb []*SetInfo.StringObjectPattern
	BeginningOfThisWeek := now.BeginningOfWeek().Format("20060102")
	EndingOfThisWeek := now.EndOfWeek().Format("20060102")
	if Filter == "0" && value == 0 {
		query := "SELECT Timesheet.Id, S.Subject_item, C.Classroom, Gn.Group_name, T.SecondName, T.FirstName, T.LastName, T2.Type, Number, Timesheet.Dates\nFROM Timesheet JOIN Subjects S on S.Id = Timesheet.Subject_item JOIN Classrooms C on C.Id = Timesheet.Classroom JOIN Groups_name Gn on Gn.Id = Timesheet.Group_name JOIN Tutors T on T.Id = Timesheet.Tutor JOIN Types T2 on T2.Id = Timesheet.Type WHERE (Dates < ? AND Dates > ?)"
		srv.Db.Select(&AllObjectsFromDb, query, EndingOfThisWeek, BeginningOfThisWeek)
	} else {
		query := "SELECT Timesheet.Id, S.Subject_item, C.Classroom, Gn.Group_name, T.SecondName, T.FirstName, T.LastName, T2.Type, Number, Timesheet.Dates\nFROM Timesheet JOIN Subjects S on S.Id = Timesheet.Subject_item JOIN Classrooms C on C.Id = Timesheet.Classroom JOIN Groups_name Gn on Gn.Id = Timesheet.Group_name JOIN Tutors T on T.Id = Timesheet.Tutor JOIN Types T2 on T2.Id = Timesheet.Type WHERE (? = ?)"
		err := srv.Db.Select(&AllObjectsFromDb, query, Filter, value)
		fmt.Println(err)
	}
	return AllObjectsFromDb
}

func GetDaysOfWeekForStrings(AllObjectsFromDb []*SetInfo.StringObjectPattern) [7]ArrayStructString {
	var SheduleTableByDayOfWeek [7]ArrayStructString
	for _, i := range AllObjectsFromDb {
		DayOfWeekByDateFromDb := i.Dates.Weekday()
		switch DayOfWeekByDateFromDb {
		case 1:
			SheduleTableByDayOfWeek[0].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[0].SubjectsOfThisDay, *i)
		case 2:
			SheduleTableByDayOfWeek[1].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[1].SubjectsOfThisDay, *i)
		case 3:
			SheduleTableByDayOfWeek[2].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[2].SubjectsOfThisDay, *i)
		case 4:
			SheduleTableByDayOfWeek[3].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[3].SubjectsOfThisDay, *i)
		case 5:
			SheduleTableByDayOfWeek[4].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[4].SubjectsOfThisDay, *i)
		case 6:
			SheduleTableByDayOfWeek[5].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[5].SubjectsOfThisDay, *i)
		case 7:
			SheduleTableByDayOfWeek[6].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[6].SubjectsOfThisDay, *i)
		}
	}
	return SheduleTableByDayOfWeek
}

func GetDaysOfWeek(AllObjectsFromDb []*SetInfo.ObjectPattern) [7]ArrayStruct {
	var SheduleTableByDayOfWeek [7]ArrayStruct
	for _, i := range AllObjectsFromDb {
		DayOfWeekByDateFromDb := i.Dates.Weekday()
		switch DayOfWeekByDateFromDb {
		case 1:
			SheduleTableByDayOfWeek[0].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[0].SubjectsOfThisDay, *i)
		case 2:
			SheduleTableByDayOfWeek[1].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[1].SubjectsOfThisDay, *i)
		case 3:
			SheduleTableByDayOfWeek[2].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[2].SubjectsOfThisDay, *i)
		case 4:
			SheduleTableByDayOfWeek[3].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[3].SubjectsOfThisDay, *i)
		case 5:
			SheduleTableByDayOfWeek[4].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[4].SubjectsOfThisDay, *i)
		case 6:
			SheduleTableByDayOfWeek[5].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[5].SubjectsOfThisDay, *i)
		case 7:
			SheduleTableByDayOfWeek[6].SubjectsOfThisDay = append(SheduleTableByDayOfWeek[6].SubjectsOfThisDay, *i)
		}
	}
	return SheduleTableByDayOfWeek
}

func AdditionSubjectToEachDay(Week [7]ArrayStruct) [7]ArrayStruct {
	var newNullObjects SetInfo.ObjectPattern
	newNullObjects.Type = -1
	newNullObjects.Tutor = -1
	newNullObjects.Subject = -1
	newNullObjects.Auditorium = -1
	newNullObjects.Group = -1
	newNullObjects.Dates = time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	newNullObjects.Id = -1
	for i := 0; i < 7; i++ {
		var MissingSubjectsMap = make(map[int32]int)
		for k := 0; k < 8; k++ {
			MissingSubjectsMap[int32(k)] = 0
		}
		for _, j := range Week[i].SubjectsOfThisDay {
			MissingSubjectsMap[j.Number]++
		}
		for w, v := range MissingSubjectsMap {
			if v == 0 {
				newNullObjects.Number = w
				Week[i].SubjectsOfThisDay = append(Week[i].SubjectsOfThisDay, newNullObjects)
			}
		}
		sort.Slice(Week[i].SubjectsOfThisDay, func(ik, jk int) bool {
			return Week[i].SubjectsOfThisDay[ik].Number < Week[i].SubjectsOfThisDay[jk].Number
		})
	}
	return Week
}

func AdditionSubjectToEachDayForStrings(Week [7]ArrayStructString) [7]ArrayStructString {
	var newNullObjects SetInfo.StringObjectPattern
	newNullObjects.Type = "-1"
	newNullObjects.LastName = "-1"
	newNullObjects.FirstName = "-1"
	newNullObjects.SecondName = "-1"
	newNullObjects.Subject = "-1"
	newNullObjects.Auditorium = "-1"
	newNullObjects.Group = "-1"
	newNullObjects.Dates = time.Date(0001, 1, 1, 0, 0, 0, 0, time.UTC)
	newNullObjects.Id = -1
	for i := 0; i < 7; i++ {
		var MissingSubjectsMap = make(map[int32]int)
		for k := 0; k < 8; k++ {
			MissingSubjectsMap[int32(k)] = 0
		}
		for _, j := range Week[i].SubjectsOfThisDay {
			MissingSubjectsMap[j.Number]++
		}
		for w, v := range MissingSubjectsMap {
			if v == 0 {
				newNullObjects.Number = w
				Week[i].SubjectsOfThisDay = append(Week[i].SubjectsOfThisDay, newNullObjects)
			}
		}
		sort.Slice(Week[i].SubjectsOfThisDay, func(ik, jk int) bool {
			return Week[i].SubjectsOfThisDay[ik].Number < Week[i].SubjectsOfThisDay[jk].Number
		})
	}
	return Week
}
