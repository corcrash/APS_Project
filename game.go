package main

import (
	"errors"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

// Game - A struct used to hold data about a single game
type Game struct {
	ID         uint32            `gorm:"column:id;not null;unique;AUTO_INCREMENT"`
	DateTime   time.Time         `gorm:"column:date_time;type:datetime;not null"`
	ChatData   string            `gorm:"column:chat_data;type:mediumtext"`
	UserScores map[string]uint32 `gorm:"-"`
}

// CreateGame - Used to create a new entry in the "games" table
func CreateGame(db *gorm.DB, game *Game) (err error) {
	if db.NewRecord(game) {
		if count := db.Save(&game).RowsAffected; count != 1 {
			log.Print("Not created: ", game)
			err = errors.New("Error creating game!")
		}
	} else {
		err = errors.New("Error checking for game!")
	}

	return err
}

// ReadGame - Used to read an entry from the "games" table
func ReadGame(db *gorm.DB, game Game) (ret Game, err error) {
	err = db.Where(&game).First(&ret).Error
	return ret, err
}

// UpdateGame - Used to update an entry in the "games" table
func UpdateGame(db *gorm.DB, game Game) (err error) {
	err = db.Model(&game).Update(map[string]interface{}{"date_time": game.DateTime, "chat_data": game.ChatData}).Error

	// for _, drawing := range game.Drawings {
	// 	drawing = {}
	// 	// TODO: Add drawing update call!
	// }

	return err
}

// DeleteGame - Used to remove an entry from the "games" table
func DeleteGame(db *gorm.DB, game Game) (err error) {
	return db.Delete(&game).Error
}
