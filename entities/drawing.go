package pictionary

// Drawing - A struct used to hold the data for a single Drawing
type Drawing struct {
	ID          uint32 `gorm:"column:id;not null;unique;AUTO_INCREMENT"`
	roundNumber uint32 `gorm:"column:round_number;not null"`
	data        string `gorm:"column:data;type:mediumtext;not null"`
	game        Game   `gorm:"ForeignKey:id;not null"`
	user        User   `gorm:"ForeignKey:id;not null"`
	answer      string `gorm:"column:answer;type:varchar(255)"`
	term        Term   `gorm:"ForeignKey:id;not null"`
}
