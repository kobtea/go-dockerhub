package dockerhub

import "time"

type Organization struct {
	ID            string    `json:"id"`
	OrgName       string    `json:"orgname"`
	FullName      string    `json:"full_name"`
	Location      string    `json:"location"`
	Company       string    `json:"company"`
	ProfileURL    string    `json:"profile_url"`
	DateJoined    time.Time `json:"date_joined"`
	GravatarURL   string    `json:"gravatar_url"`
	GravatarEmail string    `json:"gravatar_email"`
	Type          string    `json:"type"`
	Badge         string    `json:"badge"`
	IsActive      bool      `json:"is_active"`
}

type OrgGroup struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	MemberCount int    `json:"member_count"`
}

type OrgMember struct {
	ID            string    `json:"id"`
	UUID          string    `json:"uuid"`
	Username      string    `json:"username"`
	FullName      string    `json:"full_name"`
	Location      string    `json:"location"`
	Company       string    `json:"company"`
	ProfileURL    string    `json:"profile_url"`
	DateJoined    time.Time `json:"date_joined"`
	GravatarUrl   string    `json:"gravatar_url"`
	GravatarEmail string    `json:"gravatar_email"`
	Type          string    `json:"type"`
	IsAdmin       bool      `json:"is_admin"`
	IsStaff       bool      `json:"is_staff"`
	Email         string    `json:"email"`
	Role          string    `json:"role"`
	Groups        []string  `json:"groups"`
	IsGuest       bool      `json:"is_guest"`
	PrimaryEmail  string    `json:"primary_email"`
}
