package Filter

import (
	srv "timesheet/Service"
	"timesheet/SetInfo"
)

var SheduleTable []SetInfo.ObjectPattern

func FilterFunction(Filter string, value int) []SetInfo.ObjectPattern {
	query := "SELECT * FROM Timesheet WHERE (? = ?)"
	srv.Db.Select(&SheduleTable, query)
	return SheduleTable
}
