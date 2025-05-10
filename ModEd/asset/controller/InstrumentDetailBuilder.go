package controller

import "ModEd/asset/model"

type InstrumentDetailBuilder struct {
	detail *model.InstrumentDetail
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

func (b *InstrumentDetailBuilder) WithCategoryID(catID uint) *InstrumentDetailBuilder {
	b.detail.CategoryID = catID
	return b
}

func (b *InstrumentDetailBuilder) WithEstimatedPrice(price float64) *InstrumentDetailBuilder {
	b.detail.EstimatedPrice = price
	return b
}

func (b *InstrumentDetailBuilder) WithQuantity(qty uint) *InstrumentDetailBuilder {
	b.detail.Quantity = int(qty)
	return b
}

func (b *InstrumentDetailBuilder) WithRequestID(reqID uint) *InstrumentDetailBuilder {
	b.detail.InstrumentRequestID = reqID
	return b
}

func (b *InstrumentDetailBuilder) Build() *model.InstrumentDetail {
	return b.detail
}
