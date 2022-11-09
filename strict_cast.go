package cast

import (
	"fmt"
	"github.com/toolshelf/cast/internal"
	"html/template"
	"math"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func ToInt(i interface{}) (int, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		return internal.IntToInt[int64, int](v)
	case uint:
		return internal.IntToInt[uint, int](v)
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		return internal.IntToInt[uint32, int](v)
	case uint64:
		return internal.IntToInt[uint64, int](v)
	case float32:
		return internal.FloatToInt[float32, int](v)
	case float64:
		return internal.FloatToInt[float64, int](v)
	case string:
		return internal.StringToInt[int](v)
	case bool:
		return internal.BoolToNumber[int](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int.String())
	}
}

func ToInt8(i interface{}) (int8, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return internal.IntToInt[int, int8](v)
	case int8:
		return v, nil
	case int16:
		return internal.IntToInt[int16, int8](v)
	case int32:
		return internal.IntToInt[int32, int8](v)
	case int64:
		return internal.IntToInt[int64, int8](v)
	case uint:
		return internal.IntToInt[uint, int8](v)
	case uint8:
		return internal.IntToInt[uint8, int8](v)
	case uint16:
		return internal.IntToInt[uint16, int8](v)
	case uint32:
		return internal.IntToInt[uint32, int8](v)
	case uint64:
		return internal.IntToInt[uint64, int8](v)
	case float32:
		return internal.FloatToInt[float32, int8](v)
	case float64:
		return internal.FloatToInt[float64, int8](v)
	case string:
		return internal.StringToInt[int8](v)
	case bool:
		return internal.BoolToNumber[int8](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int8.String())
	}
}

func ToInt16(i interface{}) (int16, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return internal.IntToInt[int, int16](v)
	case int8:
		return int16(v), nil
	case int16:
		return v, nil
	case int32:
		return internal.IntToInt[int32, int16](v)
	case int64:
		return internal.IntToInt[int64, int16](v)
	case uint:
		return internal.IntToInt[uint, int16](v)
	case uint8:
		return int16(v), nil
	case uint16:
		return internal.IntToInt[uint16, int16](v)
	case uint32:
		return internal.IntToInt[uint32, int16](v)
	case uint64:
		return internal.IntToInt[uint64, int16](v)
	case float32:
		return internal.FloatToInt[float32, int16](v)
	case float64:
		return internal.FloatToInt[float64, int16](v)
	case string:
		return internal.StringToInt[int16](v)
	case bool:
		return internal.BoolToNumber[int16](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int16.String())
	}
}

func ToInt32(i interface{}) (int32, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return internal.IntToInt[int, int32](v)
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		return internal.IntToInt[int64, int32](v)
	case uint:
		return internal.IntToInt[uint, int32](v)
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		return internal.IntToInt[uint32, int32](v)
	case uint64:
		return internal.IntToInt[uint64, int32](v)
	case float32:
		return internal.FloatToInt[float32, int32](v)
	case float64:
		return internal.FloatToInt[float64, int32](v)
	case string:
		return internal.StringToInt[int32](v)
	case bool:
		return internal.BoolToNumber[int32](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int32.String())
	}
}

func ToInt64(i interface{}) (int64, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		return internal.IntToInt[uint, int64](v)
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		return internal.IntToInt[uint64, int64](v)
	case float32:
		return internal.FloatToInt[float32, int64](v)
	case float64:
		return internal.FloatToInt[float64, int64](v)
	case string:
		return internal.StringToInt[int64](v)
	case bool:
		return internal.BoolToNumber[int64](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Int64.String())
	}
}

func ToUint(i interface{}) (uint, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return internal.IntToUint[int, uint](v)
	case int8:
		return uint(v), nil
	case int16:
		return uint(v), nil
	case int32:
		return internal.IntToUint[int32, uint](v)
	case int64:
		return internal.IntToUint[int64, uint](v)
	case uint:
		return v, nil
	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint64:
		return internal.IntToUint[uint64, uint](v)
	case float32:
		return internal.FloatToUint[float32, uint](v)
	case float64:
		return internal.FloatToUint[float64, uint](v)
	case string:
		return internal.StringToUint[uint](v)
	case bool:
		return internal.BoolToNumber[uint](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint.String())
	}
}

