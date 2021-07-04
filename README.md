# rcollect

一个用于收集某个github账号当前star的所有仓库的工具。

## 编译

```shell
go build -o ./rollect ./cmd/...
```

## 如何使用

```shell
$ ./rcollect -h
collect github star repositorys, and export to csv file.

Usage:
  rcollect [flags]

Flags:
  -a, --account string   github account
  -h, --help             help for rcollect
  -o, --output string    file output path (default "./")
```