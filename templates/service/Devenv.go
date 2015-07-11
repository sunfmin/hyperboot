package service

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"strings"
)

func Devenv(data ServiceData) string {
	var _buffer bytes.Buffer

	var upperModule = strings.ToUpper(data.Module)
	var lowerModule = strings.ToLower(data.Module)

	_buffer.WriteString("\nexport ")
	_buffer.WriteString(gorazor.HTMLEscape((upperModule)))
	_buffer.WriteString("_MYSQL_PORT_3306_TCP_ADDR=localhost\nexport ")
	_buffer.WriteString(gorazor.HTMLEscape((upperModule)))
	_buffer.WriteString("_MYSQL_ENV_MYSQL_ROOT_PASSWORD=\nexport ")
	_buffer.WriteString(gorazor.HTMLEscape((upperModule)))
	_buffer.WriteString("_MYSQL_PORT_3306_TCP_PORT=3306\nexport ")
	_buffer.WriteString(gorazor.HTMLEscape((upperModule)))
	_buffer.WriteString("_MYSQL_DATABASE=")
	_buffer.WriteString(gorazor.HTMLEscape((lowerModule)))
	_buffer.WriteString("_development\nexport ")
	_buffer.WriteString(gorazor.HTMLEscape((upperModule)))
	_buffer.WriteString("_ENV=development\nexport ")
	_buffer.WriteString(gorazor.HTMLEscape((upperModule)))
	_buffer.WriteString("_WEBDIR=./www")

	return _buffer.String()
}
