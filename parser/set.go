package parser

import kard "github.com/souleiman/warden/card"

type set struct {
	kard.BaseSet
	Cards       []card`json:"cards"`       // The cards in the set in Object format
}

func (set *set) clean() {
	cards := set.Cards
	for i, _ := range cards {
		cards[i].clean()
	}
}

func (s set) convert() kard.Set {
	var conv_set kard.Set
	conv_set.BaseSet = s.BaseSet
	conv_set.Cards = make([]kard.Card, len(s.Cards))

	for i, card := range s.Cards {
		conv_set.Cards[i] = card.Card
	}
	return conv_set
}

func (s set) process() {
	s.clean()
	w := s.convert().CreateWriter()
	w.WriteToDatabase()
	w.Close()
}
