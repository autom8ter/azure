package azure

import (
	"golang.org/x/oauth2"
)

var defaultScopes = []string{"User.Read"}

func OAuth2Config(clientId string, redirect string, secret string, scopes ...string) *oauth2.Config {
	if len(scopes) == 0 {
		scopes = defaultScopes
	}
	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: secret, // no client secret
		RedirectURL:  redirect,

		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.microsoftonline.com/common/oauth2/authorize",
			TokenURL: "https://login.microsoftonline.com/common/oauth2/token",
		},

		Scopes: scopes,
	}
}

func NewConfig(clientId string, redirect string, secret string, scopes ...string) *oauth2.Config {
	if len(scopes) == 0 {
		scopes = []string{"User.Read"}
	}
	cfg := &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.microsoftonline.com/common/oauth2/authorize",
			TokenURL: "https://login.microsoftonline.com/common/oauth2/token",
		},
		RedirectURL: redirect,
		Scopes:      scopes,
	}
	return cfg
}
