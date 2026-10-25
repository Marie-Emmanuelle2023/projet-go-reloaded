package tests

import (
	"go-reloaded/rules"
	"testing"
)

func TestCasse(t *testing.T) {
	cas := []struct {
		entree  string
		attendu string
	}{
		{"bonjour (up, 10)", "BONJOUR"},
		{"finalement (cap) , cela n'aura pas lieu", "Finalement, cela n'aura pas lieu"},
		{"Je LUI (low) ai transmis le message", "Je lui ai transmis le message"},
		{"Salut tu vas bien ? (up, 2)", "Salut tu VAS BIEN?"},
	}
	for _, c := range cas {
		obtenu := rules.Traitement(c.entree)
		if obtenu != c.attendu {
			t.Errorf("Traitement(%q) = %q, attendu %q", c.entree, obtenu, c.attendu)
		}
	}
}
