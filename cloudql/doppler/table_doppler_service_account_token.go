package doppler

import (
	"context"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerServiceAccountToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_service_account_token",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListServiceAccountToken,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    opengovernance.GetServiceAccountToken,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the service account token."},
			{Name: "slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Slug"), Description: "The slug identifier for the service account token."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The timestamp when the service account token was created."},
			{Name: "last_seen_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LastSeenAt"), Description: "The timestamp when the service account token was last seen."},
			{Name: "expires_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ExpiresAt"), Description: "The expiration timestamp of the service account token."},
		}),
	}
}
