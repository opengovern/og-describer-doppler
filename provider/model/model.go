//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

import "time"

type Metadata struct{}

type ProjectListResponse struct {
	Page     int                  `json:"page"`
	Projects []ProjectDescription `json:"projects"`
}

type ProjectGetResponse struct {
	Project ProjectDescription `json:"project"`
}

type ProjectDescription struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type ProjectRoleListResponse struct {
	Roles []ProjectRoleDescription `json:"roles"`
}

type ProjectRoleGetResponse struct {
	Role ProjectRoleDescription `json:"role"`
}

type ProjectRoleDescription struct {
	Name         string    `json:"name"`
	Permissions  []string  `json:"permissions"`
	Identifier   string    `json:"identifier"`
	CreatedAt    time.Time `json:"created_at"`
	IsCustomRole bool      `json:"is_custom_role"`
}

type Role struct {
	Identifier string `json:"identifier"`
}

type ProjectMemberListResponse struct {
	Members []ProjectMemberDescription `json:"members"`
}

type ProjectMemberDescription struct {
	Type                  string  `json:"type"`
	Slug                  string  `json:"slug"`
	Role                  Role    `json:"role"`
	AccessAllEnvironments bool    `json:"access_all_environments"`
	Environments          *string `json:"environments"`
}

type EnvironmentListResponse struct {
	Environments []EnvironmentDescription `json:"environments"`
	Page         int                      `json:"page"`
}

type EnvironmentDescription struct {
	ID             string     `json:"id"`
	Slug           string     `json:"slug"`
	Name           string     `json:"name"`
	InitialFetchAt *time.Time `json:"initial_fetch_at"`
	CreatedAt      time.Time  `json:"created_at"`
	Project        string     `json:"project"`
}

type ConfigListResponse struct {
	Configs []ConfigDescription `json:"configs"`
	Page    int                 `json:"page"`
}

type ConfigDescription struct {
	Name           string     `json:"name"`
	Root           bool       `json:"root"`
	Inheritable    bool       `json:"inheritable"`
	Inheriting     bool       `json:"inheriting"`
	Inherits       []string   `json:"inherits"`
	Locked         bool       `json:"locked"`
	InitialFetchAt *time.Time `json:"initial_fetch_at"`
	LastFetchAt    *time.Time `json:"last_fetch_at"`
	CreatedAt      time.Time  `json:"created_at"`
	Environment    string     `json:"environment"`
	Project        string     `json:"project"`
	Slug           string     `json:"slug"`
}

type ValueType struct {
	Type string `json:"type"`
}

type SecretListResponse struct {
	Secrets map[string]SecretDescription `json:"secrets"`
}

type SecretDescription struct {
	Raw                string    `json:"raw"`
	Computed           string    `json:"computed"`
	Note               string    `json:"note"`
	RawVisibility      string    `json:"rawVisibility"`
	ComputedVisibility string    `json:"computedVisibility"`
	RawValueType       ValueType `json:"rawValueType"`
	ComputedValueType  ValueType `json:"computedValueType"`
}

type Sync struct {
	Slug         string `json:"slug"`
	Enabled      bool   `json:"enabled"`
	LastSyncedAt string `json:"lastSyncedAt"`
	Project      string `json:"project"`
	Config       string `json:"config"`
	Integration  string `json:"integration"`
}

type IntegrationListResponse struct {
	Integrations []IntegrationDescription `json:"integrations"`
}

type IntegrationGetResponse struct {
	Integration IntegrationDescription `json:"integration"`
}

type IntegrationDescription struct {
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Kind    string `json:"kind"`
	Enabled bool   `json:"enabled"`
	Syncs   []Sync `json:"syncs,omitempty"`
}

type TrustIPListResponse struct {
	IPs []string `json:"ips"`
}

type TrustIPDescription struct {
	IP string `json:"ip"`
}

type ServiceTokenListResponse struct {
	Tokens []ServiceTokenDescription `json:"tokens"`
}

type ServiceTokenDescription struct {
	Name        string     `json:"name"`
	Slug        string     `json:"slug"`
	CreatedAt   time.Time  `json:"created_at"`
	Config      string     `json:"config"`
	Environment string     `json:"environment"`
	Project     string     `json:"project"`
	ExpiresAt   *time.Time `json:"expires_at"`
}

type DefaultProjectRole struct {
	Identifier string `json:"identifier"`
}

type GroupListResponse struct {
	Groups []GroupDescription `json:"groups"`
}

type GroupGetResponse struct {
	Group GroupDescription `json:"group"`
}

type GroupDescription struct {
	Name               string             `json:"name"`
	Slug               string             `json:"slug"`
	CreatedAt          time.Time          `json:"created_at"`
	DefaultProjectRole DefaultProjectRole `json:"default_project_role"`
}

type WorkplaceRole struct {
	Name         string    `json:"name"`
	Permissions  []string  `json:"permissions"`
	Identifier   string    `json:"identifier"`
	CreatedAt    time.Time `json:"created_at"`
	IsCustomRole bool      `json:"is_custom_role"`
	IsInlineRole bool      `json:"is_inline_role"`
}

type ServiceAccountListResponse struct {
	ServiceAccounts []ServiceAccountDescription `json:"service_accounts"`
}

type ServiceAccountGetResponse struct {
	ServiceAccount ServiceAccountDescription `json:"service_account"`
}

type ServiceAccountDescription struct {
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	CreatedAt     time.Time     `json:"created_at"`
	WorkplaceRole WorkplaceRole `json:"workplace_role"`
}

type ServiceAccountTokenListResponse struct {
	APITokens []ServiceAccountTokenDescription `json:"api_tokens"`
}

type ServiceAccountTokenDescription struct {
	Name       string    `json:"name"`
	Slug       string    `json:"slug"`
	CreatedAt  time.Time `json:"created_at"`
	LastSeenAt time.Time `json:"last_seen_at"`
	ExpiresAt  time.Time `json:"expires_at"`
}

type WorkplaceListResponse struct {
	Workplace WorkplaceDescription `json:"workplace"`
}

type WorkplaceDescription struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	BillingEmail  string `json:"billing_email"`
	SecurityEmail string `json:"security_email"`
}

type User struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}

type WorkplaceUserListResponse struct {
	WorkplaceUsers []WorkplaceUserDescription `json:"workplace_users"`
	Page           int                        `json:"page"`
}

type WorkplaceUserGetResponse struct {
	WorkplaceUser WorkplaceUserDescription `json:"workplace_user"`
}

type WorkplaceUserDescription struct {
	ID        string `json:"id"`
	Access    string `json:"access"`
	CreatedAt string `json:"created_at"`
	User      User   `json:"user"`
}

type WorkplaceRoleListResponse struct {
	Roles []WorkplaceRoleDescription `json:"roles"`
}

type WorkplaceRoleGetResponse struct {
	Role WorkplaceRoleDescription `json:"role"`
}

type WorkplaceRoleDescription struct {
	Name         string   `json:"name"`
	Permissions  []string `json:"permissions"`
	Identifier   string   `json:"identifier"`
	CreatedAt    string   `json:"created_at"`
	IsCustomRole bool     `json:"is_custom_role"`
	IsInlineRole bool     `json:"is_inline_role"`
}
