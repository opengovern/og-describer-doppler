package doppler

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerServiceAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_service_account",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the service account."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug identifier for the service account."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "The timestamp when the service account was created."},
			{Name: "workplace_role", Type: proto.ColumnType_JSON, Description: "The workplace role associated with the service account."},
		},
	}
}