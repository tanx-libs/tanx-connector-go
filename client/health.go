package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Payload string `json:"payload"`
}

/*
To get started with integrating tanX for your Dapp, we suggest you test your connectivity to our REST APIs.
Health endpoint can be used to check the status of the API.
*/
func (c *Client) Health(ctx context.Context) (HealthResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.healthURL.String(), nil)
	if err != nil {
		return HealthResponse{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return HealthResponse{}, err
	}
	defer resp.Body.Close()

	var healthResponse HealthResponse
	err = json.NewDecoder(resp.Body).Decode(&healthResponse)
	if healthResponse.Status == ERROR {
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return HealthResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(healthResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return HealthResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(healthResponse.Message),
			}
		}

		return HealthResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, healthResponse.Message)

	} else if err != nil {
		return HealthResponse{}, &ErrJSONDecoding{Err: err}
	}

	return healthResponse, nil
}
