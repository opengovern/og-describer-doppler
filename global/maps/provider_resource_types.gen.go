package maps

import (
	describer2 "github.com/opengovern/og-describer-doppler/discovery/describers"
	model "github.com/opengovern/og-describer-doppler/discovery/pkg/models"
	"github.com/opengovern/og-describer-doppler/discovery/provider"
	"github.com/opengovern/og-describer-doppler/global"
)

var ResourceTypes = map[string]model.ResourceType{

	"Doppler/Config": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/Config",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListConfigs),
		GetDescriber:    nil,
	},

	"Doppler/Environment": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/Environment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListEnvironments),
		GetDescriber:    nil,
	},

	"Doppler/Group": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/Group",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListGroups),
		GetDescriber:    provider.DescribeSingleByDoppler(describer2.GetGroup),
	},

	"Doppler/Integration": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/Integration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListIntegrations),
		GetDescriber:    provider.DescribeSingleByDoppler(describer2.GetIntegration),
	},

	"Doppler/Project/Member": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/Project/Member",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListProjectMembers),
		GetDescriber:    nil,
	},

	"Doppler/Project/Role": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/Project/Role",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListProjectRoles),
		GetDescriber:    provider.DescribeSingleByDoppler(describer2.GetProjectRole),
	},

	"Doppler/Project": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/Project",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListProjects),
		GetDescriber:    provider.DescribeSingleByDoppler(describer2.GetProject),
	},

	"Doppler/Secret": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/Secret",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListSecrets),
		GetDescriber:    nil,
	},

	"Doppler/ServiceAccount": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/ServiceAccount",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListServiceAccounts),
		GetDescriber:    provider.DescribeSingleByDoppler(describer2.GetServiceAccount),
	},

	"Doppler/ServiceAccount/Token": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/ServiceAccount/Token",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListServiceAccountTokens),
		GetDescriber:    nil,
	},

	"Doppler/ServiceToken": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/ServiceToken",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListServiceTokens),
		GetDescriber:    nil,
	},

	"Doppler/TrustIP": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/TrustIP",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListIPs),
		GetDescriber:    nil,
	},

	"Doppler/WorkPlace": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/WorkPlace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListWorkplaces),
		GetDescriber:    nil,
	},

	"Doppler/WorkPlace/Role": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/WorkPlace/Role",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListWorkplaceRoles),
		GetDescriber:    provider.DescribeSingleByDoppler(describer2.GetWorkplaceRole),
	},

	"Doppler/WorkPlace/User": {
		IntegrationType: global.IntegrationName,
		ResourceName:    "Doppler/WorkPlace/User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   provider.DescribeListByDoppler(describer2.ListWorkplaceUsers),
		GetDescriber:    provider.DescribeSingleByDoppler(describer2.GetWorkplaceUser),
	},
}
