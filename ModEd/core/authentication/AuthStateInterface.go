// Wrote by MEP-1001
package authentication

import "context"

type AuthState interface {
	GetContext() context.Context
	GetMiddleware() *Middleware
	SetContext(ctx context.Context)
}
