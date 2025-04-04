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

func ListWorkplaceUsers(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		if err := processWorkplaceUsers(ctx, handler, dopplerChan, &wg); err != nil {
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

func GetWorkplaceUser(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	workplaceUser, err := processWorkPlaceUser(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	user := provider.User{
		Email:           workplaceUser.User.Email,
		Username:        workplaceUser.User.Username,
		Name:            workplaceUser.User.Name,
		ProfileImageURL: workplaceUser.User.ProfileImageURL,
	}
	value := models.Resource{
		ID:   workplaceUser.ID,
		Name: workplaceUser.User.Username,
		Description: provider.WorkplaceUserDescription{
			ID:        workplaceUser.ID,
			Access:    workplaceUser.Access,
			CreatedAt: workplaceUser.CreatedAt,
			User:      user,
		},
	}
	return &value, nil
}

func processWorkplaceUsers(ctx context.Context, handler *resilientbridge.ResilientBridge, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var workplaceUsers []provider.WorkplaceUserJSON
	var workplaceUserListResponse provider.WorkplaceUserListResponse
	baseURL := "/v3/workplace/users"
	page := 1

	for {
		params := url.Values{}
		params.Set("page", fmt.Sprintf("%d", page))
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

		if err = json.Unmarshal(resp.Data, &workplaceUserListResponse); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
		}

		if len(workplaceUserListResponse.WorkplaceUsers) == 0 {
			break
		}

		workplaceUsers = append(workplaceUsers, workplaceUserListResponse.WorkplaceUsers...)

		page++
	}

	for _, workplaceUser := range workplaceUsers {
		wg.Add(1)
		go func(workplaceUser provider.WorkplaceUserJSON) {
			defer wg.Done()
			user := provider.User{
				Email:           workplaceUser.User.Email,
				Username:        workplaceUser.User.Username,
				Name:            workplaceUser.User.Name,
				ProfileImageURL: workplaceUser.User.ProfileImageURL,
			}
			value := models.Resource{
				ID:   workplaceUser.ID,
				Name: workplaceUser.User.Username,
				Description: provider.WorkplaceUserDescription{
					ID:        workplaceUser.ID,
					Access:    workplaceUser.Access,
					CreatedAt: workplaceUser.CreatedAt,
					User:      user,
				},
			}
			dopplerChan <- value
		}(workplaceUser)
	}
	return nil
}

func processWorkPlaceUser(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*provider.WorkplaceUserJSON, error) {
	var workplaceUserGetResponse provider.WorkplaceUserGetResponse
	baseURL := "/v3/workplace/users/"

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

	if err = json.Unmarshal(resp.Data, &workplaceUserGetResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &workplaceUserGetResponse.WorkplaceUser, nil
}
