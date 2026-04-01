package unsafex

import (
	"bytes"
	"testing"
)

func TestUnsafeConverters(t *testing.T) {
	if got := Bytes2String([]byte("你好")); got != "你好" {
		t.Fatalf("Bytes2String returned %q, want %q", got, "你好")
	}

	if got := String2Bytes("hello"); !bytes.Equal(got, []byte("hello")) {
		t.Fatalf("String2Bytes returned %q, want %q", got, []byte("hello"))
	}
}
