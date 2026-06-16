# stringx

`stringx` 是一个面向 Go 项目的字符串工具集，覆盖命名风格转换、文本处理、严格校验、业务编号校验、脱敏、随机字符串、英文单复数以及零拷贝转换等常见场景。

这个仓库目前采用“根包兼容层 + 功能子包”的结构：

- `stringx`：兼容层，适合已有项目直接接入
- `caseconv`：命名风格转换
- `textx`：通用文本处理
- `validatex`：严格校验与解析
- `randomx`：随机字符串与随机数辅助
- `inflectx`：英文单复数规则管理
- `maskx`：脱敏能力
- `unsafex`：显式 `unsafe` 能力

## 安装

```bash
go get github.com/gtkit/stringx
```

## 推荐引入方式

- 新项目推荐按功能直接引入子包，依赖边界更清晰
- 旧项目可以继续使用根包 `stringx`，不需要一次性迁移
- 脱敏优先用 `maskx`
- 显式 `unsafe` 能力优先用 `unsafex`

## 快速开始

### 根包示例

```go
package main

import (
	"fmt"
	"time"

	"github.com/gtkit/stringx"
)

func main() {
	fmt.Println(stringx.ToCamel("user_profile"))             // 转为大驼峰：UserProfile
	fmt.Println(stringx.NormalizeSpace("  hello   world  ")) // 压缩空白：hello world
	fmt.Println(stringx.IsEmailStrict("user@例子.中国"))       // 严格邮箱校验：true

	info, _ := stringx.ParseChinaIDCard(
		"11010519491231002X",
		time.Date(2026, 3, 28, 0, 0, 0, 0, time.UTC),
	)
	fmt.Println(info.ProvinceCode) // 省级区划代码：11
	fmt.Println(info.Gender)       // 性别：female
	fmt.Println(info.Age)          // 年龄：76

	fmt.Println(stringx.MaskPhone("13800138000")) // 手机号脱敏：138****8000
}
```

### 子包示例

```go
package main

import (
	"fmt"

	"github.com/gtkit/stringx/caseconv"
	"github.com/gtkit/stringx/maskx"
	"github.com/gtkit/stringx/validatex"
)

func main() {
	fmt.Println(caseconv.ToKebab("HelloWorld123")) // hello-world-123
	fmt.Println(maskx.MaskEmail("alice@example.com")) // a***e@example.com

	err := validatex.ValidateChinaMobilePhone("(+86) 138-0013-8000")
	fmt.Println(err == nil) // true
}
```

## 根包方法详解

下面的说明以根包 `stringx` 为主。大多数同名子包方法语义一致，只是导入路径不同。

### 1. 命名风格转换

- `ToCamel(s string) string`：把 `snake_case`、`kebab-case`、空格分隔等格式转换为大驼峰。
- `ToLowerCamel(s string) string`：把字符串转换为小驼峰。
- `ToSnake(s string) string`：把驼峰或混合命名转换为 `snake_case`。
- `ToUpperSnake(s string) string`：把字符串转换为 `UPPER_SNAKE_CASE`。
- `ToKebab(s string) string`：把字符串转换为 `kebab-case`。

```go
fmt.Println(stringx.ToCamel("user_profile"))     // UserProfile
fmt.Println(stringx.ToLowerCamel("user_profile")) // userProfile
fmt.Println(stringx.ToSnake("HTTPServer"))        // http_server
fmt.Println(stringx.ToUpperSnake("userProfile"))  // USER_PROFILE
fmt.Println(stringx.ToKebab("Version2Build3"))    // version-2-build-3
```

### 2. 首字符处理与空白判断

- `LowerFirst(input string) string`：把首个 rune 转为小写，其余部分保持不变。
- `UpperFirst(input string) string`：把首个 rune 转为大写，其余部分保持不变。
- `FirstLowerCase(input string) string`：返回首个 rune 的小写形式。
- `FirstUpperCase(input string) string`：返回首个 rune 的大写形式。
- `IsEmpty(str string) bool`：判断是否为空字符串。
- `IsNotEmpty(str string) bool`：判断是否为非空字符串。
- `IsBlank(str string) bool`：去掉首尾空白后是否为空。
- `IsNotBlank(str string) bool`：去掉首尾空白后是否仍有内容。
- `IsAllBlank(strs ...string) bool`：多个字符串是否全部为空白。
- `IsAllNotBlank(strs ...string) bool`：多个字符串是否全部为非空白。
- `Ternary[T any](cond bool, a, b T) T`：泛型三元表达式辅助函数。

