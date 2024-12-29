//go:generate go run ../../pkg/sdk/runable/steampipe_es_client_generator/main.go -pluginPath ../../steampipe-plugin-REPLACEME/REPLACEME -file $GOFILE -output ../../pkg/sdk/es/resources_clients.go -resourceTypesFile ../resource_types/resource-types.json

// Implement types for each resource

package model

type Metadata struct{}

type ProjectListResponse struct {
	Page     int           `json:"page"`
	Projects []ProjectJSON `json:"projects"`
}

type ProjectGetResponse struct {
	Project ProjectJSON `json:"project"`
}

type ProjectJSON struct {
	ID          string `json:"id"`
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
}

type ProjectDescription struct {
	ID          string
	Slug        string
	Name        string
	Description string
	CreatedAt   string
}

type ProjectRoleListResponse struct {
	Roles []ProjectRoleJSON `json:"roles"`
}

type ProjectRoleGetResponse struct {
	Role ProjectRoleJSON `json:"role"`
}

type ProjectRoleJSON struct {
	Name         string   `json:"name"`
	Permissions  []string `json:"permissions"`
	Identifier   string   `json:"identifier"`
	CreatedAt    string   `json:"created_at"`
	IsCustomRole bool     `json:"is_custom_role"`
}

type ProjectRoleDescription struct {
	Name         string
	Permissions  []string
	Identifier   string
	CreatedAt    string
	IsCustomRole bool
}

type RoleJSON struct {
	Identifier string `json:"identifier"`
}

type Role struct {
	Identifier string
}

type ProjectMemberListResponse struct {
	Members []ProjectMemberJSON `json:"members"`
}

type ProjectMemberJSON struct {
	Type                  string   `json:"type"`
	Slug                  string   `json:"slug"`
	Role                  RoleJSON `json:"role"`
	AccessAllEnvironments bool     `json:"access_all_environments"`
	Environments          *string  `json:"environments"`
}

type ProjectMemberDescription struct {
	Type                  string
	Slug                  string
	Role                  Role
	AccessAllEnvironments bool
	Environments          *string
}

type EnvironmentListResponse struct {
	Environments []EnvironmentJSON `json:"environments"`
	Page         int               `json:"page"`
}

type EnvironmentJSON struct {
	ID             string  `json:"id"`
	Slug           string  `json:"slug"`
	Name           string  `json:"name"`
	InitialFetchAt *string `json:"initial_fetch_at"`
	CreatedAt      string  `json:"created_at"`
	Project        string  `json:"project"`
}

type EnvironmentDescription struct {
	ID             string
	Slug           string
	Name           string
	InitialFetchAt *string
	CreatedAt      string
	Project        string
}

type ConfigListResponse struct {
	Configs []ConfigJSON `json:"configs"`
	Page    int          `json:"page"`
}

type ConfigJSON struct {
	Name           string   `json:"name"`
	Root           bool     `json:"root"`
	Inheritable    bool     `json:"inheritable"`
	Inheriting     bool     `json:"inheriting"`
	Inherits       []string `json:"inherits"`
	Locked         bool     `json:"locked"`
	InitialFetchAt *string  `json:"initial_fetch_at"`
	LastFetchAt    *string  `json:"last_fetch_at"`
	CreatedAt      string   `json:"created_at"`
	Environment    string   `json:"environment"`
	Project        string   `json:"project"`
	Slug           string   `json:"slug"`
}

type ConfigDescription struct {
	Name           string
	Root           bool
	Inheritable    bool
	Inheriting     bool
	Inherits       []string
	Locked         bool
	InitialFetchAt *string
	LastFetchAt    *string
	CreatedAt      string
	Environment    string
	Project        string
	Slug           string
}

type ValueTypeJSON struct {
	Type string `json:"type"`
}

type ValueType struct {
	Type string
}

type SecretListResponse struct {
	Secrets map[string]SecretJSON `json:"secrets"`
}

type SecretJSON struct {
	Raw                string        `json:"raw"`
	Computed           string        `json:"computed"`
	Note               string        `json:"note"`
	RawVisibility      string        `json:"rawVisibility"`
	ComputedVisibility string        `json:"computedVisibility"`
	RawValueType       ValueTypeJSON `json:"rawValueType"`
	ComputedValueType  ValueTypeJSON `json:"computedValueType"`
}

type SecretDescription struct {
	Raw                string
	Computed           string
	Note               string
	RawVisibility      string
	ComputedVisibility string
	RawValueType       ValueType
	ComputedValueType  ValueType
}

type SyncJSON struct {
	Slug         string `json:"slug"`
	Enabled      bool   `json:"enabled"`
	LastSyncedAt string `json:"lastSyncedAt"`
	Project      string `json:"project"`
	Config       string `json:"config"`
	Integration  string `json:"integration"`
}

type Sync struct {
	Slug         string
	Enabled      bool
	LastSyncedAt string
	Project      string
	Config       string
	Integration  string
}

type IntegrationListResponse struct {
	Integrations []IntegrationJSON `json:"integrations"`
}

type IntegrationGetResponse struct {
	Integration IntegrationJSON `json:"integration"`
}

