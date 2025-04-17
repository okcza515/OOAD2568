package util
import (
	"flag"
	"fmt"
	"os"
)

func ValidateFlags(fs *flag.FlagSet, required []string) {
	var missing []string
	for _, name := range required {
		f := fs.Lookup(name)
		if f == nil || f.Value.String() == "" {
			missing = append(missing, name)
		}
	}
	if len(missing) > 0 {
		fmt.Printf("Error: Missing required flags: %v\n", missing)
		fs.Usage()
		os.Exit(1)
	}
}
