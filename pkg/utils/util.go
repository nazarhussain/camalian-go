package util

import (
	"fmt"
	"math"
	"reflect"
)

// Find min and max values from given values
func FindMinMax(a ...interface{}) (interface{}, interface{}) {
	var max interface{}
	var min interface{}

	switch a[0].(type) {
	case int:
		max = 0
		min = math.MaxInt16

		for _, value := range a {
			if value.(int) < min.(int) {
				min = value
			}
			if value.(int) > max.(int) {
				max = value
			}
		}

	case uint8:
		max = uint8(0)
		min = uint8(255)

		for _, value := range a {
			if value.(uint8) < min.(uint8) {
				min = value
			}
			if value.(uint8) > max.(uint8) {
				max = value
			}
		}

	case uint16:
		max = uint16(0)
		min = uint16(math.MaxUint16)

		for _, value := range a {
			if value.(uint16) < min.(uint16) {
				min = value
			}
			if value.(uint16) > max.(uint16) {
				max = value
			}
		}

	case float32:
		max = float32(0.0)
		min = float32(math.MaxFloat32)

		for _, value := range a {
			if value.(float32) < min.(float32) {
				min = value
			}
			if value.(float32) > max.(float32) {
				max = value
			}
		}

	case float64:
		max = float64(0.0)
		min = float64(math.MaxFloat32)

		for _, value := range a {
			if value.(float64) < min.(float64) {
				min = value
			}
			if value.(float64) > max.(float64) {
				max = value
			}
		}

	default:
		panic(fmt.Sprintf("Can not find min/max for type %T", a[0]))
	}

	return min, max
}

// Find an object from a slice with given selector function
func Find(slice interface{}, f func(value interface{}) bool) int {
	s := reflect.ValueOf(slice)
	if s.Kind() == reflect.Slice {
		for index := 0; index < s.Len(); index++ {
			if f(s.Index(index).Interface()) {
				return index
			}
		}
	}
	return -1
}

// Calculate sum for the whole slice values
func Sum(vs []interface{}) interface{} {
	sum := 0.0

	for _, v := range vs {
		switch i := v.(type) {
		case int:
			sum += float64(i)
		case uint:
			sum += float64(i)
		case uint8:
			sum += float64(i)
		case uint16:
			sum += float64(i)
		case float32:
			sum += float64(i)
		case float64:
			sum += i
		default:
			fmt.Printf("Mismatch Type %T\n", v)
		}
	}

	return sum
}

// Collect object from slice by given collection function
func Collect(vs []interface{}, f func(interface{}) interface{}) []interface{} {
	vsm := make([]interface{}, len(vs))

	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
