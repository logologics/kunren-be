package jisho_test

import (
	"fmt"
	"testing"

	d "github.com/logologics/kunren-be/internal/domain"
	jd "github.com/logologics/kunren-be/internal/extDict/jisho"
)

func TestConvert(t *testing.T) {
	jr := readRawJishoResponse(t, "test/jisho.home.json")
	sr, err := jd.Convert(jr)

	l := len(sr.Words)
	fmt.Print(l)
	if err != nil {
		t.Errorf("No err expected, but got %v", err)
	}

	if l != 20 {
		t.Errorf("Expected words, but got %v", len(sr.Words))
	} 

	w4 := sr.Words[4]
	if w4.DateCreated.IsZero() {
		t.Errorf("Date not updated")
	}

	if w4.Language != d.Japanese {
		t.Errorf("Language should be Japanese")
	}

	if w4.Source != "jisho" {
		t.Errorf("Source should be jisho")
	}

 	l4 := w4.Lemma
	 if l4.Key != "帰る" {
		t.Errorf("L4 expected key 帰る got %v", l4.Key)

	}

	if len(l4.Meanings) != 3 {
		t.Errorf("Wrong number of Meaning, expected 3, but got %v", len(l4.Meanings))
	}

	if l4.Lexeme != "帰る" {
		t.Errorf("Wrong Japanese[0].Lexeme, expected 帰る, but got %v", l4.Lexeme)
	}

	if l4.Reading != "かえる" {
		t.Errorf("Wrong Japanese[0].Reading, expected , but got %v", l4.Reading)
	}

	if len(l4.Meanings[0].POS) != 2 {
		t.Errorf("Wrong number of POS, expected 3, but got %v", len(l4.Meanings[0].POS))
	}

	if len(l4.Meanings[0].Translations) != 4 {
		t.Errorf("Wrong number of Translations, expected 4, but got %v", len(l4.Meanings[0].Translations))
	}

	if len(l4.Meanings[0].POS) != 2 {
		t.Errorf("Wrong number of POS, expected 2, but got %v", len(l4.Meanings[0].POS))
	}

	if l4.Meanings[0].POS[0] != "Godan verb with ru ending" {
		t.Errorf("Wrong POS, expected `Godan verb with ru ending`, but got %v", l4.Meanings[0].POS[0])
	}

	if l4.Meanings[0].Translations[1].Text != "to come home" {
		t.Errorf("Wrong Translation, expected `to come home`, but got %v", l4.Meanings[0].Translations[1].Text)
	}

	if l4.Meanings[0].Translations[1].Language != d.English {
		t.Errorf("Lang should be en but got %v", l4.Meanings[0].Translations[1].Language)
	}

}
