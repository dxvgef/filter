package filter

import "net/url"

// 是有效的URL
func (obj *Object) IsURL(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	_, obj.err = url.Parse(obj.rawValue)
	return obj
}

// 是有效的RequestURI
func (obj *Object) IsRequestURI(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	_, obj.err = url.ParseRequestURI(obj.rawValue)
	return obj
}

// 是有效的URL Query
func (obj *Object) IsURLQuery(customError ...string) *Object {
	if obj.err != nil || obj.rawValue == "" {
		return obj
	}
	_, obj.err = url.ParseQuery(obj.rawValue)
	return obj
}
