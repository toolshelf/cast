package cast

import (
	"errors"
	"fmt"
	"github.com/toolshelf/cast/internal"
	"math"
	"strconv"
	"unsafe"
)

var errNegativeNotAllowed = errors.New("unable to cast negative value")
var errOutOfRange = errors.New("out of range")
var errIncludeDecimals = errors.New("float number include decimals")

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
		return IntToInt[int64, int](s)
	case uint:
		return IntToInt[uint, int](s)
	case uint8:
		return int(s), nil
	case uint16:
		return int(s), nil
	case uint32:
		return IntToInt[uint32, int](s)
	case uint64:
		return IntToInt[uint64, int](s)
	case float32:
		return FloatToInt[float32, int](s)
	case float64:
		return FloatToInt[float64, int](s)
	case string:
		return StringToInt[int](s)
	case bool:
		return BoolToInt[int](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

func ToInt8(i interface{}) (int8, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return IntToInt[int, int8](s)
	case int8:
		return s, nil
	case int16:
		return IntToInt[int16, int8](s)
	case int32:
		return IntToInt[int32, int8](s)
	case int64:
		return IntToInt[int64, int8](s)
	case uint:
		return IntToInt[uint, int8](s)
	case uint8:
		return IntToInt[uint8, int8](s)
	case uint16:
		return IntToInt[uint16, int8](s)
	case uint32:
		return IntToInt[uint32, int8](s)
	case uint64:
		return IntToInt[uint64, int8](s)
	case float32:
		return FloatToInt[float32, int8](s)
	case float64:
		return FloatToInt[float64, int8](s)
	case string:
		return StringToInt[int8](s)
	case bool:
		return BoolToInt[int8](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int8", i, i)
	}
}

func ToInt16(i interface{}) (int16, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return IntToInt[int, int16](s)
	case int8:
		return int16(s), nil
	case int16:
		return s, nil
	case int32:
		return IntToInt[int32, int16](s)
	case int64:
		return IntToInt[int64, int16](s)
	case uint:
		return IntToInt[uint, int16](s)
	case uint8:
		return int16(s), nil
	case uint16:
		return IntToInt[uint16, int16](s)
	case uint32:
		return IntToInt[uint32, int16](s)
	case uint64:
		return IntToInt[uint64, int16](s)
	case float32:
		return FloatToInt[float32, int16](s)
	case float64:
		return FloatToInt[float64, int16](s)
	case string:
		return StringToInt[int16](s)
	case bool:
		return BoolToInt[int16](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int16", i, i)
	}
}

func ToInt32(i interface{}) (int32, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return IntToInt[int, int32](s)
	case int8:
		return int32(s), nil
	case int16:
		return int32(s), nil
	case int32:
		return s, nil
	case int64:
		return IntToInt[int64, int32](s)
	case uint:
		return IntToInt[uint, int32](s)
	case uint8:
		return int32(s), nil
	case uint16:
		return int32(s), nil
	case uint32:
		return IntToInt[uint32, int32](s)
	case uint64:
		return IntToInt[uint64, int32](s)
	case float32:
		return FloatToInt[float32, int32](s)
	case float64:
		return FloatToInt[float64, int32](s)
	case string:
		return StringToInt[int32](s)
	case bool:
		return BoolToInt[int32](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int32", i, i)
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
		return IntToInt[uint, int64](s)
	case uint8:
		return int64(s), nil
	case uint16:
		return int64(s), nil
	case uint32:
		return int64(s), nil
	case uint64:
		return IntToInt[uint64, int64](s)
	case float32:
		return FloatToInt[float32, int64](s)
	case float64:
		return FloatToInt[float64, int64](s)
	case string:
		return StringToInt[int64](s)
	case bool:
		return BoolToInt[int64](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int64", i, i)
	}
}

func ToUint(i interface{}) (uint, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return IntToUint[int, uint](s)
	case int8:
		return uint(s), nil
	case int16:
		return uint(s), nil
	case int32:
		return IntToUint[int32, uint](s)
	case int64:
		return IntToUint[int64, uint](s)
	case uint:
		return s, nil
	case uint8:
		return uint(s), nil
	case uint16:
		return uint(s), nil
	case uint32:
		return uint(s), nil
	case uint64:
		return IntToUint[uint64, uint](s)
	case float32:
		return FloatToUint[float32, uint](s)
	case float64:
		return FloatToUint[float64, uint](s)
	case string:
		return StringToUint[uint](s)
	case bool:
		return BoolToInt[uint](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint", i, i)
	}
}

func ToUint8(i interface{}) (uint8, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return IntToUint[int, uint8](s)
	case int8:
		return IntToUint[int8, uint8](s)
	case int16:
		return IntToUint[int16, uint8](s)
	case int32:
		return IntToUint[int32, uint8](s)
	case int64:
		return IntToUint[int64, uint8](s)
	case uint:
		return IntToUint[uint, uint8](s)
	case uint8:
		return s, nil
	case uint16:
		return IntToUint[uint16, uint8](s)
	case uint32:
		return IntToUint[uint32, uint8](s)
	case uint64:
		return IntToUint[uint64, uint8](s)
	case float32:
		return FloatToUint[float32, uint8](s)
	case float64:
		return FloatToUint[float64, uint8](s)
	case string:
		return StringToUint[uint8](s)
	case bool:
		return BoolToInt[uint8](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint8", i, i)
	}
}

func ToUint16(i interface{}) (uint16, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return IntToUint[int, uint16](s)
	case int8:
		return IntToUint[int8, uint16](s)
	case int16:
		return IntToUint[int16, uint16](s)
	case int32:
		return IntToUint[int32, uint16](s)
	case int64:
		return IntToUint[int64, uint16](s)
	case uint:
		return IntToUint[uint, uint16](s)
	case uint8:
		return uint16(s), nil
	case uint16:
		return s, nil
	case uint32:
		return IntToUint[uint32, uint16](s)
	case uint64:
		return IntToUint[uint64, uint16](s)
	case float32:
		return FloatToUint[float32, uint16](s)
	case float64:
		return FloatToUint[float64, uint16](s)
	case string:
		return StringToUint[uint16](s)
	case bool:
		return BoolToInt[uint16](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint16", i, i)
	}
}

func ToUint32(i interface{}) (uint32, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return IntToUint[int, uint32](s)
	case int8:
		return IntToUint[int8, uint32](s)
	case int16:
		return IntToUint[int16, uint32](s)
	case int32:
		return IntToUint[int32, uint32](s)
	case int64:
		return IntToUint[int64, uint32](s)
	case uint:
		return IntToUint[uint, uint32](s)
	case uint8:
		return uint32(s), nil
	case uint16:
		return uint32(s), nil
	case uint32:
		return s, nil
	case uint64:
		return IntToUint[uint64, uint32](s)
	case float32:
		return FloatToUint[float32, uint32](s)
	case float64:
		return FloatToUint[float64, uint32](s)
	case string:
		return StringToUint[uint32](s)
	case bool:
		return BoolToInt[uint32](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint32", i, i)
	}
}

func ToUint64(i interface{}) (uint64, error) {
	i = internal.Indirect(i)
	switch s := i.(type) {
	case int:
		return IntToUint[int, uint64](s)
	case int8:
		return IntToUint[int8, uint64](s)
	case int16:
		return IntToUint[int16, uint64](s)
	case int32:
		return IntToUint[int32, uint64](s)
	case int64:
		return IntToUint[int64, uint64](s)
	case uint:
		return IntToUint[uint, uint64](s)
	case uint8:
		return uint64(s), nil
	case uint16:
		return uint64(s), nil
	case uint32:
		return uint64(s), nil
	case uint64:
		return s, nil
	case float32:
		return FloatToUint[float32, uint64](s)
	case float64:
		return FloatToUint[float64, uint64](s)
	case string:
		return StringToUint[uint64](s)
	case bool:
		return BoolToInt[uint64](s), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to uint64", i, i)
	}
}

type SignedInt interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

func IntToInt[Input SignedInt | UnsignedInt, Output SignedInt](in Input) (ret Output, err error) {
	switch unsafe.Sizeof(ret) {
	case 1:
		if math.MinInt8 <= in && in <= math.MaxInt8 {
			ret = Output(in)
		} else {
			err = errOutOfRange
		}
	case 2:
		if math.MinInt16 <= in && in <= math.MaxInt16 {
			ret = Output(in)
		} else {
			err = errOutOfRange
		}
	case 4:
		if math.MinInt32 <= in && in <= math.MaxInt32 {
			ret = Output(in)
		} else {
			err = errOutOfRange
		}
	case 8:
		ret = Output(in)
	}

	return ret, err
}

func IntToUint[Input SignedInt | UnsignedInt, Output UnsignedInt](in Input) (ret Output, err error) {
	switch unsafe.Sizeof(ret) {
	case 1:
		if 0 <= in && in <= math.MaxUint8 {
			ret = Output(in)
		} else {
			err = errOutOfRange
		}
	case 2:
		if 0 <= in && in <= math.MaxUint16 {
			ret = Output(in)
		} else {
			err = errOutOfRange
		}
	case 4:
		if 0 <= in && in <= math.MaxUint32 {
			ret = Output(in)
		} else {
			err = errOutOfRange
		}
	case 8:
		if 0 <= in && in <= math.MaxUint64 {
			ret = Output(in)
		} else {
			err = errOutOfRange
		}
	}

	return ret, err
}

func FloatToInt[Input Float, Output SignedInt](in Input) (ret Output, err error) {
	integer, frac := math.Modf(float64(in))
	if frac != 0.0 {
		return ret, errIncludeDecimals
	}

	switch unsafe.Sizeof(ret) {
	case 1:
		if math.MinInt8 <= integer && integer <= math.MaxInt8 {
			ret = Output(integer)
		} else {
			err = errOutOfRange
		}
	case 2:
		if math.MinInt16 <= integer && integer <= math.MaxInt16 {
			ret = Output(integer)
		} else {
			err = errOutOfRange
		}
	case 4:
		if math.MinInt32 <= integer && integer <= math.MaxInt32 {
			ret = Output(integer)
		} else {
			err = errOutOfRange
		}
	case 8:
		if math.MinInt64 <= integer && integer <= math.MaxInt64 {
			ret = Output(integer)
		} else {
			err = errOutOfRange
		}
	}

	return ret, err
}

func FloatToUint[Input Float, Output UnsignedInt](in Input) (ret Output, err error) {
	integer, frac := math.Modf(float64(in))
	if frac != 0.0 {
		return ret, errIncludeDecimals
	}

	switch unsafe.Sizeof(ret) {
	case 1:
		if 0 <= integer && integer <= math.MaxUint8 {
			ret = Output(integer)
		} else {
			err = errOutOfRange
		}
	case 2:
		if 0 <= integer && integer <= math.MaxUint16 {
			ret = Output(integer)
		} else {
			err = errOutOfRange
		}
	case 4:
		if 0 <= integer && integer <= math.MaxUint32 {
			ret = Output(integer)
		} else {
			err = errOutOfRange
		}
	case 8:
		if 0 <= integer && integer <= math.MaxUint64 {
			ret = Output(integer)
		} else {
			err = errOutOfRange
		}
	}

	return ret, err
}

func BoolToInt[Output SignedInt | UnsignedInt](in bool) Output {
	if in {
		return 1
	}
	return 0
}

func StringToInt[Output SignedInt](in string) (ret Output, err error) {
	if v, err := strconv.ParseInt(in, 0, 0); err == nil {
		return IntToInt[int64, Output](v)
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to %T", in, in, ret)
}

func StringToUint[Output UnsignedInt](in string) (ret Output, err error) {
	if v, err := strconv.ParseInt(in, 0, 0); err == nil {
		return IntToUint[int64, Output](v)
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to %T", in, in, ret)
}
