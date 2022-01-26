package models

import "time"

type Patient struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	IC        string    `json:"ic" gorm:"not null"`
	DOB       time.Time `json:"dob" gorm:"not null"`
	Email     string    `json:"email" gorm:"not null"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
}
