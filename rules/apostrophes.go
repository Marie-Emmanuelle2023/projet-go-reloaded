package rules

// FixApostrophes colle chaque paire d'apostrophes aux mots qu'elle entoure
func FixApostrophes(mots []string) []string {
	var resultat []string
	ouverte := false
	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		if mot == "'" {
			if !ouverte && i+1 < len(mots) {
				mots[i+1] = "'" + mots[i+1]
				ouverte = true
				continue
			}
			if ouverte && len(resultat) > 0 {
				resultat[len(resultat)-1] += "'"
				ouverte = false
				continue
			}
		}
		resultat = append(resultat, mot)
	}
	return resultat
}
