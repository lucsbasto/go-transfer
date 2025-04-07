package externals

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"go-transfer/internal/domain/port"
	"io"
	"net/http"
)

type NotificationServiceImpl struct {
	baseURL string
}

func NewNotificationService(baseURL string) port.NotificationService {
	return &NotificationServiceImpl{
		baseURL: baseURL,
	}
}

type NotificationRequest struct {
	ReceiverID int64   `json:"receiverID"`
	Amount     float64 `json:"amount"`
}

func (s *NotificationServiceImpl) Notify(ctx context.Context, receiverID int64, amount float64) error {
	reqBody := NotificationRequest{
		ReceiverID: receiverID,
		Amount:     amount,
	}

	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.baseURL, bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	return fmt.Errorf("erro ao enviar notificação, status code: %d", resp.StatusCode)
}
