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

/*
{
  "status": "success",
  "message": "Successful",
  "payload": {
    "name": "USER NAME",
    "customer_id": "27",
    "img": null,
    "username": "guthal",
    "stark_key": "0x70f41ce6797eb444c9dc95a907....8aa592adf8f1fe3ab75317f7096d38"
  }
}
*/
type ProfileResponse struct {
	Status  Status         `json:"status"`
	Message string         `json:"message"`
	Payload ProfilePayload `json:"payload"`
}

/*
Retrieve details of a user’s portfolio i.e., the profit and loss incurred in the various markets.
Please note that this is a Private 🔒 rroute which means it needs to be authorized by the account initiating this request.
*/
func (c *Client) Profile(ctx context.Context) (ProfileResponse, error) {
	err := c.CheckAuth()
	if err != nil {
		return ProfileResponse{}, err
	}

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
	
	if profileResponse.Status == ERROR {
		// handling 4xx and 5xx errors
		if resp.StatusCode >= 400 && resp.StatusCode < 500 {
			return ProfileResponse{}, &ErrClient{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(profileResponse.Message),
			}
		} else if resp.StatusCode >= 500 {
			return ProfileResponse{}, &ErrServer{
				Status: resp.StatusCode,
				Err:    fmt.Errorf(profileResponse.Message),
			}
		}

		return ProfileResponse{}, fmt.Errorf("status: %d\nerror: %s", resp.StatusCode, profileResponse.Message)
	
	} else if err != nil {
		return ProfileResponse{}, &ErrJSONDecoding{Err: err}
	}
	
	

	return profileResponse, nil
}