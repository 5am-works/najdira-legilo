package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"najdira_legilo/agoj"
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
	menuo: for {
		fmt.Printf("Kion vi volas fari? (krei, trovi, listi)\n")
		eniro, _ := eniroLegilo.ReadString('\n')
		switch strings.TrimSpace(eniro) {
		case "krei":
			agoj.Krei(db)
		case "listi":
			agoj.Listi(db)
		case "trovi":
			agoj.Trovi(db)
		case "eliri":
			break menuo
		default:
			fmt.Printf("Mi ne komprenis\n")
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
