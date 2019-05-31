package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"najdira_legilo/modeloj"

	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	db := prepari()
	defer db.Close()

	eniroLegilo := bufio.NewReader(os.Stdin)
	fmt.Printf("Kion vi volas fari? (krei, redakti, listi)\n")
	eniro, _ := eniroLegilo.ReadString('\n')
	switch strings.TrimSpace(eniro) {
	case "krei":
		fmt.Printf("Kio estas la signifo?\n")
		signifo, _ := eniroLegilo.ReadString('\n')
		signifo = strings.TrimSpace(signifo)
		var s modeloj.Signifo
		if db.Where("signifo = ?", signifo).Find(&s).RecordNotFound() {
			fmt.Printf("La signifo \"%v\" ne ekzistas. Krei ĝin?", signifo)
			eniro, _ := eniroLegilo.ReadString('\n')
			if strings.TrimSpace(eniro) == "j" {
				db.Create(&modeloj.Signifo{Signifo: signifo})
			}
		} else {
			fmt.Printf("La signifo \"%v\" jam ekzistas.", signifo)
		}
	case "listi":
		fmt.Printf("Kion vi volas listi? (signifojn)\n")
		eniro, _ := eniroLegilo.ReadString('\n')
		if strings.TrimSpace(eniro) == "signifojn" {
			var signifoj []*modeloj.Signifo
			db.Find(&signifoj)
			for _, s := range signifoj {
				fmt.Printf("%v %v\n", s.ID, s.Signifo)
			}
		}
	}
}

func prepari() *gorm.DB {
	fmt.Printf("Konektas al la datumejo... ")

	db, err := gorm.Open(
		"postgres",
		os.Args[1],
	)

	if err != nil {
		fmt.Printf("%v\n", err.Error())
		panic("Ne eblas konekti al la datumejo")
	}

	fmt.Printf("Finita\nKontrolas la version de la datumejo... ")

	db.AutoMigrate(&modeloj.Signifo{})
	db.AutoMigrate(&modeloj.Vorto{})

	fmt.Printf("Finita\n")

	return db
}
