package producers

import "context"

type Producer interface {
	PublishMessage(ctx context.Context, message []byte) error
}
