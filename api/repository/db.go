package repository

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func SetDatabase(database *gorm.DB) {
	db = database
}

//Init database connect
func Connect(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *gorm.DB {
	connString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	db, err := gorm.Open(Dbdriver, connString)
	if err != nil {
		panic("cannot connect to the database")
	}
	db.SingularTable(true)
	return db
}
