package api

import (
	"encoding/hex"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	d "github.com/logologics/kunren-be/internal/domain"
	r "github.com/logologics/kunren-be/internal/repo"
	mongo "github.com/logologics/kunren-be/internal/repo/mongo"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
	"github.com/sirupsen/logrus"
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

func decodeKeys(sessionKeys []d.SessionKey) ([][]byte, error) {
	keys := make([][]byte, 2*len(sessionKeys))
	cnt := 0
	for _, sk := range sessionKeys {
		key, err := hex.DecodeString(sk.AuthKey)
		if err != nil {
			return nil, err
		}
		keys[cnt] = key
		cnt++
		key, err = hex.DecodeString(sk.EncryptionKey)
		if err != nil {
			return nil, err
		}
		keys[cnt] = key
		cnt++
	}
	return keys, nil
}

func initAuth(c *d.Config) (*ProviderIndex, error) {

	maxAge := 86400 * 30 // 30 days
	isProd := false      // Set to true when serving over https
	sessionKeys, err := decodeKeys(c.Auth.SessionKeys)
	if err != nil {
		return nil, err
	}
	store := sessions.NewCookieStore(sessionKeys...)
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

func sendError(c *gin.Context, status int, err error, msg string) {
	logrus.Errorf("Sending error %v: %v", msg, err)
	c.JSON(status, gin.H{"msg": msg})
}
