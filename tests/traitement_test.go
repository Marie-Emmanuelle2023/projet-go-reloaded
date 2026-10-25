package tests

import (
	"go-reloaded/rules"
	"testing"
)

func TestProcess(t *testing.T) {
	cas := []struct {
		entree  string
		attendu string
	}{
		{"Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			"Simply add 66 and 2 and you will see the result is 68."},
		{"There is no greater agony than bearing a untold story inside you.",
			"There is no greater agony than bearing an untold story inside you."},
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"hi (cap, 6)", "Hi"},
		{"ligne une\nligne   deux", "ligne une\nligne deux"},
	}
	for _, c := range cas {
		obtenu := rules.Traitement(c.entree)
		if obtenu != c.attendu {
			t.Errorf("\nentrée  : %q\nobtenu  : %q\nattendu : %q", c.entree, obtenu, c.attendu)
		}
	}
}
