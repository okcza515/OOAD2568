package util

import (
	"ModEd/asset/util"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func GetDateTimeInput(prompt string) (time.Time, error) {
	text := util.GetStringInput(prompt)
	parsedTime, err := time.Parse("2006-01-02 15:04:05", text)
	if err != nil {
		fmt.Println("Invalid date format.")
		return time.Time{}, err
	}
	return parsedTime, nil
}

func truncateStr(str string, maxLen int) string {
	if len(str) > maxLen {
		return str[:maxLen-3] + "..."
	}
	return str
}

type ColumnConfig struct {
	Header    string
	FieldName string
	Width     int
	Truncate  bool
	DateFmt   string
}

func PrintHeader(columns []ColumnConfig) {
	for _, col := range columns {
		fmt.Printf("%-*s ", col.Width, col.Header)
	}
	fmt.Println()
	for _, col := range columns {
		fmt.Printf("%-*s ", col.Width, strings.Repeat("-", len(col.Header)))
	}
	fmt.Println()
}

func PrintRow(data interface{}, columns []ColumnConfig) {
	v := reflect.ValueOf(data)

	for _, col := range columns {
		field := v.FieldByName(col.FieldName)
		if !field.IsValid() {
			fmt.Printf("%-*s ", col.Width, "")
			continue
		}

		var strVal string
		switch field.Kind() {
		case reflect.String:
			strVal = field.String()
		case reflect.Int, reflect.Int64:
			strVal = fmt.Sprintf("%d", field.Int())
		case reflect.Float64:
			strVal = fmt.Sprintf("%.2f", field.Float())
		case reflect.Struct:
			if t, ok := field.Interface().(time.Time); ok && col.DateFmt != "" {
				strVal = t.Format(col.DateFmt)
			}
		default:
			strVal = fmt.Sprintf("%v", field.Interface())
		}

		if col.Truncate {
			strVal = truncateStr(strVal, col.Width)
		}
		fmt.Printf("%-*s ", col.Width, strVal)
	}
	fmt.Println()
}

func ManageScreenWrapper(fn func() error) func() error {
	return func() error {
		err := fn()
		util.PressEnterToContinue()
		util.ClearScreen()
		return err
	}
}