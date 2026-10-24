package rules

import "strings"

func Traitement(texte string) string{
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes{
		lignes[i] = traitementLigne(ligne)
	}
	return strings.Join(lignes, "\n")
}

func traitementLigne(ligne string) string{
	mots := strings.Fields(ligne)
	var resultat []string
	for _,mot := range mots{
		switch mot{
		case "(hex)":
			ConvertLast(resultat,16)
		case "(bin)":
			ConvertLast(resultat,2)
		default:
			resultat = append(resultat, mot)
		}
	}
	return strings.Join(resultat, " ")
}