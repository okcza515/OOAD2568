package util

import "ModEd/utils/deserializer"

func CommonDeserializer(path string) deserializer.FileDeserializer {
	deserializer, err := deserializer.NewFileDeserializer(path)
		if err != nil {
			panic(err)
		}
	return *deserializer
}