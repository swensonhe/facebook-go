package facebook

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	*http.Client
}

const (
	urlScheme = "https"
	urlHost   = "graph.facebook.com"
	urlPath   = "/v2.11"

	paramKeyFields      = "fields"
	paramKeyAccessToken = "access_token"
)

func NewClient() *Client {
	return &Client{
		Client: &http.Client{},
	}
}

func (c *Client) GetMe(accessToken string, fields ...string) (*User, error) {
	u := baseUrl()
	u.Path += "/me"

	q := u.Query()
	params := []string{}
	for _, field := range fields {
		params = append(params, field)
	}
	q.Set(paramKeyFields, strings.Join(params, ","))
	q.Set(paramKeyAccessToken, accessToken)

	u.RawQuery = q.Encode()
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, ErrInternal
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	user, err := parseUser(res)
	if err != nil {
		return nil, ErrInternal
	}

	return user, nil
}

func baseUrl() url.URL {
	return url.URL{
		Scheme: urlScheme,
		Host:   urlHost,
		Path:   urlPath,
	}
}

func parseUser(res *http.Response) (*User, error) {
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
