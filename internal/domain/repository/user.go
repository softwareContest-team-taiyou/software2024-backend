package repository

type UserEntity struct {
	ID   string `gorm:"primaryKey"`
	Name string `gorm:"type:varchar(255)"`
}