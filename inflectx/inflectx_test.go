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

func TestLiteralCustomWords(t *testing.T) {
	plurals := GetPlural()
	singulars := GetSingular()
	irregulars := GetIrregular()
	uncountables := GetUncountable()
	t.Cleanup(func() {
		SetPlural(plurals)
		SetSingular(singulars)
		SetIrregular(irregulars)
		SetUncountable(uncountables)
	})

	AddUncountable("foo.bar")
	AddIrregular("c++", "c++es")

	if got := Plural("foo.bar"); got != "foo.bar" {
		t.Fatalf("Plural returned %q, want %q", got, "foo.bar")
	}
	if got := Plural("fooXbar"); got != "fooXbars" {
		t.Fatalf("Plural returned %q, want %q", got, "fooXbars")
	}
	if got := Plural("c++"); got != "c++es" {
		t.Fatalf("Plural returned %q, want %q", got, "c++es")
	}
	if got := Singular("c++es"); got != "c++" {
		t.Fatalf("Singular returned %q, want %q", got, "c++")
	}
}
