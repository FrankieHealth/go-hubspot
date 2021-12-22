package hubspot

import (
	"fmt"
	"net/http"
)

var Key string

type Authenticator interface {
	SetAuthentication(r *http.Request) error
}

type AuthMethod func(c *Client)

func SetOAuth(config *OAuthConfig) AuthMethod {
	return func(c *Client) {
		c.authenticator = &OAuth{
			retriever: &OAuthTokenManager{
				oauthPath:  fmt.Sprintf("%s/%s", c.baseURL.String(), oauthTokenPath),
				HTTPClient: c.HTTPClient,
				Config:     config,
			},
		}
	}
}

func SetAPIKey(key string) AuthMethod {
	return func(c *Client) {
		c.authenticator = &APIKey{
			apikey: key,
		}
	}
}

type OAuth struct {
	retriever OAuthTokenRetriever
}

func SetAuthentication(r *http.Request) error {
	r.Header.Set("Authorization", "Bearer "+Key)
	return nil
}

func (o *OAuth) SetAuthentication(r *http.Request) error {
	r.Header.Set("Authorization", "Bearer "+Key)
	return nil
}

type APIKey struct {
	apikey string
}

func (a *APIKey) SetAuthentication(r *http.Request) error {
	if Key != "" {
		q := r.URL.Query()
		q.Set("hapikey", a.apikey)
		r.URL.RawQuery = q.Encode()
	}
	return nil
}
