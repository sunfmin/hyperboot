package web

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Main(data WebData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\npackage main\n\nimport (\n\t\"")
	_buffer.WriteString(gorazor.HTMLEscape(data.WebPkg))
	_buffer.WriteString("\"\n\t\"")
	_buffer.WriteString(gorazor.HTMLEscape(data.HttpImplPkg))
	_buffer.WriteString("\"\n\t\"log\"\n\t\"net/http\"\n)\n\nfunc main() {\n\n\t")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("apihttpimpl.AddToMux(\"/api\", http.DefaultServeMux)\n\n\thttp.Handle(\"/\", web.Routes())\n\n\tlog.Println(\"Starting Server at 9000.\")\n\terr := http.ListenAndServe(\":9000\", nil)\n\tif err != nil {\n\t\tlog.Fatal(\"ListenAndServe: \", err)\n\t}\n}")

	return _buffer.String()
}
