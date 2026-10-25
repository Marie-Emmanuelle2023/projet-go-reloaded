package tests

import (
	"go-reloaded/rules"
	"testing"
)

func TestPonctuation(t *testing.T) {
	cas := []struct {
		entree  string
		attendu string
	}{
		{"I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{"I was thinking ... You were right", "I was thinking... You were right"},
		{"Ready, set, go (up) !", "Ready, set, GO!"},
		{"what do you think ?", "what do you think?"},
		{"hello ;world", "hello; world"},
	}
	for _, c := range cas {
		obtenu := rules.Traitement(c.entree)
		if obtenu != c.attendu {
			t.Errorf("Process(%q) = %q, attendu %q", c.entree, obtenu, c.attendu)
		}
	}
}
