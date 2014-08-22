package card

import (
	"github.com/fiam/gounidecode/unidecode"
	"regexp"
	"strings"
)

type Set struct {
	Name        string`json:"name"`        // The name of the set
	Code        string`json:"code"`        // The code name of the set
	Type        string`json:"type"`        // The type this set is in
	release     string`json:"releaseDate"` // When the set was released (YYYY-MM-DD)
	Cards       []Card`json:"cards"`       // The cards in the set in Object format
	cleaned     bool                       //Flag indicating set has been cleaned
}

func (set *Set) Clean() {
	if set.cleaned {
		return
	}

	cards := set.Cards
	for i, kard := range cards {
		card := &cards[i]

		card.Ascii = unidecode.Unidecode(kard.Name)
		if kard.Id == 0 {
			card.Id = -1;
		}

		if len(kard.Type) == 0 {
			card.Type = *card.OriginalType
		}

		if kard.Text == nil && kard.Flavor == nil {
			card.Text = card.OriginalText
		}

		if kard.Mana != nil {
			*card.Mana = manaScooper(*kard.Mana)
		}

		if kard.Text != nil {
			*card.Text = manaScooper(*kard.Text)
		}
	}

	set.cleaned = true
}

func manaScooper(mana_text string) string {
	multi_re, _ := regexp.Compile("{(.{3})}")
	mana_re := strings.NewReplacer("{", "", "}", "")

	if matches := multi_re.FindAllStringSubmatch(mana_text, -1); len(matches) != 0 {

		for _, match := range matches {
			mana_text = strings.Replace(mana_text, "{"+match[1]+"}", "("+match[1]+")", 1)
		}
	}
	return mana_re.Replace(mana_text)
}
