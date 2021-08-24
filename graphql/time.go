package graphql

import (
	"errors"
	"io"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func MarshalTime(t time.Time) Marshaler {
	if t.IsZero() {
		return Null
	}

	return WriterFunc(func(w io.Writer) {
		io.WriteString(w, strconv.Quote(t.Format(time.RFC3339)))
	})
}

func UnmarshalTime(v interface{}) (time.Time, error) {
	switch v := v.(type) {
	case string:
		return time.Parse(time.RFC3339, v)
	case *timestamppb.Timestamp:
		return v.Value
	}
	return time.Time{}, errors.New("time should be RFC3339 formatted string")
}
