package parser

import (
	"github.com/fiam/gounidecode/unidecode"
	"regexp"
	"strings"
)

type card struct {
	Id            int`json:"multiverseid"`              //The multiverseid of the card on Wizard's Gatherer web page.
	Name          string`json:"name"`                   //The card name. For split, double-faced and flip cards, just the name of one side of the card.
	Ascii         string`json:"ascii"`                  //The card name in ASCII format
	Names         *[]string`json:"names"`               //Only used for split, flip and dual cards. Will contain all the names on this card, front or back.
	image         *string`json:"imageName"`             //The mtgimage.com file name for this card.
	Type          string`json:"type"`                   // The card type. This is the type you would see on the card if printed today. Uses UTF-8 u"\u2014"
	Types         []string`json:"types"`                //The types of the card. These appear to the left of the dash in a card type.
	SuperTypes    *[]string`json:"supertypes"`          //The supertypes of the card. These appear to the far left of the card type.
	SubTypes      *[]string`json:"subtypes`             //The subtypes of the card. These appear to the right of the dash in a card type.
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
	Colors        *[]string`json:"colors"`              //The card colors.
	Legalities    *map[string]string`json:"legalities"` //Which formats this card is legal, restricted or banned in.
	Rulings       *[]ruling`json:"rulings"`             //Each Ruling object has 'date' and 'text' keys.
	Printings     *[]string`json:"printings"`           //The sets that this card was printed in.

	OriginalText  *string`json:"originalText"` //The original text on the card at the time it was printed. (Used for Lands)
	OriginalType  *string`json:"originalType"` //The original type on the card at the time it was printed. (Used for Tokens)
}

type ruling struct {
	Date string`json:"date"` //When the rule was set (YYYY-MM-DD)
	Text string`json:"text"` //The rule defined
}

func (card *card) clean() {
	card.Ascii = unidecode.Unidecode(card.Name)
	if card.Id == 0 {
		card.Id = -1;
	}
	if len(card.Type) == 0 {
		card.Type = *card.OriginalType
	}
	if card.Text == nil && card.Flavor == nil {
		card.Text = card.OriginalText
	}
	if card.Mana != nil {
		*card.Mana = mana_scooper(*card.Mana)
	}
	if card.Text != nil {
		*card.Text = mana_scooper(*card.Text)
	}
}

func mana_scooper(mana_text string) string {
	multi_re, _ := regexp.Compile("{(.{3})}")
	mana_re := strings.NewReplacer("{", "", "}", "")

	if matches := multi_re.FindAllStringSubmatch(mana_text, -1); len(matches) != 0 {

		for _, match := range matches {
			mana_text = strings.Replace(mana_text, "{"+match[1]+"}", "("+match[1]+")", 1)
		}
	}
	return mana_re.Replace(mana_text)
}
