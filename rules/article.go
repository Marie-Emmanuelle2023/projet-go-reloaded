package rules

import "strings"

func FixArticles(mots []string) {
	for i := 0; i+1 < len(mots); i++ {
		if (mots[i] == "a" || mots[i] == "A") && startsWithVowel(mots[i+1]) {
			mots[i] = mots[i] + "n"
		}
	}
}

func startsWithVowel(mot string) bool {
	if mot == "" {
		return false
	}
	return strings.ContainsRune("aeiouhAEIOUH", rune(mot[0]))
}
