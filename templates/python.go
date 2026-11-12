package templates

import "fmt"

func Python(payload, method string) string {
	switch method {
	case "base64":
		return fmt.Sprintf(`exec(__import__('base64').b64decode('%s').decode())`, payload)
	case "hex":
		return fmt.Sprintf(`exec(bytes.fromhex('%s').decode())`, payload)
	case "unicode":
		return fmt.Sprintf(`exec('%s')`, payload)
	}
	return ""
}
