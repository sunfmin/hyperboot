package react

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func ConfiguredApiJs(data ReactData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\nwindow.jQuery=require(\"jquery\");\nrequire(\"./")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("api.js\");\nwindow.")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("api.baseurl = \"/api\";\nmodule.exports = new ")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("api.AppService();")

	return _buffer.String()
}
