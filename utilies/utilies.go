package utilies

import (
	"reflect"
)

func Max(a, b interface{}) (interface{}, bool) {
	va := reflect.ValueOf(a)
	vb := reflect.ValueOf(b)

	if va.Kind() != vb.Kind() {
		return 0, false
	}

	switch va.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ca := va.Int()
		cb := vb.Int()
		if ca > cb {
			return ca, true
		}
		return cb, true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		ca := va.Uint()
		cb := vb.Uint()
		if ca > cb {
			return ca, true
		}
		return cb, true
	case reflect.Float32, reflect.Float64:
		ca := va.Float()
		cb := vb.Float()
		if ca > cb {
			return ca, true
		}
		return cb, true
	}
	return 0, false
}
