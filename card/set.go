package card

type BaseSet struct {
	Name        string`json:"name"`        // The name of the set
	Code        string`json:"code"`        // The code name of the set
	Type        string`json:"type"`        // The type this set is in
	release     string`json:"releaseDate"` // When the set was released (YYYY-MM-DD)
}

type Set struct {
	BaseSet
	Cards       []Card`json:"cards"` // The cards in the set in Object format
}
