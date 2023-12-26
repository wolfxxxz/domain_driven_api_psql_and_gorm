package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Lecture struct {
	gorm.Model
	ID          *uuid.UUID `json:"id" gorm:"primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Speaker     string     `json:"speaker"`
	Date        time.Time  `json:"date"`
	Location    string     `json:"location"`
	Duration    int        `json:"duration"`
	Students    []*User    `gorm:"many2many:lecture_students;" json:"lecture_students"`
}
