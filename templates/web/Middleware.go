package web

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Middleware(data WebData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\npackage web\n\nimport (\n\t\"")
	_buffer.WriteString(gorazor.HTMLEscape(data.TemplatePkg))
	_buffer.WriteString("\"\n\t\"fmt\"\n\t\"log\"\n\t\"net/http\"\n\t\"os\"\n\t\"runtime\"\n)\n\n// Recovery is a Negroni middleware that recovers from any panics and writes a 500 if there was one.\ntype ProductionRecovery struct {\n\tLogger     *log.Logger\n\tPrintStack bool\n\tStackAll   bool\n\tStackSize  int\n}\n\n// NewRecovery returns a new instance of Recovery\nfunc NewProductionRecovery() *ProductionRecovery {\n\treturn &ProductionRecovery{\n\t\tLogger:     log.New(os.Stdout, \"[web] \", 0),\n\t\tPrintStack: true,\n\t\tStackAll:   false,\n\t\tStackSize:  1024 * 8,\n\t}\n}\n\nfunc (rec *ProductionRecovery) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {\n\tdefer func() {\n\t\tif err := recover(); err != nil {\n\t\t\trw.WriteHeader(http.StatusInternalServerError)\n\t\t\tstack := make([]byte, rec.StackSize)\n\t\t\tstack = stack[:runtime.Stack(stack, rec.StackAll)]\n\n\t\t\tf := \"PANIC: %s\\n%s\\n%s\\n%+v\\n\"\n\t\t\terrdesc := fmt.Sprintf(f, err, stack, r.URL.RawQuery, r.Header)\n\t\t\trec.Logger.Println(errdesc)\n\n\t\t\tfmt.Fprint(rw, templates.P500(errdesc))\n\n\t\t}\n\t}()\n\n\tnext(rw, r)\n}")

	return _buffer.String()
}
