package stringx

import (
	"strings"
	"sync"
	"testing"
)

func TestPluralAndSingular(t *testing.T) {
	cases := []struct {
		name        string
		singular    string
		plural      string
		upperPlural string
	}{
		{name: "person", singular: "person", plural: "people", upperPlural: "PEOPLE"},
		{name: "child", singular: "child", plural: "children", upperPlural: "CHILDREN"},
		{name: "bus", singular: "bus", plural: "buses", upperPlural: "BUSES"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if got := Plural(tc.singular); got != tc.plural {
				t.Fatalf("Plural(%q) = %q, want %q", tc.singular, got, tc.plural)
			}

			if got := Singular(tc.plural); got != tc.singular {
				t.Fatalf("Singular(%q) = %q, want %q", tc.plural, got, tc.singular)
			}

			if got := Plural(strings.ToUpper(tc.singular)); got != tc.upperPlural {
				t.Fatalf("Plural(%q) = %q, want %q", strings.ToUpper(tc.singular), got, tc.upperPlural)
			}
		})
	}
}

func TestAddIrregular(t *testing.T) {
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

	AddIrregular("analysis", "analyses")

	if got := Plural("analysis"); got != "analyses" {
		t.Fatalf("Plural(%q) = %q, want %q", "analysis", got, "analyses")
	}

	if got := Singular("analyses"); got != "analysis" {
		t.Fatalf("Singular(%q) = %q, want %q", "analyses", got, "analysis")
	}
}

func TestSetRulesWithExportedFields(t *testing.T) {
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

	SetPlural(RegularSlice{{Find: "(quiz)$", Replace: "${1}zes"}})
	SetSingular(RegularSlice{{Find: "(quiz)zes$", Replace: "${1}"}})
	SetIrregular(IrregularSlice{{Singular: "mouse", Plural: "mice"}})
	SetUncountable([]string{"metadata"})

	if got := Plural("quiz"); got != "quizzes" {
		t.Fatalf("Plural(%q) = %q, want %q", "quiz", got, "quizzes")
	}

	if got := Singular("quizzes"); got != "quiz" {
		t.Fatalf("Singular(%q) = %q, want %q", "quizzes", got, "quiz")
	}

	if got := Plural("mouse"); got != "mice" {
		t.Fatalf("Plural(%q) = %q, want %q", "mouse", got, "mice")
	}

	if got := Singular("mice"); got != "mouse" {
		t.Fatalf("Singular(%q) = %q, want %q", "mice", got, "mouse")
	}

	if got := Plural("metadata"); got != "metadata" {
		t.Fatalf("Plural(%q) = %q, want %q", "metadata", got, "metadata")
	}
}

func TestInflectionConcurrentReadWrite(t *testing.T) {
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

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				_ = Plural("person")
				_ = Singular("people")
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			AddUncountable("metadata")
		}
	}()

	wg.Wait()
}
