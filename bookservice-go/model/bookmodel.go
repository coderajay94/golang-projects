package model

import (
	"fmt"
	"log"

	"github.com/coderajay94/bookservice-go/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	*gorm.Model
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func Init() {
	log.Println("Init function called")
	config.ConnectMySQLDB()
	db = config.GetMySQLDB()
	db.AutoMigrate(&Book{})
}

// this is different in writting in v2 of gorm
func (b *Book) CreateBook() error {
	if db == nil {
		return fmt.Errorf("DB connection is nil")
	}
	result := db.Create(&b)
	return result.Error
}

func GetAllBooks() ([]Book, error) {
	var books []Book
	err := db.Find(&books).Error
	return books, err
}

func GetBookByID(id int64) (*Book, error) {
	var book Book
	err := db.First(&book, id).Error
	return &book, err
}

func (b *Book) UpdateBook(id int64) error {
	b.ID = id
	return db.Model(&Book{}).Where("id = ?", id).Updates(b).Error
}

func DeleteBook(id int64) error {
	return db.Delete(&Book{}, id).Error
}
