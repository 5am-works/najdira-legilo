package agoj

import (
	"github.com/jinzhu/gorm"
	"najdira_legilo/modeloj"

	"bufio"
	"fmt"
	"os"
	"strings"
)

func Listi(db *gorm.DB) {
	eniroLegilo := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Kion vi volas listi? (signifojn)\n")
		eniro, _ := eniroLegilo.ReadString('\n')
		switch strings.TrimSpace(eniro) {
		case "signifojn":
			var signifoj []*modeloj.Signifo
			db.Find(&signifoj)
			for _, s := range signifoj {
				fmt.Printf("%v %v\n", s.ID, s.Signifo)
			}
		case "eliri":
			return
		default:
			fmt.Printf("Mi ne komprenis\n")
		}
	}
}
