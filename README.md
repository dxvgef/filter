# filter

golang 的数据过滤包，由 **数据输入、格式化、校验、输出** 几个部份组成。

- 每个步骤都可以自定义错误消息
- 过滤结果可以自动赋值到变量
- 支持批量操作，**合并错误处理代码**
- 可将多条过滤规则封装成函数，便于复用

> github.com/dxvgef/filter/v2

请参考单元测试代码获得使用帮助，其它问题请在 Issues 里提出。

---

## 字符串过滤

### 输入函数
- `FromString()` 输入`string`类型的数据

### 格式化函数

- `ToUpper` 字母转为大写
- `ToLower` 字母转为小写
- `Trim` 删除左右的指定字符
- `TrimSpace` 删除左右的空格
- `TrimLeft` 删除左边的指定字符串
- `TrimRight` 删除右边的指定字符串
- `TrimPrefix` 删除指定的前缀字符串
- `TrimSuffix` 删除指定的后缀字符串
- `Replace` 替换指定的字符串，可指定替换次数
- `ReplaceAll` 替换指定的字符串，替换所有
- `RemoveSpace` 删除字符串中所有出现的空格
- `Base64StdEncode` Base64 std 编码
- `Base64StdDecode` Base64 std 解码
- `Base64RawStdEncode` Base64 raw std 编码
- `Base64RawStdDecode` Base64 raw std 解码
- `Base64URLEncode` Base64 URL 编码
- `Base64URLDecode` Base64 URL 解码
- `Base64RawURLEncode` Base64 raw URL 编码
- `Base64RawURLDecode` Base64 raw URL 解码
- `HTMLEscape` 编码成HTML中显示的字符
- `HTMLUnescape` HTMLEscape的解码函数
- `URLPathEscape` 编码成能作为URL路径传输的字符
- `URLPathUnescape` URLPathEscape的解码函数
- `URLQueryEscape` 编码成能作为URL查询参数传输的字符
- `URLQueryUnescape` URLQueryEscape的解码函数

### 校验函数

- `Require` 参数不能为零值
- `Length` 必须等于指定的长度
- `MinLength` 不能小于指定长度
- `UTF8MinLength` 不能小于UTF8编码的指定长度
- `MaxLength` 不能大于指定长度
- `UTF8MaxLength` 不能大于UTF8编码的指定长度
- `IsBool` 必须是布尔值
- `IsLower` 必须是小写字母
- `IsUpper` 必须是大写字母
- `IsLetter` 必须是字母
- `IsUnsigned` 必须是无符号整数类型
- `IsLowerOrNumber` 必须是小写字母或数字
- `IsUpperOrNumber` 必须是大写字母或数字
- `IsLetterOrNumber` 必须是字母或数字
- `IsChinese` 必须是汉字
- `IsMail` 必须是电邮地址
- `IsIP` 必须是IPv4/v6地址
- `IsTCPAddr` 必须是IP or Host:Port格式
- `IsMAC` 必须是MAC地址
- `IsJSON` 必须是有效的JSON格式
- `IsChinaIDNumber` 必须是中国大陆地区身份证号码
- `IsSQLObject` 必须是有效的SQL对象名(库、表、字段)
- `IsSQLObjects` 必须是有效的SQL对象名集合
- `IsUUID` 必须是UUID格式
- `IsULID` 必须是ULID格式
- `IsURL` 必须是URL格式
- `HasLetter` 必须有字母
- `HasLower` 必须有小写字母
- `HasUpper` 必须有大写字母
- `HasNumber` 必须有数字
- `HasSymbol` 必须有符号
- `HasPrefix` 必须有指定的前缀字符串
- `HasSuffix` 必须有指定的后缀字符串
- `Contains` 必须包含指定的多个字符串
- `AllowedChars` 只允许存在指定的字符
- `AllowedSymbols` 只允许存在指定的符号（只校验符号）
- `AllowedStrings` 只允许存在指定的字符串（枚举）
- `DeniedCharts` 禁止存在指定的字符
- `DeniedSymbols` 禁止存在指定的符号（只校验符号）
- `DeniedStrings` 禁止存在指定的多个字符串

### 自定义处理函数

- `Custom` 自定义字符串处理函数，详见 `CustomStringFunc`

### 类型转换函数

- `ToInteger` 转为`int64`类型的对象，可继续使用整数过滤函数

### 输出函数

- `Error()` 返回过滤中出现的错误（`error`类型）
- `Value()` 返回过滤后的值
- `DefaultValue` 如果过滤失败则返回默认值，否则返回过滤后的值
- `Result()` 返回 `Value(), Error()`
- `Set()` 通过反射将处理结果值赋值到`string`类型的变量

---

## 整数/无符号整数过滤

### 输入函数
- `FromInteger()` 输入`int64`类型的数据

### 校验函数

- `Require` 不能为零值
- `Min` 不能小于指定值
- `Max` 不能大于指定值
- `Equal` 必须等于指定值
- `NotEqual` 不能等于指定值
- `AllowedValues` 只允许是指定的多个值
- `DeniedValues` 禁止是指定的多个值

### 自定义处理函数

- `Custom` 自定义整数值处理函数，详见 `CustomIntegerFunc`

### 输出函数

- `Error()` 返回过滤中出现的错误（`error`类型）
- `Result()` 返回 `int64, error`
- `Set()` 通过反射将处理结果值赋值到所有整数和无符符号整数类型的变量
- `Int()` 返回`uint`类型的数值
- `DefaultInt()` 返回`int`类型的数值，如果失败则返回指定的默认值
- `Int8()` 返回`int8`类型的数值
- `DefaultInt8()` 返回`int8`类型的数值，如果失败则返回指定的默认值
- `Int16()` 返回`int16`类型的数值
- `DefaultInt16()` 返回`int16`类型的数值，如果失败则返回指定的默认值
- `Int32()` 返回`int32`类型的数值
- `DefaultInt32()` 返回`int32`类型的数值，如果失败则返回指定的默认值
- `Int64()` 返回`int64`类型的数值
- `DefaultInt64()` 返回`int64`类型的数值，如果失败则返回指定的默认值
- `Uint()` 返回`uint`类型的数值
- `DefaultUint()` 返回`uint`类型的数值，如果失败则返回指定的默认值
- `Uint8()` 返回`uint8`类型的数值
- `DefaultUint8()` 返回`uint8`类型的数值，如果失败则返回指定的默认值
- `Uint16()` 返回`uint16`类型的数值
- `DefaultUint16()` 返回`uint16`类型的数值，如果失败则返回指定的默认值
- `Uint32()` 返回`uint32`类型的数值
- `DefaultUint32()` 返回`uint32`类型的数值，如果失败则返回指定的默认值
- `Uint64()` 返回`uint64`类型的数值
- `DefaultUint64()` 返回`uint64`类型的数值，如果失败则返回指定的默认值
