package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerWorkplaceUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_workplace_user",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWorkplaceUser,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetWorkplaceUser,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the workplace user."},
			{Name: "access", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Access"), Description: "The access level of the user within the workplace."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The timestamp when the workplace user was created."},
			{Name: "user", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.User"), Description: "Details of the user associated with the workplace."},
		}),
	}
}
