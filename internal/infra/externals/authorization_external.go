package externals

import (
	"context"
	"encoding/json"
	"go-transfer/internal/domain/port"
	"net/http"
)

type AuthorizationServiceImpl struct {
	baseURL string
}

func NewAuthorizationService(baseURL string) port.AuthorizationService {
	return &AuthorizationServiceImpl{
		baseURL: baseURL,
	}
}

type AuthorizationResponse struct {
	Status string `json:"status"`
	Data   struct {
		Authorization bool `json:"authorization"`
	} `json:"data"`
}

func (s *AuthorizationServiceImpl) Authorize(ctx context.Context) (bool, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, s.baseURL, nil)
	if err != nil {
		return false, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var authResp AuthorizationResponse
	err = json.NewDecoder(resp.Body).Decode(&authResp)
	if err != nil {

		return false, err
	}
	statusSuccess := "success"
	return authResp.Status == statusSuccess && authResp.Data.Authorization, nil
}
