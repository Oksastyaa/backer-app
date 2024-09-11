package user

import (
	"time"
)

type User struct {
	ID             int       `gorm:"primaryKey"`
	Name           string    `gorm:"type:varchar(100)"`
	Occupation     string    `gorm:"type:varchar(100)"`
	Email          string    `gorm:"type:varchar(100);uniqueIndex"`
	PasswordHash   string    `gorm:"type:varchar(100)"`
	AvatarFileName string    `gorm:"type:varchar(255)"`
	Role           string    `gorm:"type:varchar(100)"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`
}
