package card

import (
	"encoding/json"
	"fmt"
)

type Set struct {
	Name        string`json:"name"`           // The name of the set
	Code        string`json:"code"`           // The code name of the set
	Block       string`json:"block"`          // The block this set is in
	release     string`json:"releaseDate"`    // When the set was released (YYYY-MM-DD)
	RawCards    json.RawMessage`json:"cards"` // Maintain the raw integrity of the card in json format
	Cards       []Card                        // The cards in the set in Object format
}

func (set Set) Transform() {
	fmt.Println("Begin Transformation of RawCards -> Cards")
}
