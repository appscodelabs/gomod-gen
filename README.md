# gomod-gen

## Generate / Update go.mod.json

```console
./kubernetes/v1.18.3/update.sh
```

## Generate / Update go.mod file for any Go module

```console
cd $module_root
gomod-gen --desired-gomod=https://github.com/appscodelabs/gomod-gen/raw/master/kubernetes/v1.18.3/go.mod
go mod tidy
go mod vendor # if vendor wanted
```
