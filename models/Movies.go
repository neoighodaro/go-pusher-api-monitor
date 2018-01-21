package models

import (
	"github.com/jinzhu/gorm"
	//sqlite3 import
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//Movies - movies model object
type Movies struct {
	gorm.Model
	Name               string `sql:"not null;"`
	Genre, Description string
	Year               int64
}

var db, _ = gorm.Open("sqlite3", "./db/gorm.db")

//Get - get all movies
func (m Movies) Get() []Movies {
	var movies []Movies
	db.Find(&movies)

	return movies
}

//GetByID - get movie by id
func (m Movies) GetByID(ID int64) Movies {
	db.First(&m, ID)

	return m
}

//Create - create new movie
func (m Movies) Create() Movies {
	db.Create(&m)

	return m
}

//Delete - delete movie
func (m Movies) Delete() Movies {
	db.Delete(&m)

	return m
}

//Edit - edit movie
func (m Movies) Edit() Movies {
	db.Save(&m)

	return m
}

//Validate - validate movies
func (m Movies) Validate() bool {

	if m == (Movies{}) || m.Name == "" {
		return false
	}

	return true
}
