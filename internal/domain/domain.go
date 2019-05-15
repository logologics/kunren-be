package domain

// Welcome contains the Kunren welcome message
type Welcome struct {
	Version string `json:"version"`
	Hello   string `json:"message"`
}

// Https support
type Https struct {
	Enabled  bool   `json:"enabled"`
	CertPath string `json:"certPath"`
	KeyPath  string `json:"keyPath"`
}

type Question struct {
	Question string `json:"q"`
	Answer string `json:"a"`
	Features []string `json:"fs"`
}

type Questions struct {
	Questions map[string]Question `json:"qs"`
}
