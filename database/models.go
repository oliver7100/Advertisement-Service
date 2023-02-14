package database

type Advertisement struct {
	ID        int `gorm:"primaryKey;"`
	Image     string
	Email     string
	Activated bool `gorm:"default:false;"`
}
