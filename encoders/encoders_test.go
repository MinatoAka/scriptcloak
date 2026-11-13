package encoders

import "testing"

func TestBase64(t *testing.T) {
	got := Base64("hello")
	want := "aGVsbG8="
	if got != want {
		t.Fatalf("Base64(\"hello\") = %q, want %q", got, want)
	}
}

func TestHex(t *testing.T) {
	got := Hex("hello")
	want := "68656c6c6f"
	if got != want {
		t.Fatalf("Hex(\"hello\") = %q, want %q", got, want)
	}
}

func TestUnicode(t *testing.T) {
	got := Unicode("A🙂B")
	want := "\\u0041🙂\\u0042"
	if got != want {
		t.Fatalf("Unicode(\"A🙂B\") = %q, want %q", got, want)
	}
}
