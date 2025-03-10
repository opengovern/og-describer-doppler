package doppler

import (
	"context"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerProjectRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_project_role",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProjectRole,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("identifier"),
			Hydrate:    opengovernance.GetProjectRole,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the project role."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Permissions"), Description: "The list of permissions associated with the role."},
			{Name: "identifier", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Identifier"), Description: "The unique identifier for the project role."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The timestamp when the project role was created."},
			{Name: "is_custom_role", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.IsCustomRole"), Description: "Indicates whether the role is a custom role."},
		}),
	}
}
