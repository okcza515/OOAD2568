// MEP-1014
package procurement

// import (
// 	model "ModEd/asset/model/Procurement"

// 	"gorm.io/gorm"
// )

// type AcceptanceTestController struct {
// 	Connector *gorm.DB
// }

// func CreateAcceptanceTestController(connector *gorm.DB) *AcceptanceTestController {
// 	acceptanceTest := AcceptanceTestController{Connector: connector}
// 	connector.AutoMigrate(&model.AcceptanceTest{})
// 	return &acceptanceTest
// }

// func (acceptanceTest AcceptanceTestController) ListAll() ([]model.AcceptanceTest, error) {
// 	tests := []model.AcceptanceTest{}
// 	result := acceptanceTest.Connector.
// 		Select("AcceptanceTestID").Find(&tests)
// 	return tests, result.Error
// }

// func (acceptanceTest AcceptanceTestController) GetByID(id uint) (AcceptanceTestID string) ([]model.AcceptanceTest, error) {
// 	t := &model.AcceptanceTest{}
// 	result := acceptanceTest.Connector.Where("AcceptanceTestID = ?", id).First(t)
// 	return t, result.Error
// }

// func (acceptanceTest AcceptanceTestController) Create(t *model.AcceptanceTest) error {
// 	return acceptanceTest.Connector.Create(t).Error
// }

// func (acceptanceTest AcceptanceTestController) Update(AcceptanceTestID uint, updatedData map[string]interface{}) error {
// 	return acceptanceTest.Connector.Model(&model.AcceptanceTest{}).Where("AcceptanceTestID = ?", AcceptanceTestID).Updates(updatedData).Error
// }

// func (acceptanceTest AcceptanceTestController) DeleteByID(AcceptanceTestID uint) error {
// 	return acceptanceTest.Connector.Where("AcceptanceTestID = ?", AcceptanceTestID).Delete(&model.AcceptanceTest{}).Error
// }
