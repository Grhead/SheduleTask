package FilterPack

import (
	"github.com/jinzhu/now"
	"sort"
	"time"
	srv "timesheet/Service"
	"timesheet/SetInfo"
)

type ArrayStruct struct {
	SubjectsOfThisDay []SetInfo.ObjectPattern
}

func FilterFunction(Filter string, value int32) []*SetInfo.ObjectPattern {
	var AllObjectsFromDb []*SetInfo.ObjectPattern
	BeginningOfThisWeek := now.BeginningOfWeek()
	EndingOfThisWeek := now.EndOfWeek()
	if Filter == "0" && value == 0 {
		query := "SELECT * FROM Timesheet WHERE (Dates < ? AND Dates > ?)"
		srv.Db.Select(&AllObjectsFromDb, query, EndingOfThisWeek, BeginningOfThisWeek)
	} else {
		query := "SELECT * FROM Timesheet WHERE (? = ?)"
		srv.Db.Select(&AllObjectsFromDb, query)
	}
	return AllObjectsFromDb
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
		for _, v := range MissingSubjectsMap {
			if v == 0 {
				Week[i].SubjectsOfThisDay = append(Week[i].SubjectsOfThisDay, newNullObjects)
			}
		}
		sort.Slice(Week[i].SubjectsOfThisDay, func(ik, jk int) bool {
			return Week[i].SubjectsOfThisDay[ik].Number < Week[i].SubjectsOfThisDay[jk].Number
		})
	}
	return Week
}
