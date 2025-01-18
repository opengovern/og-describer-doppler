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

func ListWorkplaces(ctx context.Context, handler *resilientbridge.ResilientBridge, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	dopplerChan := make(chan models.Resource)
	errorChan := make(chan error, 1) // Buffered channel to capture errors

	go func() {
		defer close(dopplerChan)
		defer close(errorChan)
		if err := processWorkplaces(ctx, handler, dopplerChan, &wg); err != nil {
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

func processWorkplaces(ctx context.Context, handler *resilientbridge.ResilientBridge, dopplerChan chan<- models.Resource, wg *sync.WaitGroup) error {
	var workplaceListResponse provider.WorkplaceListResponse
	baseURL := "/v3/workplace"

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

	if err = json.Unmarshal(resp.Data, &workplaceListResponse); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}
	workplace := workplaceListResponse.Workplace
	wg.Add(1)
	go func(workplace provider.WorkplaceJSON) {
		defer wg.Done()
		value := models.Resource{
			ID:   workplace.ID,
			Name: workplace.Name,
			Description: provider.WorkplaceDescription{
				ID:            workplace.ID,
				Name:          workplace.Name,
				BillingEmail:  workplace.BillingEmail,
				SecurityEmail: workplace.SecurityEmail,
			},
		}
		dopplerChan <- value
	}(workplace)
	return nil
}
