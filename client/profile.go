package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ProfilePayload struct {
	Name       string `json:"name"`
	CustomerID string `json:"customer_id"`
	Img        []byte `json:"img"`
	Username   string `json:"username"`
	StarkKey   string `json:"stark_key"`
}

type ProfileResponse struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Payload ProfilePayload `json:"payload"`
}

// Retrieve details of a user’s portfolio
func (c *Client) Profile(ctx context.Context) (ProfileResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.profileURL.String(), nil)
	if err != nil {
		return ProfileResponse{}, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.jwtToken))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return ProfileResponse{}, err
	}
	defer resp.Body.Close()

	var profileResponse ProfileResponse
	err = json.NewDecoder(resp.Body).Decode(&profileResponse)
	if err != nil {
		return ProfileResponse{}, err
	}

	return profileResponse, nil
}



