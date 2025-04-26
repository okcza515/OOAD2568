package cli

// Wrote by MEP-1012

type MenuState interface {
	Render()
	HandleUserInput(input string) error
}
