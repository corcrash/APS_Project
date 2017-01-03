package pictionary

// Term - A struct used to hold the data for a single Term
type Term struct {
	ID          uint32 `gorm:"column:id;not null;unique;AUTO_INCREMENT"`
	term        string `gorm:"column:term;type:varchar(255);not null"`
	explanation string `gorm:"column:explanation;type:varchar(255);not null"`
}
