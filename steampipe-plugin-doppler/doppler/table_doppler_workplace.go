package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerWorkplace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_workplace",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWorkplace,
		},
		Get: &plugin.GetConfig{
			Hydrate: opengovernance.GetWorkplace,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the workplace."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the workplace."},
			{Name: "billing_email", Type: proto.ColumnType_STRING, Description: "The billing email associated with the workplace."},
			{Name: "security_email", Type: proto.ColumnType_STRING, Description: "The security email associated with the workplace."},
		}),
	}
}
