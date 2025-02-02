package interfaces

import (
	"context"
	"ecommerce-order/external"
)

type IExternal interface {
	GetProfile(ctx context.Context, token string) (external.Profile, error)
	ProduceKafkaMessage(ctx context.Context, topic string, data []byte) error
}
