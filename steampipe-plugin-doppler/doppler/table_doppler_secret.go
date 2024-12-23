package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_secret",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListSecret,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("computed"),
			Hydrate:    opengovernance.GetSecret,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "raw", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Raw"), Description: "The raw value of the secret."},
			{Name: "computed", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Computed"), Description: "The computed value of the secret."},
			{Name: "note", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Note"), Description: "A note associated with the secret."},
			{Name: "raw_visibility", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.RawVisibility"), Description: "The visibility of the raw value of the secret."},
			{Name: "computed_visibility", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ComputedVisibility"), Description: "The visibility of the computed value of the secret."},
			{Name: "raw_value_type", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.RawValueType"), Description: "The type of the raw value of the secret."},
			{Name: "computed_value_type", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.ComputedValueType"), Description: "The type of the computed value of the secret."},
		}),
	}
}
