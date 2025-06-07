# filter

golang 的数据处理包，由 **数据输入、格式化、校验、输出** 几个部份组成。

- 采用链式写法执行每个处理函数
- 每个处理函数都可以自定义失败消息
- 可将多个处理函数封装成一个函数，便于复用
- 处理结果可以自动赋值到变量
- 支持批量处理数据

## 注意：

校验函数必须是字面语义上的条件成立，才算通过校验，例如：
- `Equals(1)`的通过条件是值等于 1
- `NotEquals(1)`的通过条件是值不等于 1

请参考单元测试代码获得使用帮助，其它问题请在 Issues 里提出。

## 安装
> github.com/dxvgef/filter/v3

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

| 分类          | 函数名                   | 描述                                      |
|-------------|-----------------------|-----------------------------------------|
| **输入函数**    |                       |                                         |
|             | `FromString`          | 输入`string`类型的数据                         |
| **格式化函数**   |                       |                                         |
|             | `ToUpper`             | 字母转为大写                                  |
|             | `ToLower`             | 字母转为小写                                  |
|             | `Trim`                | 删除左右的指定字符串                              |
|             | `TrimSpace`           | 删除左右所有的空格                               |
|             | `TrimLeft`            | 删除左边所有指定的字符串                            |
|             | `TrimRight`           | 删除右边所有指定的字符串                            |
|             | `TrimPrefix`          | 删除指定的前缀字符串                              |
|             | `TrimSuffix`          | 删除指定的后缀字符串                              |
|             | `Replace`             | 替换指定的字符串，可指定替换次数                        |
|             | `ReplaceAll`          | 替换所有指定的字符串                              |
|             | `RemoveSpace`         | 删除所有空格                                  |
|             | `Base64StdEncode`     | 使用 base64.StdEncoding 编码                |
|             | `Base64StdDecode`     | 使用 base64.StdEncoding 解码                |
|             | `Base64RawStdEncode`  | 使用 base64.RawStdEncoding 编码             |
|             | `Base64RawStdDecode`  | 使用 base64.RawStdEncoding 解码             |
|             | `Base64URLEncode`     | 使用 base64.URLEncoding 编码                |
|             | `Base64URLDecode`     | 使用 base64.URLEncoding 解码                |
|             | `Base64RawURLEncode`  | 使用 base64.RawURLEncoding 编码             |
|             | `Base64RawURLDecode`  | 使用 base64.RawURLEncoding 解码             |
|             | `HTMLEscape`          | 使用 html.EscapeString 编码                 |
|             | `HTMLUnescape`        | 使用 html.UnescapeString 解码               |
|             | `URLPathEscape`       | 使用 url.PathEscape 编码                    |
|             | `URLPathUnescape`     | 使用 url.PathUnescape 解码                  |
|             | `URLQueryEscape`      | 使用 url.QueryEscape 编码                   |
|             | `URLQueryUnescape`    | 使用 url.QueryUnescape 解码                 |
| **校验函数**    |                       |                                         |
|             | `Equals`              | 等于                                      |
|             | `NotEquals`           | 不等于                                     |
|             | `Contains`            | 包含了指定的字符串                               |
|             | `NotContains`         | 没有包含指定的字符串                              |
|             | `LengthEquals`        | 长度等于                                    |
|             | `LengthNotEquals`     | 长度不等于                                   |
|             | `LengthMin`           | 长度最小值                                   |
|             | `LengthMax`           | 长度最大值                                   |
|             | `LengthRange`         | 长度范围                                    |
|             | `UTF8LengthEquals`    | UTF8编码长度等于                              |
|             | `UTF8LengthNotEquals` | UTF8编码长度不等于                             |
|             | `UTF8LengthMain`      | UTF8编码长度最小值                             |
|             | `UTF8LengthMax`       | UTF8编码长度最大值                             |
|             | `UTF8LengthRange`     | UTF8编码长度范围                              |
|             | `Enum`                | 只能是指定列表中的值                              |
|             | `Block`               | 不能是指定列表中的值                              |
|             | `AllowedChars`        | 只能有指定列表中的字符                             |
|             | `BlockChars`          | 不能有指定列表中的字符                             |
|             | `AllowedSymbols`      | 只能有指定列表中的符号                             |
|             | `BlockSymbols`        | 不能有指定列表中的符号                             |
|             | `HasLetter`           | 包含了字母(不区分大小写)                           |
|             | `HasLower`            | 包含小写字母                                  |
|             | `HasUpper`            | 包含了大写字母                                 |
|             | `HasNumber`           | 包含了数字                                   |
|             | `HasSymbol`           | 包含了符号                                   |
|             | `HasPrefix`           | 包含了指定的前缀字符串                             |
|             | `HasSuffix`           | 包含了指定的后缀字符串                             |
|             | `IsLower`             | 是小写字母                                   |
|             | `IsUpper`             | 是大写字母                                   |
|             | `IsLetter`            | 是大小写字母                                  |
|             | `IsLowerOrNumber`     | 是小写字母或数字                                |
|             | `IsUpperOrNumber`     | 是大写字母或数字                                |
|             | `IsLetterOrNumber`    | 是大小写字母或数字                               |
|             | `IsChinese`           | 是汉字                                     |
|             | `IsMail`              | 是电邮地址                                   |
|             | `IsIPv4`              | 是IPv4地址                                 |
|             | `IsIPv6`              | 是IPv6地址                                 |
|             | `IsIP`                | 是IPv4或IPv6地址                            |
|             | `IsTCPAddr`           | 是 IP:Port 格式                            |
|             | `IsMAC`               | 是MAC地址                                  |
|             | `IsJSON`              | 是JSON格式                                 |
|             | `IsSQLObject`         | 是有效的SQL对象名                              |
|             | `IsURL`               | 是有效的URL                                 |
|             | `IsUUID`              | 是UUID格式                                 |
|             | `IsULID`              | 是ULID格式                                 |
|             | `IsChineseIDCard`     | 中国大陆地区身份证号码                             |
|             | `IsHexColor`          | 是HEX颜色值（不含#）                            |
| **自定义处理函数** |                       |                                         |
|             | `Custom`              | 自定义字符串处理函数                              |
| **类型转换函数**  |                       |                                         |
|             | `ToStringSlice`       | 转为`StringSliceType`类型 (string)，须指定分隔符   |
|             | `ToInteger`           | 转为`IntegerType`类型 (int64)               |
|             | `ToIntegerSlice`      | 转为`IntegerSliceType`类型 ([]int64)，须指定分隔符 |
|             | `ToBoolean`           | 转为`Boolean`类型 (bool)                    |
|             | `ToBooleanSlice`      | 转为`BooleanSliceType`类型 ([]bool)，须指定分隔符  |
|             | `ToFloat`             | 转为`FloatType`类型 (float64)               |
|             | `ToFloatSlice`        | 转为`FloatSliceType`类型 ([]flat64)，须提定分隔符  |
| **输出函数**    |                       |                                         |
|             | `Error`               | 错误消息，`error`类型                          |
|             | `Value`               | 处理后的值，`string`类型                        |
|             | `Result`              | `Value(), Error()`                      |
|             | `DefaultValue`        | 没有错误时返回`Value()`，否则返回指定值                |
|             | `Set`                 | 赋值到`string`类型的变量，并返回处理结果(`Error()`)     |

