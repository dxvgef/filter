package filter

type Str struct {
	name         string  // 变量名称，用于拼接错误消息
	rawValue     string  // 原始值
	currentValue string  // 当前值
	err          error   // 错误
	int64        int64   // 转成int64类型的值，用于做数值运算，减少类型转换的次数
	float64      float64 // 转成float64类型的值，用于做数值运算，减少类型转换的次数
	require      bool    // 不能为零值
	sep          string  // 分隔符
}

type StrType interface {
	Require() StrType
	JoinStr(...string) StrType
	JoinBytes(...[]byte) StrType
	JoinByte(...byte) StrType
	JoinRune(...rune) StrType
	Trim(string) StrType
	TrimSpace() StrType
	TrimLeft(string) StrType
	TrimRight(string) StrType
	TrimPrefix(string) StrType
	TrimSuffix(string) StrType
	RemoveSpace(string) StrType
	ToUpper() StrType
	ToLower() StrType
	SnakeCaseToCamelCase() StrType
	SnakeCaseToPascalCase() StrType
	CamelCaseToSnakeCase() StrType
	Base64StdEncode() StrType
	Base64StdDecode(...string) StrType
	Base64RawStdEncode() StrType
	Base64RawStdDecode(...string) StrType
	Base64URLEncode() StrType
	Base64URLDecode(...string) StrType
	Base64RawURLEncode() StrType
	Base64RawURLDecode(...string) StrType
	HTMLUnescape() StrType
	HTMLEscape() StrType
	URLPathUnescape(...string) StrType
	URLPathEscape() StrType
	URLQueryUnescape(...string) StrType
	URLQueryEscape() StrType
	HasLetter(...string) StrType
	HasLower(...string) StrType
	HasUpper(...string) StrType
	HasDigit(...string) StrType
	HasSymbol(...string) StrType
	Contains(string, ...string) StrType
	HasString(string, ...string) StrType
	HasPrefix(string, ...string) StrType
	HasSuffix(string, ...string) StrType
	Equal(string, ...string) StrType
	IsBool(...string) StrType
	IsLower(...string) StrType
	IsUpper(...string) StrType
	IsLetter(...string) StrType
	IsDigit(...string) StrType
	IsLowerOrDigit(...string) StrType
	IsUpperOrDigit(...string) StrType
	IsLetterOrDigit(...string) StrType
	IsChinese(...string) StrType
	IsChinaTel(...string) StrType
	IsMail(...string) StrType
	IsIP(...string) StrType
	IsJSON(...string) StrType
	IsChinaIDNumber(...string) StrType
	IsChineseMobile(...string) StrType
	IsSqlObject(...string) StrType
	IsSqlObjects(string, ...string) StrType
	IsURL(...string) StrType
	IsUUID(...string) StrType
	MinLength(int, ...string) StrType
	MinUTF8Length(int, ...string) StrType
	MaxLength(int, ...string) StrType
	MaxUTF8Length(int, ...string) StrType
	UpdateInt64(...string) StrType
	UpdateFloat64(...string) StrType
	MinInteger(int64, ...string) StrType
	MaxInteger(int64, ...string) StrType
	MinFloat(float64, ...string) StrType
	MaxFloat(float64, ...string) StrType
	SetSeparator(string) StrType
	EnumString([]string, ...string) StrType
	EnumInt([]int, ...string) StrType
	EnumInt32([]int32, ...string) StrType
	EnumInt64([]int64, ...string) StrType
	EnumFloat32([]float32, ...string) StrType
	EnumFloat64([]float64, ...string) StrType
	DenyString([]string, ...string) StrType
	DenyInt([]int, ...string) StrType
	DenyInt32([]int32, ...string) StrType
	DenyInt64([]int64, ...string) StrType
	DenyFloat32([]float32, ...string) StrType
	DenyInFloat64([]float64, ...string) StrType
	String() (string, error)
	DefaultString(string) string
	SliceString() ([]string, error)
	DefaultSliceString([]string) []string
	Int() (int, error)
	DefaultInt(int) int
	SliceInt() ([]int, error)
	DefaultSliceInt([]int) []int
	Uint() (uint, error)
	DefaultUint(uint) uint
	SliceUint() ([]uint, error)
	DefaultSliceUint([]uint) []uint
	Int8() (int8, error)
	DefaultInt8(int8) int8
	SliceInt8() ([]int8, error)
	DefaultSliceInt8([]int8) []int8
	Uint8() (uint8, error)
	DefaultUint8(uint8) uint8
	SliceUint8() ([]uint8, error)
	DefaultSliceUint8([]uint8) []uint8
	Int16() (int16, error)
	DefaultInt16(int16) int16
	SliceInt16() ([]int16, error)
	DefaultSliceInt16([]int16) []int16
	Uint16() (uint16, error)
	DefaultUint16(uint16) uint16
	SliceUint16() ([]uint16, error)
	DefaultSliceUint16([]uint16) []uint16
	Int32() (int32, error)
	DefaultInt32(int32) int32
	SliceInt32() ([]int32, error)
	DefaultSliceInt32([]int32) []int32
	Uint32() (uint32, error)
	DefaultUint32(uint32) uint32
	SliceUint32() ([]uint32, error)
	DefaultSliceUint32([]uint32) []uint32
	Int64() (int64, error)
	DefaultInt64(int64) int64
	SliceInt64() ([]int64, error)
	DefaultSliceInt64([]int64) []int64
	Uint64() (uint64, error)
	DefaultUint64(uint64) uint64
	SliceUint64() ([]uint64, error)
	DefaultSliceUint64([]uint64) []uint64
	Float32() (float32, error)
	DefaultFloat32(float32) float32
	SliceFloat32() ([]float32, error)
	DefaultSliceFloat32([]float32) []float32
	Float64() (float64, error)
	DefaultFloat64(float64) float64
	SliceFloat64() ([]float64, error)
	DefaultSliceFloat64([]float64) []float64
	Bool() (bool, error)
	DefaultBool(bool) bool
	SliceBool() ([]bool, error)
	DefaultSliceBool([]bool) []bool
	Assign(interface{}, ...string) error
	Error() error
}
