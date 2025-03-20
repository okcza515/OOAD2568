package deserializer

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/cockroachdb/errors"
	"github.com/gocarina/gocsv"
)

type FileDeserializer struct {
	Path string
}

func NewFileDeserializer(path string) (*FileDeserializer, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, errors.Wrapf(err, "file does not exist: %s", path)
	}
	return &FileDeserializer{Path: path}, nil
}

func (d *FileDeserializer) Deserialize(data interface{}) error {
	ext := strings.ToLower(filepath.Ext(d.Path))
	switch ext {
	case ".csv":
		return d.DeserializeCSV(data)
	case ".json":
		return d.DeserializeJSON(data)
	default:
		return errors.Newf("unsupported file format: %s", ext)
	}
}

func (d *FileDeserializer) DeserializeCSV(data interface{}) error {
	file, err := os.OpenFile(d.Path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "failed to open CSV file: %s", d.Path)
	}
	defer file.Close()

	if err := gocsv.UnmarshalFile(file, data); err != nil {
		return errors.Wrap(err, "failed to unmarshal CSV file")
	}

	return nil
}

func (d *FileDeserializer) DeserializeJSON(data interface{}) error {
	file, err := os.OpenFile(d.Path, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return errors.Wrapf(err, "failed to open JSON file: %s", d.Path)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(data); err != nil {
		return errors.Wrap(err, "failed to unmarshal JSON file")
	}

	return nil
}
