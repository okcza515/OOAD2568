package model

type PresentationType string

const (
	InitialPresentation PresentationType = "Proposal"
	MidtermPresentation PresentationType = "Midterm"
	FinalPresentation   PresentationType = "Final"
)
