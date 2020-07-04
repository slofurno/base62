package base62

import (
	"crypto/rand"
	"testing"
)

func TestEncoding(t *testing.T) {
	enc := New(defaultEncoding)

	for i := 0; i < 32; i++ {
		src := make([]byte, i)
		rand.Read(src)

		s := enc.EncodeToString(src)

		if len(s) < i || len(s) > 2*i {
			t.Fatalf("unexpected encoded string length: %s (%d)", s, len(s))
		}

		decoded := enc.DecodeString(s)

		if string(src) != string(decoded) {
			t.Fatalf("expected bytes: %v, got: %v\n", src, decoded)
		}
	}
}
