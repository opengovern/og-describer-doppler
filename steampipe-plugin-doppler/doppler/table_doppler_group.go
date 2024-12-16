package doppler

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_group",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the group."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug identifier for the group."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the group was created."},
			{Name: "default_project_role", Type: proto.ColumnType_JSON, Description: "The default project role associated with the group."},
		},
	}
}