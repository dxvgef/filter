# filter

golang的数据过滤包，由 **数据输入、格式化、校验、输出** 几个部份组成。

- 每个步骤都可以自定义错误消息
- 过滤结果可以自动赋值到变量
- 支持批量操作，**合并错误处理代码**
- 可将过滤规则封装成函数，便于复用

> github.com/dxvgef/filter/v2

## 示例

请参考`example_test.go`的单元测试代码，如需其它帮助请在[Issues](https://github.com/dxvgef/filter/issues)里提出。

---

## 字符串处理

#### 输入函数
- `FromString()` 输入`string`类型的数据

#### 格式化函数
- 
- `ToUpper` 字母转为大写
- `ToLower` 字母转为小写
- `Trim` 删除左右的指定字符
- `TrimSpace` 删除左右的空格，是 `Trim(" ")` 函数的封装
- `TrimLeft` 删除所有出现在左边的指定字符串
- `TrimRight` 删除所有出现在右边的指定字符串
- `TrimPrefix` 删除指定的前缀字符串
- `TrimSuffix` 删除指定的后缀字符串
- `Replace` 替换指定的字符串，可指定替换次数
- `ReplaceAll` 替换指定的字符串，替换所有
- `RemoveSpace` 删除字符串中所有出现的空格

### 编/解码函数
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

#### 校验函数
- `Require` 参数不能为零值
- `Equal` 匹配两个字符串相等
- `MatchLength` 匹配长度
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
- `IsUnsigned` 是无符号数值
- `IsLowerOrNumber` 是小写字母或数字
- `IsUpperOrNumber` 是大写字母或数字
- `IsLetterOrNumber` 是字母或数字
- `IsChinese` 是汉字
- `IsMail` 是电邮地址
- `IsIP` 是IPv4/v6地址
- `IsTCPAddr` 是IP:Port格式
- `IsMAC` 是MAC地址
- `IsJSON` 是有效的JSON格式
- `IsChinaIDNumber` 是中国大陆地区身份证号码
- `IsSQLObject` 是SQL对象名(库、表、字段)
- `IsSQLObjects` 是SQL对象名集合
- `IsUUID` 是UUID格式
- `IsURL` 是URL格式
- `HasLetter` 必须包含字母
- `HasLower` 必须包含小写字母
- `HasUpper` 必须包含大写字母
- `HasNumber` 必须包含数字
- `HasSymbol` 必须包含符号
- `HasPrefix` 必须包含指定的前缀字符串
- `HasSuffix` 必须包含指定的后缀字符串
- `HasStr` 必须包含指定的字符串
- `EnumStr` 仅允许[]string中的值
- `EnumInt` 仅允许[]int中的值
- `EnumInt32` 仅允许[]int32中的值
- `EnumInt64` 仅允许[]int64中的值
- `EnumFloat32` 仅允许[]float32中的值
- `EnumFloat64` 仅允许[]float64中的值
- `EnumStrSlice` 将数据转为[]string,并检查其元素是否存在于指定的[]string中
- `EnumIntSlice` 将数据转为[]int,并检查其元素是否存在于指定的[]int中
- `DenyStr` 阻止[]string中的值
- `DenyInt` 阻止[]int中的值
- `DenyInt32` 阻止[]int32中的值
- `DenyInt64` 阻止[]int64中的值
- `DenyFloat32` 阻止[]float32中的值
- `DenyFloat64` 阻止[]float64中的值
- `DenyOtherSymbol` 阻止指定之外的符号

#### 输出函数

###### 类型转换
- `String` 转为string类型
- `DefaultString` 转为string类型，出错则返回默认值
- `SliceString` 转为[]string类型
- `DefaultSliceString` 转为[]string类型，出错则返回默认值
- `Int` 转为int类型
- `DefaultInt` 转为int类型，出错则返回默认值
- `SliceInt` 转为[]int类型
- `DefaultSliceInt` 转为[]int类型，出错则返回默认值
- `Int8` 转为int8类型
- `DefaultInt8` 转为int8类型，出错则返回默认值
- `SliceInt8` 转为[]int8类型
- `DefaultSliceInt8` 转为[]int8类型，出错则返回默认值
- `Int16` 转为int16类型
- `DefaultInt16` 转为int16类型，出错则返回默认值
- `SliceInt16` 转为[]int16类型
- `DefaultSliceInt16` 转为[]int16类型，出错则返回默认值
- `Int32` 转为int32类型
- `DefaultInt32` 转为int32类型，出错则返回默认值
- `SliceInt32` 转为[]int32类型
- `DefaultSliceInt32` 转为[]int32类型，出错则返回默认值
- `Int64` 转为int64类型
- `DefaultInt64` 转为int64类型，出错则返回默认值
- `SliceInt64` 转为[]int64类型
- `DefaultSliceInt64` 转为[]int64类型，出错则返回默认值
- `Uint` 转为uint类型
- `DefaultUint` 转为uint类型，出错则返回默认值
- `SliceUint` 转为[]uint类型
- `DefaultSliceUint` 转为[]uint类型，出错则返回默认值
- `Uint8` 转为uint8类型
- `DefaultUint8` 转为uint8类型，出错则返回默认值
- `SliceUint8` 转为[]uint8类型
- `DefaultSliceUint8` 转为[]uint8类型，出错则返回默认值
- `Uint16` 转为uint16类型
- `DefaultUint16` 转为uint16类型，出错则返回默认值
- `SliceUint16` 转为[]uint16类型
- `DefaultSliceUint16` 转为[]uint16类型，出错则返回默认值
- `Uint32` 转为uint32类型
- `DefaultUint32` 转为uint32类型，出错则返回默认值
- `SliceUint32` 转为[]uint32类型
- `DefaultSliceUint32` 转为[]uint32类型，出错则返回默认值
- `Uint64` 转为uint64类型
- `DefaultUint64` 转为uint64类型，出错则返回默认值
- `SliceUint64` 转为[]uint64类型
- `DefaultSliceUint64` 转为[]uint64类型，出错则返回默认值
- `Float32` 转为float32类型
- `DefaultFloat32` 转为float32类型，出错则返回默认值
- `SliceFloat32` 转为[]float32类型
- `DefaultSliceFloat32` 转为[]float32类型，出错则返回默认值
- `Float64` 转为float64类型
- `DefaultFloat64` 转为float64类型，出错则返回默认值
- `SliceFloat64` 转为[]float64类型
- `DefaultSliceFloat64` 转为[]float64类型，出错则返回默认值
- `Bool` 转为bool类型
- `DefaultBool` 转为bool类型，出错则返回默认值

###### 赋值 
- `Set()` 通过反射将处理结果值赋值到`string`类型的变量

###### 结果
- `Error()` 返回处理过程中出更的错误（`error`类型）
- `Value()` 返回已处理的值
- `Result()` 返回 `Value(), Error()`

###### 自定义处理函数

- `Custom` 自定义字符串处理函数，详见 `CustomStringFunc`
