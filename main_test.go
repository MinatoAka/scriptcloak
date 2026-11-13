package main

import "testing"

func TestSupportedCombos(t *testing.T) {
	tests := []struct {
		name   string
		lang   string
		method string
		want   bool
	}{
		{name: "javascript base64", lang: "js", method: "base64", want: true},
		{name: "javascript hex", lang: "js", method: "hex", want: true},
		{name: "javascript unicode", lang: "js", method: "unicode", want: true},
		{name: "python base64", lang: "python", method: "base64", want: true},
		{name: "python hex", lang: "python", method: "hex", want: true},
		{name: "python unicode", lang: "python", method: "unicode", want: true},
		{name: "shell base64", lang: "shell", method: "base64", want: true},
		{name: "shell hex", lang: "shell", method: "hex", want: true},
		{name: "php base64", lang: "php", method: "base64", want: true},
		{name: "php hex", lang: "php", method: "hex", want: true},
		{name: "shell unicode is unsupported", lang: "shell", method: "unicode", want: false},
		{name: "php unicode is unsupported", lang: "php", method: "unicode", want: false},
		{name: "unknown language", lang: "ruby", method: "base64", want: false},
		{name: "unknown method", lang: "js", method: "rot13", want: false},
		{name: "empty language", lang: "", method: "base64", want: false},
		{name: "empty method", lang: "js", method: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := supportedCombos(tt.lang, tt.method); got != tt.want {
				t.Errorf("supportedCombos(%q, %q) = %t, want %t", tt.lang, tt.method, got, tt.want)
			}
		})
	}
}
