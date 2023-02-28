package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tkwang0530/music-streaming/internal/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) *AuthHandler {
	return &AuthHandler{
		cfg: cfg,
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	config := h.getOAuthConfig()
	url := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (h *AuthHandler) Callback(c *gin.Context) {
	code := c.Query("code")

	config := h.getOAuthConfig()
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to exchange token"})
		return
	}

	// TODO: Save token to database or session
	fmt.Println(token)

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in"})
}

func (h *AuthHandler) getOAuthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     h.cfg.OAuth.Google.ClientID,
		ClientSecret: h.cfg.OAuth.Google.ClientSecret,
		RedirectURL:  h.cfg.OAuth.Google.RedirectURI,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     google.Endpoint,
	}
}
