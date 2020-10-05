package domain

type Config struct {
	Address string  `json:"address"`
	KunrenFe string `json:"kunrenfe"`
	Https `json:"https"`
	DB
}

// Https support
type Https struct {
	Enabled  bool   `json:"enabled"`
	CertPath string `json:"certPath"`
	KeyPath  string `json:"keyPath"`
}

type DB struct {
	Type string `json:"type"`
	URL string `json:"url"`
}
