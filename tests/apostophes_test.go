package tests

import (
	"go-reloaded/rules"
	"testing"
)

func TestApostrophes(t *testing.T) {
	cas := []struct {
		entree  string
		attendu string
	}{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '",
			"As Elton John said: 'I am the most well-known homosexual in the world'"},
		{"he said ' hello ' , then left", "he said 'hello', then left"},
		{"don't touch", "don't touch"},
	}
	for _, c := range cas {
		obtenu := rules.Traitement(c.entree)
		if obtenu != c.attendu {
			t.Errorf("Process(%q) = %q, attendu %q", c.entree, obtenu, c.attendu)
		}
	}
}
