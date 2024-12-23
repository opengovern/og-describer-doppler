package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerConfig(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_config",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListConfig,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetConfig,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the configuration."},
			{Name: "root", Type: proto.ColumnType_BOOL, Description: "Indicates whether this is a root configuration."},
			{Name: "inheritable", Type: proto.ColumnType_BOOL, Description: "Indicates if the configuration is inheritable."},
			{Name: "inheriting", Type: proto.ColumnType_BOOL, Description: "Indicates if the configuration is inheriting from another."},
			{Name: "inherits", Type: proto.ColumnType_JSON, Description: "List of configurations this one inherits from."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Description: "Indicates if the configuration is locked."},
			{Name: "initial_fetch_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the initial fetch."},
			{Name: "last_fetch_at", Type: proto.ColumnType_TIMESTAMP, Description: "Timestamp of the last fetch."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "When the configuration was created."},
			{Name: "environment", Type: proto.ColumnType_STRING, Description: "The environment associated with the configuration."},
			{Name: "project", Type: proto.ColumnType_STRING, Description: "The project associated with the configuration."},
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "The slug identifier for the configuration."},
		}),
	}
}