```go
fmt.Println(stringx.LowerFirst("GoLang"))   // goLang
fmt.Println(stringx.UpperFirst("golang"))   // Golang
fmt.Println(stringx.IsBlank(" \t\n "))      // true
fmt.Println(stringx.IsAllNotBlank("go", "lang")) // true
fmt.Println(stringx.Ternary(true, "A", "B"))     // A
```

### 3. 截取、切分与拼接

- `Substr(s string, start int, length ...int) string`：按 rune 下标截取字符串，支持负索引。
- `SubByte(str string, length int) string`：按字节长度安全截取 UTF-8 字符串，避免截断到非法边界。
- `Slug(str, separator string) string`：把空格替换为指定分隔符。
- `NormalizeSpace(str string) string`：去掉首尾空白，并把连续空白折叠为单个空格。
- `Char(str string) []string`：按 rune 切分为字符串切片。
- `Escape(s string) string`：得到适合放入 Go 双引号字面量中的转义文本。
- `Reverse(s string) string`：按 rune 反转字符串。
- `BuilderJoin(strs []string) string`：使用 `strings.Builder` 高效拼接。
- `BufferJoin(strs []string) string`：使用 `bytes.Buffer` 高效拼接。
- `Partition(str, sep string) (head, match, tail string)`：按第一次命中分隔。
- `LastPartition(str, sep string) (head, match, tail string)`：按最后一次命中分隔。
- `Insert(dst, src string, index int) string`：按 rune 索引插入子串。
- `WordSplit(str string) []string`：把文本拆成单词切片。

```go
fmt.Println(stringx.Substr("你好世界", 1, 2))         // 好世
fmt.Println(stringx.SubByte("你好世界", 3))            // 你
fmt.Println(stringx.NormalizeSpace(" a \t b \n c ")) // a b c
fmt.Println(stringx.Char("Go语言"))                   // [G o 语 言]
fmt.Println(stringx.Reverse("hello 世界"))            // 界世 olleh

head, match, tail := stringx.Partition("hello=world", "=")
fmt.Println(head, match, tail) // hello = world
```

### 4. 长度、宽度与排版

- `Len(str string) int`：返回 UTF-8 rune 数量，不是字节数。
- `RuneWidth(r rune) int`：返回单个 rune 的显示宽度。
- `Width(str string) int`：返回整串在等宽字体下的大致显示宽度。
- `WordCount(str string) int`：统计单词数量。
- `LeftJustify(str string, length int, pad string) string`：在右侧补齐。
- `RightJustify(str string, length int, pad string) string`：在左侧补齐。
- `Center(str string, length int, pad string) string`：在左右两侧补齐。

```go
fmt.Println(stringx.Len("你好"))                 // 2
fmt.Println(stringx.Width("A你好"))             // 5
fmt.Println(stringx.WordCount("hello, world"))  // 2
fmt.Println(stringx.LeftJustify("go", 5, "."))  // go...
fmt.Println(stringx.RightJustify("go", 5, ".")) // ...go
fmt.Println(stringx.Center("go", 6, "."))       // ..go..
```

### 5. 字符映射与压缩

- `type Translator`：可复用的字符映射器。
- `NewTranslator(from, to string) *Translator`：根据模式创建映射器。
- `(*Translator).Translate(str string) string`：使用预编译规则转换字符串。
- `(*Translator).TranslateRune(r rune) (result rune, translated bool)`：转换单个 rune。
- `(*Translator).HasPattern() bool`：判断是否存在有效模式。
- `Translate(str, from, to string) string`：一次性转换。
- `Delete(str, pattern string) string`：删除命中的字符。
- `Count(str, pattern string) int`：统计命中数量。
- `Squeeze(str, pattern string) string`：压缩连续重复字符。

```go
tr := stringx.NewTranslator("aeiou", "12345")
fmt.Println(tr.Translate("hello"))          // h2ll4
fmt.Println(stringx.Delete("hello", "aeiou")) // hll
fmt.Println(stringx.Count("hello", "aeiou"))  // 2
fmt.Println(stringx.Squeeze("hello   world", " ")) // hello world
```

### 6. 轻量业务校验

- `IsEmail(email string) bool`：适合表单场景的轻量邮箱正则校验。
- `IsPhone(phone string) bool`：适合中国大陆手机号表单场景的轻量校验，只接受 `1[3-9]` 开头的 11 位数字。

