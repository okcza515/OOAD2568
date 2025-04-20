package commands

import (
    "ModEd/hr/controller"
    "ModEd/hr/util"
    "fmt"
    "os"
)

func (c *MigrateStudentsCommand) Run(args []string) {
    db := util.OpenDatabase(*util.DatabasePath)

    if err := controller.MigrateStudentsToHR(db); err != nil {
        fmt.Printf("Migration failed: %v\n", err)
        os.Exit(1)
    }

    fmt.Println("Migration completed successfully!")
}