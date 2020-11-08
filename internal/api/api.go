package api

import (
	"net/http"

	"github.com/gorilla/sessions"
	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"
	mongo "github.com/logologics/kunren-be/internal/repo/mongo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	log "github.com/sirupsen/logrus"
	mp "go.mongodb.org/mongo-driver/bson/primitive"
)

type ProviderIndex struct {
	Providers    []string
	ProvidersMap map[string]string
}

// Env is the api env
type Env struct {
	Config *d.Config
	Repo   r.Repo
	d.User
	ProviderIndex *ProviderIndex
}

// AppHandlerFunc that returns error
type AppHandlerFunc func(http.ResponseWriter, *http.Request) error

// ServeHTTP calls
func (fn AppHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		switch v := err.(type) {
		case HTTPError:
			log.WithFields(log.Fields{"loc": "ServeHttp", "err": v, "ctx": v.Context}).Error(v.Message)
			v.SendError(w)
		default:
			log.WithFields(log.Fields{"loc": "ServeHttp", "err": err}).Error("Unexpected error")
			http.Error(w, "Unexpected server error", http.StatusInternalServerError)
		}
	}
}

func CreateEnv(c *d.Config) (*Env, error) {
	repo, err := createRepo(c)
	if err != nil {
		return nil, err
	}

	uID, err := mp.ObjectIDFromHex("000000000000000000000005")
	if err != nil {
		log.Fatal("Can't create default user: %v", err)
	}

	u := d.User{Email: "alex@alex.com", ID: uID}

	pI, err := initAuth(c)
	if err != nil {
		log.Fatal("Can't init auth: %v", err)
	}

	return &Env{Config: c, Repo: repo, User: u, ProviderIndex: pI}, nil

}

// CreateRepo creates a new repo (only mongo supported)
func createRepo(c *d.Config) (r.Repo, error) {
	return mongo.Connect(c)
}

func initAuth(c *d.Config) (*ProviderIndex, error) {
	
	key := "scjhbafhgkiasgb,sjvba,sdjvhdsg,vhb" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30                        // 30 days
	isProd := false                             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd

	gothic.Store = store
	googleProv := c.Auth.Providers["google"]

	goth.UseProviders(
		google.New(
			googleProv.ClientID,
			googleProv.ClientSecret,
			"http://localhost:9876/auth/callback/google"))

	m := make(map[string]string)
	m["google"] = "Google"
	keys := []string{"google"}
	return &ProviderIndex{Providers: keys, ProvidersMap: m}, nil
}
