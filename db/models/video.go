package models

//Video model
type Video struct {
	ID     int    `gorm:"primary_key;auto_increment"`
	Title  string `gorm:"size:255;not null"`
	URL    string `gorm:"size:255;not null"`
	Author User   `gorm:"foreignkey:UserID"`
}
