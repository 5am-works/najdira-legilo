package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"fmt"
	"os"
)

type Signifo struct {
	gorm.Model
	Signifo string `gorm:"UNIQUE"`
}

func (Signifo) TableName() string {
	return "signifoj"
}

func main() {
	db, err := gorm.Open(
		"postgres",
		os.Args[1],
	)

	if err != nil {
		fmt.Printf(err.Error())
		panic("Ne eblas konekti al la datumejo")
	}

	defer db.Close()

	db.AutoMigrate(&Signifo{})

	var signifo Signifo
	db.Find(&signifo, 1)

	fmt.Printf("%+v", signifo)
}
