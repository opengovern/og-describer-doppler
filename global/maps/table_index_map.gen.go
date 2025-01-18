package maps

import (
	"github.com/opengovern/og-describer-doppler/discovery/pkg/es"
)

var ResourceTypesToTables = map[string]string{
	"Doppler/Config":               "doppler_config",
	"Doppler/Environment":          "doppler_environment",
	"Doppler/Group":                "doppler_group",
	"Doppler/Integration":          "doppler_integration",
	"Doppler/Project/Member":       "doppler_project_member",
	"Doppler/Project/Role":         "doppler_project_role",
	"Doppler/Project":              "doppler_project",
	"Doppler/Secret":               "doppler_secret",
	"Doppler/ServiceAccount":       "doppler_service_account",
	"Doppler/ServiceAccount/Token": "doppler_service_account_token",
	"Doppler/ServiceToken":         "doppler_service_token",
	"Doppler/TrustIP":              "doppler_trust_ip",
	"Doppler/WorkPlace":            "doppler_workplace",
	"Doppler/WorkPlace/Role":       "doppler_workplace_role",
	"Doppler/WorkPlace/User":       "doppler_workplace_user",
}

var ResourceTypeToDescription = map[string]interface{}{
	"Doppler/Config":               opengovernance.Config{},
	"Doppler/Environment":          opengovernance.Environment{},
	"Doppler/Group":                opengovernance.Group{},
	"Doppler/Integration":          opengovernance.Integration{},
	"Doppler/Project/Member":       opengovernance.ProjectMember{},
	"Doppler/Project/Role":         opengovernance.ProjectRole{},
	"Doppler/Project":              opengovernance.Project{},
	"Doppler/Secret":               opengovernance.Secret{},
	"Doppler/ServiceAccount":       opengovernance.ServiceAccount{},
	"Doppler/ServiceAccount/Token": opengovernance.ServiceAccountToken{},
	"Doppler/ServiceToken":         opengovernance.ServiceToken{},
	"Doppler/TrustIP":              opengovernance.TrustIP{},
	"Doppler/WorkPlace":            opengovernance.Workplace{},
	"Doppler/WorkPlace/Role":       opengovernance.WorkplaceRole{},
	"Doppler/WorkPlace/User":       opengovernance.WorkplaceUser{},
}

var TablesToResourceTypes = map[string]string{
	"doppler_config":                "Doppler/Config",
	"doppler_environment":           "Doppler/Environment",
	"doppler_group":                 "Doppler/Group",
	"doppler_integration":           "Doppler/Integration",
	"doppler_project_member":        "Doppler/Project/Member",
	"doppler_project_role":          "Doppler/Project/Role",
	"doppler_project":               "Doppler/Project",
	"doppler_secret":                "Doppler/Secret",
	"doppler_service_account":       "Doppler/ServiceAccount",
	"doppler_service_account_token": "Doppler/ServiceAccount/Token",
	"doppler_service_token":         "Doppler/ServiceToken",
	"doppler_trust_ip":              "Doppler/TrustIP",
	"doppler_workplace":             "Doppler/WorkPlace",
	"doppler_workplace_role":        "Doppler/WorkPlace/Role",
	"doppler_workplace_user":        "Doppler/WorkPlace/User",
}
