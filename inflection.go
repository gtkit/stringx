package stringx

import "github.com/gtkit/stringx/inflectx"

type Regular = inflectx.Regular
type Irregular = inflectx.Irregular
type RegularSlice = inflectx.RegularSlice
type IrregularSlice = inflectx.IrregularSlice

// AddPlural 添加一条复数变形规则。
func AddPlural(find, replace string) { inflectx.AddPlural(find, replace) }

// AddSingular 添加一条单数变形规则。
func AddSingular(find, replace string) { inflectx.AddSingular(find, replace) }

// AddIrregular 添加一条不规则变形映射。
func AddIrregular(singular, plural string) { inflectx.AddIrregular(singular, plural) }

// AddUncountable 添加不可数名词规则。
func AddUncountable(values ...string) { inflectx.AddUncountable(values...) }

// GetPlural 返回当前复数变形规则。
func GetPlural() RegularSlice { return inflectx.GetPlural() }

// GetSingular 返回当前单数变形规则。
func GetSingular() RegularSlice { return inflectx.GetSingular() }

// GetIrregular 返回当前不规则变形映射。
func GetIrregular() IrregularSlice { return inflectx.GetIrregular() }

// GetUncountable 返回当前不可数名词规则。
func GetUncountable() []string { return inflectx.GetUncountable() }

// SetPlural 设置复数变形规则。
func SetPlural(inflections RegularSlice) { inflectx.SetPlural(inflections) }

// SetSingular 设置单数变形规则。
func SetSingular(inflections RegularSlice) { inflectx.SetSingular(inflections) }

// SetIrregular 设置不规则变形映射。
func SetIrregular(inflections IrregularSlice) { inflectx.SetIrregular(inflections) }

// SetUncountable 设置不可数名词规则。
func SetUncountable(inflections []string) { inflectx.SetUncountable(inflections) }

// Plural 将单词转换为复数形式。
func Plural(str string) string { return inflectx.Plural(str) }

// Singular 将单词转换为单数形式。
func Singular(str string) string { return inflectx.Singular(str) }
