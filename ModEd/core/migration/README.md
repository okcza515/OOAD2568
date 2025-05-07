# Core Migration Manager

This module provides a centralized way to handle database migrations and seed data using GORM (with SQLite) for the ModEd system.

## Adding New Migration Strategies

To enable migrations for a new module:

1. Implement a MigrationStrategy for the module.
```go
package migration

// MEP-0069 Example

import (
	"ModEd/example/model"
)

type ExampleMigrationStrategy struct {
}

func (s *ExampleMigrationStrategy) GetModels() []interface{} {
	return []interface{}{
		&model.Example1{},
		&model.Example2{},
		&model.Example3{},
	}
}

```

2. Register it in `newMigrationManager()` by replacing `nil` with your` new strategy`:

```go
migrationMap[core.MODULE_EXAMPLE] = &ExampleMigrationStrategy{}
```

## Get the MigrationManager Instance
1. Use the singleton pattern:

```go
mgr := migration.GetInstance()
```

2. Set the Database Path (Optional)

```go
mgr.SetPathDB("path/to/your/custom.db")
```

If not set, the default is: `data/ModEd.bin`

3. Register and Migrate Modules

Select which modules you want to migrate by calling MigrateModule().
Example for migrating the Asset and Curriculum modules:

```go
mgr.MigrateModule(core.MODULE_ASSET).
    MigrateModule(core.MODULE_CURRICULUM)
```

4. Build (Initialize) the Database

```go
db, err := mgr.BuildDB()
if err != nil {
    panic("Failed to build DB: " + err.Error())
}
```

## Add Seed Data

```go
mgr.AddSeedData("path/to/seed_file.json", &YourModelStruct{})

err = mgr.LoadSeedData()
    if err != nil {
    panic("Failed to load seed data: " + err.Error())
}
```

##  Drop All Tables

```go
err = mgr.DropAllTables()
if err != nil {
    panic("Failed to drop tables: " + err.Error())
}
```

## Full Flow

```go
mgr := migration.GetInstance().
    SetPathDB("data/mydatabase.db").
    MigrateModule(core.MODULE_ASSET).
    MigrateModule(core.MODULE_CURRICULUM)

db, err := mgr.BuildDB()
if err != nil {
    panic(err)
}

mgr.AddSeedData("seeds/assets.json", &[]Asset{})
err = mgr.LoadSeedData()
if err != nil {
    panic(err)
}

```
