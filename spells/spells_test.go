package spells

import (
	"testing"
)

func TestFindSpell(t *testing.T) {
	s, err := Find("Accio")
	if err != nil {
		t.Errorf("find spell failed %s", err)
	}

	if s.Name != "Accio" {
		t.Errorf("Spell name not accio")
	}
}

func TestRandomSpell(t *testing.T) {
	spells := RandomSpells(2)
	if len(spells) != 2 {
		t.Errorf("len(spells) != 2")
	}
}
