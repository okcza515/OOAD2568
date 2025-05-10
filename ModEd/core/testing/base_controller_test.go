package core_test

import (
	"ModEd/core"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ChildModel struct {
	gorm.Model
	TestModelID uint
	Value       string
}

type TestModel struct {
	gorm.Model
	Name   string
	Childs []ChildModel `gorm:"foreignKey:TestModelID"`
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

	if err := db.AutoMigrate(&TestModel{}, &ChildModel{}); err != nil {
		panic(err)
	}
	return db, dbName
}

func cleanup(db *gorm.DB, dbName string) {
	// os.Remove(dbName)
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

	assert.Equal(t, testData.ID, result.ID)
	assert.Equal(t, "TestName", result.Name)
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

func TestInsertNonPointerShouldFail(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	testData := TestModel{Name: "FailInsert"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when using non-pointer in Insert, but did not panic")
		}
	}()
	// แปลงเป็น interface แต่ intentionally wrong
	_ = controller.Insert(any(testData).(*TestModel))
}

func TestUpdateByIDNonPointerShouldFail(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	testData := TestModel{Name: "Old"}
	db.Create(&testData)

	updated := TestModel{Model: gorm.Model{ID: testData.ID}, Name: "New"}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic when using non-pointer in UpdateByID, but did not panic")
		}
	}()
	_ = controller.UpdateByID(any(updated).(*TestModel))
}

func TestDeleteByIDNonPointerInstanceInternalHandling(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)
	testData := TestModel{Name: "DeleteMe"}
	db.Create(&testData)

	err := controller.DeleteByID(testData.ID)
	assert.NoError(t, err)

	var result TestModel
	db.First(&result, testData.ID)
	assert.Zero(t, result.ID)
}

func TestRetrieveByIDWithPreloads(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)

	parent := TestModel{Name: "Parent"}
	err := controller.Insert(&parent)
	assert.NoError(t, err)

	child1 := ChildModel{TestModelID: parent.ID, Value: "Child1"}
	child2 := ChildModel{TestModelID: parent.ID, Value: "Child2"}
	db.Create(&child1)
	db.Create(&child2)

	result, err := controller.RetrieveByID(parent.ID, "Childs")
	assert.NoError(t, err)

	assert.Equal(t, "Parent", result.Name)
	assert.Len(t, result.Childs, 2)
	assert.Equal(t, "Child1", result.Childs[0].Value)
	assert.Equal(t, "Child2", result.Childs[1].Value)
}

func TestListWithPreloads(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	controller := core.NewBaseController[*TestModel](db)

	parent1 := TestModel{Name: "Parent1"}
	parent2 := TestModel{Name: "Parent2"}
	err := controller.Insert(&parent1)
	assert.NoError(t, err)
	err = controller.Insert(&parent2)
	assert.NoError(t, err)

	db.Create(&ChildModel{TestModelID: parent1.ID, Value: "Child1-1"})
	db.Create(&ChildModel{TestModelID: parent1.ID, Value: "Child1-2"})
	db.Create(&ChildModel{TestModelID: parent2.ID, Value: "Child2-1"})

	results, err := controller.List(map[string]interface{}{}, "Childs")
	assert.NoError(t, err)
	assert.Len(t, results, 2)

	for _, parent := range results {
		switch parent.Name {
		case "Parent1":
			assert.Len(t, parent.Childs, 2)
			assert.Equal(t, "Child1-1", parent.Childs[0].Value)
			assert.Equal(t, "Child1-2", parent.Childs[1].Value)
		case "Parent2":
			assert.Len(t, parent.Childs, 1)
			assert.Equal(t, "Child2-1", parent.Childs[0].Value)
		default:
			t.Errorf("Unexpected parent name: %s", parent.Name)
		}
	}
}

func TestFromCSV(t *testing.T) {
	db, dbName := Init()
	defer cleanup(db, dbName)

	// Prepare CSV content
	csvContent := `Name
CSVTest1
CSVTest2
CSVTest3
`
	// Write it to a temporary file
	csvFile := "data.csv"
	err := os.WriteFile(csvFile, []byte(csvContent), 0644)
	assert.NoError(t, err)
	defer os.Remove(csvFile)

	// Create controller and run FromCSV
	controller := core.NewBaseController[*TestModel](db)
	controller.FromCSV(csvFile)

	// Verify data
	var results []TestModel
	db.Find(&results)
	assert.Len(t, results, 3)
	assert.Equal(t, "CSVTest1", results[0].Name)
	assert.Equal(t, "CSVTest2", results[1].Name)
	assert.Equal(t, "CSVTest3", results[2].Name)
}
