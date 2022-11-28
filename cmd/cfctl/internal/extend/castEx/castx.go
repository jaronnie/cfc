package castEx

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/spf13/cast"
)

// ToFloat64SliceE casts an interface to a []float64 type.
func ToFloat64SliceE(i interface{}) ([]float64, error) {
	if i == nil {
		return []float64{}, fmt.Errorf("unable to cast %#v of type %T to []float4", i, i)
	}

	switch v := i.(type) {
	case []float64:
		return v, nil
	}

	kind := reflect.TypeOf(i).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		s := reflect.ValueOf(i)
		a := make([]float64, s.Len())
		for j := 0; j < s.Len(); j++ {
			val, err := cast.ToFloat64E(s.Index(j).Interface())
			if err != nil {
				return []float64{}, fmt.Errorf("unable to cast %#v of type %T to []float64", i, i)
			}
			a[j] = val
		}
		return a, nil
	default:
		return []float64{}, fmt.Errorf("unable to cast %#v of type %T to []float64", i, i)
	}
}

// ToStringMapSliceE casts an interface to a map[string][]interface{} type.
func ToStringMapSliceE(i interface{}) ([]map[string]interface{}, error) {
	var m []map[string]interface{}

	switch v := i.(type) {
	case []map[string]interface{}:
		return v, nil
	case []string:
		for _, i := range v {
			var mm map[string]interface{}
			err := jsonStringToObject(i, &mm)
			if err != nil {
				return nil, err
			}
			m = append(m, mm)
		}
		return m, nil
	default:
		return m, fmt.Errorf("unable to cast %#v of type %T to []map[string]interface{}", i, i)
	}
}

// jsonStringToObject attempts to unmarshall a string as JSON into
// the object passed as pointer.
func jsonStringToObject(s string, v interface{}) error {
	data := []byte(s)
	return json.Unmarshal(data, v)
}
