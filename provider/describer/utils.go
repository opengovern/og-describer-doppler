package describer

import (
	"encoding/json"
	"fmt"
	"github.com/opengovern/og-describer-doppler/provider/model"
	resilientbridge "github.com/opengovern/resilient-bridge"
	"net/url"
)

func getProjects(handler *resilientbridge.ResilientBridge) ([]model.ProjectDescription, error) {
	var projects []model.ProjectDescription
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
			return nil, fmt.Errorf("request execution failed: %w", err)
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
		}

		if err = json.Unmarshal(resp.Data, &projectListResponse); err != nil {
			return nil, fmt.Errorf("error parsing response: %w", err)
		}

		if len(projectListResponse.Projects) == 0 {
			break
		}

		projects = append(projects, projectListResponse.Projects...)

		page++
	}

	return projects, nil
}

func getConfigs(handler *resilientbridge.ResilientBridge, projectID string) ([]model.ConfigDescription, error) {
	var configs []model.ConfigDescription
	var configListResponse model.ConfigListResponse
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
			return nil, fmt.Errorf("request execution failed: %w", err)
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
		}

		if err = json.Unmarshal(resp.Data, &configListResponse); err != nil {
			return nil, fmt.Errorf("error parsing response: %w", err)
		}

		if len(configListResponse.Configs) == 0 {
			break
		}

		configs = append(configs, configListResponse.Configs...)

		page++
	}

	return configs, nil
}

func getServiceAccounts(handler *resilientbridge.ResilientBridge) ([]model.ServiceAccountDescription, error) {
	var accounts []model.ServiceAccountDescription
	var accountListResponse model.ServiceAccountListResponse
	baseURL := "/v3/workplace/service_accounts"
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
			return nil, fmt.Errorf("request execution failed: %w", err)
		}

		if resp.StatusCode >= 400 {
			return nil, fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
		}

		if err = json.Unmarshal(resp.Data, &accountListResponse); err != nil {
			return nil, fmt.Errorf("error parsing response: %w", err)
		}

		if len(accountListResponse.ServiceAccounts) == 0 {
			break
		}

		accounts = append(accounts, accountListResponse.ServiceAccounts...)

		page++
	}

	return accounts, nil
}
