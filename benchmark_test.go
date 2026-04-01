package stringx

import "testing"

func BenchmarkToCamel(b *testing.B) {
	for b.Loop() {
		_ = ToCamel("many_many_words")
	}
}

func BenchmarkToLowerCamel(b *testing.B) {
	for b.Loop() {
		_ = ToLowerCamel("many_many_words")
	}
}

func BenchmarkToSnake(b *testing.B) {
	for b.Loop() {
		_ = ToSnake("ManyManyWords123")
	}
}

func BenchmarkBuilderJoin(b *testing.B) {
	values := []string{"hello", " ", "world", "!", " stringx"}
	for b.Loop() {
		_ = BuilderJoin(values)
	}
}

func BenchmarkNormalizeSpace(b *testing.B) {
	for b.Loop() {
		_ = NormalizeSpace("  hello \t world \n  golang  ")
	}
}

func BenchmarkIsChinaIDCard(b *testing.B) {
	for b.Loop() {
		_ = IsChinaIDCard("11010519491231002X")
	}
}

func BenchmarkIsBankCard(b *testing.B) {
	for b.Loop() {
		_ = IsBankCard("4532015112830366")
	}
}
