package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	d "github.com/logologics/kunren-be/internal/domain"

	"github.com/markbates/goth/gothic"
)

var kunrenUserKey = "kunren_user"

// Callback implements the goth callback
func (e *Env) Callback(c *gin.Context) {
	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		sendError(c, http.StatusBadRequest, err, "Oauth callback failed")
		return
	}

	if err = e.storeUserInSession(user.Email, c); err != nil {
		sendError(c, http.StatusInternalServerError, err, "Oauth callback failed")
		return
	}

	c.Redirect(http.StatusTemporaryRedirect, e.Config.FrontEndURL)
}

func (e *Env) storeUserInSession(email string, c *gin.Context) error {
	user, err := e.Repo.LoadUserByEmail(email)
	//  if user is not found
	if err != nil {
		return fmt.Errorf("User is not allowed: %v", err)
	}

	gothic.StoreInSession(kunrenUserKey, user.Email, c.Request, c.Writer)

	return nil
}

// Logout implements the goth logout
func (e *Env) Logout(c *gin.Context) {
	gothic.Logout(c.Writer, c.Request)
	c.Redirect(http.StatusTemporaryRedirect, e.Config.FrontEndURL)
}

// Authorize implements the goth flow initializer
func (e *Env) Authorize(c *gin.Context) {
	// try to get the user without re-authenticating
	if user, err := gothic.CompleteUserAuth(c.Writer, c.Request); err == nil {
		if err = e.storeUserInSession(user.Email, c); err != nil {
			sendError(c, http.StatusInternalServerError, err, "Reauthorization  failed")
		}

		c.Redirect(http.StatusTemporaryRedirect, e.Config.FrontEndURL)
	}

	gothic.BeginAuthHandler(c.Writer, c.Request)
}

// Session returns the users email address
// if the user is logged in
func (e *Env) Session(c *gin.Context) {
	email, err := gothic.GetFromSession(kunrenUserKey, c.Request)
	if err != nil {
		sendError(c, http.StatusUnauthorized, err, "No user in session")
		return
	}

	c.JSON(http.StatusOK, d.Session{Email: email, Nonce: "Not used"})
}
