package dockerhub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersLoginOutput struct {
	Token string `json:"token"`
}

func (c *Client) UsersLogin() (*UsersLoginOutput, error) {
	payload, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/users/login", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var output UsersLoginOutput
	if err = json.Unmarshal(body, &output); err != nil {
		return nil, err
	}
	return &output, nil
}
