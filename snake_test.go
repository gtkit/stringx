package stringx

import "testing"

func TestToSnake(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "lower camel", in: "testCase", want: "test_case"},
		{name: "upper camel", in: "TestCase", want: "test_case"},
		{name: "spaces", in: "Test Case", want: "test_case"},
		{name: "trimmed", in: " Test Case ", want: "test_case"},
		{name: "plain", in: "test", want: "test"},
		{name: "already snake", in: "test_case", want: "test_case"},
		{name: "empty", in: "", want: ""},
		{name: "many words", in: "ManyManyWords", want: "many_many_words"},
		{name: "mixed separators", in: "AnyKind of_string", want: "any_kind_of_string"},
		{name: "numbers", in: "numbers2and55with000", want: "numbers_2_and_55_with_000"},
		{name: "json", in: "JSONData", want: "json_data"},
		{name: "user id", in: "userID", want: "user_id"},
		{name: "upper boundary", in: "AAAbbb", want: "aa_abbb"},
		{name: "digit boundary 1", in: "1A2", want: "1_a_2"},
		{name: "digit boundary 2", in: "A1B", want: "a_1_b"},
		{name: "alternating", in: "A1A2A3", want: "a_1_a_2_a_3"},
		{name: "alternating spaced", in: "A1 A2 A3", want: "a_1_a_2_a_3"},
		{name: "ab mixed digits", in: "AB1AB2AB3", want: "ab_1_ab_2_ab_3"},
		{name: "ab mixed digits spaced", in: "AB1 AB2 AB3", want: "ab_1_ab_2_ab_3"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ToSnake(tc.in)
			if got != tc.want {
				t.Fatalf("ToSnake(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}

func TestToUpperSnake(t *testing.T) {
	got := ToUpperSnake("ToUpperSnake")
	if got != "TO_UPPER_SNAKE" {
		t.Fatalf("ToUpperSnake(%q) = %q, want %q", "ToUpperSnake", got, "TO_UPPER_SNAKE")
	}
}

func TestToKebab(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{name: "camel", in: "TestCase", want: "test-case"},
		{name: "acronym", in: "JSONData", want: "json-data"},
		{name: "spaces", in: "hello world", want: "hello-world"},
		{name: "numbers", in: "Version2Build3", want: "version-2-build-3"},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ToKebab(tc.in)
			if got != tc.want {
				t.Fatalf("ToKebab(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