---

## <span id="stringslice">字符串切片处理</span>

| 分类         | 函数名                   | 描述                                    |
|--------------|-----------------------|---------------------------------------|
| **输入函数** |                       |                                       |
|              | `FromStringSlice`     | 输入`[]string`类型的数据                     |
| **格式化函数** |                       |                                       |
|              | `Trim`                | 删除每个元素值左右的指定字符串                       |
|              | `TrimSpace`           | 删除每个元素值左右所有的空格                        |
|              | `TrimLeft`            | 删除每个元素值左边所有指定的字符串                     |
|              | `TrimRight`           | 删除每个元素值右边所有指定的字符串                     |
|              | `TrimPrefix`          | 删除每个元素值指定的前缀字符串                       |
|              | `TrimSuffix`          | 删除每个元素值指定的后缀字符串                       |
|              | `DeleteEmpty`         | 删除空字符串的元素                             |
|              | `RemoveSpace`         | 删除每个元素值中的所有空格                         |
|              | `Base64StdEncode`     | 使用 base64.StdEncoding 对每个元素值编码        |
|              | `Base64StdDecode`     | 使用 base64.StdEncoding 对每个元素值解码        |
|              | `Base64RawStdEncode`  | 使用 base64.RawStdEncoding 对每个元素值编码     |
|              | `Base64RawStdDecode`  | 使用 base64.RawStdEncoding 对每个元素值解码     |
|              | `Base64URLEncode`     | 使用 base64.URLEncoding 对每个元素值编码        |
|              | `Base64URLDecode`     | 使用 base64.URLEncoding 对每个元素值解码        |
|              | `Base64RawURLEncode`  | 使用 base64.RawURLEncoding 对每个元素值编码     |
|              | `Base64RawURLDecode`  | 使用 base64.RawURLEncoding 对每个元素值解码     |
|              | `HTMLEscape`          | 使用 html.EscapeString 对每个元素值编码         |
|              | `HTMLUnescape`        | 使用 html.UnescapeString 对每个元素值解码       |
|              | `URLPathEscape`       | 使用 url.PathEscape 对每个元素值编码            |
|              | `URLPathUnescape`     | 使用 url.PathUnescape 对每个元素值解码          |
|              | `URLQueryEscape`      | 使用 url.QueryEscape 对每个元素值编码           |
|              | `URLQueryUnescape`    | 使用 url.QueryUnescape 对每个元素值解码         |
| **校验函数** |                       |                                       |
|              | `AllContains`         | 每个元素值都包含指定的子字符串                       |
|              | `AllNotContains`      | 每个元素值都不包含指定的子字符串                      |
|              | `ContainString`       | 至少有一个元素值包含指定的子字符串                     |
|              | `BlockString`         | 至少有一个元素值不包含指定的子字符串                    |
|              | `Enum`                | 每个元素都只能是指定列表中的值                       |
|              | `Block`               | 每个元素都不能是指定列表中的值                       |
|              | `AllowedChars`        | 每个元素都只能有指定列表中的字符                      |
|              | `BlockChars`          | 每个元素都不能有指定列表中的字符                      |
|              | `AllowedSymbols`      | 每个元素值都只能有指定列表中的符号                     |
|              | `BlockSymbols`        | 每个元素值是不能有指定列表中的符号                     |
|              | `HasLetter`           | 每个元素值都包含字母                            |
|              | `HasLower`            | 每个元素值都包含小写字母                          |
|              | `HasUpper`            | 每个元素值都包含大写字母                          |
|              | `HasNumber`           | 每个元素值都包含数字                            |
|              | `HasSymbol`           | 每个元素值都包含符号                            |
|              | `HasPrefix`           | 每个元素值都有指定的前缀字符串                       |
|              | `HasSuffix`           | 每个元素值都有指定的后缀字符串                       |
|              | `CountMin`            | 元素数量最小值                               |
|              | `CountMax`            | 元素数量最大值                               |
|              | `CountEquals`         | 元素数量等于                                |
|              | `CountNotEquals`      | 元素数量不等于                               |
|              | `LengthEquals`        | 每个元素值的长度都等于                           |
|              | `LengthNotEquals`     | 每个元素值的长度都不等于                          |
|              | `LengthMin`           | 每个元素值的长度最小值                           |
|              | `LengthMax`           | 每个元素值的长度最大值                           |
|              | `LengthRange`         | 每个元素值的长度范围                            |
|              | `UTF8LengthEquals`    | 每个元素值的UTF-8编码长度都等于                    |
|              | `UTF8LengthNotEquals` | 每个元素值的UTF-8编码长度都不等于                   |
|              | `UTF8LengthMin`       | 每个元素值的UTF-8编码长度最小值                    |
|              | `UTF8LengthMax`       | 每个元素值的UTF-8编码长度最大值                    |
|              | `UTF8LengthRange`     | 每个元素值的UTF-8编码长度范围                     |
|              | `IsLower`             | 每个元素值都是小写字母                           |
|              | `IsUpper`             | 每个元素值都是大写字母                           |
|              | `IsLetter`            | 每个元素值都是字母                             |
|              | `IsLowerOrNumber`     | 每个元素值都是小写字母或数字                        |
|              | `IsUpperOrNumber`     | 每个元素值都是大写字母或数字                        |
|              | `IsLetterOrNumber`    | 每个元素值都是字母或数字                          |
|              | `IsChinese`           | 每个元素值都是汉字                             |
|              | `IsMail`              | 每个元素值都是电邮地址                           |
|              | `IsIPv4`              | 每个元素值都是IPv4地址                         |
|              | `IsIPv6`              | 每个元素值都是IPv6地址                         |
|              | `IsIP`                | 每个元素值都是IPv4或IPv6地址                    |
|              | `IsTCPAddr`           | 每个元素值都是 IP:Port 格式                    |
|              | `IsMAC`               | 每个元素值都是MAC地址                          |
|              | `IsSQLObject`         | 每个元素值都是有效的SQL对象名                      |
|              | `IsUUID`              | 每个元素值都是UUID格式                         |
|              | `IsULID`              | 每个元素值都是ULID格式                         |
|              | `IsURL`               | 每个元素值都是URL格式                          |
|              | `IsChineseIDCard`     | 每个元素值都是中国大陆身份证号码                      |
| **自定义处理函数** |                       |                                       |
|              | `Custom`              | 自定义字符串处理函数，详见 `CustomStringSliceFunc` |
| **输出函数** |                       |                                       |
|             | `Error`               | 错误消息，`error`类型                        |
|             | `Value`               | 处理后的值，`[]string`类型                    |
|             | `Result`              | `Value(), Error()`                    |
|             | `DefaultValue`        | 没有错误时返回`Value()`，否则返回指定值              |
|             | `Set`                 | 赋值到`[]string`类型的变量，并返回处理结果(`Error()`) |


