package provider

import (
	model "github.com/opengovern/og-describer-doppler/pkg/sdk/models"
	"github.com/opengovern/og-describer-doppler/provider/configs"
	"github.com/opengovern/og-describer-doppler/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"Doppler/Config": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/Config",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListConfigs),
		GetDescriber:    nil,
	},

	"Doppler/Environment": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/Environment",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListEnvironments),
		GetDescriber:    nil,
	},

	"Doppler/Group": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/Group",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListGroups),
		GetDescriber:    DescribeSingleByDoppler(describer.GetGroup),
	},

	"Doppler/Integration": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/Integration",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListIntegrations),
		GetDescriber:    DescribeSingleByDoppler(describer.GetIntegration),
	},

	"Doppler/Project/Member": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/Project/Member",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListProjectMembers),
		GetDescriber:    nil,
	},

	"Doppler/Project/Role": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/Project/Role",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListProjectRoles),
		GetDescriber:    DescribeSingleByDoppler(describer.GetProjectRole),
	},

	"Doppler/Project": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/Project",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListProjects),
		GetDescriber:    DescribeSingleByDoppler(describer.GetProject),
	},

	"Doppler/Secret": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/Secret",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListSecrets),
		GetDescriber:    nil,
	},

	"Doppler/ServiceAccount": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/ServiceAccount",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListServiceAccounts),
		GetDescriber:    DescribeSingleByDoppler(describer.GetServiceAccount),
	},

	"Doppler/ServiceAccount/Token": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/ServiceAccount/Token",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListServiceAccountTokens),
		GetDescriber:    nil,
	},

	"Doppler/ServiceToken": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/ServiceToken",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListServiceTokens),
		GetDescriber:    nil,
	},

	"Doppler/TrustIP": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/TrustIP",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListIPs),
		GetDescriber:    nil,
	},

	"Doppler/WorkPlace": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/WorkPlace",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListWorkplaces),
		GetDescriber:    nil,
	},

	"Doppler/WorkPlace/Role": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/WorkPlace/Role",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListWorkplaceRoles),
		GetDescriber:    DescribeSingleByDoppler(describer.GetWorkplaceRole),
	},

	"Doppler/WorkPlace/User": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "Doppler/WorkPlace/User",
		Tags:            map[string][]string{},
		Labels:          map[string]string{},
		Annotations:     map[string]string{},
		ListDescriber:   DescribeListByDoppler(describer.ListWorkplaceUsers),
		GetDescriber:    DescribeSingleByDoppler(describer.GetWorkplaceUser),
	},
}
