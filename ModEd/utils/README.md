## Utility Function Document

### Deserializer
This function converts JSON or CSV to string so it can be parsed to GORM easier.

#### How to use:
1. Import package.
```go
//1. Import
import "ModEd/utils/deserializer"
```
2. Set path to file.
```go
deserializer, err := deserializer NewFileDeserializer(path)
if err != nil {
	//error handling
}
```
3. `courses` will be automatically generate based on given type and data.

```go
var courses []*model.Course
if err := deserializer.Deserialize(&courses); err != nil {
	//error handling
}
```
4. There are also other `deserialize` method to use

- `Deserialize()` - automatically select between CSV and JSON.
- `DeserializeCSV()` - manually select CSV.
- `DeserializeJSON()` - manually select JSON.

