package dockerhub

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *Client) GetOrg(org string) (*Organization, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/orgs/%s", c.HostURL, org), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var output Organization
	if err = json.Unmarshal(body, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type ListOrgGroupsPageOutput struct {
	Count    int        `json:"count"`
	Next     string     `json:"next,omitempty"`
	Previous string     `json:"previous,omitempty"`
	Results  []OrgGroup `json:"results"`
}

func (c *Client) ListOrgGroupsPage(org string, opt *PageOption) ([]OrgGroup, *PageOption, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/orgs/%s/groups", c.HostURL, org), nil)
	if err != nil {
		return nil, nil, err
	}
	if opt.Page != 0 {
		req.URL.Query().Add("page", strconv.Itoa(opt.Page))
	}
	if opt.PageSize != 0 {
		req.URL.Query().Add("page_size", strconv.Itoa(opt.PageSize))
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var output ListOrgGroupsPageOutput
	if err = json.Unmarshal(body, &output); err != nil {
		return nil, nil, err
	}

	var next PageOption
	if len(output.Next) != 0 {
		values, err := url.ParseQuery(output.Next)
		if err != nil {
			return nil, nil, err
		}
		if np := values.Get("page"); len(np) != 0 {
			if next.Page, err = strconv.Atoi(np); err != nil {
				return nil, nil, err
			}
		}
		if nps := values.Get("page_size"); len(nps) != 0 {
			if next.PageSize, err = strconv.Atoi(nps); err != nil {
				return nil, nil, err
			}
		}
	}
	return output.Results, &next, nil
}

func (c *Client) ListOrgGroups(org string) ([]OrgGroup, error) {
	var res []OrgGroup
	var opt PageOption
	for true {
		group, next, err := c.ListOrgGroupsPage(org, &opt)
		if err != nil {
			return nil, err
		}
		res = append(res, group...)
		opt = *next
		if next.Empty() {
			break
		}
	}
	return res, nil
}

func (c *Client) GetOrgGroup(org, group string) (*OrgGroup, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/orgs/%s/groups/%s", c.HostURL, org, group), nil)
	if err != nil {
		return nil, err
	}
	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	var output OrgGroup
	if err = json.Unmarshal(body, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

type OrgGroupInput struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (c *Client) CreateOrgGroup(org string, opt *OrgGroupInput) error {
	payload, err := json.Marshal(opt)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/orgs/%s/groups", c.HostURL, org), strings.NewReader(string(payload)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) UpdateOrgGroup(org, group string, opt *OrgGroupInput) error {
	payload, err := json.Marshal(opt)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/orgs/%s/groups/%s", c.HostURL, org, group), strings.NewReader(string(payload)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteOrgGroup(org, group string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/orgs/%s/groups/%s", c.HostURL, org, group), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

type ListOrgMembersPageOutput struct {
	Count    int         `json:"count"`
	Next     string      `json:"next,omitempty"`
	Previous string      `json:"previous,omitempty"`
	Results  []OrgMember `json:"results"`
}

func (c *Client) ListOrgMembersPage(org string, opt *PageOption) ([]OrgMember, *PageOption, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/orgs/%s/members", c.HostURL, org), nil)
	if err != nil {
		return nil, nil, err
	}
	if opt.Page != 0 {
		req.URL.Query().Add("page", strconv.Itoa(opt.Page))
	}
	if opt.PageSize != 0 {
		req.URL.Query().Add("page_size", strconv.Itoa(opt.PageSize))
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var output ListOrgMembersPageOutput
	if err = json.Unmarshal(body, &output); err != nil {
		return nil, nil, err
	}

	var next PageOption
	if len(output.Next) != 0 {
		values, err := url.ParseQuery(output.Next)
		if err != nil {
			return nil, nil, err
		}
		if np := values.Get("page"); len(np) != 0 {
			if next.Page, err = strconv.Atoi(np); err != nil {
				return nil, nil, err
			}
		}
		if nps := values.Get("page_size"); len(nps) != 0 {
			if next.PageSize, err = strconv.Atoi(nps); err != nil {
				return nil, nil, err
			}
		}
	}
	return output.Results, &next, nil
}

func (c *Client) ListOrgMembers(org string) ([]OrgMember, error) {
	var res []OrgMember
	var opt PageOption
	for true {
		group, next, err := c.ListOrgMembersPage(org, &opt)
		if err != nil {
			return nil, err
		}
		res = append(res, group...)
		opt = *next
		if next.Empty() {
			break
		}
	}
	return res, nil
}

type inviteOrgMembersPayload struct {
	Org      string   `json:"org"`
	Team     string   `json:"team"`
	Invitees []string `json:"invitees"`
	DryRun   bool     `json:"dry_run"`
}

func (c *Client) InviteOrgMembers(org, group string, usernameOrEmail []string) error {
	payload, err := json.Marshal(inviteOrgMembersPayload{org, group, usernameOrEmail, false})
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/invites/bulk", c.HostURL), strings.NewReader(string(payload)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) AddOrgGroupMember(org, group, usernameOrEmail string) error {
	payload := fmt.Sprintf(`{"member": "%s"}`, usernameOrEmail)
	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/orgs/%s/groups/%s/members", c.HostURL, org, group), strings.NewReader(payload))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteOrgGroupMember(org, group, username string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/orgs/%s/groups/%s/members/%s", c.HostURL, org, group, username), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) DeleteOrgMember(org, username string) error {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/orgs/%s/members/%s", c.HostURL, org, username), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req)
	if err != nil {
		return err
	}
	return nil
}
