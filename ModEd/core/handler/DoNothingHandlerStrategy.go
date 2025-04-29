package handler

// Wrote by MEP-1012

type DoNothingHandlerStrategy struct {
}

func (handler DoNothingHandlerStrategy) Execute() error {
	return nil
}
