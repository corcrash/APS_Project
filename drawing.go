package main

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
)

// Drawing - A struct used to hold the data for a single Drawing
type Drawing struct {
	ID          uint32 `gorm:"column:id;not null;unique;AUTO_INCREMENT"`
	RoundNumber uint32 `gorm:"column:round_number;not null"`
	Data        string `gorm:"column:data;type:mediumtext;not null"`
	Answer      string `gorm:"column:answer;type:varchar(255)"`
	Term        Term   `gorm:"column:term_id;ForeignKey:id;not null"`
	Game        Game   `gorm:"column:game_id;ForeignKey:id"`
	User        User   `gorm:"column:user_id;ForeignKey:id"`
}

// CreateDrawing - Used to create a new entry in the "drawing" table
func CreateDrawing(db *gorm.DB, drawing *Drawing) (err error) {
	if db.NewRecord(drawing) {
		if count := db.Save(&drawing).RowsAffected; count != 1 {
			log.Print("Not created: ", drawing)
			err = errors.New("Error creating drawing!")
		}
	} else {
		err = errors.New("Error checking for drawing!")
	}

	return err
}

// ReadDrawing - Used to read an entry from the "drawing" table
func ReadDrawing(db *gorm.DB, drawing Drawing) (ret Drawing, err error) {
	err = db.Where(&drawing).First(&ret).Error
	return ret, err
}

// UpdateDrawing - Used to update an entry in the "drawing" table
func UpdateDrawing(db *gorm.DB, drawing Drawing) (err error) {
	return db.Model(&drawing).Update(map[string]interface{}{"answer": drawing.Answer, "round_number": drawing.RoundNumber, "data": drawing.Data}).Error
}

// DeleteDrawing - Used to remove an entry grom the "drawing" table
func DeleteDrawing(db *gorm.DB, drawing Drawing) (err error) {
	return db.Delete(&drawing).Error
}
