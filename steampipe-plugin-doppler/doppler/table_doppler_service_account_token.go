package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerServiceAccountToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_service_account_token",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListServiceAccountToken,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetServiceAccountToken,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the service account token."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug identifier for the service account token."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the service account token was created."},
			{Name: "last_seen_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the service account token was last seen."},
			{Name: "expires_at", Type: proto.ColumnType_TIMESTAMP, Description: "The expiration timestamp of the service account token."},
		}),
	}
}
