package web

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Routes(data WebData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\npackage web\n\nimport (\n\t\"")
	_buffer.WriteString(gorazor.HTMLEscape(data.ConfigPkg))
	_buffer.WriteString("\"\n\t\"github.com/codegangsta/negroni\"\n\t\"github.com/gorilla/mux\"\n\t\"github.com/phyber/negroni-gzip/gzip\"\n\t\"net/http\"\n\t\"os\"\n\t\"path\"\n)\n\nfunc Routes() (n *negroni.Negroni) {\n\n\trouter := mux.NewRouter()\n\n\t// router.HandleFunc(\"/p/{articleId}.html\", Article)\n\n\tn = negroni.New()\n\tn.Use(gzip.Gzip(gzip.DefaultCompression))\n\n\tif config.Env != \"production\" {\n\t\tn.Use(negroni.NewRecovery())\n\t} else {\n\t\tn.Use(NewProductionRecovery())\n\t}\n\n\twebdir := path.Clean(config.WebDir)\n\tif _, err := os.Stat(webdir); err != nil {\n\t\tpanic(err)\n\t\treturn\n\t}\n\n\tstatic := &negroni.Static{\n\t\tDir:       http.Dir(webdir),\n\t\tPrefix:    \"/")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("\",\n\t\tIndexFile: \"index.html\",\n\t}\n\n\tn.Use(static)\n\tn.Use(negroni.NewLogger())\n\tn.UseHandler(router)\n\treturn\n}")

	return _buffer.String()
}
