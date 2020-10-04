package jisho_test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	jd "github.com/logologics/kunren-be/internal/extDict/jisho"
)

func TestJDomain(t *testing.T) {
	r := readJishoResponse(t, "test/jisho.home.json")

	if len(r.JLemmas) != 20 {
		t.Errorf("Wrong number of lemmas, expected 20, but got %v", len(r.JLemmas))
	}

	l4 := r.JLemmas[4]

	if l4.Key != "帰る" {
		t.Errorf("L4 expected key 帰る got %v", l4.Key)

	}

	if len(l4.Japanese) != 3 {
		t.Errorf("Wrong number of Japanese, expected 3, but got %v", len(l4.Japanese))
	}

	if len(l4.Meanings) != 3 {
		t.Errorf("Wrong number of Meaning, expected 3, but got %v", len(l4.Meanings))
	}

	if l4.Japanese[0].Lexeme != "帰る" {
		t.Errorf("Wrong Japanese[0].Lexeme, expected 帰る, but got %v", l4.Japanese[0].Lexeme)
	}

	if l4.Japanese[0].Reading != "かえる" {
		t.Errorf("Wrong Japanese[0].Reading, expected , but got %v", l4.Japanese[0].Reading)
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

	if l4.Meanings[0].Translations[1] != "to come home" {
		t.Errorf("Wrong Translation, expected `to come home`, but got %v", l4.Meanings[0].Translations[1])
	}
}

func expectedResponse() jd.JishoResponse {
	return jd.JishoResponse{}
}

func readRawJishoResponse(t *testing.T, path string) []byte {

	// read file
	f, err := os.Open(path)
	if err != nil {
		t.Errorf("Could not load json file: %v", err)
	}
	defer f.Close()

	// read bytes
	byteValue, err := ioutil.ReadAll(f)
	if err != nil {
		t.Errorf("Could not read bytes from json: %v", err)
	}

	return byteValue
}

func readJishoResponse(t *testing.T, path string) jd.JishoResponse {

	byteValue := readRawJishoResponse(t, path)

	//unmarshal
	var resp = jd.JishoResponse{}
	if err := json.Unmarshal(byteValue, &resp); err != nil {
		t.Errorf("Could not parse file : %v", err)
	}

	return resp
}
