package config

import (
	"bytes"
)

func MysqlVars() string {
	var _buffer bytes.Buffer
	_buffer.WriteString("MySQLHost          \tstring\nMySQLPort          \tstring\nMySQLRootPassword   string\nMySQLDatabase \t\tstring")

	return _buffer.String()
}
