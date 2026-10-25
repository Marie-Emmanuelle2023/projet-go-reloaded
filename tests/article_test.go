package tests

import (
	"go-reloaded/rules"
	"testing"
)

func TestArticle(t *testing.T) {
	cas := []struct {
		mots    string
		attendu string
	}{
		{"a hoster", "an hoster"},
		{"A eagle", "An eagle"},
		{"a obstruction", "an obstruction"},
		{"An article", "An article"},
	}
	for _, c := range cas {
		obtenu := rules.Traitement(c.mots)
		if obtenu != c.attendu {
			t.Errorf("Process (%q) = %q, attendu %q", c.mots, obtenu, c.attendu)
		}

	}
}
