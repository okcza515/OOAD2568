package core_test

import (
	"ModEd/core"
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestModel struct {
	gorm.Model
	ID   uint
	Name string
}

func (m TestModel) GetID() uint {
	return 0
}
func (m TestModel) ToString() string {
	return ""
}
func (m TestModel) Validate() error {
	return nil
}
func (m TestModel) ToCSVRow() string {
	return ""
}
func (m TestModel) FromCSV(raw string) error {
	return nil
}
func (m TestModel) ToJSON() string {
	return ""
}
func (m TestModel) FromJSON(raw string) error {
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

	controller := core.NewBaseController("test", db)
	testData := TestModel{Name: "TestName"}
	err := controller.Insert(testData)
	assert.NoError(t, err)
}

func TestRetrieveByID(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController("test", db)

	testData := TestModel{Name: "TestName"} // Remove explicit ID
	db.Create(&testData)

	result, err := controller.RetrieveByID(testData.ID) // Use the auto-generated ID

	if resultCasted, ok := (*result).(TestModel); ok {
		assert.NoError(t, err)
		assert.Equal(t, testData.ID, resultCasted.ID)
		assert.Equal(t, "TestName", resultCasted.Name)
	} else {
		assert.NoError(t, errors.New("Failed to cast type"))
	}
}

func TestUpdateByID(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController("test", db)
	testData := TestModel{Name: "OldName"}
	db.Create(&testData)

	updatedData := TestModel{Name: "UpdatedName"}
	err := controller.UpdateByID(updatedData)
	assert.NoError(t, err)

	var result TestModel
	db.First(&result, testData.ID)
	assert.Equal(t, "UpdatedName", result.Name)
}

func TestDeleteByID(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController("test", db)
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

	controller := core.NewBaseController("test", db)
	db.Create(&TestModel{Name: "TestName"})
	db.Create(&TestModel{Name: "TestName"})

	result, err := controller.List(map[string]interface{}{"name": "TestName"})
	assert.NoError(t, err)
	assert.Len(t, result, 2)
}

func TestListPagination(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController("test", db)
	db.Create(&TestModel{Name: "TestName"})
	db.Create(&TestModel{Name: "TestName"})
	db.Create(&TestModel{Name: "TestName"})

	result, err := controller.ListPagination(map[string]interface{}{"name": "TestName"}, 1, 2)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, int64(3), result.TotalCount)
}
