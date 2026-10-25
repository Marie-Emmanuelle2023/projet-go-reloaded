package rules

import "strings"

// Traitement corrige le texte ligne par ligne
func Traitement(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		lignes[i] = traitementLigne(ligne)
	}
	return strings.Join(lignes, "\n")
}

// traitementLigne applique les balises puis les corrections de mise en forme
func traitementLigne(ligne string) string {
	mots := strings.Fields(ligne)
	var resultat []string
	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		switch mot {
		case "(hex)":
			ConvertLast(resultat, 16)
		case "(bin)":
			ConvertLast(resultat, 2)
		case "(up)", "(low)", "(cap)":
			applyCase(resultat, mot[1:len(mot)-1], 1)
		case "(up,", "(low,", "(cap,":
			// Le nombre est dans le mot suivant, ex : "2)"
			if i+1 < len(mots) && applyCaseN(resultat, mot, mots[i+1]) {
				i++
			} else {
				resultat = append(resultat, mot)
			}
		default:
			resultat = append(resultat, mot)
		}
	}
	// L'ordre est important : apostrophes, ponctuation, articles
	resultat = FixApostrophes(resultat)
	resultat = FixPonctuation(resultat)
	FixArticles(resultat)
	return strings.Join(resultat, " ")
}
