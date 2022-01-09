package models

import (
	Conf "book-management/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"time"
)

type BookDetail struct {
	Id int `gorm:"primary_key" json:"id"`
	BookNum string `gorm:"varchar(5);not null" json:"book_num"`
	BookName string `gorm:"varchar(20);not null" json:"book_name"`
	Author string `gorm:"varchar(20);not null" json:"author"`
	Publisher string `gorm:"varchar(50);not null" json:"publisher"`
	Category string `gorm:"varchar(20);not null" json:"category"`
	Collection int `gorm:"not null" json:"collection"`
	Lend int `gorm:"not null" json:"lend"`
}

type Borrow struct {
	Id int `gorm:"primary_key" json:"id"`
	BookId int `gorm:"not null" json:"book_id"`
	UserId int `gorm:"not null" json:"user_id"`
	Count int `gorm:"not null" json:"count"`
}

type User struct {
	Id int `gorm:"primary_key" json:"id"`
	Name string `gorm:"varchar(20);not null" json:"name"`
	Password string `gorm:"varchar(100);not null" json:"password"`
	RoleId int `gorm:"not null" json:"role_id"`
}

var db *gorm.DB
var dbConfig = Conf.Conf.Db

func GetGorm() *gorm.DB {
	return db
}

func init() {
	var dbArgs = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbConfig.UserName,
		dbConfig.PassWord,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DbName,
		dbConfig.Charset,
	)
	var err error
	db, err = gorm.Open(dbConfig.EngineName,dbArgs)
	if err != nil {
		panic(err)
	}
	db.DB().SetConnMaxLifetime(4 * time.Hour)
	db.AutoMigrate(&BookDetail{}, &Borrow{}, &User{})
}
