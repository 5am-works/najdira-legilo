package iloj

import (
	"testing"
	"reflect"
)

func TestDividi(t *testing.T) {
	vortoj := []string{
		"luma",
		"limeri",
		"lirana",
		"silimana",
	}
	silaboj := [][]string{
		[]string{"lu", "ma"},
		[]string{"li", "me", "ri"},
		[]string{"li", "ra", "na"},
		[]string{"si", "li", "ma", "na"},
	}

	for i, v := range vortoj {
		s, err := Dividi(v)
		if err != nil {
			t.Errorf("Eraro okazis por %v: %v\n", v, err)
		}
		if !reflect.DeepEqual(s, silaboj[i]) {
			t.Errorf("La rezulto por %v, %v, ne estas %v.\n", v, s, silaboj[i])
		}
	}
}