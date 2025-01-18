package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/models"
	"github.com/opengovern/og-describer-doppler/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"net/url"
	"sync"
)

func ListSecrets(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	projects, err := getProjects(handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		for _, project := range projects {
			configs, err := getConfigs(handler, project.ID)
			if err != nil {
				errorChan <- err
			}
			for _, config := range configs {
				if err := processSecrets(ctx, handler, project.ID, config.Name, dopplerChan, &wg); err != nil {
					errorChan <- err // Send error to the error channel
				}
			}
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

func processSecrets(ctx context.Context, handler *resilientbridge.ResilientBridge, projectID, configName string, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var secretListResponse provider.SecretListResponse
	baseURL := "/v3/configs/config/secrets"

	params := url.Values{}
	params.Set("project", projectID)
	params.Set("config", configName)
	params.Set("include_dynamic_secrets", "false")
	params.Set("include_managed_secrets", "true")
	finalURL := fmt.Sprintf("%s?%s", baseURL, params.Encode())

	req := &resilientbridge.NormalizedRequest{
		Method:   "GET",
		Endpoint: finalURL,
		Headers:  map[string]string{"accept": "application/json"},
	}

	resp, err := handler.Request("doppler", req)
	if err != nil {
		return fmt.Errorf("request execution failed: %w", err)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
	}

	if err = json.Unmarshal(resp.Data, &secretListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for key, secret := range secretListResponse.Secrets {
		wg.Add(1)
		go func(key string, secret provider.SecretJSON) {
			defer wg.Done()
			rawValueType := provider.ValueType{
				Type: secret.RawValueType.Type,
			}
			computedValueType := provider.ValueType{
				Type: secret.ComputedValueType.Type,
			}
			value := models.Resource{
				ID:   secret.Computed,
				Name: key,
				Description: provider.SecretDescription{
					Raw:                secret.Raw,
					Computed:           secret.Computed,
					Note:               secret.Note,
					RawVisibility:      secret.RawVisibility,
					ComputedVisibility: secret.ComputedVisibility,
					RawValueType:       rawValueType,
					ComputedValueType:  computedValueType,
				},
			}
			dopplerChan <- value
		}(key, secret)
	}
	return nil
}
