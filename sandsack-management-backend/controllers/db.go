package controllers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type App struct {
	DB *gorm.DB
}

//var (
//	Server = os.Getenv("DATABASE_HOST")
//	Port = os.Getenv("DATABASE_PORT")
//	Database = os.Getenv("DATABASE_NAME")
//	User = os.Getenv("DATABASE_USER")
//	Password = os.Getenv("DATABASE_PASSWORD")
//)

var (
	Server   = "ls-f219193e8a5ca88cde62a0adc5375d7587a2b8eb.c0uxyt5nyufx.eu-central-1.rds.amazonaws.com"
	Port     = "5432"
	Database = "dbfeuerwehr"
	User     = "dbmasteruser"
	Password = "6CgB7rEAvXKaZkQtqu6|wjq}w{_<LH{0\n\n\n"
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

	a.DB, err = gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	//a.DB, err = gorm.Open("postgres", connString)
	if err != nil {
		//log.Println("Connection to database failed:", err.Error())
		log.Fatal("Connection to database failed: ", err.Error())
	}
}
