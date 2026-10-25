package rules

import (
	"strconv"
	"strings"
	"unicode"
)

// Capitalize met la première lettre en majuscule et le reste en minuscules
func Capitalize(mot string) string {
	if mot == "" {
		return mot
	}
	lettres := []rune(strings.ToLower(mot))
	lettres[0] = unicode.ToUpper(lettres[0])
	return string(lettres)
}

// applyCase applique up, low ou cap aux n derniers mots
func applyCase(mots []string, mode string, n int) {
	for i := len(mots) - 1; i >= 0 && n > 0; i-- {
		// La ponctuation seule ne compte pas comme un mot
		if countPonctuation(mots[i]) == len(mots[i]) {
			continue
		}
		switch mode {
		case "up":
			mots[i] = strings.ToUpper(mots[i])
		case "low":
			mots[i] = strings.ToLower(mots[i])
		case "cap":
			mots[i] = Capitalize(mots[i])
		}
		n--
	}
}

// applyCaseN gère les balises avec nombre, ex : "(up," + "2)"
// Renvoie false si la balise est invalide
func applyCaseN(mots []string, tag string, count string) bool {
	if !strings.HasSuffix(count, ")") {
		return false
	}
	n, err := strconv.Atoi(strings.TrimSuffix(count, ")"))
	if err != nil || n < 0 {
		return false
	}
	applyCase(mots, tag[1:len(tag)-1], n)
	return true
}
