package entity

import "time"

type User struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	FirstName string `gorm:"not null" json:"first_name" validate:"required"`
	LastName  string `gorm:"not null" json:"last_name" validate:"required"`
	Email     string `gorm:"not null;unique" json:"email" validate:"required,email"`
	Password  string `gorm:"not null" json:"password" validate:"required,min=8"`

	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"default:NULL"`
}
