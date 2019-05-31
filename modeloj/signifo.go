package modeloj

import (
	"github.com/jinzhu/gorm"
)

type Signifo struct {
	gorm.Model
	Signifo string `gorm:"UNIQUE"`
}

func (Signifo) TableName() string {
	return "signifoj"
}
