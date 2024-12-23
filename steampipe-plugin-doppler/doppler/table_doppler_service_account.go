package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerServiceAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_service_account",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListServiceAccount,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("slug"),
			Hydrate:    opengovernance.GetServiceAccount,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the service account."},
			{Name: "slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Slug"), Description: "The slug identifier for the service account."},
			{Name: "created_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CreatedAt"), Description: "The timestamp when the service account was created."},
			{Name: "workplace_role", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.WorkplaceRole"), Description: "The workplace role associated with the service account."},
		}),
	}
}
