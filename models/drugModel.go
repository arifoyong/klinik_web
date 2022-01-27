package models

type Drug struct {
	Drug_id    uint     `json:"drug_id"`
	Name       *string  `json:"name"`
	Unit_price *float64 `json:"unit_price"`
}
