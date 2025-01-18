package doppler

import (
	"context"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerProjectMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_project_member",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListProjectMember,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    opengovernance.GetProjectMember,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "The type of the project member."},
			{Name: "slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Slug"), Description: "The slug of the project member."},
			{Name: "role", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Role"), Description: "The role assigned to the project member."},
			{Name: "access_all_environments", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.AccessAllEnvironments"), Description: "Indicates whether the member has access to all environments."},
			{Name: "environments", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Environments"), Description: "The environments the member has access to."},
		}),
	}
}
