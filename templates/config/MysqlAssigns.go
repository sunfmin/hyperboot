package config

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func MysqlAssigns(upcaseModule string) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\nMySQLHost = envOrPanic(\"")
	_buffer.WriteString(gorazor.HTMLEscape((upcaseModule)))
	_buffer.WriteString("_MYSQL_PORT_3306_TCP_ADDR\", false)\nMySQLPort = envOrPanic(\"")
	_buffer.WriteString(gorazor.HTMLEscape((upcaseModule)))
	_buffer.WriteString("_MYSQL_PORT_3306_TCP_PORT\", false)\nMySQLRootPassword = envOrPanic(\"")
	_buffer.WriteString(gorazor.HTMLEscape((upcaseModule)))
	_buffer.WriteString("_MYSQL_ENV_MYSQL_ROOT_PASSWORD\", true)\nMySQLDatabase = envOrPanic(\"")
	_buffer.WriteString(gorazor.HTMLEscape((upcaseModule)))
	_buffer.WriteString("_MYSQL_DATABASE\", false)")

	return _buffer.String()
}
