# viperfile [![Go Report Card](https://goreportcard.com/badge/github.com/ttgint/viperfile)](https://goreportcard.com/report/github.com/ttgint/viperfile)

Preconfigured Viper for reading file config

> go get github.com/ttgint/viperfile

or

> dep ensure -add github.com/ttgint/viperfile

then

`viperfile.ReadLocal("config", &Config)`

where `Config` is a pointer to the struct that the config will be bound
