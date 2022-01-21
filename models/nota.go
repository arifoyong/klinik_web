package models

import "time"

type Nota struct {
	ID             uint      `json:"id" gorm:"primary_key"`
	Date           time.Time `json:"date"`
	Nota_id        uint      `json:"nota_id"`
	Drug_name      string    `json:"drug_name"`
	Drug_type      string    `json:"drug_type"`
	Box_qty        uint      `json:"box_qty"`
	Vol_box        uint      `json:"vol_box"`
	Vol_unit       uint      `json:"vol_unit"`
	Box_cost       float32   `json:"box_cost"`
	Disc_percent   float32   `json:"disc_percent"`
	Disc_val       float32   `json:"disc_value"`
	Spdisc_percent float32   `json:"spdisc_percent"`
	Spdisc_val     float32   `json:"spdisc_value"`
	Ttl_cost_disc  float32   `json:"ttl_cost_disc"`
	Tax            float32   `json:"tax"`
	Ttl_cost       float32   `json:"ttl_cost"`
	Unit_cost      float32   `json:"unit_cost"`
	Vendor         string    `json:"vendor"`
}
