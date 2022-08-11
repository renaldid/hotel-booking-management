package app

import (
	"database/sql"
	"github.com/renaldid/hotel_booking_management.git/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	//db, err := sql.Open("mysql", "root@tcp(localhost:3306)/booking_management_system")
	db, err := sql.Open("mysql", "sql6507898:TvznJXBLjy@tcp(sql6.freemysqlhosting.net:3306)/sql6507898")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
