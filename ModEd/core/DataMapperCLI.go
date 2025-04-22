package core

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

type LoadStrategy interface {
	Load(recordType reflect.Type) (interface{}, error)
}

type DisplayStrategy interface {
	Display(results interface{})
}

type CSVLoadStrategy struct {
	FilePath string
}

func (s *CSVLoadStrategy) Load(recordType reflect.Type) (interface{}, error) {
	mapper := &CSVMapper[any]{Path: s.FilePath}
	deserializeMethod := reflect.ValueOf(mapper).MethodByName("Deserialize")
	result := deserializeMethod.Call(nil)[0].Interface()
	return result, nil
}

type JSONLoadStrategy struct {
	FilePath string
}

func (s *JSONLoadStrategy) Load(recordType reflect.Type) (interface{}, error) {
	mapper := &JSONMapper[any]{Path: s.FilePath}
	deserializeMethod := reflect.ValueOf(mapper).MethodByName("Deserialize")
	result := deserializeMethod.Call(nil)[0].Interface()
	return result, nil
}

type TextDisplayStrategy struct{}

func (s *TextDisplayStrategy) Display(results interface{}) {
	resultsValue := reflect.ValueOf(results)

	if resultsValue.Kind() != reflect.Slice {
		fmt.Println("Error: Results are not in expected format")
		return
	}

	length := resultsValue.Len()
	fmt.Printf("Found %d records:\n", length)

	for i := 0; i < length; i++ {
		record := resultsValue.Index(i).Interface()
		fmt.Printf("Record %d: %+v\n", i+1, record)
	}
}

type JSONDisplayStrategy struct{}

func (s *JSONDisplayStrategy) Display(results interface{}) {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

type LoadStrategyFactory struct{}

func (f *LoadStrategyFactory) CreateLoadStrategy(filePath string) (LoadStrategy, error) {
	ext := filepath.Ext(filePath)

	if ext == ".csv" {
		return &CSVLoadStrategy{FilePath: filePath}, nil
	} else if ext == ".json" {
		return &JSONLoadStrategy{FilePath: filePath}, nil
	}

	return nil, fmt.Errorf("unsupported file extension: %s", ext)
}

type DisplayStrategyFactory struct{}

func (f *DisplayStrategyFactory) CreateDisplayStrategy(format string) DisplayStrategy {
	if format == "json" {
		return &JSONDisplayStrategy{}
	}
	return &TextDisplayStrategy{}
}

type DataMapperCLI struct {
	FilePath       string
	OutputFormat   string
	loadFactory    LoadStrategyFactory
	displayFactory DisplayStrategyFactory
}

func NewDataMapperCLI() *DataMapperCLI {
	return &DataMapperCLI{
		loadFactory:    LoadStrategyFactory{},
		displayFactory: DisplayStrategyFactory{},
	}
}

func (cli *DataMapperCLI) ParseArgs() {
	flag.StringVar(&cli.FilePath, "file", "", "Path to data file (.csv or .json)")
	flag.StringVar(&cli.OutputFormat, "format", "text", "Output format (text or json)")
	flag.Parse()

	if cli.FilePath == "" {
		fmt.Println("Error: File path is required")
		cli.PrintUsage()
		os.Exit(1)
	}
}

func (cli *DataMapperCLI) PrintUsage() {
	fmt.Println("Usage: program -file=<file_path> [-format=<format>]")
	fmt.Println("Options:")
	fmt.Println("  -file     Path to data file (.csv or .json)")
	fmt.Println("  -format   Output format: 'text' (default) or 'json'")
}

func (cli *DataMapperCLI) Run(recordType reflect.Type) {
	cli.ParseArgs()

	loadStrategy, err := cli.loadFactory.CreateLoadStrategy(cli.FilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	displayStrategy := cli.displayFactory.CreateDisplayStrategy(cli.OutputFormat)

	results, err := loadStrategy.Load(recordType)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	displayStrategy.Display(results)
}