---

## <span id="integer">整数处理</span>

| 分类         | 函数名             | 描述                                |
|--------------|-----------------|-----------------------------------|
| **输入函数** |                 |                                   |
|              | `FromInteger`   | 输入`int64`类型的数据                    |
| **格式化函数** |                 |                                   |
|              | `Replace`       | 替换指定的值                            |
| **校验函数** |                 |                                   |
|              | `Equals`        | 等于                                |
|              | `NotEquals`     | 不等于                               |
|              | `Min`           | 最小值                               |
|              | `Max`           | 最大值                               |
|              | `Range`         | 范围                                |
|              | `Enum`          | 只能是指定列表中的值                        |
|              | `Block`           | 不能是指定列表中的值                        |
| **自定义处理函数** |                 |                                   |
|              | `Custom`        | 自定义整数值处理函数，详见 `CustomIntegerFunc` |
| **输出函数** |                 |                                   |
|             | `Error`         | 错误消息，`error`类型                    |
|             | `Value`         | 处理后的值，`int64`类型                   |
|             | `Result`        | `Value(), Error()`                |
|             | `DefaultValue`  | 没有错误时返回`Value()`，否则返回指定值          |
|             | `Set`           | 赋值到整数类型的变量，并返回处理结果(`Error()`)     |
|              | `Int`           | 返回`uint`类型的数值                     |
|              | `DefaultInt`    | 返回`int`类型的数值，失败返回指定的值             |
|              | `Int8`          | 返回`int8`类型的数值                     |
|              | `DefaultInt8`   | 返回`int8`类型的数值，失败返回指定的值            |
|              | `Int16`         | 返回`int16`类型的数值                    |
|              | `DefaultInt16`  | 返回`int16`类型的数值，失败返回指定的值           |
|              | `Int32`         | 返回`int32`类型的数值                    |
|              | `DefaultInt32`  | 返回`int32`类型的数值，失败返回指定的值           |
|              | `Int64`         | 返回`int64`类型的数值                    |
|              | `DefaultInt64`  | 返回`int64`类型的数值，失败返回指定的值           |
|              | `Uint`          | 返回`uint`类型的数值                     |
|              | `DefaultUint`   | 返回`uint`类型的数值，失败返回指定的值            |
|              | `Uint8`         | 返回`uint8`类型的数值                    |
|              | `DefaultUint8`  | 返回`uint8`类型的数值，失败返回指定的值           |
|              | `Uint16`        | 返回`uint16`类型的数值                   |
|              | `DefaultUint16` | 返回`uint16`类型的数值，失败返回指定的值          |
|              | `Uint32`        | 返回`uint32`类型的数值                   |
|              | `DefaultUint32` | 返回`uint32`类型的数值，失败返回指定的值          |
|              | `Uint64`        | 返回`uint64`类型的数值                   |
|              | `DefaultUint64` | 返回`uint64`类型的数值，失败返回指定的值          |

