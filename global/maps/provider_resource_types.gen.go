package maps

import (
	"github.com/opengovern/og-describer-doppler/discovery/describers"
	model "github.com/opengovern/og-describer-doppler/discovery/pkg/models"
	"github.com/opengovern/og-describer-doppler/discovery/provider"
	"github.com/opengovern/og-describer-doppler/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
)

var ResourceTypes = map[string]model.ResourceType{

	"Doppler/Config": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/Config",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListConfigs),
		GetDescriber:    nil,
	},

	"Doppler/Environment": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/Environment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListEnvironments),
		GetDescriber:    nil,
	},

	"Doppler/Group": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/Group",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListGroups),
		GetDescriber:    provider.DescribeSingleByDoppler(describers.GetGroup),
	},

	"Doppler/Integration": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/Integration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListIntegrations),
		GetDescriber:    provider.DescribeSingleByDoppler(describers.GetIntegration),
	},

	"Doppler/Project/Member": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/Project/Member",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListProjectMembers),
		GetDescriber:    nil,
	},

	"Doppler/Project/Role": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/Project/Role",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListProjectRoles),
		GetDescriber:    provider.DescribeSingleByDoppler(describers.GetProjectRole),
	},

	"Doppler/Project": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/Project",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListProjects),
		GetDescriber:    provider.DescribeSingleByDoppler(describers.GetProject),
	},

	"Doppler/Secret": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/Secret",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListSecrets),
		GetDescriber:    nil,
	},

	"Doppler/ServiceAccount": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/ServiceAccount",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListServiceAccounts),
		GetDescriber:    provider.DescribeSingleByDoppler(describers.GetServiceAccount),
	},

	"Doppler/ServiceAccount/Token": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/ServiceAccount/Token",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListServiceAccountTokens),
		GetDescriber:    nil,
	},

	"Doppler/ServiceToken": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/ServiceToken",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListServiceTokens),
		GetDescriber:    nil,
	},

	"Doppler/TrustIP": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/TrustIP",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListIPs),
		GetDescriber:    nil,
	},

	"Doppler/WorkPlace": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/WorkPlace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListWorkplaces),
		GetDescriber:    nil,
	},

	"Doppler/WorkPlace/Role": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/WorkPlace/Role",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListWorkplaceRoles),
		GetDescriber:    provider.DescribeSingleByDoppler(describers.GetWorkplaceRole),
	},

	"Doppler/WorkPlace/User": {
		IntegrationType: constants.IntegrationName,
		ResourceName:    "Doppler/WorkPlace/User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describers.ListWorkplaceUsers),
		GetDescriber:    provider.DescribeSingleByDoppler(describers.GetWorkplaceUser),
	},
}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"Doppler/Config": {
		Name:            "Doppler/Config",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/Environment": {
		Name:            "Doppler/Environment",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/Group": {
		Name:            "Doppler/Group",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/Integration": {
		Name:            "Doppler/Integration",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/Project/Member": {
		Name:            "Doppler/Project/Member",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/Project/Role": {
		Name:            "Doppler/Project/Role",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/Project": {
		Name:            "Doppler/Project",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/Secret": {
		Name:            "Doppler/Secret",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/ServiceAccount": {
		Name:            "Doppler/ServiceAccount",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/ServiceAccount/Token": {
		Name:            "Doppler/ServiceAccount/Token",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/ServiceToken": {
		Name:            "Doppler/ServiceToken",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/TrustIP": {
		Name:            "Doppler/TrustIP",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/WorkPlace": {
		Name:            "Doppler/WorkPlace",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/WorkPlace/Role": {
		Name:            "Doppler/WorkPlace/Role",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},

	"Doppler/WorkPlace/User": {
		Name:            "Doppler/WorkPlace/User",
		IntegrationType: constants.IntegrationName,
		Description:     "",
	},
}

var ResourceTypesList = []string{
	"Doppler/Config",
	"Doppler/Environment",
	"Doppler/Group",
	"Doppler/Integration",
	"Doppler/Project/Member",
	"Doppler/Project/Role",
	"Doppler/Project",
	"Doppler/Secret",
	"Doppler/ServiceAccount",
	"Doppler/ServiceAccount/Token",
	"Doppler/ServiceToken",
	"Doppler/TrustIP",
	"Doppler/WorkPlace",
	"Doppler/WorkPlace/Role",
	"Doppler/WorkPlace/User",
}
