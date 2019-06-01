package modeloj

import (
	"github.com/jinzhu/gorm"
	"najdira_legilo/iloj"
	"najdira_legilo/gramatiko"

	"fmt"
)

type Vorto struct {
	gorm.Model
	Vorto                string `gorm:"UNIQUE"`
	SignifoID            uint `gorm:"NOT NULL"`
	Signifo              Signifo
	Ecoj uint8 `gorm:"NOT NULL";`
	Radikoj              []*Vorto `gorm:"many2many:radikoj;association_jointable_foreignkey:radiko_id"`
}

func (Vorto) TableName() string {
	return "vortoj"
}

func (v *Vorto) Montri() {
	silaboj, _ := iloj.Dividi(v.Vorto)
	fmt.Printf(
		"%v (%v, %v): %v\n",
		iloj.Normaligi(v.Vorto),
		len(silaboj),
		iloj.Tipo(silaboj[len(silaboj)-1]),
		v.Signifo.Signifo,
	)
	fmt.Printf("Prepozicioj: ")
	for i, p := range gramatiko.Prepozicioj {
		if (v.Ecoj >> uint(i)) & 1 == 1 {
			fmt.Printf("%v ", p)
		}
	}
	fmt.Printf("\n")
	fmt.Printf("Radikoj: ")
	for _, r := range v.Radikoj {
		fmt.Printf("%v ", iloj.Normaligi(r.Vorto))
	}
	fmt.Printf("\n")
}