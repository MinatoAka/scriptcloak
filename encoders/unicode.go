package encoders

import (
	"fmt"
	"strings"
)

// Unicode mengubah setiap karakter menjadi \uXXXX.
// Hanya mendukung BMP (Basic Multilingual Plane, < 0x10000).
// Karakter di atas BMP akan di-skip (jarang dipakai di script pada umumnya).
func Unicode(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r > 0xFFFF {
			b.WriteRune(r) // biarkan apa adanya
			continue
		}
		fmt.Fprintf(&b, "\\u%04X", r)
	}
	return b.String()
}
