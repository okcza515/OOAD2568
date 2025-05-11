package controller

import (
	model "ModEd/asset/model"
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

func (b *InstrumentRequestBuilder) AddInstrumentDetailBuilder(detail *InstrumentDetailBuilder) *InstrumentRequestBuilder {
	b.req.Instruments = append(b.req.Instruments, *detail.Build())
	b.req.TotalEstimatedPrice += detail.price * float64(detail.quantity)
	return b
}

func (b *InstrumentRequestBuilder) Build() *model.InstrumentRequest {
	return b.req
}

// ---------------------------------------------
// Sub-builder for InstrumentDetail (Optional)
// ---------------------------------------------
type InstrumentDetailBuilder struct {
	detail   *model.InstrumentDetail
	price    float64
	quantity int
}

func NewInstrumentDetailBuilder() *InstrumentDetailBuilder {
	return &InstrumentDetailBuilder{detail: &model.InstrumentDetail{}}
}

func (b *InstrumentDetailBuilder) WithLabel(label string) *InstrumentDetailBuilder {
	b.detail.InstrumentLabel = label
	return b
}

func (b *InstrumentDetailBuilder) WithDescription(desc string) *InstrumentDetailBuilder {
	b.detail.Description = &desc
	return b
}

func (b *InstrumentDetailBuilder) WithCategoryID(id uint) *InstrumentDetailBuilder {
	b.detail.CategoryID = id
	return b
}

func (b *InstrumentDetailBuilder) WithEstimatedPrice(price float64) *InstrumentDetailBuilder {
	b.price = price
	b.detail.EstimatedPrice = price
	return b
}

func (b *InstrumentDetailBuilder) WithQuantity(qty int) *InstrumentDetailBuilder {
	b.quantity = qty
	b.detail.Quantity = qty
	return b
}

func (b *InstrumentDetailBuilder) WithRequestID(reqID uint) *InstrumentDetailBuilder {
	b.detail.InstrumentRequestID = reqID
	return b
}

func (b *InstrumentDetailBuilder) Build() *model.InstrumentDetail {
	return b.detail
}
