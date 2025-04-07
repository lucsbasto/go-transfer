package port

import "context"

type NotificationService interface {
	Notify(ctx context.Context, receiverID int64, amount float64) error
}
