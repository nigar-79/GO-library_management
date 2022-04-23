package models

import (
	"time"

	models "sdbgo.in/sdbgo/src/models/base"
)

//BookBindingPublication book binding and publication details
type BookBindingPublication struct {
	PublicationID   int       `gorm:"primary_key;size=50" json:"publication_id"`
	PublisherName   string    `gorm:"size=50" json:"publisher_name"`
	PublicationYear string    `gorm:"size=10" json:"publication_year"`
	PublishedDate   time.Time `gorm:"size=50" json:"published_date"`
	CopyrightYear   string    `gorm:"size=10" json:"copyright_year"`
	Edition         string    `gorm:"size=10" json:"edition"`
	BindingName     string    `gorm:"size=200" json:"binding_name"`
	BindingType     string    `gorm:"size=50" json:"binding_type"`
	models.Base
}

// BookBinding binding details
func (BookBindingPublication) BookBinding() string {
	return "lib_mgmt.book_binding_publication"
}

// BookPublication publication details
func (BookBindingPublication) BookPublication() string {
	return "lib_mgmt.book_binding_publication"
}
