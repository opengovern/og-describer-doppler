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

func ListServiceAccountTokens(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors
	accounts, err := getServiceAccounts(handler)
	if err != nil {
		return nil, err
	}

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		for _, account := range accounts {
			if err := processServiceAccountTokens(ctx, handler, account.Slug, dopplerChan, &wg); err != nil {
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

func processServiceAccountTokens(ctx context.Context, handler *resilientbridge.ResilientBridge, serviceAccountSlug string, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var tokens []provider.ServiceAccountTokenJSON
	var tokenListResponse provider.ServiceAccountTokenListResponse
	baseURL := "/v3/workplace/service_accounts/service_account/service_account/"
	page := 1
	perPage := "20"

	for {
		params := url.Values{}
		params.Set("page", fmt.Sprintf("%d", page))
		params.Set("per_page", perPage)
		finalURL := fmt.Sprintf("%s%s/tokens?%s", baseURL, serviceAccountSlug, params.Encode())

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

		if err = json.Unmarshal(resp.Data, &tokenListResponse); err != nil {
			return fmt.Errorf("error parsing response: %w", err)
		}

		if len(tokenListResponse.APITokens) == 0 {
			break
		}

		tokens = append(tokens, tokenListResponse.APITokens...)

		page++
	}

	for _, token := range tokens {
		wg.Add(1)
		go func(token provider.ServiceAccountTokenJSON) {
			defer wg.Done()
			value := models.Resource{
				ID:   token.Slug,
				Name: token.Name,
				Description: provider.ServiceAccountTokenDescription{
					Name:       token.Name,
					Slug:       token.Slug,
					CreatedAt:  token.CreatedAt,
					ExpiresAt:  token.ExpiresAt,
					LastSeenAt: token.LastSeenAt,
				},
			}
			dopplerChan <- value
		}(token)
	}
	return nil
}