```go
fmt.Println(stringx.IsEmail("user@example.com")) // true
fmt.Println(stringx.IsPhone("13800138000"))      // true
fmt.Println(stringx.IsPhone("+86 13800138000"))  // false
```

### 7. 严格协议校验

- `ParseEmailAddress(address string) (*mail.Address, error)`：使用 `net/mail` 解析邮箱地址。
- `IsEmailStrict(address string) bool`：严格邮箱校验。
- `IsURL(raw string) bool`：严格绝对 URL 校验。
- `IsHTTPURL(raw string) bool`：仅允许 `http` / `https` 的绝对 URL。
- `IsHostname(host string) bool`：主机名校验，内置 IDNA 支持。
- `IsIP(text string) bool`：IP 校验。
- `IsIPv4(text string) bool`：IPv4 校验。
- `IsIPv6(text string) bool`：IPv6 校验。
- `IsCIDR(text string) bool`：CIDR 校验。
- `IsMAC(text string) bool`：MAC 地址校验。
- `IsUUID(text string) bool`：UUID 文本格式校验。
- `IsUUIDv4(text string) bool`：UUID v4 校验。
- `IsE164(text string) bool`：E.164 电话号码语法校验。

```go
fmt.Println(stringx.IsEmailStrict("user@例子.中国"))     // true
fmt.Println(stringx.IsHTTPURL("https://例子.中国"))      // true
fmt.Println(stringx.IsIPv4("127.0.0.1"))               // true
fmt.Println(stringx.IsUUIDv4("550e8400-e29b-41d4-a716-446655440000")) // true
fmt.Println(stringx.IsE164("+8613800138000"))          // true
```

### 8. 错误型严格验证 API

这些方法和 `Is*` 系列语义一致，但会返回详细错误原因，适合接口层、表单层和日志层使用。

- `ValidateEmailStrict(address string) error`
- `ValidateURL(raw string) error`
- `ValidateHTTPURL(raw string) error`
- `ValidateHostname(host string) error`
- `ValidateIP(text string) error`
- `ValidateIPv4(text string) error`
- `ValidateIPv6(text string) error`
- `ValidateCIDR(text string) error`
- `ValidateMAC(text string) error`
- `ValidateUUID(text string) error`
- `ValidateUUIDv4(text string) error`
- `ValidateE164(text string) error`

```go
if err := stringx.ValidateHTTPURL("ftp://example.com"); err != nil {
	fmt.Println(err) // 输出明确错误原因
}

if err := stringx.ValidateIPv4("2001:db8::1"); err != nil {
	fmt.Println(err) // address is not IPv4
}
```

### 9. 中国业务编号与解析

- `ValidateChinaIDCard(text string) error`：校验 18 位二代居民身份证。
- `IsChinaIDCard(text string) bool`：布尔版身份证校验。
- `ParseChinaIDCard(text string, now time.Time) (*ChinaIDCardInfo, error)`：解析身份证中的出生日期、年龄、性别、省级区划代码。
- `ValidateChinaUSCC(text string) error`：校验统一社会信用代码。
- `IsChinaUSCC(text string) bool`：布尔版统一社会信用代码校验。
- `ValidateBankCard(text string) error`：使用 Luhn 算法校验银行卡号。
- `IsBankCard(text string) bool`：布尔版银行卡校验。
- `type Gender`：身份证解析中的性别类型，包含 `GenderUnknown`、`GenderMale`、`GenderFemale`。
- `type ChinaIDCardInfo`：身份证解析结果结构体，包含 `Number`、`ProvinceCode`、`BirthDate`、`Age`、`Gender`。

```go
info, err := stringx.ParseChinaIDCard(
	"11010519491231002X",
	time.Date(2026, 3, 28, 0, 0, 0, 0, time.UTC),
)
if err != nil {
	panic(err)
}

fmt.Println(info.BirthDate.Format("2006-01-02")) // 1949-12-31
fmt.Println(info.Gender)                         // female
fmt.Println(info.Age)                            // 76

fmt.Println(stringx.IsChinaUSCC("91350211M000100Y46")) // true
fmt.Println(stringx.IsBankCard("4532015112830366"))    // true
```

### 10. 脱敏

- `MaskPhone(text string) string`：手机号脱敏，默认保留前三后四。
- `MaskEmail(text string) string`：邮箱脱敏，保留域名。
- `MaskName(text string) string`：姓名脱敏。
- `MaskChinaIDCard(text string) string`：身份证脱敏，保留前 6 位和后 4 位。
- `MaskBankCard(text string) string`：银行卡脱敏，保留前 6 位和后 4 位。

