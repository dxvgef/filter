package filter

type Str struct {
	name         string // 变量名称，用于拼接错误消息
	rawValue     string // 原始值
	currentValue string // 当前值
	err          error  // 错误
	require      bool   // 不能为零值
	requireErr   string
}

type StrType interface {
	// 清理
	Trim(string) StrType
	TrimSpace() StrType
	TrimLeft(string) StrType
	TrimRight(string) StrType
	TrimPrefix(string) StrType
	TrimSuffix(string) StrType
	RemoveSpace() StrType
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
	JoinStr(...string) StrType
	JoinBytes(...[]byte) StrType
	JoinByte(...byte) StrType
	JoinRune(...rune) StrType

	// 校验
	Require(...string) StrType
	Equal(string, ...string) StrType

	HasLetter(...string) StrType
	HasLower(...string) StrType
	HasUpper(...string) StrType
	HasDigit(...string) StrType
	HasSymbol(...string) StrType
	HasString(string, ...string) StrType
	HasPrefix(string, ...string) StrType
	HasSuffix(string, ...string) StrType

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
	IsChinaMobile(...string) StrType
	IsSQLObject(...string) StrType
	IsSQLObjects(string, ...string) StrType
	IsURL(...string) StrType
	IsUUID(...string) StrType

	MinLength(int, ...string) StrType
	MinUTF8Length(int, ...string) StrType
	MaxLength(int, ...string) StrType
	MaxUTF8Length(int, ...string) StrType

	MinInteger(int64, ...string) StrType
	MaxInteger(int64, ...string) StrType
	MinFloat(float64, ...string) StrType
	MaxFloat(float64, ...string) StrType

	EnumString([]string, ...string) StrType
	EnumInt([]int, ...string) StrType
	EnumInt8([]int8, ...string) StrType
	EnumInt16([]int16, ...string) StrType
	EnumInt32([]int32, ...string) StrType
	EnumInt64([]int64, ...string) StrType
	EnumFloat32([]float32, ...string) StrType
	EnumFloat64([]float64, ...string) StrType

	DenyString([]string, ...string) StrType
	DenyInt([]int, ...string) StrType
	DenyInt8([]int8, ...string) StrType
	DenyInt16([]int16, ...string) StrType
	DenyInt32([]int32, ...string) StrType
	DenyInt64([]int64, ...string) StrType
	DenyFloat32([]float32, ...string) StrType
	DenyFloat64([]float64, ...string) StrType

	// 输出
	String() (string, error)
	DefaultString(string) string
	SliceString(string) ([]string, error)
	DefaultSliceString(string, []string) []string
	Int() (int, error)
	DefaultInt(int) int
	SliceInt(string) ([]int, error)
	DefaultSliceInt(string, []int) []int
	Uint() (uint, error)
	DefaultUint(uint) uint
	SliceUint(string) ([]uint, error)
	DefaultSliceUint(string, []uint) []uint
	Int8() (int8, error)
	DefaultInt8(int8) int8
	SliceInt8(string) ([]int8, error)
	DefaultSliceInt8(string, []int8) []int8
	Uint8() (uint8, error)
	DefaultUint8(uint8) uint8
	SliceUint8(string) ([]uint8, error)
	DefaultSliceUint8(string, []uint8) []uint8
	Int16() (int16, error)
	DefaultInt16(int16) int16
	SliceInt16(string) ([]int16, error)
	DefaultSliceInt16(string, []int16) []int16
	Uint16() (uint16, error)
	DefaultUint16(uint16) uint16
	SliceUint16(string) ([]uint16, error)
	DefaultSliceUint16(string, []uint16) []uint16
	Int32() (int32, error)
	DefaultInt32(int32) int32
	SliceInt32(string) ([]int32, error)
	DefaultSliceInt32(string, []int32) []int32
	Uint32() (uint32, error)
	DefaultUint32(uint32) uint32
	SliceUint32(string) ([]uint32, error)
	DefaultSliceUint32(string, []uint32) []uint32
	Int64() (int64, error)
	DefaultInt64(int64) int64
	SliceInt64(string) ([]int64, error)
	DefaultSliceInt64(string, []int64) []int64
	Uint64() (uint64, error)
	DefaultUint64(uint64) uint64
	SliceUint64(string) ([]uint64, error)
	DefaultSliceUint64(string, []uint64) []uint64
	Float32() (float32, error)
	DefaultFloat32(float32) float32
	SliceFloat32(string) ([]float32, error)
	DefaultSliceFloat32(string, []float32) []float32
	Float64() (float64, error)
	DefaultFloat64(float64) float64
	SliceFloat64(string) ([]float64, error)
	DefaultSliceFloat64(string, []float64) []float64
	Bool() (bool, error)
	DefaultBool(bool) bool
	SliceBool(string) ([]bool, error)
	DefaultSliceBool(string, []bool) []bool

	// 赋值到普通变量
	Set(interface{}, ...string) error
	// 赋值到切片变量
	SetSlice(interface{}, string, ...string) error

	// 结果
	Error() error
}
