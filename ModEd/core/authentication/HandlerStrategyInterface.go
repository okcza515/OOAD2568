// Wrote by MEP-1001
package authentication

type AuthHandlerStrategy interface {
	Execute() error
}
