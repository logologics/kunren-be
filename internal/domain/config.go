package domain

type Config struct {
	Address  string `json:"address"`
	KunrenFe string `json:"kunrenfe"`
	Https    `json:"https"`
	DB       `json:"db"`
	Auth     Auth `json:"https"`
}

type Auth struct {
	Providers map[string]Provider `json:"providers"`
}

type Provider struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

// Https support
type Https struct {
	Enabled  bool   `json:"enabled"`
	CertPath string `json:"certPath"`
	KeyPath  string `json:"keyPath"`
}

type DB struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
