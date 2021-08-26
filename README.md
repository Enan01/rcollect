# rcollect

还在为一页一页查找自己GitHub star的项目而烦恼吗？

快来使用 rcollect，一个用于收集某个github账号当前star的所有仓库的工具，提升你的幸福感❤️❤️。

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
