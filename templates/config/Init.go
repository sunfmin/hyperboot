package config

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
	"strings"
)

func Init(data ConfigData) string {
	var _buffer bytes.Buffer

	var upcaseModule = strings.ToUpper(data.Module)

	_buffer.WriteString("\npackage config\n\nimport (\n\t\"log\"\n\t\"os\"\n\t\"strings\"\n)\n\n// this is configured from env variables\nvar (\n\tEnv                 string\n\n\t")
	if data.Database == "mysql" {

		_buffer.WriteString((MysqlVars()))

	}
	_buffer.WriteString("\n)\n\nfunc init() {\n\tEnv = envOrPanic(\"")
	_buffer.WriteString(gorazor.HTMLEscape((upcaseModule)))
	_buffer.WriteString("_ENV\", false)\n\t")
	if data.Database == "mysql" {

		_buffer.WriteString((MysqlAssigns(upcaseModule)))

	}
	_buffer.WriteString("\n}\n\nfunc envOrPanic(key string, allowEmpty bool) (r string) {\n\tr = os.Getenv(key)\n\tif r == \"\" && !allowEmpty {\n\t\tpanic(\"env \" + key + \" is not set\")\n\t}\n\tlogValue := r\n\tif strings.Contains(logValue, \"PASSWORD\") || strings.Contains(logValue, \"SECRET\") {\n\t\tlogValue = \"<HIDDEN>\"\n\t}\n\tlog.Printf(\"Configure: %s = %s\\n\", key, logValue)\n\treturn\n}")

	return _buffer.String()
}
