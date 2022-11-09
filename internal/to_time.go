package internal

import (
	"math"
	"reflect"
	"time"
	"unsafe"
)

var timeLayout = []string{
	time.RFC3339,
	"2006-01-02T15:04:05", // iso8601 without timezone
	time.RFC1123Z,
	time.RFC1123,
	time.RFC822Z,
	time.RFC822,
	time.RFC850,
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	"2006-01-02 15:04:05.999999999 -0700 MST", // Time.String()
	"2006-01-02",
	"02 Jan 2006",
	"2006-01-02T15:04:05-0700", // RFC3339 without timezone hh:mm colon
	"2006-01-02 15:04:05 -07:00",
	"2006-01-02 15:04:05 -0700",
	"2006-01-02 15:04:05Z07:00", // RFC3339 without T
	"2006-01-02 15:04:05Z0700",  // RFC3339 without T or timezone hh:mm colon
	"2006-01-02 15:04:05",
	time.Kitchen,
	time.Stamp,
	time.StampMilli,
	time.StampMicro,
	time.StampNano,
}

// StringToTime attempts to parse a string into a time.Time type using a
// predefined list of formats.  If no suitable format is found, an error is
// returned.
func StringToTime(in string) (ret time.Time, err error) {
	for _, layout := range timeLayout {
		if ret, err = time.Parse(layout, in); err == nil {
			return
		}
	}
	return ret, SyntaxError(in, reflect.TypeOf((*time.Time)(nil)).Elem().String())
}

func IntToTime[Input SignedInt | UnsignedInt](in Input) (ret time.Time, err error) {
	switch unsafe.Sizeof(in) {
	case 1:
		return time.Unix(int64(in), 0), nil
	case 2:
		return time.Unix(int64(in), 0), nil
	case 4:
		return time.Unix(int64(in), 0), nil
	case 8:
		if in > 0 && uint64(in) > uint64(math.MaxInt64) {
			return ret, RangeError(in, reflect.TypeOf((*time.Time)(nil)).Elem().String())
		} else {
			return time.Unix(int64(in), 0), nil
		}
	default:
		return ret, SyntaxError(in, reflect.TypeOf((*time.Time)(nil)).Elem().String())
	}
}
