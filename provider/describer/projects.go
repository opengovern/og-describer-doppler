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

func ListProjects(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		if err := processProjects(ctx, handler, dopplerChan, &wg); err != nil {
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

func GetProject(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	project, err := processProject(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   project.ID,
		Name: project.Name,
		Description: model.ProjectDescription{
			ID:          project.ID,
			Slug:        project.Slug,
			Name:        project.Name,
			Description: project.Description,
			CreatedAt:   project.CreatedAt,
		},
	}
	return &value, nil
}

func processProjects(ctx context.Context, handler *resilientbridge.ResilientBridge, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var projects []model.ProjectJSON
	var projectListResponse model.ProjectListResponse
	baseURL := "/v3/projects"
	page := 1
	perPage := "20"

	for {
		params := url.Values{}
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

		if err = json.Unmarshal(resp.Data, &projectListResponse); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
		}

		if len(projectListResponse.Projects) == 0 {
			break
		}

		projects = append(projects, projectListResponse.Projects...)

		page++
	}

	for _, project := range projects {
		wg.Add(1)
		go func(project model.ProjectJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   project.ID,
				Name: project.Name,
				Description: model.ProjectDescription{
					ID:          project.ID,
					Slug:        project.Slug,
					Name:        project.Name,
					Description: project.Description,
					CreatedAt:   project.CreatedAt,
				},
			}
			dopplerChan <- value
		}(project)
	}
	return nil
}

func processProject(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.ProjectJSON, error) {
	var projectGetResponse model.ProjectGetResponse
	baseURL := "/v3/projects/project"

	params := url.Values{}
	params.Set("project", resourceID)
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

	if err = json.Unmarshal(resp.Data, &projectGetResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &projectGetResponse.Project, nil
}
