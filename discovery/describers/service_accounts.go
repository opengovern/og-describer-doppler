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

func ListServiceAccounts(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		if err := processServiceAccounts(ctx, handler, dopplerChan, &wg); err != nil {
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

func GetServiceAccount(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*models.Resource, error) {
	account, err := processServiceAccount(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	workplaceRole := provider.WorkplaceRole{
		Name:         account.WorkplaceRole.Name,
		Permissions:  account.WorkplaceRole.Permissions,
		IsCustomRole: account.WorkplaceRole.IsCustomRole,
		IsInlineRole: account.WorkplaceRole.IsInlineRole,
		Identifier:   account.WorkplaceRole.Identifier,
		CreatedAt:    account.WorkplaceRole.CreatedAt,
	}
	value := models.Resource{
		ID:   account.Slug,
		Name: account.Name,
		Description: provider.ServiceAccountDescription{
			Name:          account.Name,
			Slug:          account.Slug,
			CreatedAt:     account.CreatedAt,
			WorkplaceRole: workplaceRole,
		},
	}
	return &value, nil
}

func processServiceAccounts(ctx context.Context, handler *resilientbridge.ResilientBridge, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var accounts []provider.ServiceAccountJSON
	var accountListResponse provider.ServiceAccountListResponse
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
			return fmt.Errorf("request execution failed: %w", err)
		}

		if resp.StatusCode >= 400 {
			return fmt.Errorf("error %d: %s", resp.StatusCode, string(resp.Data))
		}

		if err = json.Unmarshal(resp.Data, &accountListResponse); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
		}

		if len(accountListResponse.ServiceAccounts) == 0 {
			break
		}

		accounts = append(accounts, accountListResponse.ServiceAccounts...)

		page++
	}

	for _, account := range accounts {
		wg.Add(1)
		go func(account provider.ServiceAccountJSON) {
			defer wg.Done()
			workplaceRole := provider.WorkplaceRole{
				Name:         account.WorkplaceRole.Name,
				Permissions:  account.WorkplaceRole.Permissions,
				IsCustomRole: account.WorkplaceRole.IsCustomRole,
				IsInlineRole: account.WorkplaceRole.IsInlineRole,
				Identifier:   account.WorkplaceRole.Identifier,
				CreatedAt:    account.WorkplaceRole.CreatedAt,
			}
			value := models.Resource{
				ID:   account.Slug,
				Name: account.Name,
				Description: provider.ServiceAccountDescription{
					Name:          account.Name,
					Slug:          account.Slug,
					CreatedAt:     account.CreatedAt,
					WorkplaceRole: workplaceRole,
				},
			}
			dopplerChan <- value
		}(account)
	}
	return nil
}

func processServiceAccount(ctx context.Context, handler *resilientbridge.ResilientBridge, resourceID string) (*provider.ServiceAccountJSON, error) {
	var accountGetResponse provider.ServiceAccountGetResponse
	baseURL := "/v3/workplace/service_accounts/service_account/"

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

	if err = json.Unmarshal(resp.Data, &accountGetResponse); err != nil {
		return nil, fmt.Errorf("error parsing response: %w", err)
	}

	return &accountGetResponse.ServiceAccount, nil
}
