package web

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Index(data WebData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n<!doctype html>\n<html lang=\"en\">\n  <head>\n    <meta charset=\"utf-8\">\n    <title>Main</title>\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <link rel=\"stylesheet\" href=\"/")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("/assets/main.css\">\n  </head>\n  <body>\n    <section id=\"react\"></section>\n    <script src=\"/")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("/assets/main.js\"></script>\n  </body>\n</html>")

	return _buffer.String()
}
