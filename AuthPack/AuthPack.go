package AuthPack

import (
	"fmt"
	srv "timesheet/Service"
)

type IsAuthComplete struct {
	Id    int32  `db:"Id"`
	Login string `db:"Login"`
	Pass  string `db:"Pass"`
	Role  int32  `db:"Role"`
}

func Authorization(login string, pass string) (bool, int32) {
	var IsAccept bool
	var IsAcceptRole int32
	var result IsAuthComplete
	query := "SELECT Id, Login, Pass, Role FROM Accounts WHERE Login = ? AND Pass = ?"
	srv.Db.Get(&result, query, login, pass)
	fmt.Println(result)
	fmt.Println(result.Id)
	if result.Id == 0 {
		IsAccept = false
		IsAcceptRole = -1
	} else {
		IsAccept = true
		IsAcceptRole = result.Role
	}
	return IsAccept, IsAcceptRole
}
