package inflectx

import (
	"regexp"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

type inflection struct {
	regexp  *regexp.Regexp
	replace string
	literal bool
}

// Regular 表示一条基于正则表达式的单复数替换规则。
type Regular struct {
	Find    string
	Replace string
}

// Irregular 表示一组不规则的单复数映射关系。
type Irregular struct {
	Singular string
	Plural   string
}

// RegularSlice 是 Regular 的切片类型。
type RegularSlice []Regular

// IrregularSlice 是 Irregular 的切片类型。
type IrregularSlice []Irregular

var pluralInflections = RegularSlice{
	{"([a-z])$", "${1}s"},
	{"s$", "s"},
	{"^(ax|test)is$", "${1}es"},
	{"(octop|vir)us$", "${1}i"},
	{"(octop|vir)i$", "${1}i"},
	{"(alias|status|campus)$", "${1}es"},
	{"(bu)s$", "${1}ses"},
	{"(buffal|tomat)o$", "${1}oes"},
	{"([ti])um$", "${1}a"},
	{"([ti])a$", "${1}a"},
	{"sis$", "ses"},
	{"(?:([^f])fe|([lr])f)$", "${1}${2}ves"},
	{"(hive)$", "${1}s"},
	{"([^aeiouy]|qu)y$", "${1}ies"},
	{"(x|ch|ss|sh)$", "${1}es"},
	{"(matr|vert|ind)(?:ix|ex)$", "${1}ices"},
	{"^(m|l)ouse$", "${1}ice"},
	{"^(m|l)ice$", "${1}ice"},
	{"^(ox)$", "${1}en"},
	{"^(oxen)$", "${1}"},
	{"(quiz)$", "${1}zes"},
	{"(drive)$", "${1}s"},
}

var singularInflections = RegularSlice{
	{"s$", ""},
	{"(ss)$", "${1}"},
	{"(n)ews$", "${1}ews"},
	{"([ti])a$", "${1}um"},
	{"((a)naly|(b)a|(d)iagno|(p)arenthe|(p)rogno|(s)ynop|(t)he)(sis|ses)$", "${1}sis"},
	{"(^analy)(sis|ses)$", "${1}sis"},
	{"([^f])ves$", "${1}fe"},
	{"(hive)s$", "${1}"},
	{"(tive)s$", "${1}"},
	{"([lr])ves$", "${1}f"},
	{"([^aeiouy]|qu)ies$", "${1}y"},
	{"(s)eries$", "${1}eries"},
	{"(m)ovies$", "${1}ovie"},
	{"(c)ookies$", "${1}ookie"},
	{"(x|ch|ss|sh)es$", "${1}"},
	{"^(m|l)ice$", "${1}ouse"},
	{"(bus|campus)(es)?$", "${1}"},
	{"(o)es$", "${1}"},
	{"(shoe)s$", "${1}"},
	{"(cris|test)(is|es)$", "${1}is"},
	{"^(a)x[ie]s$", "${1}xis"},
	{"(octop|vir)(us|i)$", "${1}us"},
	{"(alias|status)(es)?$", "${1}"},
	{"^(ox)en", "${1}"},
	{"(vert|ind)ices$", "${1}ex"},
	{"(matr)ices$", "${1}ix"},
	{"(quiz)zes$", "${1}"},
	{"(database)s$", "${1}"},
	{"(drive)s$", "${1}"},
}

var irregularInflections = IrregularSlice{
	{"person", "people"},
	{"man", "men"},
	{"child", "children"},
	{"sex", "sexes"},
	{"move", "moves"},
	{"ombie", "ombies"},
	{"goose", "geese"},
	{"foot", "feet"},
	{"moose", "moose"},
	{"tooth", "teeth"},
	{"leaf", "leaves"},
	{"chassis", "chassis"},
}

var uncountableInflections = []string{
	"equipment", "information", "rice", "money", "species", "series", "fish", "sheep",
	"jeans", "police", "milk", "salt", "time", "water", "paper", "food", "art", "cash",
	"music", "help", "luck", "oil", "progress", "rain", "research", "shopping", "software", "traffic",
}

var (
	compiledPluralMaps   []inflection
	compiledSingularMaps []inflection
	inflectionMu         sync.RWMutex
)

func init() {
	compile()
}

// AddPlural 添加一条复数规则。
func AddPlural(find, replace string) {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	pluralInflections = append(pluralInflections, Regular{Find: find, Replace: replace})
	compileLocked()
}

// AddSingular 添加一条单数规则。
func AddSingular(find, replace string) {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	singularInflections = append(singularInflections, Regular{Find: find, Replace: replace})
	compileLocked()
}

// AddIrregular 添加一条不规则映射。
func AddIrregular(singular, plural string) {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	irregularInflections = append(irregularInflections, Irregular{Singular: singular, Plural: plural})
	compileLocked()
}

// AddUncountable 添加不可数名词。
func AddUncountable(values ...string) {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	uncountableInflections = append(uncountableInflections, values...)
	compileLocked()
}

// GetPlural 返回当前复数规则副本。
func GetPlural() RegularSlice {
	inflectionMu.RLock()
	defer inflectionMu.RUnlock()
	return cloneRegularSlice(pluralInflections)
}

// GetSingular 返回当前单数规则副本。
func GetSingular() RegularSlice {
	inflectionMu.RLock()
	defer inflectionMu.RUnlock()
	return cloneRegularSlice(singularInflections)
}

// GetIrregular 返回当前不规则映射副本。
func GetIrregular() IrregularSlice {
	inflectionMu.RLock()
	defer inflectionMu.RUnlock()
	return cloneIrregularSlice(irregularInflections)
}

// GetUncountable 返回当前不可数名词集合副本。
func GetUncountable() []string {
	inflectionMu.RLock()
	defer inflectionMu.RUnlock()
	return cloneStringSlice(uncountableInflections)
}

// SetPlural 替换复数规则。
func SetPlural(inflections RegularSlice) {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	pluralInflections = cloneRegularSlice(inflections)
	compileLocked()
}

// SetSingular 替换单数规则。
func SetSingular(inflections RegularSlice) {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	singularInflections = cloneRegularSlice(inflections)
	compileLocked()
}

// SetIrregular 替换不规则规则。
func SetIrregular(inflections IrregularSlice) {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	irregularInflections = cloneIrregularSlice(inflections)
	compileLocked()
}

// SetUncountable 替换不可数名词集合。
func SetUncountable(inflections []string) {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	uncountableInflections = cloneStringSlice(inflections)
	compileLocked()
}

// Plural 将英文单词转为复数。
func Plural(str string) string {
	inflectionMu.RLock()
	defer inflectionMu.RUnlock()
	for _, inf := range compiledPluralMaps {
		if inf.regexp.MatchString(str) {
			return inf.replaceAll(str)
		}
	}
	return str
}

// Singular 将英文单词转为单数。
func Singular(str string) string {
	inflectionMu.RLock()
	defer inflectionMu.RUnlock()
	for _, inf := range compiledSingularMaps {
		if inf.regexp.MatchString(str) {
			return inf.replaceAll(str)
		}
	}
	return str
}

func compile() {
	inflectionMu.Lock()
	defer inflectionMu.Unlock()
	compileLocked()
}

func compileLocked() {
	compiledPluralMaps = resetInflections(
		compiledPluralMaps,
		len(uncountableInflections)+len(irregularInflections)*3+len(pluralInflections)*3,
	)
	compiledSingularMaps = resetInflections(
		compiledSingularMaps,
		len(uncountableInflections)+len(irregularInflections)*3+len(singularInflections)*3,
	)

	for _, uncountable := range uncountableInflections {
		inf := inflection{regexp: regexp.MustCompile("(?i)^(" + regexp.QuoteMeta(uncountable) + ")$"), replace: "${1}"}
		compiledPluralMaps = append(compiledPluralMaps, inf)
		compiledSingularMaps = append(compiledSingularMaps, inf)
	}

	for _, value := range irregularInflections {
		compiledPluralMaps = append(compiledPluralMaps,
			inflection{regexp: regexp.MustCompile(regexp.QuoteMeta(strings.ToUpper(value.Singular)) + "$"), replace: strings.ToUpper(value.Plural), literal: true},
			inflection{regexp: regexp.MustCompile(regexp.QuoteMeta(titleFirst(value.Singular)) + "$"), replace: titleFirst(value.Plural), literal: true},
			inflection{regexp: regexp.MustCompile(regexp.QuoteMeta(value.Singular) + "$"), replace: value.Plural, literal: true},
		)
	}

	for _, value := range irregularInflections {
		compiledSingularMaps = append(compiledSingularMaps,
			inflection{regexp: regexp.MustCompile(regexp.QuoteMeta(strings.ToUpper(value.Plural)) + "$"), replace: strings.ToUpper(value.Singular), literal: true},
			inflection{regexp: regexp.MustCompile(regexp.QuoteMeta(titleFirst(value.Plural)) + "$"), replace: titleFirst(value.Singular), literal: true},
			inflection{regexp: regexp.MustCompile(regexp.QuoteMeta(value.Plural) + "$"), replace: value.Singular, literal: true},
		)
	}

	for i := len(pluralInflections) - 1; i >= 0; i-- {
		value := pluralInflections[i]
		compiledPluralMaps = append(compiledPluralMaps,
			inflection{regexp: regexp.MustCompile(strings.ToUpper(value.Find)), replace: strings.ToUpper(value.Replace)},
			inflection{regexp: regexp.MustCompile(value.Find), replace: value.Replace},
			inflection{regexp: regexp.MustCompile("(?i)" + value.Find), replace: value.Replace},
		)
	}

	for i := len(singularInflections) - 1; i >= 0; i-- {
		value := singularInflections[i]
		compiledSingularMaps = append(compiledSingularMaps,
			inflection{regexp: regexp.MustCompile(strings.ToUpper(value.Find)), replace: strings.ToUpper(value.Replace)},
			inflection{regexp: regexp.MustCompile(value.Find), replace: value.Replace},
			inflection{regexp: regexp.MustCompile("(?i)" + value.Find), replace: value.Replace},
		)
	}
}

func (inf inflection) replaceAll(str string) string {
	if inf.literal {
		return inf.regexp.ReplaceAllStringFunc(str, func(string) string {
			return inf.replace
		})
	}
	return inf.regexp.ReplaceAllString(str, inf.replace)
}

func resetInflections(dst []inflection, size int) []inflection {
	if cap(dst) < size {
		return make([]inflection, 0, size)
	}
	return dst[:0]
}

func cloneRegularSlice(src RegularSlice) RegularSlice {
	dst := make(RegularSlice, len(src))
	copy(dst, src)
	return dst
}

func cloneIrregularSlice(src IrregularSlice) IrregularSlice {
	dst := make(IrregularSlice, len(src))
	copy(dst, src)
	return dst
}

func cloneStringSlice(src []string) []string {
	dst := make([]string, len(src))
	copy(dst, src)
	return dst
}

func titleFirst(s string) string {
	if s == "" {
		return s
	}
	r, size := utf8.DecodeRuneInString(s)
	return string(unicode.ToUpper(r)) + s[size:]
}
