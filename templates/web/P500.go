package web

import (
	"bytes"
)

func P500(data WebData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("@{\n\timport (\n\t\t// \"strings\"\n\t)\n\n\tvar errdesc string\n}\n\n<h1>500</h1>\n\n<p> 网站爆掉了，已通知程序猿 </p>\n\n<div id=\"errdesc\" class=\"hidden\">\n\t@errdesc\n</div>")

	return _buffer.String()
}
