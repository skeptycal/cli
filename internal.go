package cli

import (
	"errors"
	"fmt"
	"strconv"
)

func simpleEncode(b byte) string {
	// return ansiPrefix + string([]byte{b}) + ansiSuffix
	return fmt.Sprintf(fmtANSI, b)
}

func byteEncode(b Any) (byte, error) {
	switch v := b.(type) {
	case byte:
		return v, nil
	case int:
		return byte(v & 255), nil
	case uint:
		return byte(v & 255), nil
	case float32, float64:
		return byte(v.(float64)), nil
	case string:
		i, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return 0, err
		}
		return byte(i), nil
	default:
		return 0, errors.New("encoding error: unable to convert input to byte")
	}
}
