package controller

import (
	model "ModEd/asset/model"
	"time"
)

// --- TORBuilder ---

type TORBuilder struct {
	tor *model.TOR
}

func NewTORBuilder() *TORBuilder {
	return &TORBuilder{
		tor: &model.TOR{},
	}
}

func (b *TORBuilder) WithInstrumentRequestID(id uint) *TORBuilder {
	b.tor.InstrumentRequestID = id
	return b
}

func (b *TORBuilder) WithScope(scope string) *TORBuilder {
	b.tor.Scope = scope
	return b
}

func (b *TORBuilder) WithDeliverables(deliverables string) *TORBuilder {
	b.tor.Deliverables = deliverables
	return b
}

func (b *TORBuilder) WithTimeline(timeline string) *TORBuilder {
	b.tor.Timeline = timeline
	return b
}

func (b *TORBuilder) WithCommittee(committee string) *TORBuilder {
	b.tor.Committee = committee
	return b
}

func (b *TORBuilder) WithStatus(status model.TORStatus) *TORBuilder {
	b.tor.Status = status
	return b
}

func (b *TORBuilder) WithCreatedAt(t time.Time) *TORBuilder {
	b.tor.CreatedAt = t
	return b
}

func (b *TORBuilder) Build() *model.TOR {
	return b.tor
}

// --- ProcurementBuilder ---

type ProcurementBuilder struct {
	procurement *model.Procurement
}

func NewProcurementBuilder() *ProcurementBuilder {
	return &ProcurementBuilder{
		procurement: &model.Procurement{},
	}
}

func (b *ProcurementBuilder) WithTOR(tor *model.TOR) *ProcurementBuilder {
	b.procurement.TORID = tor.TORID
	return b
}

func (b *ProcurementBuilder) WithApproverID(approverID uint) *ProcurementBuilder {
	b.procurement.ApproverID = &approverID
	return b
}

func (b *ProcurementBuilder) WithStatus(status model.ProcurementStatus) *ProcurementBuilder {
	b.procurement.Status = status
	return b
}

func (b *ProcurementBuilder) WithApprovalTime(t time.Time) *ProcurementBuilder {
	b.procurement.ApprovalTime = &t
	return b
}

func (b *ProcurementBuilder) WithCreatedAt(t time.Time) *ProcurementBuilder {
	b.procurement.CreatedAt = t
	return b
}

func (b *ProcurementBuilder) Build() *model.Procurement {
	return b.procurement
}

// Optional: Combined builder return if needed
func (b *ProcurementBuilder) BuildWithTOR(tor *model.TOR) (*model.TOR, *model.Procurement) {
	b.procurement.TORID = tor.TORID
	return tor, b.procurement
}
