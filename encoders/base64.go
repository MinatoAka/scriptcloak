package encoders

import "encoding/base64"

func Base64(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
