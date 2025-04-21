package procurement

import (
	asset "ModEd/asset/model"
	model "ModEd/asset/model/Procurement"
	"testing"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&model.InstrumentRequest{}, &model.InstrumentDetail{}, &asset.Category{})
	if err != nil {
		t.Fatalf("failed to migrate schema: %v", err)
	}

	return db
}

func TestCreateInstrumentRequest(t *testing.T) {
	db := setupTestDB(t)
	controller := CreateInstrumentRequestController(db)

	request := &model.InstrumentRequest{
		Status:       model.StatusDraft,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DepartmentID: 1,
	}

	err := controller.CreateInstrumentRequest(request)
	if err != nil {
		t.Errorf("CreateInstrumentRequest failed: %v", err)
	}

	var check model.InstrumentRequest
	err = db.First(&check, request.InstrumentRequestID).Error
	if err != nil {
		t.Errorf("Could not retrieve created request: %v", err)
	}
}

func TestAddInstrumentToRequest(t *testing.T) {
	db := setupTestDB(t)
	controller := CreateInstrumentRequestController(db)

	// create dummy category
	category := asset.Category{CategoryName: "Test Category"}
	if err := db.Create(&category).Error; err != nil {
		t.Fatalf("Failed to create category: %v", err)
	}

	// create dummy request
	request := &model.InstrumentRequest{
		Status:       model.StatusDraft,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		DepartmentID: 1,
	}
	if err := controller.CreateInstrumentRequest(request); err != nil {
		t.Fatalf("CreateInstrumentRequest failed: %v", err)
	}

	description := "Test Instrument"
	detail := &model.InstrumentDetail{
		InstrumentLabel: "Label 1",
		Description:     &description,
		Quantity:        1,
		CategoryID:      category.ID,
	}

	err := controller.AddInstrumentToRequest(request.InstrumentRequestID, detail)
	if err != nil {
		t.Errorf("AddInstrumentToRequest failed: %v", err)
	}

	var found model.InstrumentDetail
	if err := db.First(&found, detail.InstrumentDetailID).Error; err != nil {
		t.Errorf("Failed to retrieve instrument detail: %v", err)
	}
}
