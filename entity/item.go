package entity

import "time"

type Item struct {
	ID        int        `gorm:"primary_key" json:"id"`
	Name      string     `json:"name"`
	Category  string     `json:"category"`
	CreatedAt time.Time  `json:"created_at"`
}