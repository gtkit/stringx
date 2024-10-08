# stringx
### 随机字符串
```
// RandomString 生成长度为 length 的随机字符串.
RandomString(length int) string

// SecRandomString 生成长度为 length 的安全随机字符串.
func SecRandomString(length int) (string, error)

// RandomNumber 生成长度为 length 随机数字字符串.
func RandomNumber(length int) string
```
### 三元运算符.
```
/**
 * Example:
 	func main() {
	  fmt.Println(Ter(true, 1, 2)) // 1
	  fmt.Println(Ter(false, 1, 2)) // 2
	}.
*/
func Ternary[T any](cond bool, a, b T) T 
```

### 时间字符串
```
// MillTime 获取当前时间戳(毫秒)字符串.
func MillTime() string

// UnixTime 获取当前时间戳(秒)字符串.
func UnixTime() string

```

### 字符串操作
```
// CamelToSnake 驼峰命名转成蛇形命名，如果不是驼峰命名，则转成对应小写字符串
// UserAgent → user_agent
// User → user.
func CamelToSnake(camelCase string) string

// SnakeToCamel 蛇形命名转成驼峰命名，如果不是蛇形命名，则转成对应小写字符串
// user_agent → UserAgent
// user → user.
func SnakeToCamel(snakeCase string) string

// LowerFirst 将字符串的第一个字母转换为小写
// UserAgent → userAgent.
func LowerFirst(input string) string

// UpperFirst 将首字母转换为大写
// user → User.
func UpperFirst(input string) string

// Ufirst 首字母大写.
func Ufirst(s string) string 

// Lfirst 首字母小写.
func Lfirst(s string) string 

// FirstLowerCase 获取字符串的首字母小写.
// user → u.
// User → u.
func FirstLowerCase(input string) string

// FirstUpperCase 获取字符串的首字母大写.
// user → U.
// User → U.
func FirstUpperCase(input string) string

// IsEmpty 判断字符串是否为空.
func IsEmpty(str string) bool

// IsNotEmpty 判断字符串是否不为空.
func IsNotEmpty(str string) bool

// IsBlank 判断字符串是否为空白.
func IsBlank(str string) bool

// IsNotBlank 判断字符串是否不为空白.
func IsNotBlank(str string) bool

// IsAllNotBlank 判断字符串是否全部不为空白.
func IsAllNotBlank(strs ...string) bool

// Substr 截取字符串.
func Substr(s string, start int, strlength ...int) string

// Slug 空格转换为指定分隔符.
func Slug(str, separator string) string

// Char returns a char slice
func Char(str string) []strin

// Escape 转义字符串.
func Escape(s string) string

// SubByte 截取指定长度的字节.
func SubByte(str string, length int) string

/**
 * 生成随机字符串
 * @param length 字符串长度
 * @param chartype 字符串类型，可选值："l"、"u"、"lu"、"n"、"lun"、"sc"、"all"，默认值为"lun"
 * @return 随机字符串
 * l 小写字母
 * u 大写字母
 * lu 小写字母+大写字母
 * n 数字
 * lun 小写字母+大写字母+数字
 * sc 特殊字符
 * all 所有字符
 */
Random(length int, chartype ...string) string

// SecRandom 生成长度为 length 的安全随机字符串.
func SecRandom(length int) (string, error)

// RandomN 生成长度为 length 随机数字字符串.
func RandomN(length int) string 
```

### 字符串拼接
```
// BuilderJoin joins a list of strings into a single string using a strings.Builder.
func BuilderJoin(strs []string) string

// BufferJoin joins a list of strings into a single string using a bytes.Buffer.
func BufferJoin(strs []string) string
```
