package graphql

import (
	"fmt"
	"io"
	"strings"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

func MarshalBoolean(b bool) Marshaler {
	return WriterFunc(func(w io.Writer) {
		if b {
			w.Write(trueLit)
		} else {
			w.Write(falseLit)
		}
	})
}

func UnmarshalBoolean(v interface{}) (bool, error) {
	switch v := v.(type) {
	case string:
		return strings.ToLower(v) == "true", nil
	case int:
		return v != 0, nil
	case bool:
		return v, nil
	case *wrapperspb.Bool:
		return v.Value, nil
	default:
		return false, fmt.Errorf("%T is not a bool", v)
	}
}
