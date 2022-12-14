package internal

import (
	"errors"
	"math"
	"reflect"
	"strconv"
	"unsafe"
)

type SignedInt interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedInt interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type Float interface {
	float32 | float64
}

type Number interface {
	SignedInt | UnsignedInt | Float
}

var MaxSafeInteger32 = 1<<24 - 1
var MinSafeInteger32 = -MaxSafeInteger32
var MaxSafeInteger64 = 1<<53 - 1
var MinSafeInteger64 = -MaxSafeInteger64

func IntToInt[Input SignedInt | UnsignedInt, Output SignedInt](in Input) (ret Output, err error) {
	switch unsafe.Sizeof(ret) {
	case 1:
		if in <= 0 && math.MinInt8 <= int64(in) || in > 0 && uint64(in) <= math.MaxInt8 {
			ret = Output(in)
		} else {
			err = RangeError(in, reflect.Int8.String())
		}
	case 2:
		if in <= 0 && math.MinInt16 <= int64(in) || in > 0 && uint64(in) <= math.MaxInt16 {
			ret = Output(in)
		} else {
			err = RangeError(in, reflect.Int16.String())
		}
	case 4:
		if in <= 0 && math.MinInt32 <= int64(in) || in > 0 && uint64(in) <= math.MaxInt32 {
			ret = Output(in)
		} else {
			err = RangeError(in, reflect.TypeOf(ret).String())
		}
	case 8:
		if in <= 0 || in > 0 && uint64(in) <= math.MaxInt64 {
			ret = Output(in)
		} else {
			err = RangeError(in, reflect.TypeOf(ret).String())
		}
	}

	return ret, err
}

func IntToUint[Input SignedInt | UnsignedInt, Output UnsignedInt](in Input) (ret Output, err error) {
	switch unsafe.Sizeof(ret) {
	case 1:
		if 0 <= in && uint64(in) <= math.MaxUint8 {
			ret = Output(in)
		} else {
			err = RangeError(in, reflect.Uint8.String())
		}
	case 2:
		if 0 <= in && uint64(in) <= math.MaxUint16 {
			ret = Output(in)
		} else {
			err = RangeError(in, reflect.Uint16.String())
		}
	case 4:
		if 0 <= in && uint64(in) <= math.MaxUint32 {
			ret = Output(in)
		} else {
			err = RangeError(in, reflect.TypeOf(ret).String())
		}
	case 8:
		if 0 <= in {
			ret = Output(in)
		} else {
			err = RangeError(in, reflect.TypeOf(ret).String())
		}
	}

	return ret, err
}

func FloatToInt[Input Float, Output SignedInt](in Input) (ret Output, err error) {
	integer, frac := math.Modf(float64(in))
	if frac != 0.0 {
		return ret, SyntaxError(in, reflect.TypeOf(ret).String())
	}
	if integer < float64(MinSafeInteger64) || float64(MaxSafeInteger64) < integer {
		return ret, SafeError(in, reflect.TypeOf(ret).String())
	}

	switch unsafe.Sizeof(ret) {
	case 1:
		if math.MinInt8 <= integer && integer <= math.MaxInt8 {
			ret = Output(integer)
		} else {
			err = RangeError(in, reflect.Int8.String())
		}
	case 2:
		if math.MinInt16 <= integer && integer <= math.MaxInt16 {
			ret = Output(integer)
		} else {
			err = RangeError(in, reflect.Int16.String())
		}
	case 4:
		if math.MinInt32 <= integer && integer <= math.MaxInt32 {
			ret = Output(integer)
		} else {
			err = RangeError(in, reflect.TypeOf(ret).String())
		}
	case 8:
		ret = Output(integer)
	}

	return ret, err
}

func FloatToUint[Input Float, Output UnsignedInt](in Input) (ret Output, err error) {
	integer, frac := math.Modf(float64(in))
	if frac != 0.0 {
		return ret, SyntaxError(in, reflect.TypeOf(ret).String())
	}
	if integer < 0 {
		return ret, RangeError(in, reflect.TypeOf(ret).String())
	}
	if integer > float64(MaxSafeInteger64) {
		return ret, SafeError(in, reflect.TypeOf(ret).String())
	}

	switch unsafe.Sizeof(ret) {
	case 1:
		if 0 <= integer && integer <= math.MaxUint8 {
			ret = Output(integer)
		} else {
			err = RangeError(in, reflect.Uint8.String())
		}
	case 2:
		if 0 <= integer && integer <= math.MaxUint16 {
			ret = Output(integer)
		} else {
			err = RangeError(in, reflect.Uint16.String())
		}
	case 4:
		if 0 <= integer && integer <= math.MaxUint32 {
			ret = Output(integer)
		} else {
			err = RangeError(in, reflect.TypeOf(ret).String())
		}
	case 8:
		ret = Output(integer)
	}

	return ret, err
}

