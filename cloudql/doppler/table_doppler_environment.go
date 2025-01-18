package doppler

import (
	"context"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerEnvironment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_environment",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListEnvironment,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    opengovernance.GetEnvironment,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the environment."},
			{Name: "slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Slug"), Description: "The slug of the environment."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the environment."},
			{Name: "initial_fetch_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.InitialFetchAt"), Description: "The timestamp of the initial fetch for the environment."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The timestamp when the environment was created."},
			{Name: "project", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Project"), Description: "The project associated with the environment."},
		}),
	}
}
