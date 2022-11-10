package cast

import "time"

func DirectToInt(i interface{}, expect ...int) int {
	v, err := ToInt(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToInt8(i interface{}, expect ...int8) int8 {
	v, err := ToInt8(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToInt16(i interface{}, expect ...int16) int16 {
	v, err := ToInt16(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToInt32(i interface{}, expect ...int32) int32 {
	v, err := ToInt32(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToInt64(i interface{}, expect ...int64) int64 {
	v, err := ToInt64(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToUint(i interface{}, expect ...uint) uint {
	v, err := ToUint(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToUint8(i interface{}, expect ...uint8) uint8 {
	v, err := ToUint8(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToUint16(i interface{}, expect ...uint16) uint16 {
	v, err := ToUint16(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToUint32(i interface{}, expect ...uint32) uint32 {
	v, err := ToUint32(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToUint64(i interface{}, expect ...uint64) uint64 {
	v, err := ToUint64(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToBool(i interface{}, expect ...bool) bool {
	v, err := ToBool(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToFloat32(i interface{}, expect ...float32) float32 {
	v, err := ToFloat32(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToFloat64(i interface{}, expect ...float64) float64 {
	v, err := ToFloat64(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToString(i interface{}, expect ...string) string {
	v, err := ToString(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToTime(i interface{}, expect ...time.Time) time.Time {
	v, err := ToTime(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}

func DirectToDuration(i interface{}, expect ...time.Duration) time.Duration {
	v, err := ToDuration(i)
	if err != nil && len(expect) != 0 {
		return expect[0]
	}
	return v
}
