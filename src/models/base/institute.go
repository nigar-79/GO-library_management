package models

import (
	"time"

	"sdbgo.in/sdbgo/src/models"
)

// Institution details
type Institution struct {
	InstituteID             string    `gorm:"primary_key;size=20" json:"institute_id"`
	Institutename           string    `gorm:"size:100" json:"institute_name"`
	InstituteType           string    `gorm:"size:100" json:"institute_type"`
	ContactPersonSalutation string    `gorm:"size:10" json:"contact_person_salutation"`
	ContactPersonName       string    `gorm:"size:100" json:"contact_person_name"`
	ContactPersonEmailID    string    `gorm:"size:400" json:"contact_person_email_id"`
	ContactPersonPhoneNo    string    `gorm:"size:200" json:"contact_person_phone_no"`
	WebURL                  string    `gorm:"size:255" json:"web_url"`
	LogoPath                string    `gorm:"size:255" json:"logo_path"`
	AddressLine1            string    `gorm:"size:500" json:"address_line1"`
	AddressLine2            string    `gorm:"size:500" json:"address_line2"`
	AddressLine3            string    `gorm:"size:500" json:"address_line3"`
	City                    string    `gorm:"size:50" json:"city"`
	State                   string    `gorm:"size:50" json:"state"`
	Country                 string    `gorm:"size:50" json:"country"`
	CountryCode             string    `gorm:"size:50" json:"country_code"`
	PostalCode              string    `gorm:"size:50" json:"postal_code"`
	TelNo                   string    `gorm:"size:50" json:"tel_no"`
	FaxNo                   string    `gorm:"size:50" json:"fax_no"`
	InstituteStartDate      time.Time `gorm:"size:20" json:"institute_start_date"`
	AreaInAcres             float64   `gorm:"size:200" json:"area_in_acres"`
	Status                  string    `gorm:"size:20" json:"status"`
	models.Base
}

// TableName TABLE NAME
func (Institution) TableName() string {
	return "base.institute"
}
