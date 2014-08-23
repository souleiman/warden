package parser

import (
	"github.com/fiam/gounidecode/unidecode"
	"regexp"
	"strings"
	kard "github.com/souleiman/warden/card"
)

type card struct {
	kard.Card

	OriginalText  *string`json:"originalText"` //The original text on the card at the time it was printed. (Used for Lands)
	OriginalType  *string`json:"originalType"` //The original type on the card at the time it was printed. (Used for Tokens)
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
