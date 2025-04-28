package core

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type DataMapperCLI struct {
	Args []string
}

func NewDataMapperCLI(args []string) *DataMapperCLI {
	return &DataMapperCLI{
		Args: args,
	}
}

func (cli *DataMapperCLI) Run() error {
	if len(cli.Args) < 2 {
		cli.printUsage()
		return errors.New("insufficient arguments")
	}

	command := cli.Args[0]

	switch command {
	case "import":
		return cli.handleImport(cli.Args[1:])
	case "export":
		return cli.handleExport(cli.Args[1:])
	default:
		cli.printUsage()
		return fmt.Errorf("unknown command: %s", command)
	}
}

func (cli *DataMapperCLI) handleImport(args []string) error {
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	inputFile := importCmd.String("file", "", "Path to input file (CSV or JSON)")
	modelType := importCmd.String("type", "", "Type of model to import")

	err := importCmd.Parse(args)
	if err != nil {
		return err
	}

	if *inputFile == "" || *modelType == "" {
		fmt.Println("Error: Both file and type parameters are required")
		fmt.Println("Usage: import -file=<filepath> -type=<modeltype>")
		return errors.New("missing required parameters")
	}

	mapper, err := CreateMapper[interface{}](*inputFile)
	if err != nil {
		return err
	}

	fmt.Printf("Importing data from %s...\n", *inputFile)
	data := mapper.Deserialize()
	fmt.Printf("Successfully imported %d records\n", len(data))

	return nil
}

func (cli *DataMapperCLI) handleExport(args []string) error {
	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
	outputFile := exportCmd.String("file", "", "Path to output file (CSV or JSON)")
	modelType := exportCmd.String("type", "", "Type of model to export")

	err := exportCmd.Parse(args)
	if err != nil {
		return err
	}

	if *outputFile == "" || *modelType == "" {
		fmt.Println("Error: Both file and type parameters are required")
		fmt.Println("Usage: export -file=<filepath> -type=<modeltype>")
		return errors.New("missing required parameters")
	}

	fmt.Printf("Exporting data to %s...\n", *outputFile)
	// Fix: Implementation would connect to database and export data
	fmt.Println("Export functionality to be implemented based on database connection")

	return nil
}

func (cli *DataMapperCLI) printUsage() {
	execName := filepath.Base(os.Args[0])
	fmt.Printf("Usage: %s <command> [options]\n\n", execName)
	fmt.Println("Commands:")
	fmt.Println("  import -file=<filepath> -type=<modeltype>  Import data from CSV or JSON file")
	fmt.Println("  export -file=<filepath> -type=<modeltype>  Export data to CSV or JSON file")
	fmt.Println("\nExamples:")
	fmt.Println("  import -file=students.csv -type=Student")
	fmt.Println("  export -file=instructors.json -type=Instructor")
	fmt.Println("  list -type=Faculty")
}

func DataMapperCLIMain() {
	cli := NewDataMapperCLI(os.Args[1:])
	err := cli.Run()
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
