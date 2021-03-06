package models

import (
	"fmt"
	"time"

	jira "Atlassian/jira"
)

type MembersFromGroup struct {
	Values     []jira.User `json:"values" structs:"values"`
	StartAt    int         `json:"startAt" structs:"startAt"`
	MaxResults int         `json:"maxResults" structs:"maxResults"`
	Total      int         `json:"total" structs:"total"`
}

type ConfigType int

const (
	JIRA ConfigType = iota
	GMAIL
	SDESK
)

type Config struct {
	JiraUsername       string
	JiraApiToken       string
	JiraAddress        string
	JiraProjectGroups  string
	JiraProfilePage    string
	JiraNextButton     string
	JiraUseCaseJQL     string
	JiraSearchBySubj   string
	JiraOrganizationId string
	GmailDomainParts   string
	GmailAdminEmail    string
	GmailAdminPassword string
	SDeskIssueTypeName string
	SDeskProjectKey    string
}

type User struct {
	AccountID    string
	DisplayName  string
	EmailAddress string
	LastLogin    time.Time
}

func (u User) String() string {
	return fmt.Sprint(u.AccountID, u.DisplayName, u.EmailAddress)
}
