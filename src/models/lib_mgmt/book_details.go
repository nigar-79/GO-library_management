package models

type BookDetails struct {
	BookID               int    `gorm:"primary_key;size=50" json:"book_id"`
	ISBN                 string `gorm:"size=20" json:"isbn"`
	BookTitle            string `gorm:"size=20" json:"book_title"`
	Language             string `gorm:"size=50" json:"language_name"`
	TotalNoOfCopies      int    `gorm:"size=20" json:"total_no_of_copies"`
	CategoryID           int    `gorm:"References:CategoryID;size=50" json:"category_id"`
	BindingPublicationID int    `gorm:"References:PublicationID;size=50" json:"binding_publication_id"`
	ShelfID              int    `gorm:"References:shelf_id;size=50" json:"shelf_id"`
}
