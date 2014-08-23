package parser

import kard "github.com/souleiman/warden/card"

type set struct {
	kard.BaseSet
	Cards       []card`json:"cards"`       // The cards in the set in Object format
}

func (set *set) Clean() {
	cards := set.Cards
	for i, _ := range cards {
		cards[i].clean()
	}
}