---

## <span id="integerslice">整数切片处理</span>

| 分类         | 函数名                | 描述                                 |
|--------------|--------------------|------------------------------------|
| **输入函数** |                    |                                    |
|              | `FromIntegerSlice` | 输入`[]int64`类型的数据                   |
| **校验函数** |                    |                                    |
|              | `CountMin`         | 元素数量最小值                            |
|              | `CountMax`         | 元素数量最大值                            |
|              | `CountEquals`      | 元素数量等于                             |
|              | `CountNotEquals`   | 元素数量不等于                            |
|              | `Contains`         | 存在指定值的元素                           |
|              | `NotContains`      | 不存在指定值的元素                          |
|              | `Min`              | 元素的最小值                             |
|              | `Max`              | 元素的最大值                             |
|              | `Range`            | 元素的值范围                             |
|              | `Enum`             | 元素只能是指定的值                        |
|              | `Block`            | 元素不能是指定的值                      |
| **自定义处理函数** |                    |                                    |
|              | `Custom`           | 自定义整数值处理函数，详见 `CustomIntegerSliceFunc` |
| **输出函数** |                    |                                    |
|             | `Error`            | 错误消息，`error`类型                     |
|             | `Value`            | 处理后的值，`[]int64`类型                  |
|             | `Result`           | `Value(), Error()`                 |
|             | `Set`              | 赋值到整数切片类型的变量，并返回处理结果(`Error()`)    |
|              | `Int`              | 返回`[]uint`类型的数值                    |
|              | `DefaultInt`       | 返回`[]int`类型的数值，失败返回指定的值            |
|              | `Int8`             | 返回`[]int8`类型的数值                    |
|              | `DefaultInt8`      | 返回`[]int8`类型的数值，失败返回指定的值           |
|              | `Int16`            | 返回`[]int16`类型的数值                   |
|              | `DefaultInt16`     | 返回`[]int16`类型的数值，失败返回指定的值          |
|              | `Int32`            | 返回`[]int32`类型的数值                   |
|              | `DefaultInt32`     | 返回`[]int32`类型的数值，失败返回指定的值          |
|              | `Int64`            | 返回`[]int64`类型的数值                   |
|              | `DefaultInt64`     | 返回`[]int64`类型的数值，失败返回指定的值          |
|              | `Uint`             | 返回`[]uint`类型的数值                    |
|              | `DefaultUint`      | 返回`[]uint`类型的数值，失败返回指定的值           |
|              | `Uint8`            | 返回`[]uint8`类型的数值                   |
|              | `DefaultUint8`     | 返回`[]uint8`类型的数值，失败返回指定的值          |
|              | `Uint16`           | 返回`[]uint16`类型的数值                  |
|              | `DefaultUint16`    | 返回`[]uint16`类型的数值，失败返回指定的值         |
|              | `Uint32`           | 返回`[]uint32`类型的数值                  |
|              | `DefaultUint32`    | 返回`[]uint32`类型的数值，失败返回指定的值         |
|              | `Uint64`           | 返回`[]uint64`类型的数值                  |
|              | `DefaultUint64`    | 返回`[]uint64`类型的数值，失败返回指定的值         |


