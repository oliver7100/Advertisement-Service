package database

type Advertisement struct {
	ID          int    `gorm:"primaryKey;" json:"id"`
	Image       string `json:"image"`
	Email       string `json:"email"`
	Activated   bool   `gorm:"default:false;" json:"activated"`
	Description string `json:"description"`
}
