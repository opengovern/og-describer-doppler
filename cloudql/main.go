package main

import (
	"github.com/opengovern/og-describer-doppler/cloudql/doppler"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: doppler.Plugin})
}
