// MEP-1005

package core

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"

	"gorm.io/gorm"
)

type BaseController[T RecordInterface] struct {
	db *gorm.DB
}

func NewBaseController[T RecordInterface](db *gorm.DB) *BaseController[T] {
	return &BaseController[T]{db: db}
}

func (controller *BaseController[T]) Insert(data T) error {
	return controller.db.Create(&data).Error
}

func (controller *BaseController[T]) InsertMany(data []T) error {
	return controller.db.Create(data).Error
}

func (controller *BaseController[T]) UpdateByID(data T) error {
	result := controller.db.Model(data).Where("id = ?", data.GetID()).Updates(data)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return result.Error
}

func (controller *BaseController[T]) UpdateByCondition(condition map[string]interface{}, data T) error {
	result := controller.db.Model(data).Where(condition).Updates(data)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}
	return result.Error
}

func (controller *BaseController[T]) RetrieveByID(id uint, preloads ...string) (T, error) {
	var record T
	query := controller.db

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where("id = ?", id).First(&record).Error; err != nil {
		return record, err
	}

	return record, nil
}

func (controller *BaseController[T]) RetrieveByCondition(condition map[string]interface{}, preloads ...string) (T, error) {
	var record T

	query := controller.db
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Where(condition).First(&record).Error; err != nil {
		return record, err
	}

	return record, nil
}

func (controller *BaseController[T]) DeleteByID(id uint) error {
	var record T
	result := controller.db.Where("id = ?", id).Delete(&record)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return result.Error
}

func (controller *BaseController[T]) DeleteByCondition(condition map[string]interface{}) error {
	var record T
	result := controller.db.Where(condition).Delete(&record)
	if result.RowsAffected == 0 {
		return errors.New("record not found")
	}

	return result.Error
}

func (controller *BaseController[T]) List(condition map[string]interface{}, preloads ...string) ([]T, error) {
	var records []T
	query := controller.db

	if condition != nil {
		query = query.Where(condition)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Find(&records).Error; err != nil {
		return nil, err
	}
	return records, nil
}

func (controller *BaseController[T]) ListPagination(condition map[string]interface{}, page, pageSize int, preloads ...string) ([]T, error) {
	var records []T
	var totalCount int64
	query := controller.db

	if condition != nil {
		query = query.Where(condition)
	}

	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	if err := query.Model(new(T)).Count(&totalCount).Error; err != nil {
		return nil, err
	}

	offset := (page - 1) * pageSize

	if err := query.Offset(offset).Limit(pageSize).Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (controller *BaseController[T]) FromCSV(path string) {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)

	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	if len(rows) < 2 {
		fmt.Println("CSV does not contain any data")
		return
	}

	headers := rows[0]
	var result []T

	for rowIndex, row := range rows[1:] {
		var zero T
		tType := reflect.TypeOf(zero)
		if tType.Kind() == reflect.Ptr {
			tType = tType.Elem()
		}
		newObj := reflect.New(tType).Interface()

		v := reflect.ValueOf(newObj).Elem()

		for i, header := range headers {
			if i >= len(row) {
				continue
			}

			field := v.FieldByName(header)
			if !field.IsValid() || !field.CanSet() {
				continue
			}

			value := row[i]
			switch field.Kind() {
			case reflect.String:
				field.SetString(value)
			case reflect.Int, reflect.Int64, reflect.Int32:
				if intVal, err := strconv.ParseInt(value, 10, 64); err == nil {
					field.SetInt(intVal)
				}
			case reflect.Uint, reflect.Uint64, reflect.Uint32:
				if uintVal, err := strconv.ParseUint(value, 10, 64); err == nil {
					field.SetUint(uintVal)
				}
			case reflect.Float32, reflect.Float64:
				if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
					field.SetFloat(floatVal)
				}
			case reflect.Bool:
				if boolVal, err := strconv.ParseBool(value); err == nil {
					field.SetBool(boolVal)
				}
			default:
				fmt.Printf("Unsupported field type: %s in row %d\n", field.Kind(), rowIndex+2)
			}
		}

		typedObj := newObj.(T)
		result = append(result, typedObj)
	}

	if err := controller.InsertMany(result); err != nil {
		fmt.Println("InsertMany failed:", err)
	}
}
