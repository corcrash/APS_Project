package pictionary

import (
	"time"
)

// Game - A struct used to hold data about a single game
type Game struct {
	ID        uint32    `gorm:"column:id;not null;unique;AUTO_INCREMENT"`
	dateTime  time.Time `gorm:"column:date_time;type:datetime;not null"`
	chatData  string    `gorm:"column:chat_data;type:mediumtext"`
	drawings  []Drawing `gorm:"ForeignKey:id"`
	team      uint8     `gorm:"table:user_played_game;column:team"`
	userScore uint32    `gorm:"table:user_played_game;column:user_score"`
}
