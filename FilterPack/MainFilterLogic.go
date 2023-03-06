package FilterPack

import (
	"github.com/jinzhu/now"
	srv "timesheet/Service"
	"timesheet/SetInfo"
)

type ArrayStruct struct {
	Arr []SetInfo.ObjectPattern
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
func GetDayOfWeek(AllObjectsFromDb []*SetInfo.ObjectPattern) [7]ArrayStruct {
	var SheduleTableByDayOfWeek [7]ArrayStruct
	for _, i := range AllObjectsFromDb {
		DayOfWeekByDateFromDb := i.Dates.Weekday()
		switch DayOfWeekByDateFromDb {
		case 1:
			SheduleTableByDayOfWeek[0].Arr = append(SheduleTableByDayOfWeek[0].Arr, *i)
		case 2:
			SheduleTableByDayOfWeek[1].Arr = append(SheduleTableByDayOfWeek[1].Arr, *i)
		case 3:
			SheduleTableByDayOfWeek[2].Arr = append(SheduleTableByDayOfWeek[2].Arr, *i)
		case 4:
			SheduleTableByDayOfWeek[3].Arr = append(SheduleTableByDayOfWeek[3].Arr, *i)
		case 5:
			SheduleTableByDayOfWeek[4].Arr = append(SheduleTableByDayOfWeek[4].Arr, *i)
		case 6:
			SheduleTableByDayOfWeek[5].Arr = append(SheduleTableByDayOfWeek[5].Arr, *i)
		case 7:
			SheduleTableByDayOfWeek[6].Arr = append(SheduleTableByDayOfWeek[6].Arr, *i)
		}
	}
	return SheduleTableByDayOfWeek
}
