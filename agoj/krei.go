package agoj

import (
	"github.com/jinzhu/gorm"
	"najdira_legilo/gramatiko"
	"najdira_legilo/iloj"
	"najdira_legilo/modeloj"

	"bufio"
	"fmt"
	"os"
	"strings"
)

func Krei(db *gorm.DB) {
	eniroLegilo := bufio.NewReader(os.Stdin)
	fmt.Printf("Kio estas la signifo?\n")
	signifo, _ := eniroLegilo.ReadString('\n')
	signifo = strings.TrimSpace(signifo)
	var s modeloj.Signifo
	var ekzistantaVorto modeloj.Vorto
	ekzistanta := false
	if db.Where("signifo = ?", signifo).Find(&s).RecordNotFound() {
		fmt.Printf("La signifo \"%v\" ne ekzistas. Vi kreos novan vorton: ", signifo)
		s = modeloj.Signifo{Signifo: signifo}
		db.Create(&s)
	} else {
		fmt.Printf("La signifo \"%v\" jam ekzistas. ", signifo)
		if db.Where("signifo_id = ?", s.ID).First(&ekzistantaVorto).RecordNotFound() {
			fmt.Printf("Vi kreos la unuan vorton por ĉi tiu signifo: ")
		} else {
			ekzistanta = true
			fmt.Printf("Vi kreos egalvorton de %v: ", ekzistantaVorto.Vorto)
		}
	}

	eniro, _ := eniroLegilo.ReadString('\n')
	novaVorto := strings.TrimSpace(eniro)
	silaboj, err := iloj.Dividi(novaVorto)
	if err != nil {
		fmt.Printf("La vorto havas eraron: %v\n", err.Error())
		return
	}

	demandiRadikojn := func() []*modeloj.Vorto {
		radikoj := []*modeloj.Vorto{}
		fmt.Printf("Radikoj de %v: ", novaVorto)
		eniro, _ := eniroLegilo.ReadString('\n')
		for strings.TrimSpace(eniro) != "" {
			radiko := strings.TrimSpace(eniro)
			var r modeloj.Vorto
			if db.Where("vorto = ?", radiko).First(&r).RecordNotFound() {
				fmt.Printf("La vorto %v ne ekzistas. Reprevu: ", radiko)
			} else {
				radikoj = append(radikoj, &r)
			}
			eniro, _ = eniroLegilo.ReadString('\n')
		}

		return radikoj
	}

	finaĵo := silaboj[len(silaboj)-1]

	if !ekzistanta {
		if _, ok := gramatiko.ValidajFinaĵoj[finaĵo]; ok {
			tipo := iloj.Tipo(finaĵo)
			if k, _ := iloj.Demandi(fmt.Sprintf("La vorto estas %v, kiu finas per %v. Krei la vorton? (j/_)\n", tipo, finaĵo)); k {
				radikoj := demandiRadikojn()
				var opcioj []string
				switch tipo {
				case "substantivo":
					opcioj = gramatiko.Prepozicioj
				case "verbo":
					opcioj = gramatiko.Kazoj
				}
				var elektoj uint8 = 0
				for i, o := range opcioj {
					if k2, _ := iloj.Demandi(fmt.Sprintf("%v? (j/_)", o)); k2 {
						elektoj = elektoj | (1 << uint(i))
					}
				}

				vorto := modeloj.Vorto{
					Vorto:     novaVorto,
					SignifoID: s.ID,
					Signifo:   s,
					Ecoj:      elektoj,
					Radikoj:   radikoj,
				}

				fmt.Printf("Vi kreos: ")
				vorto.Montri()

				if k2, _ := iloj.Demandi("Konfirmi?"); k2 {
					db.Create(&vorto)
				}
			}
		} else {
			fmt.Printf("La finaĵo %v ne estas valida.", silaboj[len(silaboj)-1])
		}
	} else {
		tipo := iloj.Tipo(silaboj[len(silaboj)-1])
		if tipo == "???" {
			fmt.Printf("\"%v\" ne havas validan finaĵon.\n", novaVorto)
			return
		}
		v := ekzistantaVorto.Vorto
		ekzistantaFinaĵo := v[len(v)-2:]
		ekzistantaTipo := iloj.Tipo(ekzistantaFinaĵo)
		if tipo == ekzistantaTipo {
			radikoj := demandiRadikojn()
			vorto := modeloj.Vorto{
				Vorto:     novaVorto,
				SignifoID: s.ID,
				Signifo:   s,
				Ecoj:      ekzistantaVorto.Ecoj,
				Radikoj:   radikoj,
			}
			fmt.Printf("Vi kreos: ")
			vorto.Montri()
			fmt.Printf("Konfirmi? (j/_) ")
			eniro, _ = eniroLegilo.ReadString('\n')
			if strings.TrimSpace(eniro) == "j" {
				db.Create(&vorto)
			} else {
				fmt.Printf("Nuligis.\n")
			}
		} else {
			fmt.Printf(
				"La tipo de %v (%v) ne estas tiu de %v (%v).\n",
				novaVorto,
				tipo,
				ekzistantaVorto,
				ekzistantaTipo,
			)
		}
	}
}
