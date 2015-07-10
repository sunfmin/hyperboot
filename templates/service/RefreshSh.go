package service

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func RefreshSh(data ServiceData) string {
	var _buffer bytes.Buffer

	var comment = "\n# "
	if data.IsApp {
		comment = "\n"
	}

	_buffer.WriteString("\nhypermusk -pkg=")
	_buffer.WriteString(gorazor.HTMLEscape((data.ApiPkg)))
	_buffer.WriteString(" -impl=")
	_buffer.WriteString(gorazor.HTMLEscape((data.ServicesPkg)))
	_buffer.WriteString(" -lang=server -outdir=./\ngofmt -w ./")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("apihttpimpl")
	_buffer.WriteString(gorazor.HTMLEscape((comment)))
	_buffer.WriteString("hypermusk -pkg=")
	_buffer.WriteString(gorazor.HTMLEscape((data.ApiPkg)))
	_buffer.WriteString(" -lang=javascript -outdir=./react")
	_buffer.WriteString(gorazor.HTMLEscape((comment)))
	_buffer.WriteString("gorazor ./web/templates ./web/templates")
	_buffer.WriteString(gorazor.HTMLEscape((comment)))
	_buffer.WriteString("cd react && npm install . && webpack .")

	return _buffer.String()
}
