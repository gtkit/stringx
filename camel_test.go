package stringx

import "testing"

func TestToCamel(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "snake", in: "test_case", want: "TestCase"},
		{name: "dot", in: "test.case", want: "TestCase"},
		{name: "plain", in: "test", want: "Test"},
		{name: "camel", in: "TestCase", want: "TestCase"},
		{name: "space", in: " test  case ", want: "TestCase"},
		{name: "empty", in: "", want: ""},
		{name: "many words", in: "many_many_words", want: "ManyManyWords"},
		{name: "mixed separators", in: "AnyKind of_string", want: "AnyKindOfString"},
		{name: "hyphen", in: "odd-fix", want: "OddFix"},
		{name: "numbers", in: "numbers2And55with000", want: "Numbers2And55With000"},
		{name: "acronym", in: "ID", want: "Id"},
		{name: "constant", in: "CONSTANT_CASE", want: "ConstantCase"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ToCamel(tc.in)
			if got != tc.want {
				t.Fatalf("ToCamel(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestToLowerCamel(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "hyphen", in: "foo-bar", want: "fooBar"},
		{name: "camel", in: "TestCase", want: "testCase"},
		{name: "empty", in: "", want: ""},
		{name: "mixed separators", in: "AnyKind of_string", want: "anyKindOfString"},
		{name: "dot and hyphen", in: "AnyKind.of-string", want: "anyKindOfString"},
		{name: "acronym", in: "ID", want: "id"},
		{name: "space", in: "some string", want: "someString"},
		{name: "leading space", in: " some string", want: "someString"},
		{name: "constant", in: "CONSTANT_CASE", want: "constantCase"},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			got := ToLowerCamel(tc.in)
			if got != tc.want {
				t.Fatalf("ToLowerCamel(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
