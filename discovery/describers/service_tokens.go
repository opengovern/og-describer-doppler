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

func ListServiceTokens(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
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
				if err := processServiceTokens(ctx, handler, project.ID, config.Name, dopplerChan, &wg); err != nil {
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

func processServiceTokens(ctx context.Context, handler *resilientbridge.ResilientBridge, projectID, configName string, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var tokenListResponse provider.ServiceTokenListResponse
	baseURL := "/v3/configs/config/tokens"

	params := url.Values{}
	params.Set("project", projectID)
	params.Set("config", configName)
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

	if err = json.Unmarshal(resp.Data, &tokenListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, token := range tokenListResponse.Tokens {
		wg.Add(1)
		go func(token provider.ServiceTokenJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   token.Slug,
				Name: token.Name,
				Description: provider.ServiceTokenDescription{
					Name:        token.Name,
					Slug:        token.Slug,
					CreatedAt:   token.CreatedAt,
					Config:      token.Config,
					Environment: token.Environment,
					Project:     token.Project,
					ExpiresAt:   token.ExpiresAt,
				},
			}
			dopplerChan <- value
		}(token)
	}
	return nil
}
