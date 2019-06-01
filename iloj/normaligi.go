package iloj

import "strings"

func Normaligi(vorto string) string {
	s1 := strings.ReplaceAll(vorto, "_", "")
	s2 := strings.ReplaceAll(s1, "A", "ai")
	s3 := strings.ReplaceAll(s2, "E", "ei")
	return strings.ReplaceAll(s3, "O", "ou")
}