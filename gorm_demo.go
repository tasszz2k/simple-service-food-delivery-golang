package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {

	// Connect to MySQL

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	//dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost,
		dbPort, dbName)

	fmt.Println(dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db, err)
	if err != nil {
		log.Panicln(err)
	}

	// Insert new Note
	newNote := Note{
		Title:   "My First Note",
		Content: "This is my first note",
	}
	if err := db.Create(&newNote).Error; err != nil {
		log.Panicln(err)
	}
	fmt.Println(newNote)

	// Select all Notes where id > 1
	var notes []Note
	if err := db.Where("id > ?", 1).Find(&notes).Error; err != nil {
		log.Panicln(err)
	}
	fmt.Println(notes)

}

/**
  id      INT AUTO_INCREMENT,
  title   VARCHAR(100) NOT NULL,
  content TEXT         NULL,
  image   JSON         NULL,
*/
type Note struct {
	Id      int    `json:"id,ommitempty" gorm:"column:id;primary_key;AUTO_INCREMENT"`
	Title   string `json:"title" gorm:"column:title"`
	Content string `json:"content" gorm:"column:content"`
}

func (Note) TableName() string {
	return "notes"
}
