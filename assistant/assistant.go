package assistant

import "context"

type Assistant interface {
	GenerateMessage(ctx context.Context, msgs []string) (string, error)
}
