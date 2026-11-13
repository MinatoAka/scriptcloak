package templates

import "testing"

func TestJS(t *testing.T) {
	payload := "aGVsbG8="

	if got, want := JS(payload, "base64"), `eval(new TextDecoder().decode(Uint8Array.from(atob("aGVsbG8="),c=>c.charCodeAt(0))))` ; got != want {
		t.Fatalf("JS(base64) = %q, want %q", got, want)
	}

	if got, want := JS(payload, "hex"), `eval(new TextDecoder().decode(new Uint8Array("aGVsbG8=".match(/.{1,2}/g).map(h=>parseInt(h,16)))))` ; got != want {
		t.Fatalf("JS(hex) = %q, want %q", got, want)
	}

	if got, want := JS(payload, "unicode"), `eval("aGVsbG8=")` ; got != want {
		t.Fatalf("JS(unicode) = %q, want %q", got, want)
	}

	if got := JS(payload, "invalid"); got != "" {
		t.Fatalf("JS(invalid) = %q, want empty string", got)
	}
}

func TestPython(t *testing.T) {
	payload := "aGVsbG8="

	if got, want := Python(payload, "base64"), `exec(__import__('base64').b64decode('aGVsbG8=').decode())` ; got != want {
		t.Fatalf("Python(base64) = %q, want %q", got, want)
	}

	if got, want := Python(payload, "hex"), `exec(bytes.fromhex('aGVsbG8=').decode())` ; got != want {
		t.Fatalf("Python(hex) = %q, want %q", got, want)
	}

	if got, want := Python(payload, "unicode"), `exec('aGVsbG8=')` ; got != want {
		t.Fatalf("Python(unicode) = %q, want %q", got, want)
	}

	if got := Python(payload, "invalid"); got != "" {
		t.Fatalf("Python(invalid) = %q, want empty string", got)
	}
}

func TestShell(t *testing.T) {
	payload := "aGVsbG8="

	if got, want := Shell(payload, "base64"), "eval \"$(echo 'aGVsbG8=' | base64 -d)\"" ; got != want {
		t.Fatalf("Shell(base64) = %q, want %q", got, want)
	}

	if got, want := Shell(payload, "hex"), "eval \"$(echo 'aGVsbG8=' | xxd -r -p)\"" ; got != want {
		t.Fatalf("Shell(hex) = %q, want %q", got, want)
	}

	if got, want := Shell(payload, "invalid"), "# ERROR: method not supported for shell\n" ; got != want {
		t.Fatalf("Shell(invalid) = %q, want %q", got, want)
	}
}

func TestPHP(t *testing.T) {
	payload := "aGVsbG8="

	if got, want := PHP(payload, "base64"), "<?php eval(base64_decode('aGVsbG8=')); ?>" ; got != want {
		t.Fatalf("PHP(base64) = %q, want %q", got, want)
	}

	if got, want := PHP(payload, "hex"), "<?php eval(hex2bin('aGVsbG8=')); ?>" ; got != want {
		t.Fatalf("PHP(hex) = %q, want %q", got, want)
	}

	if got, want := PHP(payload, "invalid"), "<?php echo 'ERROR: method not supported for PHP'; ?>" ; got != want {
		t.Fatalf("PHP(invalid) = %q, want %q", got, want)
	}
}
