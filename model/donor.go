package model

import "time"

type Donor struct {
	ID int `json:"id"`
	// UserID             string    `json:"-"`
	FullName           string    `json:"full_name"`
	NickName           string    `json:"nick_name"`
	Email              string    `json:"email"`
	PhoneNumber        string    `json:"phone_number"`
	Faculty            string    `json:"faculty"`
	Major              string    `json:"major"`
	ClassYear          string    `json:"class_year"`
	Activities         *string   `json:"activities"`
	CompanyName        *string   `json:"company_name"`
	Position           *string   `json:"position"`
	Domicile           string    `json:"domicile"`
	DonationPermonth   float32   `json:"donation_permonth"`
	DonationType       string    `json:"donation_type"`
	DateReminder       int       `json:"date_reminder"`
	InterestForSpeaker bool      `json:"interest_for_speaker"`
	Motivate           string    `json:"motivate"`
	Suggestion         string    `json:"suggestion"`
	Period             int       `json:"period"`
	CreatedAt          time.Time `json:"created_at"`
	CreatedBy          string    `json:"created_by"`
	UpdatedAt          time.Time `json:"updated_at"`
	UpdatedBy          string    `json:"updated_by"`
}
