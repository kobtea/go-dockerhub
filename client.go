package dockerhub

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const HostURL string = "https://hub.docker.com/v2"

type Client struct {
	HostURL    *url.URL
	HTTPClient *http.Client
	Token      string
	Auth       *Auth
}

type ClientOption struct {
	Endpoint string
}

type PageOption struct {
	Page     int
	PageSize int
}

func (p *PageOption) Empty() bool {
	return p.Page == 0
}

func NewClient(auth *Auth, option *ClientOption) (*Client, error) {
	hostUrlStr := HostURL
	if len(option.Endpoint) != 0 {
		hostUrlStr = option.Endpoint
	}
	if len(auth.Username) == 0 || len(auth.Password) == 0 {
		return nil, fmt.Errorf("need username and password")
	}
	hostUrl, err := url.ParseRequestURI(hostUrlStr)
	if err != nil {
		return nil, err
	}

	client := &Client{
		HostURL:    hostUrl,
		HTTPClient: http.DefaultClient,
		Auth:       auth,
	}
	res, err := client.UsersLogin()
	if err != nil {
		return nil, err
	}
	client.Token = res.Token
	return client, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode/100 != 2 {
		return nil, fmt.Errorf("status: %d, message: %s", res.StatusCode, body)
	}
	return body, nil
}