```go
fmt.Println(stringx.MaskPhone("13800138000"))              // 138****8000
fmt.Println(stringx.MaskEmail("alice@example.com"))        // a***e@example.com
fmt.Println(stringx.MaskName("张三"))                       // 张*
fmt.Println(stringx.MaskChinaIDCard("11010519491231002X")) // 110105********002X
fmt.Println(stringx.MaskBankCard("4532015112830366"))      // 453201******0366
```

### 11. 英文单复数与规则管理

- `type Regular`：正则替换规则，字段为 `Find`、`Replace`。
- `type Irregular`：不规则映射，字段为 `Singular`、`Plural`。
- `type RegularSlice` / `IrregularSlice`：对应切片类型。
- `AddPlural(find, replace string)`：添加复数规则。
- `AddSingular(find, replace string)`：添加单数规则。
- `AddIrregular(singular, plural string)`：添加不规则映射。
- `AddUncountable(values ...string)`：添加不可数名词。
- `GetPlural() RegularSlice` / `GetSingular() RegularSlice` / `GetIrregular() IrregularSlice` / `GetUncountable() []string`：读取当前规则副本。
- `SetPlural(...)` / `SetSingular(...)` / `SetIrregular(...)` / `SetUncountable(...)`：整体替换规则。
- `Plural(str string) string`：转为复数。
- `Singular(str string) string`：转为单数。

```go
fmt.Println(stringx.Plural("person"))   // people
fmt.Println(stringx.Singular("people")) // person

stringx.AddIrregular("analysis", "analyses")
fmt.Println(stringx.Plural("analysis")) // analyses
```

### 12. 随机字符串、随机数与时间

- `Random(length int, chartype ...string) string`：按类型生成随机字符串。
  - `l`：小写字母
  - `u`：大写字母
  - `lu`：大小写字母
  - `n`：数字
  - `lun`：字母加数字
  - `sc`：特殊字符
  - `all`：全部字符
- `SecRandom(length int) (string, error)`：使用加密安全随机源生成随机字符串。
- `RandomFromCharset(length int, charset []byte) string`：使用显式字符集生成随机字符串。
- `RandomN(length int) string`：随机数字字符串。
- `RandStr(length int) string`：随机小写字母字符串。
- `RandStrUpper(length int) string`：随机大写字母字符串。
- `RandId() string`：固定 16 位十六进制随机 ID。
- `Randn(length int) string`：已废弃，等同于 `RandomN`。
- `RandomEle[T any](slice []T) T`：从切片中随机取一个元素，空切片返回零值。
- `RandInt(min, max int) int`：返回 `[min, max)` 范围内的随机整数。
- `RUint(n int) int`：返回 `[0, n)` 范围内的随机整数。
- `RandIntHandler(maxN, count int, handler func(num, i int))`：连续生成随机数并回调，索引 `i` 逆序传入。
- `MillTime() string`：当前 Unix 毫秒时间戳字符串。
- `UnixTime() string`：当前 Unix 秒时间戳字符串。
- `type RNG`：轻量随机数命名空间类型。
- `func (RNG) Uint32() uint32`
- `func (RNG) Uint32n(maxN uint32) uint32`
- `func (RNG) Uint64() uint64`
- `func (RNG) Uint64n(maxN uint64) uint64`：返回 `[0, maxN]` 范围内的随机值。

```go
fmt.Println(stringx.Random(8))         // 随机字母数字
fmt.Println(stringx.Random(6, "n"))    // 随机数字
fmt.Println(stringx.RandStr(6))        // 随机小写字母
fmt.Println(stringx.RandId())          // 16 位十六进制 ID
fmt.Println(stringx.RandomFromCharset(6, []byte("ABC123"))) // 自定义字符集
fmt.Println(stringx.RandomEle([]string{"a", "b", "c"})) // 随机元素

fmt.Println(stringx.RandInt(10, 20)) // [10,20) 之间的整数
fmt.Println(stringx.UnixTime())      // 当前秒级时间戳

var r stringx.RNG
fmt.Println(r.Uint32n(100)) // [0,100) 之间的 uint32
```

### 13. 零拷贝转换

