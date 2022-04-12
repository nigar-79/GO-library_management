package models

import "time"

// Institution details
type Institution struct {
	InstituteID             string    `json:"institute_id"`
	Institutename           string    `json:"institute_name"`
	InstituteType           string    `json:"institute_type"`
	ContactPersonSalutation string    `json:"contact_person_salutation"`
	ContactPersonName       string    `json:"contact_person_name"`
	ContactPersonEmailID    string    `json:"contact_person_email_id"`
	ContactPersonPhoneNo    string    `json:"contact_person_phone_no"`
	WebURL                  string    `json:"web_url"`
	LogoPath                string    `json:"logo_path"`
	AddressLine1            string    `json:"address_line1"`
	AddressLine2            string    `json:"address_line2"`
	AddressLine3            string    `json:"address_line3"`
	City                    string    `json:"city"`
	State                   string    `json:"state"`
	Country                 string    `json:"country"`
	CountryCode             string    `json:"country_code"`
	PostalCode              string    `json:"postal_code"`
	TelNo                   string    `json:"tel_no"`
	FaxNo                   string    `json:"fax_no"`
	InstituteStartDate      time.Time `json:"institute_start_date"`
	AreaInAcres             float64   `json:"area_in_acres"`
	Status                  string    `json:"status"`
	Base                    `json:"base"`
}

// TableName TABLE NAME
func (Institution) TableName() string {
	return "base.institute"
}
