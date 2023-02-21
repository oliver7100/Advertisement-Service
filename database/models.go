package database

type Advertisement struct {
	ID          int     `gorm:"primaryKey;" json:"id"`
	Image       string  `json:"image"`
	Email       *string `gorm:"not null;type:varchar(255);" json:"email"`
	Activated   bool    `gorm:"not null;default:false;" json:"activated"`
	Description *string `gorm:"not null;" json:"description"`
}
