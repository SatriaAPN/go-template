package interceptor

import (
	"context"

	"github.com/google/uuid"
)

func SetRequestId(ctx *context.Context) {
	uuid := uuid.NewString()
	*ctx = context.WithValue(*ctx, "X-Request-Id", uuid)
}
