package port

import "context"

type AuthorizationService interface {
	Authorize(ctx context.Context) (bool, error)
}
