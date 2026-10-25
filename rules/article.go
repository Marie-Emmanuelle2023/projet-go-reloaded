package rules

import "strings"

// FixArticles remplace "a" par "an" devant une voyelle ou un h
func FixArticles(mots []string) {
	for i := 0; i+1 < len(mots); i++ {
		if (mots[i] == "a" || mots[i] == "A") && startsWithVowel(mots[i+1]) {
			mots[i] = mots[i] + "n"
		}
	}
}

// startsWithVowel vérifie si le mot commence par une voyelle ou un h
func startsWithVowel(mot string) bool {
	if mot == "" {
		return false
	}
	return strings.ContainsRune("aeiouhAEIOUH", rune(mot[0]))
}
