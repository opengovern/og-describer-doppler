package global

import "github.com/opengovern/og-util/pkg/integration"

const (
	IntegrationTypeLower = "doppler"                                     // example: aws, azure
	IntegrationName      = integration.Type("doppler_account")           // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL      = "github.com/opengovern/og-describers-doppler" // example: github.com/opengovern/og-describers-aws
)

type IntegrationCredentials struct {
	Token string `json:"token"`
}
