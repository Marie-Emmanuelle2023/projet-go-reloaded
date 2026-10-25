package tests

import (
	"go-reloaded/rules"
	"testing"
)

func TestConvertBase(t *testing.T) {
	cas := []struct {
		mot     string
		base    int
		attendu string
	}{
		{"1E", 16, "30"},
		{"ff", 16, "255"},
		{"10", 2, "2"},
		{"bonjour", 16, "bonjour"},
	}
	for _, c := range cas {
		obtenu := rules.ConvertBase(c.mot, c.base)
		if obtenu != c.attendu {
			t.Errorf("ConvertBase(%q, %d) = %q, attendu %q", c.mot, c.base, obtenu, c.attendu)
		}
	}
}
