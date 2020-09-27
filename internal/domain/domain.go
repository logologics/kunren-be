package domain

// Welcome contains the Kunren welcome message
type Welcome struct {
	Version string `json:"version"`
	Hello   string `json:"message"`
}

type Question struct {
	ID int `json:"id"`
	Question string `json:"q"`
	Answer string `json:"a"`
	Features []string `json:"fs"`
}

type Questions struct {
	Questions []Question `json:"qs"`
}
