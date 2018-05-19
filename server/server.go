package spellsServer

import (
	"context"

	proto "github.com/fedepaol/grpcsample/protospell"
	"github.com/fedepaol/grpcsample/spells"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

type HarryServer struct {
}

func (s HarryServer) GetSpell(ctx context.Context, spell *proto.SpellName) (*proto.Spell, error) {
	s1, err := spells.Find(spell.Name)
	if err != nil {
		return nil, err
	}

	res := proto.Spell{s1.Name, s1.Type, s1.Effect}
	return &res, nil
}

func (s HarryServer) GetRandomSpell(ctx context.Context, empty *google_protobuf.Empty) (*proto.Spell, error) {
	sp := spells.RandomSpells(1)[0]

	return &proto.Spell{
		Name:   sp.Name,
		Type:   sp.Type,
		Effect: sp.Effect}, nil
}

func (s HarryServer) GetRandomSpells(ctx context.Context, howMany *proto.RandomSpellsNumber) (*proto.SpellsList, error) {
	spellList := spells.RandomSpells(int(howMany.HowMany))

	var res proto.SpellsList
	res.Spells = make([](*proto.Spell), len(spellList))
	for indx, spell := range spellList {
		res.Spells[indx] = &proto.Spell{
			Name:   spell.Name,
			Type:   spell.Type,
			Effect: spell.Effect}
	}
	return &res, nil
}
