package cast

import "time"

func MustToInt(i interface{}) int {
	v, err := ToInt(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToInt8(i interface{}) int8 {
	v, err := ToInt8(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToInt16(i interface{}) int16 {
	v, err := ToInt16(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToInt32(i interface{}) int32 {
	v, err := ToInt32(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToInt64(i interface{}) int64 {
	v, err := ToInt64(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToUint(i interface{}) uint {
	v, err := ToUint(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToUint8(i interface{}) uint8 {
	v, err := ToUint8(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToUint16(i interface{}) uint16 {
	v, err := ToUint16(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToUint32(i interface{}) uint32 {
	v, err := ToUint32(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToUint64(i interface{}) uint64 {
	v, err := ToUint64(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToBool(i interface{}) bool {
	v, err := ToBool(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToFloat32(i interface{}) float32 {
	v, err := ToFloat32(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToFloat64(i interface{}) float64 {
	v, err := ToFloat64(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToString(i interface{}) string {
	v, err := ToString(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToTime(i interface{}) time.Time {
	v, err := ToTime(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToDuration(i interface{}) time.Duration {
	v, err := ToDuration(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToStringMapString(i interface{}) map[string]string {
	v, err := ToStringMapString(i)
	if err != nil {
		panic(err)
	}
	return v
}

func MustToStringMapAny(i interface{}) map[string]interface{} {
	v, err := ToStringMapAny(i)
	if err != nil {
		panic(err)
	}
	return v
}
