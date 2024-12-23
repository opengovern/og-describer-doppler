package doppler

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-doppler/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDopplerTrustIP(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name: "doppler_trust_ip",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListTrustIP,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("ip"),
			Hydrate:    opengovernance.GetTrustIP,
		},
		Columns: integrationColumns([]*plugin.Column{
			{Name: "ip", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.IP"), Description: "The trusted IP address."},
		}),
	}
}
