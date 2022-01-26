package models

import "time"

type Visit struct {
	ID              uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	Date            time.Time `json:"date"`
	Patient_id      uint      `json:"patient_id"`
	Problems        string    `json:"problems"`
	Diagnosis       string    `json:"diagnosis"`
	Prescription_id uint      `json:"prescription_id"`
}
