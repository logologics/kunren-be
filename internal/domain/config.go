package domain

type Config struct {
	Address string  `json:"address"`
	KunrenFe string `json:"kunrenfe"`
	Https `json:"https"`
}

// Https support
type Https struct {
	Enabled  bool   `json:"enabled"`
	CertPath string `json:"certPath"`
	KeyPath  string `json:"keyPath"`
}
