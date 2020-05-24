package db

import (
	"log"
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/mitchellh/go-homedir"
)

func init() {
	homeDir, e := homedir.Dir()
	if e != nil {
		log.Fatal(e)
	}
	dbPath := path.Join(homeDir, "users.db")
	db, e := gorm.Open("sqlite3", dbPath)
	defer db.Close()
	if e != nil {
		log.Fatal(e)
	}
	db.AutoMigrate(&User{})
	var count int
	db.Model(&User{}).Count(&count)
	if count == 0 {
		db.Create(&User{Name: "Foo", PhoneNumber: "01234 56789"})
		db.Create(&User{Name: "Bar", PhoneNumber: "01-234567-89"})
		db.Create(&User{Name: "Baz", PhoneNumber: "01234_56789"})
	}
}

type User struct {
	gorm.Model
	Name        string
	PhoneNumber string
}

type DB struct {
	db *gorm.DB
}

func New() (*DB, error) {
	homeDir, e := homedir.Dir()
	if e != nil {
		return nil, e
	}
	dbPath := path.Join(homeDir, "users.db")
	db, e := gorm.Open("sqlite3", dbPath)
	if e != nil {
		return nil, e
	}
	return &DB{db: db}, nil
}

func (o *DB) Close() error {
	return o.db.Close()
}

func (o *DB) GetAllUsers() []User {
	var users []User
	o.db.Find(&users)
	return users
}

func (o *DB) UpdateUser(u User) {
	var user User
	o.db.First(&user, u.ID)
	user.Name = u.Name
	user.PhoneNumber = u.PhoneNumber
	o.db.Save(&user)
}
