package azure

import (
	"context"
	"golang.org/x/oauth2"
	"net/http"
)

var defaultScopes = []string{"User.Read"}

type Config struct {
	ClientId string
	Redirect string
	Secret   string
	Scopes   []string
}

func (c *Config) Oauth2() *oauth2.Config {
	if len(c.Scopes) == 0 {
		c.Scopes = defaultScopes
	}
	cfg := &oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.Secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://login.microsoftonline.com/common/oauth2/authorize",
			TokenURL: "https://login.microsoftonline.com/common/oauth2/token",
		},
		RedirectURL: c.Redirect,
		Scopes:      c.Scopes,
	}
	return cfg
}

func (c *Config) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.Oauth2().AuthCodeURL(state, opts...)
}

func (c *Config) Client(ctx context.Context, code string, opts ...oauth2.AuthCodeOption) (*http.Client, error) {
	tok, err := c.Oauth2().Exchange(ctx, code, opts...)
	if err != nil {
		return nil, err
	}
	return c.Oauth2().Client(ctx, tok), nil
}
