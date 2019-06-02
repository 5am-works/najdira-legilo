package iloj

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Demandi(demando string) (bool, error) {
	fmt.Println(demando)
	eniroLegilo := bufio.NewReader(os.Stdin)
	eniro, err := eniroLegilo.ReadString('\n')
	if err != nil {
		return false, err
	}

	if strings.TrimSpace(eniro) == "j" {
		return true, nil
	} else {
		return false, nil
	}
}