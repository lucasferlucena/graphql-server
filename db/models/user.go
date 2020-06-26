package models

//User model
type User struct {
	ID   int    `gorm:"primary_key;auto_increment"`
	Name string `gorm:"size:255;not null"`
}
