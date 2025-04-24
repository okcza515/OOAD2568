package core_test

import (
	"ModEd/core"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestModel struct {
	gorm.Model
	Name string
}

func (m *TestModel) GetID() uint {
	return m.Model.ID
}
func (m *TestModel) ToString() string {
	return ""
}
func (m *TestModel) Validate() error {
	return nil
}
func (m *TestModel) ToCSVRow() string {
	return ""
}
func (m *TestModel) FromCSV(raw string) error {
	return nil
}
func (m *TestModel) ToJSON() string {
	return ""
}
func (m *TestModel) FromJSON(raw string) error {
	return nil
}

func Init() (*gorm.DB, string) {
	dbName := "test.db"
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	if err := db.AutoMigrate(&TestModel{}); err != nil {
		panic(err)
	}
	return db, dbName
}

func cleanup(db *gorm.DB, dbName string) {
	os.Remove(dbName)
}

func TestInsert(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	testData := TestModel{Name: "TestName"}
	err := controller.Insert(&testData)
	assert.NoError(t, err)
}

func TestInsertMany(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	records := []*TestModel{
		{Name: "A"},
		{Name: "B"},
		{Name: "C"},
		{Name: "D"},
		{Name: "E"},
	}

	err := controller.InsertMany(records)
	assert.NoError(t, err)
}

func TestRetrieveByID(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)

	testData := TestModel{Name: "TestName"} // Remove explicit ID
	controller.Insert(&testData)

	result, err := controller.RetrieveByID(testData.ID) // Use the auto-generated ID
	if err != nil {
		assert.NoError(t, err)
	}

	assert.Equal(t, testData.ID, (*result).ID)
	assert.Equal(t, "TestName", (*result).Name)
}

func TestUpdateByID(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	testData := TestModel{Model: gorm.Model{ID: 1}, Name: "OldName"}
	db.Create(&testData)

	updatedData := TestModel{Model: gorm.Model{ID: 1}, Name: "UpdatedName"}
	err := controller.UpdateByID(&updatedData)
	assert.NoError(t, err)
}

func TestDeleteByID(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	testData := TestModel{Name: "TestName"}
	db.Create(&testData)

	err := controller.DeleteByID(testData.ID)
	assert.NoError(t, err)

	var result TestModel
	db.First(&result, testData.ID)
	assert.Empty(t, result.ID)
}

func TestList(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	db.Create(&TestModel{Name: "TestName"})
	db.Create(&TestModel{Name: "TestName"})

	result, err := controller.List(map[string]interface{}{"name": "TestName"})
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestListPagination(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	db.Create(&TestModel{Name: "TestName"})
	db.Create(&TestModel{Name: "TestName"})
	db.Create(&TestModel{Name: "TestName"})

	result, err := controller.ListPagination(map[string]interface{}{"name": "TestName"}, 1, 3)
	assert.NoError(t, err)
	assert.Len(t, result, 3)
	assert.Equal(t, 3, len(result))
}
