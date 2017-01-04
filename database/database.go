package database

import "github.com/jinzhu/gorm"

var (
	// GlobalDB - Used to hold a global db variable if none is passed
	GlobalDB *gorm.DB
)

// CRUD - Interface used to apstract common operations
type CRUD interface {
	Create(db *gorm.DB) error
	Read(db *gorm.DB) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
}

// Create - Function used to insert a new record into the database
func Create(model CRUD) error {
	return model.Create(GlobalDB)
}

// Read - Function used to read a record from the database
func Read(model CRUD) error {
	return model.Read(GlobalDB)
}

// Update - Function used to update a record in the database
func Update(model CRUD) error {
	return model.Update(GlobalDB)
}

// Delete - Function used to remove a record from the database
func Delete(model CRUD) error {
	return model.Delete(GlobalDB)
}
