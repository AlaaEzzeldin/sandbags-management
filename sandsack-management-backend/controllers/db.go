package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type App struct {
	DB *gorm.DB
}


var (
	Server = os.Getenv("DATABASE_HOST")
	Port = os.Getenv("DATABASE_PORT")
	Database = os.Getenv("DATABASE_NAME")
	User = os.Getenv("DATABASE_USER")
	Password = os.Getenv("DATABASE_PASSWORD")
)



func (a *App) Init() {
	// postgres connection
	connString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		Server,
		Port,
		Database,
		User,
		Password,
	)
	var err error
	a.DB, err = gorm.Open("postgres", connString)

	if err != nil {
		//log.Fatal("Connection to database failed: ", err.Error())
	}


	// logging from database
	a.DB.LogMode(true)
	a.DB.SingularTable(true)

}


