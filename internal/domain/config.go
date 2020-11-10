package domain

type Config struct {
	Address  string `json:"address"`
	FrontEndURL string `json:"frontEndURL"`
	FrontEndPath   string `json:"frontEndPath"`
	Https    `json:"https"`
	DB       `json:"db"`
	Auth     Auth `json:"https"`
}

type SessionKey struct {
	AuthKey       string `json:"authKey"`
	EncryptionKey string `json:"encryptionKey"`
}

type Auth struct {
	SessionKeys []SessionKey        `json:"sessionKeys"`
	Providers   map[string]Provider `json:"providers"`
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
