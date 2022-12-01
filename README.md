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

[release](https://github.com/jaronnie/cfc/releases)

## how to use cfctl

```shell
cfctl get name -f test.toml

cfctl set_string name value -f test.toml

# if you do not set config type, we will try to set config type.
# if you have been setting config type, we will use it!
cat test.toml | cfctl get name
cat test.toml | cfctl get name -p toml
```

if you want to get more about cfctl, [click here](cmd/cfctl/README.md)

