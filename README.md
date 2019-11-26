# filter
golang开发的数据过滤器，由 **数据输入->数据清洗->数据验证->类型转换->结果输出** 四个部份组成。

- 自定义错误消息
- 过滤结果自动赋值到指定变量
- 批量过滤多个数据，自动赋值到对应变量

## 基本示例
单项数据过滤，并返回结果
```Go
password, err := filter.FromString("123", "密码").
	Trim().
	MinLength(6).MaxLength(32).
	String()
if err != nil {
    log.Println(err.Error())
    return
}
log.Println(password)
```

单项数据过滤，并自动赋值到变量
```Go
var password string
err := filter.Set(&password, filter.FromString("Abc123-", "密码").
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
err := filter.MSet(
    filter.El(&reqData.password, filter.FromString("Abc123-", "密码").
        Required().MinLength(6).MaxLength(32).HasLetter().HasUpper().HasDigit().HasSymbol(),
    ),
    filter.El(&reqData.age, filter.FromString("3", "年龄").
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
- `Separator` 指定Slice类型的分隔符

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
- `IsBool` 布尔值
- `IsLower` 小写字母
- `IsUpper` 大写字母
- `IsLetter` 字母
- `IsDigit` 数字，不能包含有任何符号或其它字符
- `IsLowerOrDigit` 小写字母或数字
- `IsUpperOrDigit` 大写字母或数字
- `IsLetterOrDigit` 字母或数字
- `IsChinese` 汉字
- `IsMail` 电邮地址
- `IsIP` IPv4/v6地址
- `IsJSON` 有效的JSON格式
- `IsTimestamp` 有效的Unix时间戳
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
- `InString` 值存在于指定的string中
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
- `Strings` 按指定分隔符，转为[]string类型
- `DefaultStrings` 按指定分隔符，转为[]string类型，如果出现错误则只返回默认值
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