---

## <span id="float">浮点值处理</span>

| 分类         | 函数名              | 描述                        |
|--------------|------------------|---------------------------|
| **输入函数** |                  |                           |
|              | `FromFloat`      | 输入`float64`类型的数据          |
| **校验函数** |                  |                           |
|              | `Equals`         | 等于                        |
|              | `NotEquals`      | 不等于                       |
|              | `Min`            | 最小值                       |
|              | `Max`            | 最大值                       |
|              | `Range`          | 范围                        |
|              | `Enum`           | 只能是指定的值                   |
|              | `Block`          | 不能是指定的值                   |
| **自定义处理函数** |                  |                           |
|              | `Custom`         | 自定义浮点值处理函数，详见 `CustomFloatFunc` |
| **输出函数** |                  |                           |
|             | `Error`          | 错误消息，`error`类型            |
|             | `Value`          | 处理后的值，`float64`类型         |
|             | `Result`         | `Value(), Error()`        |
|             | `Set`            | 赋值到浮点类型的变量，并返回处理结果(`Error()`) |
|              | `Float32`        | 返回`float32`类型的数值          |
|              | `DefaultFloat32` | 返回`float32`类型的数值，失败返回指定的值 |
|              | `Float64`        | 返回`float64`类型的数值          |
|              | `DefaultFloat64` | 返回`float64`类型的数值，失败返回指定的值 |

