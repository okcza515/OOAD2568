package utils

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
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

type MenuItem struct {
	Title    string
	Action   func(io *MenuIO)
	Children []*MenuItem
}

func (m *MenuItem) AddChild(child *MenuItem) {
	m.Children = append(m.Children, child)
}

type MenuBuilder struct {
	Root *MenuItem
	IO   *MenuIO
}

func NewMenuBuilder(root *MenuItem, reader io.Reader, writer io.Writer) *MenuBuilder {
	if reader == nil {
		reader = os.Stdin
	}
	if writer == nil {
		writer = os.Stdout
	}

	return &MenuBuilder{
		Root: root,
		IO: &MenuIO{
			Reader: bufio.NewReader(reader),
			Writer: writer,
		},
	}
}

func (mb *MenuBuilder) AddMenuPath(path []string, action func(io *MenuIO)) {
	current := mb.Root

	for _, title := range path {
		found := false
		for _, child := range current.Children {
			if child.Title == title {
				current = child
				found = true
				break
			}
		}

		if !found {
			newItem := &MenuItem{Title: title}
			current.AddChild(newItem)
			current = newItem
		}
	}

	current.Action = action
}

func (mb *MenuBuilder) Show() {
	mb.showMenu(mb.Root, "")
}

func (mb *MenuBuilder) showMenu(item *MenuItem, path string) {
	for {
		fmt.Fprintf(mb.IO.Writer, "\n%s> %s\n", path, item.Title)
		for i, child := range item.Children {
			fmt.Fprintf(mb.IO.Writer, "[%d] %s\n", i+1, child.Title)
		}
		if path != "" {
			fmt.Fprint(mb.IO.Writer, "[0] Back\n")
		} else {
			fmt.Fprint(mb.IO.Writer, "[0] Exit\n")
		}

		fmt.Fprint(mb.IO.Writer, "Select an option: ")
		choice, _ := mb.IO.ReadInput()

		index, err := strconv.Atoi(choice)
		if err != nil || index < 0 || index > len(item.Children) {
			mb.IO.Println("Invalid input, try again.")
			continue
		}

		if index == 0 {
			break
		}

		selected := item.Children[index-1]
		if len(selected.Children) > 0 {
			mb.showMenu(selected, path+"/"+selected.Title)
		} else if selected.Action != nil {
			selected.Action(mb.IO)
		} else {
			mb.IO.Println("No action assigned.")
		}
	}
}
