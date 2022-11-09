package cast

import (
	"github.com/toolshelf/cast/internal"
	"reflect"
)

func ToInt(i interface{}) (int, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return s, nil
	case int8:
		return int(s), nil
	case int16:
		return int(s), nil
	case int32:
		return int(s), nil
	case int64:
		return internal.IntToInt[int64, int](s)
	case uint:
		return internal.IntToInt[uint, int](s)
	case uint8:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint32:
		return internal.IntToInt[uint32, int](s)
	case uint64:
		return internal.IntToInt[uint64, int](s)
	case float32:
		return internal.FloatToInt[float32, int](s)
	case float64:
		return internal.FloatToInt[float64, int](s)
	case string:
		return internal.StringToInt[int](s)
	case bool:
		return internal.BoolToInt[int](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int.String())
	}
}

func ToInt8(i interface{}) (int8, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return internal.IntToInt[int, int8](s)
	case int8:
		return s, nil
	case int16:
		return internal.IntToInt[int16, int8](s)
	case int32:
		return internal.IntToInt[int32, int8](s)
	case int64:
		return internal.IntToInt[int64, int8](s)
	case uint:
		return internal.IntToInt[uint, int8](s)
	case uint8:
		return internal.IntToInt[uint8, int8](s)
	case uint16:
		return internal.IntToInt[uint16, int8](s)
	case uint32:
		return internal.IntToInt[uint32, int8](s)
	case uint64:
		return internal.IntToInt[uint64, int8](s)
	case float32:
		return internal.FloatToInt[float32, int8](s)
	case float64:
		return internal.FloatToInt[float64, int8](s)
	case string:
		return internal.StringToInt[int8](s)
	case bool:
		return internal.BoolToInt[int8](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int8.String())
	}
}

func ToInt16(i interface{}) (int16, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return internal.IntToInt[int, int16](s)
	case int8:
		return int16(s), nil
	case int16:
		return s, nil
	case int32:
		return internal.IntToInt[int32, int16](s)
	case int64:
		return internal.IntToInt[int64, int16](s)
	case uint:
		return internal.IntToInt[uint, int16](s)
	case uint8:
		return int16(s), nil
	case uint16:
		return internal.IntToInt[uint16, int16](s)
	case uint32:
		return internal.IntToInt[uint32, int16](s)
	case uint64:
		return internal.IntToInt[uint64, int16](s)
	case float32:
		return internal.FloatToInt[float32, int16](s)
	case float64:
		return internal.FloatToInt[float64, int16](s)
	case string:
		return internal.StringToInt[int16](s)
	case bool:
		return internal.BoolToInt[int16](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int16.String())
	}
}

func ToInt32(i interface{}) (int32, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return internal.IntToInt[int, int32](s)
	case int8:
		return int32(s), nil
	case int16:
		return int32(s), nil
	case int32:
		return s, nil
	case int64:
		return internal.IntToInt[int64, int32](s)
	case uint:
		return internal.IntToInt[uint, int32](s)
	case uint8:
		return int32(s), nil
	case uint16:
		return int32(s), nil
	case uint32:
		return internal.IntToInt[uint32, int32](s)
	case uint64:
		return internal.IntToInt[uint64, int32](s)
	case float32:
		return internal.FloatToInt[float32, int32](s)
	case float64:
		return internal.FloatToInt[float64, int32](s)
	case string:
		return internal.StringToInt[int32](s)
	case bool:
		return internal.BoolToInt[int32](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int32.String())
	}
}

func ToInt64(i interface{}) (int64, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return int64(s), nil
	case int8:
		return int64(s), nil
	case int16:
		return int64(s), nil
	case int32:
		return int64(s), nil
	case int64:
		return s, nil
	case uint:
		return internal.IntToInt[uint, int64](s)
	case uint8:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint64:
		return internal.IntToInt[uint64, int64](s)
	case float32:
		return internal.FloatToInt[float32, int64](s)
	case float64:
		return internal.FloatToInt[float64, int64](s)
	case string:
		return internal.StringToInt[int64](s)
	case bool:
		return internal.BoolToInt[int64](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int64.String())
	}
}

