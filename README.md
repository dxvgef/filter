# filter

golang的数据过滤包，由 **数据输入->数据清理->数据校验->类型转换->结果输出** 几个部份组成。

- 每个步骤都可以自定义错误消息
- 过滤结果可以自动赋值到变量
- 支持批量操作，无需逐一检查错误

## 数据输入

**`InputString(str, name)`**
要过滤的数据来源，目前仅支持字符串
第一个参数str为要过滤的字符串<br>
第二个参数为数据的名称，用于拼接 `{名称}: {错误消息}` 这样的错误消息

## 数据清洗

- `Trim()` 去除前后空格
- `RemoveSpace` 去除所有空格
- `ReplaceAll` 替换所有
- `ToUpper` 字母转为大写
- `ToLower` 字母转为小写
- `SnakeCaseToCamelCase` 蛇形转驼峰: hello_world => helloWorld
- `SnakeCaseToPascalCase` 蛇形转帕斯卡: hello_world => HelloWorld
- `CamelCaseToSnakeCase` 驼峰(含帕斯卡)转蛇形 helloWorld/HelloWorld => hello_world
- `HTMLEscape` 编码成HTML中显示的字符
- `HTMLUnescape` HTMLEscape的解码函数
- `URLPathEscape` 编码成能作为URL路径传输的字符
- `URLPathUnescape` URLPathEscape的解码函数
- `URLQueryEscape` 编码成能作为URL查询参数传输的字符
- `URLQueryUnescape` URLQueryEscape的解码函数
- `Base64StdEncode` Base64 std 编码
- `Base64StdDecode` Base64 std 解码
- `Base64RawStdEncode` Base64 raw std 编码
- `Base64RawStdDecode` Base64 raw std 解码
- `Base64URLEncode` Base64 URL 编码
- `Base64URLDecode` Base64 URL 解码
- `Base64RawURLEncode` Base64 raw URL 编码
- `Base64RawURLDecode` Base64 raw URL 解码
- `SetSeparator` 指定Slice类型的分隔符，配合`Strings`类型转换方法使用

## 数据验证

所有数据验证函数，都可以传入自定义错误消息，例如MinLength(""自定义错误消息")

- `Required` 必须有值（允许""之外的零值）。如果不使用此规则，当参数值为""时，数据验证默认不生效

- `MinLength` 最小长度
- `MinUTF8Length` UTF8编码最小长度
- `MaxLength` 最大长度
- `MaxUTF8Length` UTF8编码最大长度
- `MinInteger` 最小整数值
- `MaxInteger` 最大整数值
- `MinFloat` 最小浮点值
- `MaxFloat` 最大浮点值

- `IsBool` 是布尔值
- `IsLower` 是小写字母
- `IsUpper` 是大写字母
- `IsLetter` 是字母
- `IsDigit` 是数字，不能包含有任何符号或其它字符
- `IsLowerOrDigit` 是小写字母或数字
- `IsUpperOrDigit` 是大写字母或数字
- `IsLetterOrDigit` 是字母或数字
- `IsChinese` 是汉字
- `IsMail` 是电邮地址
- `IsIP` 蝇IPv4/v6地址
- `IsJSON` 是有效的JSON格式
- `IsChinaTel` 是中国大陆地区固定电话号码
- `IsChinaMobile` 是中国大陆地区手机号码
- `IsChinaIDNumber` 是中国大陆地区身份证号码
- `IsSQLObject` 是SQL对象名(库、表、字段)
- `IsSQLObjects` 是SQL对象名集合
- `IsUUID` 是UUID格式
- `IsURL` 是URL格式

- `HasLetter` 必须包含字母
- `HasLower` 必须包含小写字母
- `HasUpper` 必须包含大写字母
- `HasDigit` 必须包含数字
- `HasSymbol` 必须包含符号
- `HasPrefix` 必须包含指定的前缀字符串
- `HasSuffix` 必须包含指定的后缀字符串
- `HasString` 必须包含指定的字符串
- `Contains` 与 HasString 相同

- `EnumString` 仅允许[]string中的值
- `EnumInt` 仅允许[]int中的值
- `EnumInt32` 仅允许[]int32中的值
- `EnumInt64` 仅允许[]int64中的值
- `EnumFloat32` 仅允许[]float32中的值
- `EnumFloat64` 仅允许[]float64中的值

- `DenyString` 阻止[]string中的值
- `DenyInt` 阻止[]int中的值
- `DenyInt32` 阻止[]int32中的值
- `DenyInt64` 阻止[]int64中的值
- `DenyFloat32` 阻止[]float32中的值
- `DenyFloat64` 阻止[]float64中的值

## 类型转换

- `String` 转为string类型，并返回error
- `DefaultString` 转为string类型，如果出现错误则只返回默认值
- `Strings` 转为[]string类型，使用`Separator`方法指定元素分隔符
- `DefaultStrings` 转为[]string类型，如果出现错误则只返回默认值，使用`Separator`方法指定元素分隔符
- `Int` 转为int类型，并返回error
- `DefaultInt` 转为int类型，如果出现错误则只返回默认值
- `Int8` 转为int8类型，并返回error
- `DefaultInt8` 转为int8类型，如果出现错误则只返回默认值
- `Int16` 转为int16类型，并返回error
- `DefaultInt16` 转为int16类型，如果出现错误则只返回默认值
- `Int32` 转为int32类型，并返回error
- `DefaultInt32` 转为int32类型，如果出现错误则只返回默认值
- `Int64` 转为int64类型，并返回error
- `DefaultInt64` 转为int64类型，如果出现错误则只返回默认值
- `Float32` 转为float32类型，并返回errBor
- `DefaultFloat32` 转为float8类型，如果出现错误则只返回默认值
- `Float64` 转为float64类型，并返回error
- `DefaultFloat64` 转为float64类型，如果出现错误则只返回默认值
- `Bool` 转为bool类型，并返回error
- `DefaultBool` 转为bool类型，如果出现错误则只返回默认值

## 结果输出

除了使用类型转换函数得到过滤后的数据，还可以使用以下函数将过滤结果赋值到指定变量

- `Set` 将单个过滤结果赋值到变量
- `MSet` 将多个过滤结果赋值到对应的变量
  - `El` 用于创建`MSet`函数的`Element`入参类型
- `Silent` 静默模式。如果过滤过程中发生错误，不会返回任何错误，只适用于`El`和`Set`方法
- `Default` 默认值。如果过滤过程中发生错误，用默认值进行赋值，只适用于`El`和`Set`方法
- `Error`` 过滤结果，返回`error`类型
