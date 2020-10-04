package jisho

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	d "github.com/logologics/kunren-be/internal/domain"
)

var jishoSearchAPIURL = "https://jisho.org/api/v1/search/words?keyword="
var timeOut = time.Duration(2 * time.Second)

var httpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns:       20,
		IdleConnTimeout:    timeOut,
		DisableCompression: false,
		DisableKeepAlives:  false,
	},
	Timeout: timeOut}

func Search(query string) (d.SearchResult, error) {
	r, err := httpClient.Get(jishoSearchAPIURL + query)
	if err != nil {
		return d.SearchResult{}, err
	}

	// close the response
	defer func() {
		if r.Body != nil {
			r.Body.Close()
		}
	}()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return d.SearchResult{}, err
	}
	return Convert(body)

}

// convert a jisho json result into a kunren searchresult
func Convert(b []byte) (d.SearchResult, error) {

	//unmarshal
	var resp = JishoResponse{}
	if err := json.Unmarshal(b, &resp); err != nil {
		return d.SearchResult{}, err
	}

	sr := d.SearchResult{}
	sr.Words = convWords(resp.JLemmas)

	return sr, nil
}

func convWords(jlems []JLemma) []d.Word {
	words := make([]d.Word, len(jlems))
	for i, jl := range jlems {
		words[i] = convWord(jl)
	}
	return words
}

func convWord(jl JLemma) d.Word {
	w := d.Word{}
	w.Key = jl.Key
	w.Language = d.Japanese
	w.Lemma = convLemma(jl)
	w.Source = "jisho"
	w.DateCreated = time.Now()
	return w
}

func convLemma(jl JLemma) d.Lemma {
	l := d.Lemma{}
	l.Key = jl.Key
	l.Lexeme = jl.Japanese[0].Lexeme
	l.Reading = jl.Japanese[0].Reading
	l.Meanings = convMeanings(jl.Meanings)
	return l
}

func convMeanings(jms []JMeaning) []d.Meaning {
	ms := make([]d.Meaning, len(jms))
	for i, jm := range jms {
		ms[i] = convMeaning(jm)
	}
	return ms
}

func convMeaning(jm JMeaning) d.Meaning {
	m := d.Meaning{}
	m.POS = jm.POS
	m.Translations = convTranslations(jm.Translations)
	return m
}

func convTranslations(jtrs []string) []d.Translation {
	trs := make([]d.Translation, len(jtrs))
	for i, jtr := range jtrs {
		trs[i] = d.Translation{Language: d.English, Text: jtr}
	}
	return trs
}
