package util

import "strings"

func TrToEn(name string) string {
	trChars := []string{"ı", "ğ", "İ", "Ğ", "ç", "Ç", "ş", "Ş", "ö", "Ö", "ü", "Ü"}
	enChars := []string{"i", "g", "I", "G", "c", "C", "s", "S", "o", "O", "u", "U"}

	for i, toreplace := range trChars {
		r := strings.NewReplacer(toreplace, enChars[i])
		name = r.Replace(name)
	}

	return name
}
