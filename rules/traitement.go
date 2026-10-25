package rules

import "strings"

func Traitement(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		lignes[i] = traitementLigne(ligne)
	}
	return strings.Join(lignes, "\n")
}

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
			if i+1 < len(mots) && applyCaseN(resultat, mot, mots[i+1]) {
				i++
			} else {
				resultat = append(resultat, mot)
			}
		default:
			resultat = append(resultat, mot)
		}
	}
	resultat = FixApostrophes(resultat)
	resultat = FixPonctuation(resultat)
	FixArticles(resultat)
	return strings.Join(resultat, " ")
}
