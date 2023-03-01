package GetInfo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
)

var db *sqlx.DB

const dsn = "root:123@tcp(127.0.0.1:6603)/Shedule?charset=utf8mb4&parseTime=True"

func initDB() (err error) {
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

func CreateObjectOfShedule() {

}

func GetSubjectFromDb() []Subjects {
	initDB()
	var arrayOfSubjects []Subjects
	err := db.Select(&arrayOfSubjects, "select * from Subjects")
	if err != sql.ErrNoRows {
		fmt.Println(err)
	}
	return arrayOfSubjects
}

func GetTutorFromDb() []string {

}

func GetAuditoriumFromDb() []string {

}

func GetTypeFromDb() []string {

}

func GetGroupFromDb() []string {

}
