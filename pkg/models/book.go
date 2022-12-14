package models

import(
	"github.com/jinzhu/gorm"
	"github.com/PrathameshAnwekar/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`

}

func init(){
	config.Connect()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book{
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)

	return &getBook, db
}

func DeleteBook(id int64) Book{
	var book Book
	db.Where("ID = ?", id).Delete(&book)
	return book
}