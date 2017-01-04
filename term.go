package main

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

// Term - A struct used to hold the data for a single Term
type Term struct {
	ID          uint32 `gorm:"column:id;not null;unique;AUTO_INCREMENT"`
	Term        string `gorm:"column:term;type:varchar(255);not null"`
	Explanation string `gorm:"column:explanation;type:varchar(255);not null"`
}

// CreateTerm - Used to create a new entry in the "term" table
func CreateTerm(db *gorm.DB, term *Term) (err error) {
	if db.NewRecord(term) {
		if count := db.Save(&term).RowsAffected; count != 1 {
			log.Print("Not created: ", term)
			err = errors.New("Error creating term!")
		}
	} else {
		err = errors.New("Error checking for term!")
	}

	return err
}

// ReadTerm - Used to read an entry from the "term" table
func ReadTerm(db *gorm.DB, term Term) (ret Term, err error) {
	err = db.Where(&term).First(&ret).Error
	return ret, err
}

// UpdateTerm - Used to update an entry in the "term" table
func UpdateTerm(db *gorm.DB, term Term) (err error) {
	return db.Model(&term).Update(map[string]interface{}{"term": term.Term, "explanation": term.Explanation}).Error
}

// DeleteTerm - Used to remove an entry grom the "term" table
func DeleteTerm(db *gorm.DB, term Term) (err error) {
	return db.Delete(&term).Error
}
