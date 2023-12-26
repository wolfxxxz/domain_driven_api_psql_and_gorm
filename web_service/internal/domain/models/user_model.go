package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        *uuid.UUID `json:"id" gorm:"primaryKey"`
	Email     string     `json:"user_email"`
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Password  string     `json:"password"`
	Role      string     `json:"role"`
}
