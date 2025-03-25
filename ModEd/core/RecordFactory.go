package core

type RecordCreator func() *RecordInterface

type RecordFactory struct {
	mapper map[string]RecordCreator
}

var FactorySingleton *RecordFactory = &RecordFactory{}

func RegisterModel(name string, creator RecordCreator) {
	(*FactorySingleton).mapper[name] = creator
}

func CreateRecord(name string) *RecordInterface {
	creator := FactorySingleton.mapper[name]
	if creator == nil {
		return nil
	}
	return creator()
}
