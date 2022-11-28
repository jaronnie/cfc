# cfc

cfc means config customize.

It uses [viper](https://github.com/spf13/viper) internal encoding package. Thanks for [spf13](https://github.com/spf13).

But also you can use cfctl to get value in command line or shell scripts...

## import package

```shell
import "github.com/jaronnie/cfc"
```

## how to use cfc package

[_example](_example)

## download cfctl

```shell
go install github.com/jaronnie/cfc/cmd/cfctl@latest
```

## how to use cfctl

```shell
cfctl get name -f test.toml

cfctl set_string name value -f test.toml

# default config file type is toml, if others please use -p to specify config file type
cat test.toml | cfctl get name
```

if you want to get more about cfctl, [click here](cmd/cfctl/README.md)

