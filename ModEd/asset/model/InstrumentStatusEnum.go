package model

// MEP-1012 Asset

type InstrumentStatusEnum string

const (
	INS_AVAILABLE InstrumentStatusEnum = "Available"
	INS_BORROWED  InstrumentStatusEnum = "Borrowed"
	INS_BROKEN    InstrumentStatusEnum = "Broken"
	INS_LOST      InstrumentStatusEnum = "Lost"
	INS_SALVAGING InstrumentStatusEnum = "Salvaging"
	INS_SALVAGED  InstrumentStatusEnum = "Salvaged"
	INS_DONATED   InstrumentStatusEnum = "Donated"
)
