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

func ListProjectMembers(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
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
			if err := processProjectMembers(ctx, handler, project.ID, dopplerChan, &wg); err != nil {
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

func processProjectMembers(ctx context.Context, handler *resilientbridge.ResilientBridge, projectID string, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var projectMembers []model.ProjectMemberJSON
	var projectMemberListResponse model.ProjectMemberListResponse
	baseURL := "/v3/projects/project/members"
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

		if err = json.Unmarshal(resp.Data, &projectMemberListResponse); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
		}

		if len(projectMemberListResponse.Members) == 0 {
			break
		}

		projectMembers = append(projectMembers, projectMemberListResponse.Members...)

		page++
	}

	for _, projectMember := range projectMembers {
		wg.Add(1)
		go func(projectMember model.ProjectMemberJSON) {
			defer wg.Done()
			role := model.Role{
				Identifier: projectMember.Role.Identifier,
			}
			value := models.Resource{
				ID:   projectMember.Slug,
				Name: projectMember.Type,
				Description: model.ProjectMemberDescription{
					Type:                  projectMember.Type,
					Slug:                  projectMember.Slug,
					Role:                  role,
					AccessAllEnvironments: projectMember.AccessAllEnvironments,
					Environments:          projectMember.Environments,
				},
			}
			dopplerChan <- value
		}(projectMember)
	}
	return nil
}
