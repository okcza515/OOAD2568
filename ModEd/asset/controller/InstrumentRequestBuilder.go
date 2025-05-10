package controller

import (
	"ModEd/asset/model"
)

type InstrumentRequestBuilder struct {
	req *model.InstrumentRequest
}

func NewInstrumentRequestBuilder() *InstrumentRequestBuilder {
	return &InstrumentRequestBuilder{req: &model.InstrumentRequest{}}
}

func (b *InstrumentRequestBuilder) WithStatus(status model.InstrumentRequestStatus) *InstrumentRequestBuilder {
	b.req.Status = status
	return b
}

func (b *InstrumentRequestBuilder) WithDepartmentID(deptID uint) *InstrumentRequestBuilder {
	b.req.DepartmentID = deptID
	return b
}

func (b *InstrumentRequestBuilder) WithLinkedToTOR(linked bool) *InstrumentRequestBuilder {
	b.req.IsLinkedToTOR = linked
	return b
}

func (b *InstrumentRequestBuilder) AddInstrumentDetail(detail model.InstrumentDetail) *InstrumentRequestBuilder {
	b.req.Instruments = append(b.req.Instruments, detail)
	b.req.TotalEstimatedPrice += detail.EstimatedPrice * float64(detail.Quantity)
	return b
}

func (b *InstrumentRequestBuilder) Build() *model.InstrumentRequest {
	return b.req
}
