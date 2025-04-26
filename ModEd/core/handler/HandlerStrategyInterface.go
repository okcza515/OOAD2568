package handler

// Wrote by MEP-1012

type HandlerStrategy interface {
	Execute() error
}
