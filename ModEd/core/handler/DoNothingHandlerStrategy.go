package handler

// Wrote by MEP-1012

type DoNothingHandlerStrategy struct {
}

func (cs DoNothingHandlerStrategy) Execute() error {
	return nil
}
