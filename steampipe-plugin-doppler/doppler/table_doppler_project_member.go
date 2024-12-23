package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerProjectMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_project_member",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProjectMember,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetProjectMember,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Description: "The type of the project member."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug of the project member."},
			{Name: "role", Type: proto.ColumnType_JSON, Description: "The role assigned to the project member."},
			{Name: "access_all_environments", Type: proto.ColumnType_BOOL, Description: "Indicates whether the member has access to all environments."},
			{Name: "environments", Type: proto.ColumnType_STRING, Description: "The environments the member has access to."},
		}),
	}
}
