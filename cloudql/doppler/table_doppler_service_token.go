package doppler

import (
	"context"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerServiceToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_service_token",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListServiceToken,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    opengovernance.GetServiceToken,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the service token."},
			{Name: "slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Slug"), Description: "The slug identifier for the service token."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The timestamp when the service token was created."},
			{Name: "config", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Config"), Description: "The configuration associated with the service token."},
			{Name: "environment", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Environment"), Description: "The environment associated with the service token."},
			{Name: "project", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Project"), Description: "The project associated with the service token."},
			{Name: "expires_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ExpiresAt"), Description: "The expiration timestamp of the service token."},
		}),
	}
}
