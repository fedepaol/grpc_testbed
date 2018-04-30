package spells

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

// Spell is a spell used in Harry Potter movies.
type Spell struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Effect string `json:"effect"`
}

var spells = make(map[string]Spell)
var spellsList = make([]Spell, 0)

func init() {
	jsonSpells := loadSpells()
	for _, s := range jsonSpells {
		spells[s.Name] = s
	}
}

// Find tries to retrieve a spell by its name.
func Find(name string) (Spell, error) {
	spell, found := spells[name]
	if !found {
		return Spell{}, fmt.Errorf("%s not found", name)
	}
	return spell, nil
}

// RandomSpells return a slice of random spells.
func RandomSpells(howMany int) (res []Spell) {
	if howMany == 0 {
		return nil
	}

	res = make([]Spell, howMany)
	for i := 0; i < howMany; i++ {
		indx := rand.Intn(len(spellsList))
		res[i] = spellsList[indx]
	}
	return res
}

func loadSpells() (res []Spell) {
	raw, err := ioutil.ReadFile("./spells.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.Unmarshal(raw, &res)

	for _, spell := range res {
		spellsList = append(spellsList, spell)
	}
	return
}
