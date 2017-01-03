package pictionary

// User - Struct used to hold data about a single user
type User struct {
	ID       uint32    `gorm:"column:id;not null;unique;AUTO_INCREMENT"`
	name     string    `gorm:"column:name;type:varchar(45);not null"`
	surname  string    `gorm:"column:surname;type:varchar(45);not null"`
	email    string    `gorm:"column:email;type:varchar(255);not null;unique"`
	google   bool      `gorm:"column:google"`
	drawings []Drawing `gorm:"ForeignKey:id"`
	games    []Game    `gorm:"many2many:user_played_game"`
}
