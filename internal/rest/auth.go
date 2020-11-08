package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	d "github.com/logologics/kunren-be/internal/domain"

	"github.com/logologics/kunren-be/internal/api"
	"github.com/markbates/goth/gothic"
)

var KunrenUserKey = "kunren_user"
var fe = "http://localhost:3000/"

// Callback implements the goth callback
func (e *Env) Callback(w http.ResponseWriter, r *http.Request) error {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return api.NewHTTPBadRequest(err, "Something went wrong", "auth callback")
	}

	err = e.storeUserInSession(user.Email, w, r)
	if err != nil {
		return err
	}

	w.Header().Set("Location", fe)
	w.WriteHeader(http.StatusTemporaryRedirect)
	return nil
}

func (e *Env) storeUserInSession(email string, w http.ResponseWriter, r *http.Request) error {
	user, err := e.Repo.LoadUserByEmail(email)
	//  if user is not found
	if err != nil {
		return api.NewHTTPBadRequest(err, "User is not allowed", "auth callback")
	}

	gothic.StoreInSession(KunrenUserKey, user.Email, r, w)

	return nil
}

// Logout implements the goth logout
func (e *Env) Logout(w http.ResponseWriter, r *http.Request) error {
	gothic.Logout(w, r)
	w.Header().Set("Location", fe)
	w.WriteHeader(http.StatusTemporaryRedirect)
	return nil
}

// Authorize implements the goth flow initializer
func (e *Env) Authorize(w http.ResponseWriter, r *http.Request) error {
	// try to get the user without re-authenticating
	if user, err := gothic.CompleteUserAuth(w, r); err == nil {
		err = e.storeUserInSession(user.Email, w, r)
		if err != nil {
			return err
		}

		w.Header().Set("Location", fe)
		w.WriteHeader(http.StatusTemporaryRedirect)
		return nil
	}

	gothic.BeginAuthHandler(w, r)
	return nil
}

// Session returns the users email address
// if the user is logged in
func (e *Env) Session(w http.ResponseWriter, r *http.Request) error {
	email, err := gothic.GetFromSession(KunrenUserKey, r)
	if err != nil {
		return api.NewHTTPUnauthorized(err, "User not found in session", "auth callback")
	}

	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(d.Session{Email: email, Nonce: "Not used"})
}

// AuthenticatedMW is a middleware that check that
// user has a valid session
func (e *Env) AuthenticatedMW(next api.AppHandlerFunc) api.AppHandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) error {

		email, err := gothic.GetFromSession(KunrenUserKey, r)
		if err != nil {
			return api.NewHTTPUnauthorized(err, "User not found in session", "auth callback")
		}

		if email == "" {
			return api.NewHTTPUnauthorized(fmt.Errorf("Email empty"), "User not found in session", "auth callback")
		}
		return next(w, r)

	}
}
