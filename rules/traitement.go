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
	return strings.Join(mots, " ")
}