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

func ListGroups(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		if err := processGroups(ctx, handler, dopplerChan, &wg); err != nil {
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

func GetGroup(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	group, err := processGroup(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	defaultProjectRole := model.DefaultProjectRole{
		Identifier: group.DefaultProjectRole.Identifier,
	}
	value := models.Resource{
		ID:   group.Slug,
		Name: group.Name,
		Description: model.GroupDescription{
			Name:               group.Name,
			Slug:               group.Slug,
			CreatedAt:          group.CreatedAt,
			DefaultProjectRole: defaultProjectRole,
		},
	}
	return &value, nil
}

func processGroups(ctx context.Context, handler *resilientbridge.ResilientBridge, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var groups []model.GroupJSON
	var groupListResponse model.GroupListResponse
	baseURL := "/v3/workplace/groups"
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

		if err = json.Unmarshal(resp.Data, &groupListResponse); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
		}

		if len(groupListResponse.Groups) == 0 {
			break
		}

		groups = append(groups, groupListResponse.Groups...)

		page++
	}

	for _, group := range groups {
		wg.Add(1)
		go func(group model.GroupJSON) {
			defer wg.Done()
			defaultProjectRole := model.DefaultProjectRole{
				Identifier: group.DefaultProjectRole.Identifier,
			}
			value := models.Resource{
				ID:   group.Slug,
				Name: group.Name,
				Description: model.GroupDescription{
					Name:               group.Name,
					Slug:               group.Slug,
					CreatedAt:          group.CreatedAt,
					DefaultProjectRole: defaultProjectRole,
				},
			}
			dopplerChan <- value
		}(group)
	}
	return nil
}

func processGroup(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*model.GroupJSON, error) {
	var groupGetResponse model.GroupGetResponse
	baseURL := "/v3/workplace/groups/group/"

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

	if err = json.Unmarshal(resp.Data, &groupGetResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &groupGetResponse.Group, nil
}
