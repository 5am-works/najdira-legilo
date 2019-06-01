package agoj

import (
	"github.com/jinzhu/gorm"
	"najdira_legilo/modeloj"

	"bufio"
	"fmt"
	"os"
	"strings"
)

func Trovi(db *gorm.DB) {
	eniroLegilo := bufio.NewReader(os.Stdin)
	fmt.Printf("Eniru la vorton: ")
	eniro, _ := eniroLegilo.ReadString('\n')
	vorto := strings.TrimSpace(eniro)
	var v modeloj.Vorto
	if db.Preload("Radikoj").Preload("Signifo").Where("vorto = ?", vorto).Find(&v).RecordNotFound() {
		fmt.Printf("La vorto %v ne ekzistas.\n", vorto)
	} else {
		fmt.Printf("Jen la vorto:\n")
		v.Montri()
	}
}
