package iloj

import (
	"fmt"
)

func Dividi(vorto string) ([]string, error) {
	if len(vorto) % 2 == 1 || len(vorto) == 0 {
		return nil, fmt.Errorf(
			"La vorto \"%v\" havas nevalidan langecon: %v",
			vorto,
			len(vorto),
		)
	}

	silaboj := make([]string, len(vorto) / 2)
	for i, _ := range vorto {
		if i % 2 == 1 {
			continue
		}

		silaboj[i / 2] = vorto[i:i+2]
	}

	return silaboj, nil
}