package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerServiceToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_service_token",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListServiceToken,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetServiceToken,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the service token."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug identifier for the service token."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the service token was created."},
			{Name: "config", Type: proto.ColumnType_STRING, Description: "The configuration associated with the service token."},
			{Name: "environment", Type: proto.ColumnType_STRING, Description: "The environment associated with the service token."},
			{Name: "project", Type: proto.ColumnType_STRING, Description: "The project associated with the service token."},
			{Name: "expires_at", Type: proto.ColumnType_TIMESTAMP, Description: "The expiration timestamp of the service token."},
		},
	}
}
