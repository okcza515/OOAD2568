package seed

import (
	controller "ModEd/curriculum/controller/class"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

func CreateClassSeed(db *gorm.DB, path string) (classes []*model.Class, err error) {

	classController := controller.NewClassController(db)
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&classes); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize classes")
	}

	for _, class := range classes {
		_, err := classController.CreateClass(class)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create class")
		}
	}
	fmt.Println("Create Class Seed Successfully")
	return classes, nil
}
