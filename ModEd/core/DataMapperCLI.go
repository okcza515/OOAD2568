// MEP-1001
package core

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
)

// fix: Please Check @NoCODEEE
// ========== Strategy Pattern Interfaces ==========

// LoadStrategy defines a strategy for loading data from a file
type LoadStrategy interface {
	Load(recordType reflect.Type) (interface{}, error)
}

// DisplayStrategy defines a strategy for displaying data
type DisplayStrategy interface {
	Display(results interface{})
}

// ========== Concrete Strategy Implementations ==========

// CSVLoadStrategy implements LoadStrategy for CSV files
type CSVLoadStrategy struct {
	FilePath string
}

func (s *CSVLoadStrategy) Load(recordType reflect.Type) (interface{}, error) {
	mapper := &CSVMapper[any]{Path: s.FilePath}
	deserializeMethod := reflect.ValueOf(mapper).MethodByName("Deserialize")
	result := deserializeMethod.Call(nil)[0].Interface()
	return result, nil
}

// JSONLoadStrategy implements LoadStrategy for JSON files
type JSONLoadStrategy struct {
	FilePath string
}

func (s *JSONLoadStrategy) Load(recordType reflect.Type) (interface{}, error) {
	mapper := &JSONMapper[any]{Path: s.FilePath}
	deserializeMethod := reflect.ValueOf(mapper).MethodByName("Deserialize")
	result := deserializeMethod.Call(nil)[0].Interface()
	return result, nil
}

// TextDisplayStrategy implements DisplayStrategy for text output
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

// JSONDisplayStrategy implements DisplayStrategy for JSON output
type JSONDisplayStrategy struct{}

func (s *JSONDisplayStrategy) Display(results interface{}) {
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling to JSON: %v\n", err)
		return
	}
	fmt.Println(string(jsonData))
}

// ========== Factory Methods ==========

// LoadStrategyFactory creates the appropriate LoadStrategy based on file extension
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

// DisplayStrategyFactory creates the appropriate DisplayStrategy based on format
type DisplayStrategyFactory struct{}

func (f *DisplayStrategyFactory) CreateDisplayStrategy(format string) DisplayStrategy {
	if format == "json" {
		return &JSONDisplayStrategy{}
	}
	// Default to text display
	return &TextDisplayStrategy{}
}

// ========== Main CLI Implementation ==========

// DataMapperCLI provides a command line interface for interacting with DataMapper
type DataMapperCLI struct {
	FilePath       string
	OutputFormat   string
	loadFactory    LoadStrategyFactory
	displayFactory DisplayStrategyFactory
}

// NewDataMapperCLI creates a new CLI instance with the specified configuration
func NewDataMapperCLI() *DataMapperCLI {
	return &DataMapperCLI{
		loadFactory:    LoadStrategyFactory{},
		displayFactory: DisplayStrategyFactory{},
	}
}

// ParseArgs parses command line arguments
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

// PrintUsage prints usage information
func (cli *DataMapperCLI) PrintUsage() {
	fmt.Println("Usage: program -file=<file_path> [-format=<format>]")
	fmt.Println("Options:")
	fmt.Println("  -file     Path to data file (.csv or .json)")
	fmt.Println("  -format   Output format: 'text' (default) or 'json'")
}

// Run executes the CLI application with the specified data type
func (cli *DataMapperCLI) Run(recordType reflect.Type) {
	cli.ParseArgs()

	// Create load strategy using factory
	loadStrategy, err := cli.loadFactory.CreateLoadStrategy(cli.FilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Create display strategy using factory
	displayStrategy := cli.displayFactory.CreateDisplayStrategy(cli.OutputFormat)

	// Load data
	results, err := loadStrategy.Load(recordType)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Display results
	displayStrategy.Display(results)
}
