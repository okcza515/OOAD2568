package model

import ()

type AssetManagement struct {
    InstrumentID Instrument //[1:1]
    InstrumentLabel  Instrument //[N:1]
    Quantity Supply //[1:1]
	BorrowStatus int
	BorrowDate BorrowInstrument //[1:1]
    ReturnDate ReturnInstrument //[1:1]
    ExpectedReturnDate BorrowInstrument //[1:1]
    IsLate bool
    BorrowID BorrowInstrument //[N:1]
    RoomID Instrument //[N:1]
}