---

## <span id="floatslice">浮点切片处理</span>

| 分类         | 函数名              | 描述                                   |
|--------------|------------------|--------------------------------------|
| **输入函数** |                  |                                      |
|              | `FromFloatSlice` | 输入`[]float64`类型的数据                   |
| **校验函数** |                  |                                      |
|              | `Contains`       | 存在指定值的元素                             |
|              | `NotContains`    | 不存在指定值的元素                            |
|              | `Enum`           | 元素只能是指定的值                            |
|              | `Block`          | 元素不能是指定的值                            |
|              | `Min`            | 元素的最小值                               |
|              | `MaxMax`         | 元素的最大值                               |
|              | `Range`          | 元素的值范围                               |
|              | `CountMin`       | 元素数量最小值                              |
|              | `CountMax`       | 元素数量最大值                              |
|              | `CountEquals`    | 元素数量等于                               |
|              | `CountNotEquals` | 元素数量不等于                              |
| **自定义处理函数** |                  |                                      |
|              | `Custom`         | 自定义浮点值处理函数，详见 `CustomFloatSliceFunc` |
| **输出函数** |                  |                                      |
|             | `Error`          | 错误消息，`error`类型                       |
|             | `Value`          | 处理后的值，`[]float64`类型                  |
|             | `Result`         | `Value(), Error()`                   |
|             | `Set`            | 赋值到浮点切片类型的变量，并返回处理结果(`Error()`)      |
|              | `Float32`        | 返回`[]float32`类型的数值                   |
|              | `DefaultFloat32` | 返回`[]float32`类型的数值，失败返回指定的值          |
|              | `Float64`        | 返回`[]float64`类型的数值                   |
|              | `DefaultFloat64` | 返回`[]float64`类型的数值，失败返回指定的值          |

---

## <span id="boolean">布尔值处理</span>

| 分类         | 函数名                            | 描述                                |
|--------------|-----------------------------------|-----------------------------------|
| **输入函数** |                                   |                                   |
|              | `FromBoolean`                     | 输入`bool`类型的数据                     |
| **校验函数** |                                   |                                   |
|              | `Equals`                          | 等于指定值                             |
|              | `NotEquals`                       | 不等于指定值                            |
| **自定义处理函数** |                                   |                                   |
|              | `Custom`                          | 自定义布尔值处理函数，详见 `CustomBooleanFunc` |
| **输出函数** |                                   |                                   |
|             | `Error`                 | 错误消息，`error`类型                    |
|             | `Value`                 | 处理后的值，`bool`类型                    |
|             | `Result`                | `Value(), Error()`                |
|             | `DefaultValue`          | 没有错误时返回`Value()`，否则返回指定值          |
|             | `Set`                   | 赋值到`bool`类型的变量，并返回处理结果(`Error()`) |

---

## <span id="booleanslice">布尔切片处理</span>

| 分类         | 函数名                | 描述                                     |
|--------------|--------------------|----------------------------------------|
| **输入函数** |                    |                                        |
|              | `FromBooleanSlice` | 输入`[]bool`类型的数据                        |
| **校验函数** |                    |                                        |
|              | `Contains`         | 存在指定值的元素                               |
|              | `NotContains`      | 不存在指定值的元素                              |
|              | `CountEquals`      | 元素数量等于                                 |
|              | `CountNotEquals`   | 元素数量不等于                                |
|              | `CountMin`         | 元素数量最小值                                |
|              | `CountMax`         | 元素数量最大值                                |
| **自定义处理函数** |                    |                                        |
|              | `Custom`           | 自定义布尔值处理函数，详见 `CustomBooleanSliceFunc` |
| **输出函数** |                    |                                        |
|             | `Error`            | 错误消息，`error`类型                         |
|             | `Value`            | 处理后的值，`[]bool`类型                       |
|             | `Result`           | `Value(), Error()`                     |
|             | `DefaultValue`     | 没有错误时返回`Value()`，否则返回指定值               |
|             | `Set`              | 赋值到`[]bool`类型的变量，并返回处理结果(`Error()`)    |
