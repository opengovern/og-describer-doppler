package doppler

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDopplerWorkplace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_workplace",
		List: &plugin.ListConfig{
			Hydrate: nil,
		},
		Get: &plugin.GetConfig{
			Hydrate: nil,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the workplace."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the workplace."},
			{Name: "billing_email", Type: proto.ColumnType_STRING, Description: "The billing email associated with the workplace."},
			{Name: "security_email", Type: proto.ColumnType_STRING, Description: "The security email associated with the workplace."},
		},
	}
}
