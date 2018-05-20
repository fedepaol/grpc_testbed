package client

import (
	"context"
	"log"

	proto "github.com/fedepaol/grpcsample/protospell"
	"google.golang.org/grpc"
)

// Run is expected to do some client stuff against the server.
// In this case it retrieves a list of random spells and re-retrieves them one by one.
func Run(address string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := proto.NewHarryPotterSpellsClient(conn)

	spells, err := client.GetRandomSpells(context.TODO(), &proto.RandomSpellsNumber{
		HowMany: int32(5),
	})

	if err != nil {
		log.Fatalf("failed to retrieve random spells: %v", err)
	}

	for _, s := range spells.Spells {
		spell, err := client.GetSpell(context.TODO(), &proto.SpellName{Name: s.GetName()})

		if err != nil {
			log.Fatalf("Error in retreiving spell %s", s.GetName())
		}

		log.Println("Retrieved spell ", *spell)
	}
}
