package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerProject(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_project",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProject,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetProject,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the project."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug of the project."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the project."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "The description of the project."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the project was created."},
		}),
	}
}
