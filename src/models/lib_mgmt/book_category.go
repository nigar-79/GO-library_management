package models

//BookCategory contains details about book category
type BookCategory struct {
	CategoryID   int    `gorm:"primary_key;size=50" json:"category_id"`
	CategoryName string `gorm:"size=100" json:"category_name"`
}

//BookCategoryType contains detail about book category
func (BookCategory) BookCategoryType() string {
	return "lib_mgmt.book_category"
}
