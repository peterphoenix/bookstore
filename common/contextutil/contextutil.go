package contextutil

import "context"

type ctxKey string

const (
	ZeroUUID = "00000000-0000-0000-0000-000000000000"

	UserIDKey ctxKey = "X-User-ID"
)

func WithUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDKey, userID)
}

func GetUserID(ctx context.Context) string {
	userID := ctx.Value(UserIDKey)
	if userID != nil {
		return userID.(string)
	}
	return ZeroUUID
}
