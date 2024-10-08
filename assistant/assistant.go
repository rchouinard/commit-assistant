package assistant

import "context"

type Assistant interface {
	GenerateMessage(ctx context.Context, diffInput string) (string, error)
}
