package doppler

import (
	"context"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerWorkplace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_workplace",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWorkplace,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetWorkplace,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "The unique identifier for the workplace."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the workplace."},
			{Name: "billing_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.BillingEmail"), Description: "The billing email associated with the workplace."},
			{Name: "security_email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.SecurityEmail"), Description: "The security email associated with the workplace."},
		}),
	}
}
