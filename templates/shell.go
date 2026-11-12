package templates

import "fmt"

func Shell(payload, method string) string {
	switch method {
	case "base64":
		return fmt.Sprintf("eval \"$(echo '%s' | base64 -d)\"", payload)
	case "hex":
		return fmt.Sprintf("eval \"$(echo '%s' | xxd -r -p)\"", payload)
	}
	return "# ERROR: method not supported for shell\n"
}
