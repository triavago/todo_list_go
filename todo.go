package main

import (
	"log"
	"time"

	//"net/http"
	//"strconv"
	//"context"
	//"fmt"
	//"database/sql"
	//"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	//"github.com/aws/aws-sdk-go-v2/config"
	//"github.com/aws/aws-sdk-go-v2/feature/rds/auth"
	//_ "github.com/go-sql-driver/mysql"
)

type ToDoItem struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Title     string     `json:"title" gorm:"column:title;"`
	Status    string     `json:"status" gorm:"column:status;"`
	CreatedAt *time.Time `json:"created_at" gorm:"column:created_at;"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"column:updated_at;"`
}

func (ToDoItem) TableName() string { return "todo_items" }

func main() {

	dsn := "admin:quang123@tcp(database-1.cuuvxe1pfre7.ap-southeast-1.rds.amazonaws.com:3306)/"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	log.Print("Connected:", db)
	if err = db.Exec("CREATE DATABASE IF NOT EXISTS todo_db").Error; err != nil {
		log.Println(err)
        return
	}
	dsn = "admin:quang123@tcp(database-1.cuuvxe1pfre7.ap-southeast-1.rds.amazonaws.com:3306)/todo_db?charset=utf8mb4&parseTime=True"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil {
        log.Println(err)
        return
	}
	db.Table("todo_db").AutoMigrate(&ToDoItem{})
}