- `String2Bytes(s string) []byte`：零拷贝把字符串转为字节切片。
- `Bytes2String(b []byte) string`：零拷贝把字节切片转为字符串。
- `CloneStringBytes(s string) []byte`：复制字符串，返回可安全修改的字节切片。
- `CloneBytesString(b []byte) string`：复制字节切片，返回不受后续修改影响的字符串。

这两个方法基于 `unsafe`，返回值和原对象共享底层内存：

- `String2Bytes` 的结果只能只读，不能修改
- `Bytes2String` 得到的字符串会受原切片后续修改影响

```go
bs := stringx.String2Bytes("hello")
str := stringx.Bytes2String([]byte("world"))
safeBytes := stringx.CloneStringBytes("mutable")
safeString := stringx.CloneBytesString([]byte("snapshot"))

fmt.Println(string(bs)) // hello
fmt.Println(str)        // world
fmt.Println(safeBytes)  // [109 117 116 97 98 108 101]
fmt.Println(safeString) // snapshot
```

## 子包补充说明

### `caseconv`

`caseconv` 直接暴露命名转换能力：

- `ToCamel`
- `ToLowerCamel`
- `ToSnake`
- `ToUpperSnake`
- `ToKebab`

适合只想依赖命名风格转换的场景。

### `textx`

`textx` 直接暴露通用文本处理能力，语义与根包同名方法一致，包括：

- 首字符处理：`LowerFirst`、`UpperFirst`、`FirstLowerCase`、`FirstUpperCase`
- 空白判断：`IsEmpty`、`IsNotEmpty`、`IsBlank`、`IsNotBlank`、`IsAllBlank`、`IsAllNotBlank`
- 截取与格式化：`NormalizeSpace`、`Substr`、`SubByte`、`Slug`、`Char`、`Escape`、`Reverse`
- 拼接与切分：`BuilderJoin`、`BufferJoin`、`Partition`、`LastPartition`、`Insert`
- 统计与排版：`Len`、`RuneWidth`、`Width`、`WordCount`、`WordSplit`、`LeftJustify`、`RightJustify`、`Center`
- 转换器：`Translator`、`NewTranslator`、`Translate`、`Delete`、`Count`、`Squeeze`

### `validatex`

`validatex` 包含根包同名的严格校验与解析能力，并额外提供一个仅在子包暴露的方法：

- `ValidateChinaMobilePhone(text string) error`

这个方法使用本地规则严格校验中国手机号：

- 允许 `+86` / `86` 前缀
- 允许空格、连字符、圆括号作为格式化字符
- 标准化后必须是 11 位手机号码
- 会进一步校验常见手机号段前缀

```go
err := validatex.ValidateChinaMobilePhone("(+86) 138-0013-8000")
fmt.Println(err == nil) // true
```

### `randomx`

`randomx` 适合只引入随机相关能力，包含：

- `Random`
- `SecRandom`
- `RandomFromCharset`
- `RandomN`
- `RandStr`
- `RandStrUpper`
- `RandId`
- `Randn`
- `RandomEle`

### `inflectx`

`inflectx` 适合只引入英文单复数规则管理，包含：

- 类型：`Regular`、`Irregular`、`RegularSlice`、`IrregularSlice`
- 规则管理：`AddPlural`、`AddSingular`、`AddIrregular`、`AddUncountable`
- 规则读取：`GetPlural`、`GetSingular`、`GetIrregular`、`GetUncountable`
- 规则替换：`SetPlural`、`SetSingular`、`SetIrregular`、`SetUncountable`
- 转换：`Plural`、`Singular`

### `maskx`

`maskx` 是推荐的脱敏入口，包含：

- `MaskPhone`
- `MaskEmail`
- `MaskName`
- `MaskChinaIDCard`
- `MaskBankCard`

### `unsafex`

`unsafex` 是推荐的显式 `unsafe` 入口，包含：

- `String2Bytes`
- `Bytes2String`

## 常见使用建议

- 只需要轻量格式判断时，用 `IsEmail`、`IsPhone`
- 需要协议级校验和错误信息时，用 `Validate*` 或 `IsEmailStrict`、`IsHTTPURL` 等严格方法
- 需要结构化业务解析时，用 `ParseChinaIDCard`
- 需要脱敏时，优先直接引入 `maskx`
- 需要复用字符映射规则时，优先用 `NewTranslator`
- 需要零拷贝能力时，优先显式引入 `unsafex`

## 测试与验证

建议持续执行：

```bash
go test ./...
go test -cover ./...
go test -bench . ./...
go vet ./...
```

如果本地工具链完整，也建议执行：

```bash
go test -race ./...
```
