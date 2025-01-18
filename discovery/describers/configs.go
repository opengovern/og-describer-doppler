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

func ListConfigs(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
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
			if err := processConfigs(ctx, handler, project.ID, dopplerChan, &wg); err != nil {
				errorChan <- err // Send error to the error channel
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

func processConfigs(ctx context.Context, handler *resilientbridge.ResilientBridge, projectID string, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var configs []provider.ConfigJSON
	var configListResponse provider.ConfigListResponse
	baseURL := "/v3/configs"
	page := 1
	perPage := "20"

	for {
		params := url.Values{}
		params.Set("project", projectID)
		params.Set("page", fmt.Sprintf("%d", page))
		params.Set("per_page", perPage)
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

		if err = json.Unmarshal(resp.Data, &configListResponse); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
		}

		if len(configListResponse.Configs) == 0 {
			break
		}

		configs = append(configs, configListResponse.Configs...)

		page++
	}

	for _, config := range configs {
		wg.Add(1)
		go func(config provider.ConfigJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   config.Slug,
				Name: config.Name,
				Description: provider.ConfigDescription{
					Name:           config.Name,
					Slug:           config.Slug,
					Project:        config.Project,
					Root:           config.Root,
					Inheritable:    config.Inheritable,
					Inheriting:     config.Inheriting,
					InitialFetchAt: config.InitialFetchAt,
					Inherits:       config.Inherits,
					Locked:         config.Locked,
					LastFetchAt:    config.LastFetchAt,
					CreatedAt:      config.CreatedAt,
					Environment:    config.Environment,
				},
			}
			dopplerChan <- value
		}(config)
	}
	return nil
}
