package configs

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "doppler"                                    // example: aws, azure
	IntegrationName      = integration.Type("DOPPLER_ACCOUNT")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describer-doppler" // example: github.com/opengovern/og-describer-aws
)
