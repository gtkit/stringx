package inflectx

import (
	"strings"
	"testing"
)

func TestInflection(t *testing.T) {
	if got := Plural("person"); got != "people" {
		t.Fatalf("Plural returned %q, want %q", got, "people")
	}

	if got := Singular("people"); got != "person" {
		t.Fatalf("Singular returned %q, want %q", got, "person")
	}

	if got := Plural(strings.ToUpper("person")); got != "PEOPLE" {
		t.Fatalf("Plural upper returned %q, want %q", got, "PEOPLE")
	}

	AddIrregular("analysis", "analyses")
	if got := Plural("analysis"); got != "analyses" {
		t.Fatalf("Plural irregular returned %q, want %q", got, "analyses")
	}
}
