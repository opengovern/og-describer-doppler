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

func ListEnvironments(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
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
			if err := processEnvironments(ctx, handler, project.ID, dopplerChan, &wg); err != nil {
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

func processEnvironments(ctx context.Context, handler *resilientbridge.ResilientBridge, projectID string, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var environments []provider.EnvironmentJSON
	var environmentListResponse provider.EnvironmentListResponse
	baseURL := "/v3/environments"
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

		if err = json.Unmarshal(resp.Data, &environmentListResponse); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
		}

		if len(environmentListResponse.Environments) == 0 {
			break
		}

		environments = append(environments, environmentListResponse.Environments...)

		page++
	}

	for _, environment := range environments {
		wg.Add(1)
		go func(environment provider.EnvironmentJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   environment.Slug,
				Name: environment.Name,
				Description: provider.EnvironmentDescription{
					ID:             environment.ID,
					Slug:           environment.Slug,
					Name:           environment.Name,
					CreatedAt:      environment.CreatedAt,
					Project:        environment.Project,
					InitialFetchAt: environment.InitialFetchAt,
				},
			}
			dopplerChan <- value
		}(environment)
	}
	return nil
}
