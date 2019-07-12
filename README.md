# filter
golang开发的数据过滤器，由 **数据输入->数据清洗->数据验证->类型转换->结果输出** 四个部份组成。

- 自定义错误消息
- 过滤结果自动赋值到指定变量
- 批量过滤多个数据，自动赋值到对应变量

## 基本示例
单项数据过滤，并返回结果
```Go
password, err := FromString("123", "密码").Trim().MinLength(6).MaxLength(32).String()
if err != nil {
    log.Println(err.Error())
    return
}
log.Println(password)
```

单项数据过滤，并自动赋值到变量
```Go
var password string
err := Set(
    &password, FromString("Abc123-", "密码").
    MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol(),
)
if err != nil {
    log.Println(err.Error())
    return
}
```

多项数据过滤，并自动赋值到对应变量
```Go
var reqData struct {
    password string
    age      int16
}
err := MSet(
    El(&reqData.password,
        FromString("Abc123-", "密码").
            MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol(),
    ),
    El(&reqData.age,
        FromString("3", "年龄").
            IsDigit().MinInteger(18)),
)
if err != nil {
    log.Println(err.Error())
    return
}
log.Println("密码", reqData.password)
log.Println("年龄", reqData.age)
```

## 数据输入
**`FromString(str, name)`**
要过滤的数据来源，目前仅支持字符串
第一个参数str为来源参数值<br>
第二个参数为名，用于拼接默认错误消息

## 数据清洗
- `Trim()` 去除前后空格
- `RemoveSpace` 去除所有空格
- `ReplaceAll` 替换所有
- `SnakeCaseToCamelCase` 蛇形转驼峰: hello_world => helloWorld
- `SnakeCaseToPascalCase` 蛇形转帕斯卡: hello_world => HelloWorld
- `CamelCaseToSnakeCase` 驼峰(含帕斯卡)转蛇形 helloWorld/HelloWorld => hello_world

## 数据验证
所有数据验证函数，都可以传入自定义错误消息，例如MinLength(""自定义错误消息")
- `MinLength 最小长度` 最小长度
- `MinUTF8Length` UTF8编码最小长度
- `MaxLength` 最大长度
- `MaxUTF8Length` UTF8编码最大长度
- `MinInteger` 最小整数值
- `MaxInteger` 最大整数值
- `MinFloat` 最小浮点值
- `MaxFloat` 最大浮点值
- `IsLower` 小写字母
- `IsUpper` 大写字母
- `IsLetter` 字母
- `IsDigit` 数字
- `IsLowerOrDigit` 小写字母或数字
- `IsUpperOrDigit` 大写字母或数字
- `IsLetterOrDigit 字母或数字` 字母或数字
- `IsChinese` 汉字
- `IsMail` 电邮地址
- `IsIP` IPv4/v6地址
- `IsJSON` 有效的JSON格式
- `IsChineseTel` 中国大陆地区固定电话号码
- `IsChineseMobile` 中国大陆地区手机号码
- `IsChineseIDNumber` 中国大陆地区身份证号码
- `IsSQLobject` SQL对象名(库、表、字段)
- `IsSQLobjects` SQL对象名集合
- `IsUUID` UUID格式
- `HasLetter` 存在字母
- `HasLower` 存在小写字母
- `HasUpper` 存在大写字母
- `HasDigit` 存在数字
- `HasSymbol` 存在符号
- `HasPrefix` 存在指定的前缀字符串
- `Contains` 存在指定字符串
- `DenyStrings` 阻止存在于[]string中的值
- `DenyInts` 阻止存在于[]int中的值
- `DenyInts32` 阻止存在于[]int32中的值
- `DenyInts64` 阻止存在于[]int64中的值
- `DenyFloats32` 阻止存在于[]float32中的值
- `DenyFloats64` 阻止存在于[]float64中的值

## 类型转换
- `String` 转为string类型，并返回error
- `MustString` 转为string类型，如果失败则返回默认值
- `Strings` 按指定分隔符，转为[]string类型
- `Int` 转为int类型，并返回error
- `MustInt` 转为int类型，如果失败则返回默认值
- `Int8` 转为int8类型，并返回error
- `MustInt8` 转为int8类型，如果失败则返回默认值
- `Int16` 转为int16类型，并返回error
- `MustInt16` 转为int16类型，如果失败则返回默认值
- `Int32` 转为int32类型，并返回error
- `MustInt32` 转为int32类型，如果失败则返回默认值
- `Int64` 转为int64类型，并返回error
- `MustInt64` 转为int64类型，如果失败则返回默认值
- `Float32` 转为float32类型，并返回errBor
- `MustFloat32` 转为float8类型，如果失败则返回默认值
- `Float64` 转为float64类型，并返回error
- `MustFloat64` 转为float64类型，如果失败则返回默认值
- `Bool` 转为bool类型，并返回error
- `MustBool` 转为bool类型，如果失败则返回默认值

## 结果输出
除了使用类型转换函数得到过滤后的数据，还可以使用以下函数将过滤结果赋值到指定变量
- `Set` 将单个过滤结果赋值到变量
- `MSet` 将多个过滤结果赋值到对应的变量