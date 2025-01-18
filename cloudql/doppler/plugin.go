package doppler

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-doppler",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"doppler_config":                tableDopplerConfig(ctx),
			"doppler_environment":           tableDopplerEnvironment(ctx),
			"doppler_group":                 tableDopplerGroup(ctx),
			"doppler_integration":           tableDopplerIntegration(ctx),
			"doppler_project":               tableDopplerProject(ctx),
			"doppler_project_member":        tableDopplerProjectMember(ctx),
			"doppler_project_role":          tableDopplerProjectRole(ctx),
			"doppler_secret":                tableDopplerSecret(ctx),
			"doppler_service_account":       tableDopplerServiceAccount(ctx),
			"doppler_service_account_token": tableDopplerServiceAccountToken(ctx),
			"doppler_service_token":         tableDopplerServiceToken(ctx),
			"doppler_trust_ip":              tableDopplerTrustIP(ctx),
			"doppler_workplace":             tableDopplerWorkplace(ctx),
			"doppler_workplace_role":        tableDopplerWorkplaceRole(ctx),
			"doppler_workplace_user":        tableDopplerWorkplaceUser(ctx),
		},
	}
	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "platform_account_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}
		}
	}
	return p
}
