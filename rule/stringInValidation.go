package rule

type StringInValidator func(v string, slice [] string) bool

func stringContains(v string, slice []string) bool {
	for _, e := range slice {
		if e == v {
			return true
		}
	}
	return false
}
func NewStringInValidate(slice []string, validator StringInValidator, message string) Rule {
	return Rule{
		Message:             message,
		StringValueValidate: validator,
		StringValue:         slice,
	}

}
