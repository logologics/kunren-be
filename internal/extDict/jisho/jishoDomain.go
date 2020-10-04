package jisho

type JishoResponse struct {
	JLemmas []JLemma `json:"data"`
}

type JLemma struct {
	Key      string     `json:"slug"`
	Japanese []Japanese `json:"japanese"`
	Meanings []JMeaning `json:"senses"`
}

type Japanese struct {
	Reading string `json:"reading"`
	Lexeme  string `json:"word"`
}

type JMeaning struct {
	POS         []string `json:"parts_of_speech"`
	Translations []string `json:"english_definitions"`
}