type IntegrationJSON struct {
	Slug    string     `json:"slug"`
	Name    string     `json:"name"`
	Type    string     `json:"type"`
	Kind    string     `json:"kind"`
	Enabled bool       `json:"enabled"`
	Syncs   []SyncJSON `json:"syncs,omitempty"`
}

type IntegrationDescription struct {
	Slug    string
	Name    string
	Type    string
	Kind    string
	Enabled bool
	Syncs   []Sync
}

type TrustIPListResponse struct {
	IPs []string `json:"ips"`
}

type TrustIPDescription struct {
	IP string
}

type ServiceTokenListResponse struct {
	Tokens []ServiceTokenJSON `json:"tokens"`
}

type ServiceTokenJSON struct {
	Name        string  `json:"name"`
	Slug        string  `json:"slug"`
	CreatedAt   string  `json:"created_at"`
	Config      string  `json:"config"`
	Environment string  `json:"environment"`
	Project     string  `json:"project"`
	ExpiresAt   *string `json:"expires_at"`
}

type ServiceTokenDescription struct {
	Name        string
	Slug        string
	CreatedAt   string
	Config      string
	Environment string
	Project     string
	ExpiresAt   *string
}

type DefaultProjectRoleJSON struct {
	Identifier string `json:"identifier"`
}

type DefaultProjectRole struct {
	Identifier string
}

type GroupListResponse struct {
	Groups []GroupJSON `json:"groups"`
}

type GroupGetResponse struct {
	Group GroupJSON `json:"group"`
}

type GroupJSON struct {
	Name               string                 `json:"name"`
	Slug               string                 `json:"slug"`
	CreatedAt          string                 `json:"created_at"`
	DefaultProjectRole DefaultProjectRoleJSON `json:"default_project_role"`
}

type GroupDescription struct {
	Name               string
	Slug               string
	CreatedAt          string
	DefaultProjectRole DefaultProjectRole
}

type WorkplaceRoleJSON struct {
	Name         string   `json:"name"`
	Permissions  []string `json:"permissions"`
	Identifier   string   `json:"identifier"`
	CreatedAt    string   `json:"created_at"`
	IsCustomRole bool     `json:"is_custom_role"`
	IsInlineRole bool     `json:"is_inline_role"`
}

type WorkplaceRole struct {
	Name         string
	Permissions  []string
	Identifier   string
	CreatedAt    string
	IsCustomRole bool
	IsInlineRole bool
}

type ServiceAccountListResponse struct {
	ServiceAccounts []ServiceAccountJSON `json:"service_accounts"`
}

type ServiceAccountGetResponse struct {
	ServiceAccount ServiceAccountJSON `json:"service_account"`
}

type ServiceAccountJSON struct {
	Name          string            `json:"name"`
	Slug          string            `json:"slug"`
	CreatedAt     string            `json:"created_at"`
	WorkplaceRole WorkplaceRoleJSON `json:"workplace_role"`
}

type ServiceAccountDescription struct {
	Name          string
	Slug          string
	CreatedAt     string
	WorkplaceRole WorkplaceRole
}

type ServiceAccountTokenListResponse struct {
	APITokens []ServiceAccountTokenJSON `json:"api_tokens"`
}

type ServiceAccountTokenJSON struct {
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	CreatedAt  string `json:"created_at"`
	LastSeenAt string `json:"last_seen_at"`
	ExpiresAt  string `json:"expires_at"`
}

type ServiceAccountTokenDescription struct {
	Name       string
	Slug       string
	CreatedAt  string
	LastSeenAt string
	ExpiresAt  string
}

type WorkplaceListResponse struct {
	Workplace WorkplaceJSON `json:"workplace"`
}

type WorkplaceJSON struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	BillingEmail  string `json:"billing_email"`
	SecurityEmail string `json:"security_email"`
}

type WorkplaceDescription struct {
	ID            string
	Name          string
	BillingEmail  string
	SecurityEmail string
}

type UserJSON struct {
	Email           string `json:"email"`
	Name            string `json:"name"`
	Username        string `json:"username"`
	ProfileImageURL string `json:"profile_image_url"`
}

type User struct {
	Email           string
	Name            string
	Username        string
	ProfileImageURL string
}

type WorkplaceUserListResponse struct {
	WorkplaceUsers []WorkplaceUserJSON `json:"workplace_users"`
	Page           int                 `json:"page"`
}

type WorkplaceUserGetResponse struct {
	WorkplaceUser WorkplaceUserJSON `json:"workplace_user"`
}

type WorkplaceUserJSON struct {
	ID        string   `json:"id"`
	Access    string   `json:"access"`
	CreatedAt string   `json:"created_at"`
	User      UserJSON `json:"user"`
}

type WorkplaceUserDescription struct {
	ID        string
	Access    string
	CreatedAt string
	User      User
}

type WorkplaceRoleListResponse struct {
	Roles []WorkplaceRoleJSON `json:"roles"`
}

type WorkplaceRoleGetResponse struct {
	Role WorkplaceRoleJSON `json:"role"`
}

type WorkplaceRoleDescription struct {
	Name         string
	Permissions  []string
	Identifier   string
	CreatedAt    string
	IsCustomRole bool
	IsInlineRole bool
}
