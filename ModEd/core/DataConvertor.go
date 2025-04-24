package core

import "fmt"

type DataConvertor struct {
	Controller *StagingBaseController
	Mapper     DataMapper[RecordInterface]
}

func CreateConvertor[T RecordInterface](controller *StagingBaseController) *DataConvertor {
	convertor := DataConvertor{
		Controller: controller,
		Mapper:     nil,
	}
	return &convertor
}

func (convertor *DataConvertor) Convert(path string) error {
	if convertor.Mapper == nil {
		mapper, err := CreateMapper[RecordInterface](path)
		if err != nil {
			return err
		}
		convertor.Mapper = mapper
	}

	// records := convertor.Mapper.Deserialize()
	// for _, record := range records {
	// 	err := convertor.Controller.Insert(*record)
	// 	if err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func (convertor *DataConvertor) Show() error {
	records, err := convertor.Controller.List(nil)
	if err != nil {
		return err
	}
	for _, record := range records {
		fmt.Print(record.ToString())
	}
	return nil
}
