package modeloj

import (
	"github.com/jinzhu/gorm"
)

type Vorto struct {
	gorm.Model
	Vorto                string `gorm:"UNIQUE"`
	SignifoID            int
	Signifo              Signifo
	SurNomivnativo       bool
	SuperAkuzativo       bool
	SubDativo            bool
	ApudGenetivo         bool
	EnLokativo           bool
	AlLativo             bool
	KontraŭInstrumentalo bool
	Radikoj              []*Vorto `gorm:"MANY2MANY:radikoj,association_jointable_foreignkey:radiko_id`
}

func (Vorto) TableName() string {
	return "vortoj"
}
