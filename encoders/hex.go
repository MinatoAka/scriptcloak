package encoders

import "encoding/hex"

func Hex(s string) string {
	return hex.EncodeToString([]byte(s))
}
