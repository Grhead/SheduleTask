package Filter

import (
	srv "timesheet/Service"
	"timesheet/SetInfo"
)

type ArrayStruct struct {
	arr []SetInfo.ObjectPattern
}

var SheduleTableByDayOfWeek [7]ArrayStruct

func FilterFunction(Filter string, value int) []SetInfo.ObjectPattern {
	var AllObjectsFromDb []SetInfo.ObjectPattern
	if Filter == "0" && value == 0 {
		query := "SELECT * FROM Timesheet"
		srv.Db.Select(&AllObjectsFromDb, query)
	} else {
		query := "SELECT * FROM Timesheet WHERE (? = ?)"
		srv.Db.Select(&AllObjectsFromDb, query)
	}
	return AllObjectsFromDb
}
func GetDayOfWeek(AllObjectsFromDb []SetInfo.ObjectPattern) [7]ArrayStruct {
	for _, i := range AllObjectsFromDb {
		DayOfWeekByDateFromDb := i.Dates.Weekday()
		switch DayOfWeekByDateFromDb {
		case 1:

			SheduleTableByDayOfWeek[0].arr = append(SheduleTableByDayOfWeek[0].arr, i)

		case 2:

			SheduleTableByDayOfWeek[1].arr = append(SheduleTableByDayOfWeek[1].arr, i)

		case 3:

			SheduleTableByDayOfWeek[2].arr = append(SheduleTableByDayOfWeek[2].arr, i)

		case 4:

			SheduleTableByDayOfWeek[3].arr = append(SheduleTableByDayOfWeek[3].arr, i)

		case 5:

			SheduleTableByDayOfWeek[4].arr = append(SheduleTableByDayOfWeek[4].arr, i)

		case 6:

			SheduleTableByDayOfWeek[5].arr = append(SheduleTableByDayOfWeek[5].arr, i)

		case 7:

			SheduleTableByDayOfWeek[6].arr = append(SheduleTableByDayOfWeek[6].arr, i)

		}
	}
	return SheduleTableByDayOfWeek
}
