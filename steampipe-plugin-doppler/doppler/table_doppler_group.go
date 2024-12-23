package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_group",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListGroup,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    opengovernance.GetGroup,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the group."},
			{Name: "slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Slug"), Description: "The slug identifier for the group."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The timestamp when the group was created."},
			{Name: "default_project_role", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.DefaultProjectRole"), Description: "The default project role associated with the group."},
		}),
	}
}
