package user

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Email     string    `gorm:"size:50;not null;unique" json:"email"`
	Password  string    `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `gorm:"" json:"created_at"`
	UpdatedAt time.Time `gorm:"" json:"updated_at"`
}
