package seed

import (
	controller "ModEd/curriculum/controller/curriculum"
	"ModEd/curriculum/model"
	"ModEd/utils/deserializer"
	"fmt"

	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

func CreateCurriculumSeed(db *gorm.DB, path string) (curriculums []*model.Curriculum, err error) {

	curriculumController := controller.NewCurriculumController(db)
	deserializer, err := deserializer.NewFileDeserializer(path)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create file deserializer")
	}

	if err := deserializer.Deserialize(&curriculums); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize curriculums")
	}

	for _, curriculum := range curriculums {
		_, err := curriculumController.CreateCurriculum(curriculum)
		if err != nil {
			return nil, errors.Wrap(err, "failed to create curriculum")
		}
	}
	fmt.Println("Create Curriculum Seed Successfully")
	return curriculums, nil
}
