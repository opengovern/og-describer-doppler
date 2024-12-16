package doppler

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerEnvironment(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_environment",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the environment."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug of the environment."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the environment."},
			{Name: "initial_fetch_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp of the initial fetch for the environment."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the environment was created."},
			{Name: "project", Type: proto.ColumnType_STRING, Description: "The project associated with the environment."},
		},
	}
}
