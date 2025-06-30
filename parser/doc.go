package parser

import (
	"reflect"
	"strings"
)

type Doc string

func (d Doc) Description() string {
	return strings.SplitN(string(d), "`", 2)[0]
}

func (d Doc) ParamDoc(param string) string {
	parts := strings.Split(string(d), "`")
	if len(parts) != 3 {
		return ""
	}

	return reflect.StructTag(parts[1]).Get(param)
}
