package base62

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func TestEncoding(t *testing.T) {
	enc := New()

	for i := 0; i < 32; i++ {
		src := make([]byte, i)
		rand.Read(src)

		s := enc.EncodeToString(src)
		fmt.Println(s)

		decoded := enc.DecodeString(s)
		fmt.Println(decoded)

		if string(src) != string(decoded) {
			t.Fatalf("expected bytes: %v, got: %v\n", src, decoded)
		}
	}
}
