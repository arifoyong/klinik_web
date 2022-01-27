package models

import (
	"time"
)

type Patient struct {
	ID        uint       `json:"id"`
	Firstname *string    `json:"firstname"`
	Lastname  *string    `json:"lastname"`
	IC        *string    `json:"ic" `
	DOB       *time.Time `json:"dob" `
	Email     *string    `json:"email"`
	Phone     *string    `json:"phone"`
	Address   *string    `json:"address"`
}
