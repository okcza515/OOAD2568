// MEP-1005

package core

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type MenuIO struct {
	Reader *bufio.Reader
	Writer io.Writer
}

func (io *MenuIO) Println(msg string) {
	fmt.Fprintln(io.Writer, msg)
}

func (io *MenuIO) Print(msg string) {
	fmt.Fprint(io.Writer, msg)
}

func (io *MenuIO) ReadInput() (string, error) {
	line, err := io.Reader.ReadString('\n')
	return strings.TrimSpace(line), err
}

func (io *MenuIO) ReadInputFloat() (float64, error) {
	input, err := io.ReadInput()
	if err != nil {
		return 0, err
	}

	return strconv.ParseFloat(input, 64)
}

func (io *MenuIO) ReadInputTime() (time.Time, error) {
	input, _ := io.ReadInput()
	return time.Parse("2006-01-02", input)
}

func (io *MenuIO) ReadInputID() (uint, error) {
	input, err := io.ReadInput()
	if err != nil {
		return 0, err
	}

	if input == "-1" {
		io.Println("Cancelled.")
		return 0, errors.New("canceled")
	}

	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return uint(id), nil
}

func (io *MenuIO) ReadYesOrNo() (bool, error) {
	input, err := io.ReadInput()
	if err != nil {
		return false, err
	}

	return strings.ToLower(input) == "yes" || strings.ToLower(input) == "y", nil
}

func flattenFields(typ reflect.Type, prefix string, path []int, fieldMap map[string][]int) {
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)

		if field.PkgPath != "" {
			continue
		}

		fieldName := field.Name
		if prefix != "" {
			fieldName = prefix + "." + fieldName
		}

		if field.Anonymous && field.Type.Kind() == reflect.Struct {
			flattenFields(field.Type, prefix, append(path, i), fieldMap)
		} else {
			fieldMap[fieldName] = append(path, i)
			fieldMap[field.Name] = append(path, i)
		}
	}
}

func (io *MenuIO) PrintTableFromSlice(slice interface{}, columns []string) {
	val := reflect.ValueOf(slice)
	if val.Kind() != reflect.Slice || val.Len() == 0 {
		io.Println("Error: Expected a non-empty slice of structs")
		return
	}

	elem := val.Index(0)
	if elem.Kind() == reflect.Ptr {
		elem = elem.Elem()
	}
	typ := elem.Type()

	fieldMap := make(map[string][]int)
	flattenFields(typ, "", []int{}, fieldMap)

	var colPaths [][]int
	var validColumns []string
	for _, col := range columns {
		if path, ok := fieldMap[col]; ok {
			validColumns = append(validColumns, col)
			colPaths = append(colPaths, path)
		}
	}
	if len(validColumns) == 0 {
		io.Println("No valid columns to display.")
		return
	}

	var rows [][]string
	for i := 0; i < val.Len(); i++ {
		rowVal := val.Index(i)
		if rowVal.Kind() == reflect.Ptr {
			rowVal = rowVal.Elem()
		}

		var row []string
		for _, idxPath := range colPaths {
			field := rowVal.FieldByIndex(idxPath)
			row = append(row, fmt.Sprintf("%v", field.Interface()))
		}
		rows = append(rows, row)
	}

	io.PrintTable(validColumns, rows)
}

func (io *MenuIO) PrintTable(headers []string, rows [][]string) {
	colWidths := make([]int, len(headers))
	for i, header := range headers {
		colWidths[i] = len(header)
	}
	for _, row := range rows {
		for i, col := range row {
			if len(col) > colWidths[i] {
				colWidths[i] = len(col)
			}
		}
	}

	for i, header := range headers {
		fmt.Fprintf(io.Writer, "%-*s  ", colWidths[i], header)
	}
	fmt.Fprintln(io.Writer)

	for _, width := range colWidths {
		fmt.Fprintf(io.Writer, "%s  ", strings.Repeat("-", width))
	}
	fmt.Fprintln(io.Writer)

	for _, row := range rows {
		for i, col := range row {
			fmt.Fprintf(io.Writer, "%-*s  ", colWidths[i], col)
		}
		fmt.Fprintln(io.Writer)
	}
}
