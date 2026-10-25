package rules

import "strings"

// FixPonctuation colle la ponctuation au mot précédent
func FixPonctuation(mots []string) []string {
	var resultat []string
	for _, mot := range mots {
		n := countPonctuation(mot)
		if n > 0 && len(resultat) > 0 {
			dernier := len(resultat) - 1
			resultat[dernier] = resultat[dernier] + mot[:n]
			mot = mot[n:]
		}
		if mot != "" {
			resultat = append(resultat, mot)
		}
	}
	return resultat
}

// countPonctuation compte les signes de ponctuation au début du mot
func countPonctuation(mot string) int {
	n := 0
	for n < len(mot) && strings.ContainsRune(".,!?:;", rune(mot[n])) {
		n++
	}
	return n
}
