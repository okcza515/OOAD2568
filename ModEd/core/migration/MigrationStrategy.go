package migration

// Wrote By : MEP-1010, MEP-1012

type MigrationStrategy interface {
	GetModels() []interface{}
}
