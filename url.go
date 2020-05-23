package filter

import "net/url"

// 是有效的URL
func (obj *Object) IsURL(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	_, obj.err = url.ParseRequestURI(obj.rawValue)
	return obj
}
