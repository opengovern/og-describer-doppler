package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerWorkplaceRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_workplace_role",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWorkplaceRole,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetWorkplaceRole,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the workplace role."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Description: "The list of permissions associated with the workplace role."},
			{Name: "identifier", Type: proto.ColumnType_STRING, Description: "The unique identifier for the workplace role."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the workplace role was created."},
			{Name: "is_custom_role", Type: proto.ColumnType_BOOL, Description: "Indicates whether the role is a custom role."},
			{Name: "is_inline_role", Type: proto.ColumnType_BOOL, Description: "Indicates whether the role is an inline role."},
		},
	}
}
