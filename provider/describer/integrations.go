package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-doppler/pkg/sdk/models"
	"github.com/opengovern/og-describer-doppler/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"net/url"
	"sync"
)

func ListIntegrations(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		if err := processIntegrations(ctx, handler, dopplerChan, &wg); err != nil {
			errorChan <- err // Send error to the error channel
		}
		wg.Wait()
	}()

	var values []models.Resource
	for {
		select {
		case value, ok := <-dopplerChan:
			if !ok {
				return values, nil
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		case err := <-errorChan:
			return nil, err
		}
	}
}

func GetIntegration(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	integration, err := processIntegration(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	var syncs []model.Sync
	for _, syncObject := range integration.Syncs {
		syncs = append(syncs, model.Sync{
			Slug:         syncObject.Slug,
			Enabled:      syncObject.Enabled,
			LastSyncedAt: syncObject.LastSyncedAt,
			Project:      syncObject.Project,
			Config:       syncObject.Config,
			Integration:  syncObject.Integration,
		})
	}
	value := models.Resource{
		ID:   integration.Slug,
		Name: integration.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.IntegrationDescription{
				Slug:    integration.Slug,
				Name:    integration.Name,
				Type:    integration.Type,
				Kind:    integration.Kind,
				Enabled: integration.Enabled,
				Syncs:   syncs,
			},
		},
	}
	return &value, nil
}

func processIntegrations(ctx context.Context, handler *resilientbridge.ResilientBridge, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var integrationListResponse model.IntegrationListResponse
	baseURL := "/v3/integrations"

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: baseURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("doppler", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &integrationListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, integration := range integrationListResponse.Integrations {
		wg.Add(1)
		go func(integration model.IntegrationJSON) {
			defer wg.Done()
			var syncs []model.Sync
			for _, syncObject := range integration.Syncs {
				syncs = append(syncs, model.Sync{
					Slug:         syncObject.Slug,
					Enabled:      syncObject.Enabled,
					LastSyncedAt: syncObject.LastSyncedAt,
					Project:      syncObject.Project,
					Config:       syncObject.Config,
					Integration:  syncObject.Integration,
				})
			}
			value := models.Resource{
				ID:   integration.Slug,
				Name: integration.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.IntegrationDescription{
						Slug:    integration.Slug,
						Name:    integration.Name,
						Type:    integration.Type,
						Kind:    integration.Kind,
						Enabled: integration.Enabled,
						Syncs:   syncs,
					},
				},
			}
			dopplerChan <- value
		}(integration)
	}
	return nil
}

func processIntegration(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.IntegrationJSON, error) {
	var integrationGetResponse model.IntegrationGetResponse
	baseURL := "/v3/integrations/integration"

	params := url.Values{}
	params.Set("integration", resourceID)
	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("doppler", req)
	if err != nil {
		return nil, fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &integrationGetResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &integrationGetResponse.Integration, nil
}
