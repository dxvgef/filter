# filter

golang 处理包，由 **数据输入、格式化、校验、输出** 几个部份组成。

- 采用链式写法执行每个处理函数
- 每个处理函数都可以自定义失败消息
- 可将多个处理函数封装成一个函数，便于复用
- 处理结果可以自动赋值到变量
- 支持批量处理数据

## 注意：

校验函数必须是字面语义上的条件成立，才算通过校验，例如：

- `Is(1)`的通过条件是值等于 1
- `IsNot(1)`的通过条件是值不等于 1
- `Contains("@")`的通过条件是值包含字符 "@"
- `NotContains("@")`的通过条件是值不包含字符 "@"

请参考单元测试代码获得使用帮助，其它问题请在 Issues 里提出。

## 安装

> github.com/dxvgef/filter/v2

---

## 目录

- [字符串处理](#string)
- [字符串切片处理](#stringslice)
- [整数处理](#integer)
- [整数切片处理](#integerslice)
- [浮点处理](#float)
- [浮点切片处理](#floatslice)
- [布尔处理](#boolean)
- [布尔切片处理](#booleanslice)

## <span id="string">字符串处理</span>

| 分类         | 函数名                  | 描述                                      |
|------------|----------------------|-----------------------------------------|
| **输入函数**   | `FromString`         | 输入`string`类型                            |
| **格式化函数**  | `ToUpper`            | 同`strings.ToUpper`                      |
|            | `ToLower`            | 同`strings.ToLower`                      |
|            | `Trim`               | 同`strings.Trim`                         |
|            | `TrimSpace`          | 同`strings.TrimSpace`                    |
|            | `TrimLeft`           | 同`strings.TrimLeft`                     |
|            | `TrimRight`          | 同`strings.TrimRight`                    |
|            | `TrimPrefix`         | 同`strings.Prefix`                       |
|            | `TrimSuffix`         | 同`strings.TrimSuffix`                   |
|            | `Replace`            | 同`strings.Replace`                      |
|            | `ReplaceAll`         | `strings.ReplaceAll`                    |
|            | `RemoveSpace`        | 删除所有空格                                  |
|            | `Base64StdEncode`    | 使用 base64.StdEncoding 编码                |
|            | `Base64StdDecode`    | 使用 base64.StdEncoding 解码                |
|            | `Base64RawStdEncode` | 使用 base64.RawStdEncoding 编码             |
|            | `Base64RawStdDecode` | 使用 base64.RawStdEncoding 解码             |
|            | `Base64URLEncode`    | 使用 base64.URLEncoding 编码                |
|            | `Base64URLDecode`    | 使用 base64.URLEncoding 解码                |
|            | `Base64RawURLEncode` | 使用 base64.RawURLEncoding 编码             |
|            | `Base64RawURLDecode` | 使用 base64.RawURLEncoding 解码             |
|            | `HTMLEscape`         | 使用 html.EscapeString 编码                 |
|            | `HTMLUnescape`       | 使用 html.UnescapeString 解码               |
|            | `URLPathEscape`      | 使用 url.PathEscape 编码                    |
|            | `URLPathUnescape`    | 使用 url.PathUnescape 解码                  |
|            | `URLQueryEscape`     | 使用 url.QueryEscape 编码                   |
|            | `URLQueryUnescape`   | 使用 url.QueryUnescape 解码                 |
|            | `ToSnakeCase`        | 转为小写蛇形命名风格                              |
|            | `ToCamelCase`        | 转为小写开头的驼峰命名风格                           |
| **校验函数**   | `Is`                 | 是                                       |
|            | `IsNot`              | 不是                                      |
|            | `Contains`           | 包含                                      |
|            | `NotContains`        | 没有包含                                    |
|            | `LengthIs`           | 长度是                                     |
|            | `LengthIsNot`        | 长度不是                                    |
|            | `MinLength`          | 长度最小值                                   |
|            | `MaxLength`          | 长度最大值                                   |
|            | `LengthRange`        | 长度范围                                    |
|            | `UTF8LengthIs`       | UTF8编码的长度是                              |
|            | `UTF8LengthIsNot`    | UTF8编码的长度不是                             |
|            | `MinUTF8Length`      | UTF8编码的长度最小值                            |
|            | `MaxUTF8Length`      | UTF8编码的长度最大值                            |
|            | `UTF8LengthRange`    | UTF8编码的长度范围                             |
|            | `In`                 | 在列表中                                    |
|            | `NotIn`              | 不在列表中                                   |
|            | `AllowedChars`       | 允许的字符                                   |
|            | `BlockChars`         | 不允许的字符                                  |
|            | `AllowedSymbols`     | 允许的符号                                   |
|            | `BlockSymbols`       | 不允许的符号                                  |
|            | `HasLetter`          | 包含字母                                    |
|            | `HasLower`           | 包含小写字母                                  |
|            | `HasUpper`           | 包含大写字母                                  |
|            | `HasNumber`          | 包含数字                                    |
|            | `HasSymbol`          | 包含符号                                    |
|            | `HasPrefix`          | 包含前缀                                    |
|            | `HasSuffix`          | 包含后缀                                    |
|            | `IsLetter`           | 是字母                                     |
|            | `IsLower`            | 是小写字母                                   |
|            | `IsUpper`            | 是大写字母                                   |
|            | `IsLowerOrNumber`    | 是小写字母或数字                                |
|            | `IsUpperOrNumber`    | 是大写字母或数字                                |
|            | `IsLetterOrNumber`   | 是字母或数字                                  |
|            | `IsChinese`          | 是汉字                                     |
|            | `IsMail`             | 是电邮地址                                   |
|            | `IsIPv4`             | 是IPv4地址                                 |
|            | `IsIPv6`             | 是IPv6地址                                 |
|            | `IsIP`               | 是IPv4或IPv6地址                            |
|            | `IsTCPAddr`          | 是 IP:Port 格式                            |
|            | `IsMAC`              | 是MAC地址                                  |
|            | `IsJSON`             | 是JSON格式                                 |
|            | `IsSQLObject`        | 是有效的SQL对象名                              |
|            | `IsSQLOperator`      | 是有效的SQL运算符                              |
|            | `IsURL`              | 是有效的URL                                 |
|            | `IsUUID`             | 是UUID格式                                 |
|            | `IsULID`             | 是ULID格式                                 |
|            | `IsChineseIDCard`    | 是中国大陆地区身份证号码                            |
|            | `IsHexColor`         | 是HEX颜色值（#前缀可选）                          |
| **自定义函数**  | `Custom`             | 自定义函数，详见 `CustomStringFunc`             |
| **类型转换函数** | `ToStringSlice`      | 转为`StringSliceType`类型 (string)，须指定分隔符   |
|            | `ToInteger`          | 转为`IntegerType`类型 (int64)               |
|            | `ToIntegerSlice`     | 转为`IntegerSliceType`类型 ([]int64)，须指定分隔符 |
|            | `ToBoolean`          | 转为`Boolean`类型 (bool)                    |
|            | `ToBooleanSlice`     | 转为`BooleanSliceType`类型 ([]bool)，须指定分隔符  |
|            | `ToFloat`            | 转为`FloatType`类型 (float64)               |
|            | `ToFloatSlice`       | 转为`FloatSliceType`类型 ([]float64)，须指定分隔符 |
| **输出函数**   | `Error`              | 错误消息，`error`类型                          |
|            | `Value`              | 处理后的值，`string`类型                        |
|            | `Result`             | `Value(), Error()`                      |
|            | `DefaultValue`       | 没有错误时返回`Value()`，否则返回指定值                |
|            | `Set`                | 赋值到`string`类型的变量，并返回处理结果                |

---

## <span id="stringslice">字符串切片处理</span>

与`StringType`函数的差异：

| 分类        | 函数名               | 描述             |
|-----------|-------------------|----------------|
| **输入函数**  | `FromStringSlice` | 输入`[]string`类型 |
| **格式化函数** | `DeleteEmpty`     | 删除空字符串的元素      |
| **校验函数**  | `In`              | 在列表中           |
|           | `NotIn`           | 没在列表中          |
|           | `MinCount`        | 元素数量最小值        |
|           | `MaxCount`        | 元素数量最大值        |
|           | `CountIs`         | 元素数量等于         |
|           | `CountIsNot`      | 元素数量不等于        |
|           | `LengthIs`        | 元素长度等于         |
|           | `LengthIsNot`     | 元素长度不等于        |
|           | `MinLength`       | 元素长度最小值        |
|           | `MaxLength`       | 元素长度最大值        |
|           | `LengthRange`     | 元素长度范围         |

---

## <span id="integer">整数处理</span>

| 分类        | 函数名             | 描述                     |
|-----------|-----------------|------------------------|
| **输入函数**  | `FromInteger`   | 输入`int64`类型            |
| **格式化函数** | `Replace`       | 替换                     |
| **校验函数**  | `Is`            | 等于                     |
|           | `IsNot`         | 不等于                    |
|           | `Min`           | 最小值                    |
|           | `Max`           | 最大值                    |
|           | `Range`         | 范围                     |
|           | `In`            | 在列表中                   |
|           | `NotIn`         | 不在列表中                  |
| **自定义函数** | `Custom`        | 详见 `CustomIntegerFunc` |
| **输出函数**  | `Error`         | 错误消息，`error`类型         |
|           | `Value`         | 处理后的值，`int64`类型        |
|           | `Result`        | `Value(), Error()`     |
|           | `Set`           | 赋值到整数类型的变量，并返回处理结果     |
|           | `Int`           | 转为`int`类型              |
|           | `DefaultInt`    | 转为`int`类型，失败返回指定的值     |
|           | `Int8`          | 转为`int8`类型             |
|           | `DefaultInt8`   | 转为`int8`类型，失败返回指定的值    |
|           | `Int16`         | 转为`int16`类型            |
|           | `DefaultInt16`  | 转为`int16`类型，失败返回指定的值   |
|           | `Int32`         | 转为`int32`类型            |
|           | `DefaultInt32`  | 转为`int32`类型，失败返回指定的值   |
|           | `Int64`         | 转为`int64`类型            |
|           | `DefaultInt64`  | 转为`int64`类型，失败返回指定的值   |
|           | `Uint`          | 转为`uint`类型             |
|           | `DefaultUint`   | 转为`uint`类型，失败返回指定的值    |
|           | `Uint8`         | 转为`uint8`类型            |
|           | `DefaultUint8`  | 转为`uint8`类型，失败返回指定的值   |
|           | `Uint16`        | 转为`uint16`类型           |
|           | `DefaultUint16` | 转为`uint16`类型，失败返回指定的值  |
|           | `Uint32`        | 转为`uint32`类型           |
|           | `DefaultUint32` | 转为`uint32`类型，失败返回指定的值  |
|           | `Uint64`        | 转为`uint64`类型           |
|           | `DefaultUint64` | 转为`uint64`类型，失败返回指定的值  |

---

## <span id="integerslice">整数切片处理</span>

与`IntegerType`类型的差异

| 分类        | 函数名                  | 描述                                     |
|-----------|----------------------|----------------------------------------|
| **输入函数**  | `FromIntegerSlice`   | 输入`[]int64`类型                          |
| **校验函数**  | `MinCount`           | 元素数量最小值                                |
|           | `MaxCount`           | 元素数量最大值                                |
|           | `CountIs`            | 元素数量等于                                 |
|           | `CountIsNot`         | 元素数量不等于                                |
| **输出函数**  | `IntSlice`           | 转为`[]int`类型                            |
|           | `DefaultIntSlice`    | 转为`[]int`类型，失败返回指定的值                   |
|           | `Int8Slice`          | 转为`[]int8`类型                           |
|           | `DefaultInt8Slice`   | 转为`[]int8`类型，失败返回指定的值                  |
|           | `Int16Slice`         | 转为`[]int16`类型                          |
|           | `DefaultInt16Slice`  | 转为`[]int16`类型，失败返回指定的值                 |
|           | `Int32Slice`         | 转为`[]int32`类型                          |
|           | `DefaultInt32Slice`  | 转为`[]int32`类型，失败返回指定的值                 |
|           | `Int64Slice`         | 转为`[]int64`类型                          |
|           | `DefaultInt64Slice`  | 转为`[]int64`类型，失败返回指定的值                 |
|           | `UintSlice`          | 转为`[]uint`类型                           |
|           | `DefaultUintSlice`   | 转为`[]uint`类型，失败返回指定的值                  |
|           | `Uint8Slice`         | 转为`[]uint8`类型                          |
|           | `DefaultUint8Slice`  | 转为`[]uint8`类型，失败返回指定的值                 |
|           | `Uint16Slice`        | 转为`[]uint16`类型                         |
|           | `DefaultUint16Slice` | 转为`[]uint16`类型，失败返回指定的值                |
|           | `Uint32Slice`        | 转为`[]uint32`类型                         |
|           | `DefaultUint32Slice` | 转为`[]uint32`类型，失败返回指定的值                |
|           | `Uint64Slice`        | 转为`[]uint64`类型                         |
|           | `DefaultUint64Slice` | 转为`[]uint64`类型，失败返回指定的值                |

---

## <span id="float">浮点值处理</span>

| 分类        | 函数名              | 描述                              |
|-----------|------------------|---------------------------------|
| **输入函数**  | `FromFloat`      | 输入`float64`类型                   |
| **校验函数**  | `Is`             | 等于                              |
|           | `IsNot`          | 不等于                             |
|           | `Min`            | 最小值                             |
|           | `Max`            | 最大值                             |
|           | `Range`          | 范围                              |
|           | `In`             | 在列表中                            |
|           | `NotIn`          | 不在列表中                           |
| **自定义函数** | `Custom`         | 自定义浮点值处理函数，详见 `CustomFloatFunc` |
| **输出函数**  | `Error`          | 错误消息，`error`类型                  |
|           | `Value`          | 处理后的值，`float64`类型               |
|           | `Result`         | `Value(), Error()`              |
|           | `Set`            | 赋值到浮点类型的变量，并返回处理结果              |
|           | `Float32`        | 转为`float32`类型                   |
|           | `DefaultFloat32` | 转为`float32`类型，失败返回指定的值          |
|           | `Float64`        | 转为`float64`类型                   |
|           | `DefaultFloat64` | 转为`float64`类型，失败返回指定的值          |

---

## <span id="floatslice">浮点切片处理</span>

与`FloatType`类型的差异

| 分类       | 函数名                   | 描述                       |
|----------|-----------------------|--------------------------|
| **输入函数** | `FromFloatSlice`      | 输入`[]float64`类型          |
| **校验函数** | `MinCount`            | 元素数量最小值                  |
|          | `MaxCount`            | 元素数量最大值                  |
|          | `CountIs`             | 元素数量等于                   |
|          | `CountIsNot`          | 元素数量不等于                  |
| **输出函数** | `Float32Slice`        | 转为`[]float32`类型          |
|          | `DefaultFloat32Slice` | 转为`[]float32`类型，失败返回指定的值 |
|          | `Float64Slice`        | 转为`[]float64`类型          |
|          | `DefaultFloat64Slice` | 转为`[]float64`类型，失败返回指定的值 |

---

## <span id="boolean">布尔值处理</span>

| 分类        | 函数名            | 描述                           |
|-----------|----------------|------------------------------|
| **输入函数**  | `FromBoolean`  | 输入`bool`类型                   |
| **校验函数**  | `Is`           | 等于                           |
|           | `IsNot`        | 不等于                          |
| **自定义函数** | `Custom`       | 自定义函数，详见 `CustomBooleanFunc` |
| **输出函数**  | `Error`        | 错误消息，`error`类型               |
|           | `Value`        | 处理后的值，`bool`类型               |
|           | `Result`       | `Value(), Error()`           |
|           | `DefaultValue` | 没有错误时返回`Value()`，否则返回指定值     |
|           | `Set`          | 赋值到`bool`类型的变量，并返回处理结果       |

---

## <span id="booleanslice">布尔切片处理</span>

与`BooleanType`类型的差异

| 分类       | 函数名                | 描述           |
|----------|--------------------|--------------|
| **输入函数** | `FromBooleanSlice` | 输入`[]bool`类型 |
| **校验函数** | `CountIs`          | 元素数量等于       |
|          | `CountIsNot`       | 元素数量不等于      |
|          | `MinCount`         | 元素数量最小值      |
|          | `MaxCount`         | 元素数量最大值      |
