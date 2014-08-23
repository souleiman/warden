package card

type BaseCard struct {
	Id            int`json:"multiverseid"`              //The multiverseid of the card on Wizard's Gatherer web page.
	Name          string`json:"name"`                   //The card name. For split, double-faced and flip cards, just the name of one side of the card.
	Ascii         string`json:"ascii"`                  //The card name in ASCII format
	image         *string`json:"imageName"`             //The mtgimage.com file name for this card.
	Type          string`json:"type"`                   // The card type. This is the type you would see on the card if printed today. Uses UTF-8 u"\u2014"
	Mana          *string`json:"manaCost"`              //The mana cost of this card.
	Cost          *float64`json:"cmc"`                  // Converted mana cost.
	Text          *string`json:"text"`                  // The text of the card.
	Flavor        *string`json:"flavor"`                //The flavor text of the card.
	Layout        string`json:"layout"`                 //The card layout. Possible values: normal, split, flip, double-faced, token, plane, scheme, phenomenon, leveler, vanguard
	Power         *string`json:"power"`                 //The power of the card. This is only present for creatures.
	Toughness     *string`json:"toughness"`             //The toughness of the card. This is only present for creatures.
	Loyalty       *int`json:"loyalty"`                  //The loyalty of the card. This is only present for planeswalkers.
	Hand          *int`json:"hand"`                     // Maximum hand size modifier. Only exists for Vanguard cards.
	Life          *int`json:"life"`                     // Starting life total modifier. Only exists for Vanguard cards.
}

type Card struct {
	BaseCard

	Names         *[]string`json:"names"`               //Only used for split, flip and dual cards. Will contain all the names on this card, front or back.
	Types         []string`json:"types"`                //The types of the card. These appear to the left of the dash in a card type.
	SuperTypes    *[]string`json:"supertypes"`          //The supertypes of the card. These appear to the far left of the card type.
	SubTypes      *[]string`json:"subtypes`             //The subtypes of the card. These appear to the right of the dash in a card type.
	Colors        *[]string`json:"colors"`              //The card colors.
	Rulings       *[]ruling`json:"rulings"`             //Each Ruling object has 'date' and 'text' keys.
	Legalities    *map[string]string`json:"legalities"` //Which formats this card is legal, restricted or banned in.
	Printings     *[]string`json:"printings"`           //The sets that this card was printed in.
}

type ruling struct {
	Date string`json:"date"` //When the rule was set (YYYY-MM-DD)
	Text string`json:"text"` //The rule defined
}
