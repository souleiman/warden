package parser

type set struct {
	Name        string`json:"name"`        // The name of the set
	Code        string`json:"code"`        // The code name of the set
	Type        string`json:"type"`        // The type this set is in
	release     string`json:"releaseDate"` // When the set was released (YYYY-MM-DD)
	Cards       []card`json:"cards"`       // The cards in the set in Object format
}

func (set *set) Clean() {
	cards := set.Cards
	for i, _ := range cards {
		cards[i].clean()
	}
}
