package main

import (
	"log"

	"errors"

	"github.com/jinzhu/gorm"
)

// User - Struct used to hold data about a single user
type User struct {
	ID       int       `gorm:"column:id;primary_key" sql:"not null;unique;AUTO_INCREMENT"`
	Name     string    `gorm:"column:name" sql:"type:varchar(45);not null"`
	Surname  string    `gorm:"column:surname" sql:"type:varchar(45);not null"`
	Email    string    `gorm:"column:email" sql:"type:varchar(255);not null;unique"`
	Google   bool      `gorm:"column:google" sql:"DEFAULT:false"`
	Drawings []Drawing `gorm:"ForeignKey:id"`
	Games    []Game    `gorm:"many2many:user_played_game"`
}

// CreateUser - Used to create a new entry in the "users" table
func CreateUser(db *gorm.DB, user *User) (err error) {
	if db.NewRecord(user) {
		if count := db.Save(&user).RowsAffected; count != 1 {
			log.Print("Not created: ", user)
			err = errors.New("Error creating user!")
		}
	} else {
		err = errors.New("Error checking for user!")
	}

	return err
}

// ReadUser - Used to read an entry from the "users" table
func ReadUser(db *gorm.DB, user User) (ret User, err error) {
	err = db.Where(&user).First(&ret).Error
	return ret, err
}

// UpdateUser - Used to update an entry in the "users" table
func UpdateUser(db *gorm.DB, user User) (err error) {
	return db.Model(&user).Update(map[string]interface{}{"name": user.Name, "surname": user.Surname, "email": user.Email, "google": user.Google}).Error
}

// DeleteUser - Used to remove an entry grom the "users" table
func DeleteUser(db *gorm.DB, user User) (err error) {
	return db.Delete(&user).Error
}
