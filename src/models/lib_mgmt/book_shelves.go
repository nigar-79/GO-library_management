package models

import "sdbgo.in/sdbgo/src/models"

//BookShelves contails information about floor, row and shelf number
type BookShelves struct {
	ShelfId int `gorm:"primary_key;size=50" json:"shelf_id"`
	FloorNo int `gorm:"size=10" json:"floor_no"`
	RowNo   int `gorm:"size=10" json:"row_no"`
	ShelfNo int `gorm:"unique_key;size:10" json:"shelf_no"`
	models.Base
}

//ShelfDetails information about floor, row and shelf number
func (BookShelves) ShelfDetails() string {
	return "lib_mgmt.book_shelves"
}
