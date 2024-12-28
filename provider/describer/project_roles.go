package describer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-doppler/pkg/sdk/models"
	"github.com/opengovern/og-describer-doppler/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"sync"
)

func ListProjectRoles(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		if err := processProjectRoles(ctx, handler, dopplerChan, &wg); err != nil {
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

func GetProjectRole(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	projectRole, err := processProjectRole(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   projectRole.Identifier,
		Name: projectRole.Name,
		Description: model.ProjectRoleDescription{
			Name:         projectRole.Name,
			Permissions:  projectRole.Permissions,
			Identifier:   projectRole.Identifier,
			CreatedAt:    projectRole.CreatedAt,
			IsCustomRole: projectRole.IsCustomRole,
		},
	}
	return &value, nil
}

func processProjectRoles(ctx context.Context, handler *resilientbridge.ResilientBridge, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var projectRoleListResponse model.ProjectRoleListResponse
	baseURL := "/v3/projects/roles"

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

	if err = json.Unmarshal(resp.Data, &projectRoleListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, role := range projectRoleListResponse.Roles {
		wg.Add(1)
		go func(role model.ProjectRoleJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   role.Identifier,
				Name: role.Name,
				Description: model.ProjectRoleDescription{
					Name:         role.Name,
					Permissions:  role.Permissions,
					Identifier:   role.Identifier,
					CreatedAt:    role.CreatedAt,
					IsCustomRole: role.IsCustomRole,
				},
			}
			dopplerChan <- value
		}(role)
	}
	return nil
}

func processProjectRole(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.ProjectRoleJSON, error) {
	var projectRoleGetResponse model.ProjectRoleGetResponse
	baseURL := "/v3/projects/roles/role/"

	finalURL := fmt.Sprintf("%s%s", baseURL, resourceID)

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

	if err = json.Unmarshal(resp.Data, &projectRoleGetResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &projectRoleGetResponse.Role, nil
}
