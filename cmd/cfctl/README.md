# cfctl

> NOTICE: The problem that viper does not partition case may cause is that the value of the key that has never been set will become lowercase.

## command

Support config type which is supported by viper.

* dotenv
* hcl
* ini
* javaproperties
* json
* toml
* yaml

Use cfctl -h to get all usage.

### auto complete command

```shell
cfctl completion zsh
cfctl completion bash
cfctl completion fish
cfctl completion powershell
```

### get

```shell
cfctl get key -f test.toml
cfctl get key.0 -f test.toml
cfctl get a.b.c -f test.toml

# specify output, may unmarshal error
cfctl get "*" -f test.toml -o yaml
cfctl get x.y -f test.toml -o json

# read from stdin
cat test.toml | cfctl get key
cat test.toml | cfctl get key -p toml
```

### set

```shell
cfctl set_string key value -f test.toml
cfctl set_int key 1 -f test.toml
cfctl set_ins key 1,2,3 -f test.toml
cfctl set_object key '{"key":"value"}' -f test.toml
cfctl set_objects key '{"key":"value"},{"key":"value"}' -f test.toml

# if index not exist in key, will occurs error.
# You can use append command to do it. But if it is a empty array, you can not append.
# In this condition, you can use set_ints and so on to initialize an array.
cfctl set_int key.0 4 -f test.toml
```

### append

```shell
cfctl append key '{"name":"value2"}' -f _example/test.json

# if array in array
cfctl append key.0.name '{"xx":"value2"}' -f _example/test.json
```

### del

```shell
cfctl del nj -f _example/test.toml
cfctl del nj.jaronnie.name -f _example/test.toml
cfctl del array.0 -f _example/test.json
```

## todo
- [x] set_string
- [x] set_int
- [x] set_float
- [x] set_object
- [x] set_bool
- [x] set_strings
- [x] set_ints
- [x] set_floats
- [x] set_objects
- [x] get
- [x] append
- [x] del