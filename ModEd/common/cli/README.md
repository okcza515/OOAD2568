## Command and arguments for CommonDataCLI

1. cd to ModEd

```sh
cd ModEd
```

2. run main

```go
go run common/cli/CommonDataCLI.go --database="data/ModEd.bin" --path="data/[Filename]" <first arg> <second arg> <...> ...
// (--database is not require but --path is necessary)
```
