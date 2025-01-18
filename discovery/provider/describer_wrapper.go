package provider

import (
	"errors"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/models"
	"github.com/opengovern/og-util/pkg/describe/enums"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"github.com/opengovern/resilient-bridge/adapters"
	"golang.org/x/net/context"
	"time"
)

// DescribeListByDoppler A wrapper to pass doppler authorization to describers functions
func DescribeListByDoppler(describe func(context.Context, *resilientbridge.ResilientBridge, *models.StreamSender) ([]models.Resource, error)) models.ResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *models.StreamSender) ([]models.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		resilientBridge := resilientbridge.NewResilientBridge()

		restMaxRequests := 500
		restWindowSecs := int64(60)

		// Register Doppler provider
		resilientBridge.RegisterProvider("doppler", &adapters.DopplerAdapter{APIToken: cfg.Token}, &resilientbridge.ProviderConfig{
			UseProviderLimits:   true,
			MaxRequestsOverride: &restMaxRequests,
			WindowSecsOverride:  &restWindowSecs,
			MaxRetries:          3,
			BaseBackoff:         200 * time.Millisecond,
		})

		// Get values from describers
		var values []models.Resource
		result, err := describe(ctx, resilientBridge, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByDoppler A wrapper to pass doppler authorization to describers functions
func DescribeSingleByDoppler(describe func(context.Context, *resilientbridge.ResilientBridge, string) (*models.Resource, error)) models.SingleResourceDescriber {
	return func(ctx context.Context, cfg models.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string, stream *models.StreamSender) (*models.Resource, error) {
		ctx = WithTriggerType(ctx, triggerType)

		var err error
		// Check for the token
		if cfg.Token == "" {
			return nil, errors.New("token must be configured")
		}

		resilientBridge := resilientbridge.NewResilientBridge()

		restMaxRequests := 500
		restWindowSecs := int64(60)

		// Register Doppler provider
		resilientBridge.RegisterProvider("doppler", &adapters.DopplerAdapter{APIToken: cfg.Token}, &resilientbridge.ProviderConfig{
			UseProviderLimits:   true,
			MaxRequestsOverride: &restMaxRequests,
			WindowSecsOverride:  &restWindowSecs,
			MaxRetries:          3,
			BaseBackoff:         200 * time.Millisecond,
		})

		// Get value from describers
		value, err := describe(ctx, resilientBridge, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
