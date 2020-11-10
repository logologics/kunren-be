package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/securecookie"
)

type SessionKey struct {
	AuthKey       string `json:"authKey"`
	EncryptionKey string `json:"encryptionKey"`
}

func main() {

	k := &SessionKey{
		AuthKey:       hex.EncodeToString(securecookie.GenerateRandomKey(64)),
		EncryptionKey: hex.EncodeToString(securecookie.GenerateRandomKey(32)),
	}

	js, err := json.MarshalIndent(k, " ", "  ")
	if err != nil {
		log.Fatalf("Something went wrong %v", err)
	}
	fmt.Println(string(js))
}
