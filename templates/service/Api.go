package service

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"strings"
)

func Api(data ServiceData) string {
	var _buffer bytes.Buffer

	var lowerModule = strings.ToLower(data.Module)

	_buffer.WriteString("\npackage ")
	_buffer.WriteString(gorazor.HTMLEscape((lowerModule)))
	_buffer.WriteString("api\n\nimport (\n\t\"errors\"\n)\n\nvar (\n\tPermissionDeniedError = errors.New(\"permission denied.\")\n)\n\ntype User struct {\n\tId           string\n\tName        string\n\tUpperName   string\n}\n\n\ntype AppService interface {\n\tGetUserService(session string) (service UserService, err error)\n}\n\n\ntype UserService interface {\n\tPutUser(id int, name string) (user *User, err error)\n}")

	return _buffer.String()
}
