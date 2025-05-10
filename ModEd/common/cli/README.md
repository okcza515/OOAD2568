## Command and arguments for CommonDataCLI

1. cd to ModEd

```sh
cd ModEd
```

2. run main

```go
go run common/cli/CommonDataCLI.go common/cli/MenuItemHandlers.go common/cli/CLIFunction.go --database="data/ModEd.bin" --path="data/StudentList.csv" <first arg> <second arg> <...> ...
// (--database is not require but data's --path is necessary)
```