func ToUint(i interface{}) (uint, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return internal.IntToUint[int, uint](s)
	case int8:
		return uint(s), nil
	case int16:
		return uint(s), nil
	case int32:
		return internal.IntToUint[int32, uint](s)
	case int64:
		return internal.IntToUint[int64, uint](s)
	case uint:
		return s, nil
	case uint8:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint32:
		return uint(s), nil
	case uint64:
		return internal.IntToUint[uint64, uint](s)
	case float32:
		return internal.FloatToUint[float32, uint](s)
	case float64:
		return internal.FloatToUint[float64, uint](s)
	case string:
		return internal.StringToUint[uint](s)
	case bool:
		return internal.BoolToInt[uint](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint.String())
	}
}

func ToUint8(i interface{}) (uint8, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return internal.IntToUint[int, uint8](s)
	case int8:
		return internal.IntToUint[int8, uint8](s)
	case int16:
		return internal.IntToUint[int16, uint8](s)
	case int32:
		return internal.IntToUint[int32, uint8](s)
	case int64:
		return internal.IntToUint[int64, uint8](s)
	case uint:
		return internal.IntToUint[uint, uint8](s)
	case uint8:
		return s, nil
	case uint16:
		return internal.IntToUint[uint16, uint8](s)
	case uint32:
		return internal.IntToUint[uint32, uint8](s)
	case uint64:
		return internal.IntToUint[uint64, uint8](s)
	case float32:
		return internal.FloatToUint[float32, uint8](s)
	case float64:
		return internal.FloatToUint[float64, uint8](s)
	case string:
		return internal.StringToUint[uint8](s)
	case bool:
		return internal.BoolToInt[uint8](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint8.String())
	}
}

func ToUint16(i interface{}) (uint16, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return internal.IntToUint[int, uint16](s)
	case int8:
		return internal.IntToUint[int8, uint16](s)
	case int16:
		return internal.IntToUint[int16, uint16](s)
	case int32:
		return internal.IntToUint[int32, uint16](s)
	case int64:
		return internal.IntToUint[int64, uint16](s)
	case uint:
		return internal.IntToUint[uint, uint16](s)
	case uint8:
		return uint16(s), nil
	case uint16:
		return s, nil
	case uint32:
		return internal.IntToUint[uint32, uint16](s)
	case uint64:
		return internal.IntToUint[uint64, uint16](s)
	case float32:
		return internal.FloatToUint[float32, uint16](s)
	case float64:
		return internal.FloatToUint[float64, uint16](s)
	case string:
		return internal.StringToUint[uint16](s)
	case bool:
		return internal.BoolToInt[uint16](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint16.String())
	}
}

func ToUint32(i interface{}) (uint32, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return internal.IntToUint[int, uint32](s)
	case int8:
		return internal.IntToUint[int8, uint32](s)
	case int16:
		return internal.IntToUint[int16, uint32](s)
	case int32:
		return internal.IntToUint[int32, uint32](s)
	case int64:
		return internal.IntToUint[int64, uint32](s)
	case uint:
		return internal.IntToUint[uint, uint32](s)
	case uint8:
		return uint32(s), nil
	case uint16:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint64:
		return internal.IntToUint[uint64, uint32](s)
	case float32:
		return internal.FloatToUint[float32, uint32](s)
	case float64:
		return internal.FloatToUint[float64, uint32](s)
	case string:
		return internal.StringToUint[uint32](s)
	case bool:
		return internal.BoolToInt[uint32](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint32.String())
	}
}

func ToUint64(i interface{}) (uint64, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return internal.IntToUint[int, uint64](s)
	case int8:
		return internal.IntToUint[int8, uint64](s)
	case int16:
		return internal.IntToUint[int16, uint64](s)
	case int32:
		return internal.IntToUint[int32, uint64](s)
	case int64:
		return internal.IntToUint[int64, uint64](s)
	case uint:
		return internal.IntToUint[uint, uint64](s)
	case uint8:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint32:
		return uint64(s), nil
	case uint64:
		return s, nil
	case float32:
		return internal.FloatToUint[float32, uint64](s)
	case float64:
		return internal.FloatToUint[float64, uint64](s)
	case string:
		return internal.StringToUint[uint64](s)
	case bool:
		return internal.BoolToInt[uint64](s), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint64.String())
	}
}
