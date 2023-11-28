package externalEvents

import (
	"github.com/mehdihadeli/go-ecommerce-microservices/internal/pkg/core/messaging/types"
)

type ProductDeletedV1 struct {
	*types.Message
	ProductId string `json:"productId,omitempty"`
}
