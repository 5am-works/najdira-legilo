package iloj

import (
	"najdira_legilo/gramatiko"
)

func Tipo(finaĵo string) string {
	if _, ok := gramatiko.S1Finaĵoj[finaĵo]; ok {
		return "substantivo"
	} else if _, ok := gramatiko.V0Finaĵoj[finaĵo]; ok {
		return "verbo"
	}

	return "???"
}