func BoolToNumber[Output Number](in bool) Output {
	if in {
		return Output(1)
	}
	return Output(0)
}

func StringToInt[Output SignedInt](in string) (ret Output, err error) {
	if i, inErr := strconv.ParseInt(in, 10, 64); inErr == nil {
		if ret, inErr = IntToInt[int64, Output](i); inErr == nil {
			return ret, nil
		} else if errors.Is(inErr, ErrRange) {
			return ret, RangeError(in, reflect.TypeOf(ret).String())
		}
	} else if errors.Is(inErr, strconv.ErrRange) {
		return ret, RangeError(in, reflect.TypeOf(ret).String())
	}

	if f, inErr := strconv.ParseFloat(in, 64); inErr == nil {
		if ret, inErr = FloatToInt[float64, Output](f); inErr == nil {
			return ret, nil
		} else if errors.Is(inErr, ErrSafe) {
			return ret, SafeError(in, reflect.TypeOf(ret).String())
		} else if errors.Is(inErr, ErrRange) {
			return ret, RangeError(in, reflect.TypeOf(ret).String())
		}
	} else if errors.Is(inErr, strconv.ErrRange) {
		return ret, RangeError(in, reflect.TypeOf(ret).String())
	}

	return ret, SyntaxError(in, reflect.TypeOf(ret).String())
}

func StringToUint[Output UnsignedInt](in string) (ret Output, err error) {
	if i, inErr := strconv.ParseInt(in, 10, 64); inErr == nil {
		if ret, inErr = IntToUint[int64, Output](i); inErr == nil {
			return ret, nil
		} else if errors.Is(inErr, ErrRange) {
			return 0, RangeError(in, reflect.TypeOf(ret).String())
		}
	} else if errors.Is(inErr, strconv.ErrRange) {
		return 0, RangeError(in, reflect.TypeOf(ret).String())
	}

	if f, inErr := strconv.ParseFloat(in, 64); inErr == nil {
		if ret, inErr = FloatToUint[float64, Output](f); inErr == nil {
			return ret, nil
		} else if errors.Is(inErr, ErrSafe) {
			return 0, SafeError(in, reflect.TypeOf(ret).String())
		} else if errors.Is(inErr, ErrRange) {
			return 0, RangeError(in, reflect.TypeOf(ret).String())
		}
	} else if errors.Is(inErr, strconv.ErrRange) {
		return 0, RangeError(in, reflect.TypeOf(ret).String())
	}

	return 0, SyntaxError(in, reflect.TypeOf(ret).String())
}

func StringToFloat[Output Float](in string) (ret Output, err error) {
	var size = int(unsafe.Sizeof(ret))
	if f, inErr := strconv.ParseFloat(in, size*8); inErr == nil {
		return Output(f), nil
	} else if errors.Is(inErr, strconv.ErrRange) {
		return 0, RangeError(in, reflect.TypeOf(ret).String())
	}
	return 0, SyntaxError(in, reflect.TypeOf(ret).String())
}

func IntToFloat[Input SignedInt | UnsignedInt, Output Float](in Input) (ret Output, err error) {
	switch unsafe.Sizeof(ret) {
	case 4:
		if in < 0 && int64(in) < int64(MinSafeInteger32) || in > 0 && uint64(in) > uint64(MaxSafeInteger32) {
			err = SafeError(in, reflect.TypeOf(ret).String())
		} else {
			ret = Output(in)
		}
	case 8:
		if in < 0 && int64(in) < int64(MinSafeInteger64) || in > 0 && uint64(in) > uint64(MaxSafeInteger64) {
			err = SafeError(in, reflect.TypeOf(ret).String())
		} else {
			ret = Output(in)
		}
	}

	return ret, err
}
