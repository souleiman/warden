package card

type Set struct {
	Name        string`json:"name"`        // The name of the set
	Code        string`json:"code"`        // The code name of the set
	Block       string`json:"block"`       // The block this set is in
	release     string`json:"releaseDate"` // When the set was released (YYYY-MM-DD)
	Cards       []Card`json:"cards"`       // The cards in the set in Object format
}
