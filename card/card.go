package card

type Set struct {
	Name    string`json:"name"`        // The name of the set
	Code    string`json:"code"`        // The code name of the set
	Block   string`json:"block"`       // The block this set is in
	release string`json:"releaseDate"` // When the set was released (YYYY-MM-DD)
	Cards   []Card`json:"cards"`       // The cards in the set
}

type Card struct {
	Id                 int`json:"multiverseid"`             //The multiverseid of the card on Wizard's Gatherer web page.
	Name               string`json:"name"`                  //The card name. For split, double-faced and flip cards, just the name of one side of the card.
	Names              []string`json:"names"`               //Only used for split, flip and dual cards. Will contain all the names on this card, front or back.
	image              string`json:"imageName"`             //The mtgimage.com file name for this card.

	Type               string`json:"type"`                  // The card type. This is the type you would see on the card if printed today. Uses UTF-8 u"\u2014"
	Supertypes         []string`json:"supertypes"`          //The supertypes of the card. These appear to the far left of the card type.
	Types              []string`json:"types"`               //The types of the card. These appear to the left of the dash in a card type.
	Subtypes           []string`json:"subtypes`             //The subtypes of the card. These appear to the right of the dash in a card type.

	Mana               string`json:"manaCost"`              //The mana cost of this card.
	Cost               float64`json:"cmc"`                  // Converted mana cost.

	Text               string`json:"text"`                  // The text of the card.
	Flavor             string`json:"flavor"`                //The flavor text of the card.

	Power              string`json:"power"`                 //The power of the card. This is only present for creatures.
	Toughness          string`json:"toughness"`             //The toughness of the card. This is only present for creatures.

	Loyalty            int`json:"loyalty"`                  //The loyalty of the card. This is only present for planeswalkers.

	Layout             string`json:"layout"`                //The card layout. Possible values: normal, split, flip, double-faced, token, plane, scheme, phenomenon, leveler, vanguard
	Colors             []string`json:"colors"`              //The card colors.
	Rarity             string`json:"rarity"`                // The rarity of the card. Example values: Basic Land, Common, Uncommon, Rare, Mythic Rare, Special

	Artist             string`json:"artist"`                // The artist of the card.
	Number             string`json:"number"`                // The card number in the set.
	Variations         []int`json:"variations"`             //If a card has alternate art (for example, 4 different Forests, or the 2 Brothers Yamazaki) then each other variation's id will be listed here, NOT including the current card's id.

	Legalities         map[string]string`json:"legalities"` //Which formats this card is legal, restricted or banned in.
	Rulings            []Ruling`json:"rulings"`             //Each Ruling object has 'date' and 'text' keys.
	Printings          []string`json:"printings"`           //The sets that this card was printed in.

	TimeShifted        bool`json:"timeshifted"`             //If this card was a timeshifted card in the set.

	Hand               int`json:"hand"`                     // Maximum hand size modifier. Only exists for Vanguard cards.
	Life               int`json:"life"`                     // Starting life total modifier. Only exists for Vanguard cards.

	OriginalText       string`json:"originalText"`          //The original text on the card at the time it was printed.
	OriginalType       string`json:"originalType"`          //The original type on the card at the time it was printed.
}

type Ruling struct {
	Date string`json:"date"` //When the rule was set (YYYY-MM-DD)
	Text string`json:"text"` //The rule defined
}
