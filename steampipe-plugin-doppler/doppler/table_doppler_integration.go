package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerIntegration(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_integration",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListIntegration,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    opengovernance.GetIntegration,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Slug"), Description: "The slug identifier for the integration."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the integration."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "The type of the integration."},
			{Name: "kind", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Kind"), Description: "The kind or category of the integration."},
			{Name: "enabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Enabled"), Description: "Indicates whether the integration is enabled."},
			{Name: "syncs", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Syncs"), Description: "The list of syncs associated with the integration."},
		}),
	}
}
