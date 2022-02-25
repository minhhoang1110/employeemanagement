package types

import (
	"strconv"
	"time"
)

// UnixTime defines a timestamp encoded as epoch seconds in JSON
// swagger:strfmt datetime
type UnixTime time.Time

func TimeNow() UnixTime {
	t := time.Now()
	return *(*UnixTime)(&t)
}

func UnixTimeFromTime(t time.Time) UnixTime {
	return *(*UnixTime)(&t)
}

// MarshalJSON is used to convert the timestamp to JSON
func (t UnixTime) MarshalJSON() ([]byte, error) {
	if time.Time(t).Unix() > 0 {
		return []byte(strconv.FormatInt(time.Time(t).Unix(), 10)), nil
	}
	return []byte("0"), nil
}

// UnmarshalJSON is used to convert the timestamp from JSON
func (t *UnixTime) UnmarshalJSON(s []byte) (err error) {
	r := string(s)
	q, err := strconv.ParseInt(r, 10, 64)
	if err != nil {
		return err
	}
	*(*time.Time)(t) = time.Unix(q, 0)
	return nil
}

// String returns the RFC1123Z representation of the timestamp
func (t UnixTime) String() string {
	return t.Time().Local().Format(time.RFC1123Z)
}

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
func (t UnixTime) Unix() int64 {
	return time.Time(t).Unix()
}

// Time returns the JSON time as a time.Time
func (t UnixTime) Time() time.Time {
	return time.Time(t)
}
