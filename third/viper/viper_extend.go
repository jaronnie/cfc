package viper

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"strconv"
	"strings"
)

func SetIndex(key string, value interface{}) error { return v.SetIndex(key, value) }

func (v *Viper) SetIndex(key string, value interface{}) error {
	// If alias passed in, then set the proper override
	key = v.realKey(strings.ToLower(key))
	value = toCaseInsensitiveValue(value)

	path := strings.Split(key, v.keyDelim)
	lastKey := strings.ToLower(path[len(path)-1])
	configMap := deepSearchStrong(v.config, path[0:len(path)-1])

	switch x := configMap.(type) {
	case map[string]interface{}:
		x[lastKey] = value
	case []interface{}:
		// is this an array
		// lastKey has to be a num
		idx, err := strconv.Atoi(lastKey)
		if err == nil {
			if idx < len(x) && idx >= 0 {
				x[idx] = value
			} else {
				return errors.Errorf("not index [%d] in array", idx)
			}
		}
	}
	return nil
}

func UnSet(key string) { v.UnSet(key) }

// UnSet https://github.com/spf13/viper/pull/519
// support any type unset even if index value in array
func (v *Viper) UnSet(key string) {
	// If alias passed in, then set the proper override
	key = v.realKey(strings.ToLower(key))

	path := strings.Split(key, v.keyDelim)
	lastKey := strings.ToLower(path[len(path)-1])
	configMap := deepSearchStrong(v.config, path[0:len(path)-1])

	switch x := configMap.(type) {
	case map[string]interface{}:
		delete(x, lastKey)
	case []interface{}:
		idx, err := strconv.Atoi(lastKey)
		if err == nil {
			if idx < len(x) && idx >= 0 {
				x = append(x[:idx], x[idx+1:]...)
				split := strings.Split(key, v.keyDelim)
				for i, s := range split {
					if _, err := cast.ToIntE(s); err == nil {
						join := strings.Join(split[:i], v.keyDelim)
						v.Set(join, x)
					}
				}
			} else {
				return
			}
		}
	}
}

func deepSearchStrong(m map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return m
	}
	var currentPath string
	var stepArray = false
	var currentArray []interface{}
	var currentEntity interface{}
	for _, k := range path {
		if len(currentPath) == 0 {
			currentPath = k
		} else {
			currentPath = fmt.Sprintf("%v.%v", currentPath, k)
		}
		if stepArray {
			idx, err := strconv.Atoi(k)
			if err != nil {
				return nil
			}
			if len(currentArray) <= idx {
				return nil
			}
			m3, ok := currentArray[idx].(map[string]interface{})
			if !ok {
				return nil
			}
			// continue search from here
			m = m3
			currentEntity = m
			stepArray = false // don't support arrays of arrays
		} else {
			m2, ok := m[k]
			if !ok {
				// intermediate key does not exist
				return nil
			}
			m3, ok := m2.(map[string]interface{})
			if !ok {
				// is this an array
				m4, ok := m2.([]interface{})
				if ok {
					// continue search from here
					currentArray = m4
					currentEntity = currentArray
					stepArray = true
					m3 = nil
				} else {
					// intermediate key is a value
					return nil
				}
			} else {
				// continue search from here
				m = m3
				currentEntity = m
			}
		}
	}

	return currentEntity
}
