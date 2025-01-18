package describers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-doppler/discovery/pkg/models"
	"github.com/opengovern/og-describer-doppler/discovery/provider"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"sync"
)

func ListWorkplaceRoles(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		if err := processWorkPlaceRoles(ctx, handler, dopplerChan, &wg); err != nil {
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

func GetWorkplaceRole(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	workplaceRole, err := processWorkplaceRole(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   workplaceRole.Identifier,
		Name: workplaceRole.Name,
		Description: provider.WorkplaceRoleDescription{
			Name:         workplaceRole.Name,
			Permissions:  workplaceRole.Permissions,
			Identifier:   workplaceRole.Identifier,
			CreatedAt:    workplaceRole.CreatedAt,
			IsCustomRole: workplaceRole.IsCustomRole,
			IsInlineRole: workplaceRole.IsInlineRole,
		},
	}
	return &value, nil
}

func processWorkPlaceRoles(ctx context.Context, handler *resilientbridge.ResilientBridge, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var workplaceRoleListResponse provider.WorkplaceRoleListResponse
	baseURL := "/v3/workplace/roles"

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

	if err = json.Unmarshal(resp.Data, &workplaceRoleListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	for _, workplaceRole := range workplaceRoleListResponse.Roles {
		wg.Add(1)
		go func(workplaceRole provider.WorkplaceRoleJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   workplaceRole.Identifier,
				Name: workplaceRole.Name,
				Description: provider.WorkplaceRoleDescription{
					Name:         workplaceRole.Name,
					Permissions:  workplaceRole.Permissions,
					Identifier:   workplaceRole.Identifier,
					CreatedAt:    workplaceRole.CreatedAt,
					IsCustomRole: workplaceRole.IsCustomRole,
					IsInlineRole: workplaceRole.IsInlineRole,
				},
			}
			dopplerChan <- value
		}(workplaceRole)
	}
	return nil
}

func processWorkplaceRole(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*provider.WorkplaceRoleJSON, error) {
	var workplaceRoleGetResponse provider.WorkplaceRoleGetResponse
	baseURL := "/v3/workplace/roles/role/"

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

	if err = json.Unmarshal(resp.Data, &workplaceRoleGetResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &workplaceRoleGetResponse.Role, nil
}
