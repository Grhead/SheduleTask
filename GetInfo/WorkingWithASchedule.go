package GetInfo

import (
	"database/sql"
	"fmt"
	// Import MySQL driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	// Import SQLX package
	_ "github.com/jmoiron/sqlx"
	"time"
)

type ObjectPattern struct {
	Subject    int
	Tutor      int
	Auditorium int
	Type       int
	Group      int
	Date       time.Time
}
type Subjects struct {
	Id          int    `db:"Id"`
	SubjectItem string `db:"Subject_item"`
}

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:123@tcp(127.0.0.1:6603)/Shedule?charset=utf8mb4&parseTime=True"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

//func CreateObjectOfShedule() {
//
//}

func GetSubjectFromDb() []Subjects {
	initDB()
	var arrayOfSubjects []Subjects
	err := db.Select(&arrayOfSubjects, "select * from Subjects")
	if err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return arrayOfSubjects
}

//func GetTutorFromDb() []string {
//
//}
//
//func GetAuditoriumFromDb() []string {
//
//}
//
//func GetTypeFromDb() []string {
//
//}
//
//func GetGroupFromDb() []string {
//
//}