func ToUint8(i interface{}) (uint8, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return internal.IntToUint[int, uint8](v)
	case int8:
		return internal.IntToUint[int8, uint8](v)
	case int16:
		return internal.IntToUint[int16, uint8](v)
	case int32:
		return internal.IntToUint[int32, uint8](v)
	case int64:
		return internal.IntToUint[int64, uint8](v)
	case uint:
		return internal.IntToUint[uint, uint8](v)
	case uint8:
		return v, nil
	case uint16:
		return internal.IntToUint[uint16, uint8](v)
	case uint32:
		return internal.IntToUint[uint32, uint8](v)
	case uint64:
		return internal.IntToUint[uint64, uint8](v)
	case float32:
		return internal.FloatToUint[float32, uint8](v)
	case float64:
		return internal.FloatToUint[float64, uint8](v)
	case string:
		return internal.StringToUint[uint8](v)
	case bool:
		return internal.BoolToNumber[uint8](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint8.String())
	}
}

func ToUint16(i interface{}) (uint16, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return internal.IntToUint[int, uint16](v)
	case int8:
		return internal.IntToUint[int8, uint16](v)
	case int16:
		return internal.IntToUint[int16, uint16](v)
	case int32:
		return internal.IntToUint[int32, uint16](v)
	case int64:
		return internal.IntToUint[int64, uint16](v)
	case uint:
		return internal.IntToUint[uint, uint16](v)
	case uint8:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint32:
		return internal.IntToUint[uint32, uint16](v)
	case uint64:
		return internal.IntToUint[uint64, uint16](v)
	case float32:
		return internal.FloatToUint[float32, uint16](v)
	case float64:
		return internal.FloatToUint[float64, uint16](v)
	case string:
		return internal.StringToUint[uint16](v)
	case bool:
		return internal.BoolToNumber[uint16](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint16.String())
	}
}

func ToUint32(i interface{}) (uint32, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return internal.IntToUint[int, uint32](v)
	case int8:
		return internal.IntToUint[int8, uint32](v)
	case int16:
		return internal.IntToUint[int16, uint32](v)
	case int32:
		return internal.IntToUint[int32, uint32](v)
	case int64:
		return internal.IntToUint[int64, uint32](v)
	case uint:
		return internal.IntToUint[uint, uint32](v)
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint64:
		return internal.IntToUint[uint64, uint32](v)
	case float32:
		return internal.FloatToUint[float32, uint32](v)
	case float64:
		return internal.FloatToUint[float64, uint32](v)
	case string:
		return internal.StringToUint[uint32](v)
	case bool:
		return internal.BoolToNumber[uint32](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint32.String())
	}
}

func ToUint64(i interface{}) (uint64, error) {
	i = internal.Indirect(i)
	switch v := i.(type) {
	case int:
		return internal.IntToUint[int, uint64](v)
	case int8:
		return internal.IntToUint[int8, uint64](v)
	case int16:
		return internal.IntToUint[int16, uint64](v)
	case int32:
		return internal.IntToUint[int32, uint64](v)
	case int64:
		return internal.IntToUint[int64, uint64](v)
	case uint:
		return internal.IntToUint[uint, uint64](v)
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case float32:
		return internal.FloatToUint[float32, uint64](v)
	case float64:
		return internal.FloatToUint[float64, uint64](v)
	case string:
		return internal.StringToUint[uint64](v)
	case bool:
		return internal.BoolToNumber[uint64](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Uint64.String())
	}
}

func ToFloat32(i interface{}) (float32, error) {
	i = internal.Indirect(i)

	switch v := i.(type) {
	case float64:
		if v < -math.MaxFloat32 || v > math.MaxFloat32 {
			return 0, internal.RangeError(i, reflect.Float32.String())
		}
		return float32(v), nil
	case float32:
		return v, nil
	case int:
		return internal.IntToFloat[int, float32](v)
	case int64:
		return internal.IntToFloat[int64, float32](v)
	case int32:
		return internal.IntToFloat[int32, float32](v)
	case int16:
		return internal.IntToFloat[int16, float32](v)
	case int8:
		return internal.IntToFloat[int8, float32](v)
	case uint:
		return internal.IntToFloat[uint, float32](v)
	case uint64:
		return internal.IntToFloat[uint64, float32](v)
	case uint32:
		return internal.IntToFloat[uint32, float32](v)
	case uint16:
		return internal.IntToFloat[uint16, float32](v)
	case uint8:
		return internal.IntToFloat[uint8, float32](v)
	case string:
		return internal.StringToFloat[float32](v)
	case bool:
		return internal.BoolToNumber[float32](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Float32.String())
	}
}

