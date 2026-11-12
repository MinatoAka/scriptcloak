package templates

import "fmt"

func JS(payload, method string) string {
	switch method {
	case "base64":
		return fmt.Sprintf(
			`eval(new TextDecoder().decode(Uint8Array.from(atob("%s"),c=>c.charCodeAt(0))))`,
			payload,
		)
	case "hex":
		return fmt.Sprintf(
			`eval(new TextDecoder().decode(new Uint8Array("%s".match(/.{1,2}/g).map(h=>parseInt(h,16)))))`,
			payload,
		)
	case "unicode":
		return fmt.Sprintf(`eval("%s")`, payload)
	}
	return ""
}