func ToFloat64(i interface{}) (float64, error) {
	i = internal.Indirect(i)

	switch v := i.(type) {
	case float64:
		return v, nil
	case float32:
		return float64(v), nil
	case int:
		return internal.IntToFloat[int, float64](v)
	case int64:
		return internal.IntToFloat[int64, float64](v)
	case int32:
		return internal.IntToFloat[int32, float64](v)
	case int16:
		return internal.IntToFloat[int16, float64](v)
	case int8:
		return internal.IntToFloat[int8, float64](v)
	case uint:
		return internal.IntToFloat[uint, float64](v)
	case uint64:
		return internal.IntToFloat[uint64, float64](v)
	case uint32:
		return internal.IntToFloat[uint32, float64](v)
	case uint16:
		return internal.IntToFloat[uint16, float64](v)
	case uint8:
		return internal.IntToFloat[uint8, float64](v)
	case string:
		return internal.StringToFloat[float64](v)
	case bool:
		return internal.BoolToNumber[float64](v), nil
	default:
		return 0, internal.SyntaxError(i, reflect.Float64.String())
	}
}

func ToString(i interface{}) (string, error) {
	i = internal.IndirectToStringerOrError(i)

	switch v := i.(type) {
	case string:
		return v, nil
	case bool:
		return strconv.FormatBool(v), nil
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64), nil
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32), nil
	case int:
		return strconv.Itoa(v), nil
	case int64:
		return strconv.FormatInt(v, 10), nil
	case int32:
		return strconv.Itoa(int(v)), nil
	case int16:
		return strconv.FormatInt(int64(v), 10), nil
	case int8:
		return strconv.FormatInt(int64(v), 10), nil
	case uint:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint64:
		return strconv.FormatUint(v, 10), nil
	case uint32:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint16:
		return strconv.FormatUint(uint64(v), 10), nil
	case uint8:
		return strconv.FormatUint(uint64(v), 10), nil
	case []byte:
		return string(v), nil
	case template.HTML:
		return string(v), nil
	case template.URL:
		return string(v), nil
	case template.JS:
		return string(v), nil
	case template.CSS:
		return string(v), nil
	case template.HTMLAttr:
		return string(v), nil
	case fmt.Stringer:
		return v.String(), nil
	case error:
		return v.Error(), nil
	default:
		return "", internal.SyntaxError(i, reflect.String.String())
	}
}

func ToTime(i interface{}) (time.Time, error) {
	i = internal.Indirect(i)

	switch v := i.(type) {
	case time.Time:
		return v, nil
	case string:
		return internal.StringToTime(v)
	case int:
		return internal.IntToTime[int](v)
	case int32:
		return internal.IntToTime[int32](v)
	case int64:
		return internal.IntToTime[int64](v)
	case uint:
		return internal.IntToTime[uint](v)
	case uint32:
		return internal.IntToTime[uint32](v)
	case uint64:
		return internal.IntToTime[uint64](v)
	default:
		return time.Time{}, internal.SyntaxError(i, reflect.TypeOf((*time.Time)(nil)).Elem().String())
	}
}

func ToDuration(i interface{}) (d time.Duration, err error) {
	i = internal.Indirect(i)

	switch v := i.(type) {
	case time.Duration:
		return v, err
	case int, int64, int32, int16, int8, uint, uint64, uint32, uint16, uint8, float32, float64:
		if l, inErr := ToInt64(v); inErr == nil {
			return time.Duration(l), nil
		} else {
			return 0, internal.RangeError(i, reflect.TypeOf((*time.Duration)(nil)).Elem().String())
		}
	case string:
		// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
		if strings.ContainsAny(v, "smh nuµ") {
			if d, err = time.ParseDuration(v); err == nil {
				return d, err
			}
		} else {
			if d, err = time.ParseDuration(v + "ns"); err == nil {
				return d, err
			}
		}
	}

	return 0, internal.SyntaxError(i, reflect.TypeOf((*time.Duration)(nil)).Elem().String())
}
