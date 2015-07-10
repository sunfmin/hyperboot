package services

import (
	"fmt"
	"github.com/sunfmin/hyperboot/templates/react"
	"github.com/sunfmin/hyperboot/templates/web"
	"go/build"
	"path/filepath"
	"strings"
)

func CreateWeb(rootPath string, module string, withoutExistOutput bool) (err error) {
	f := filepath.Join(module, "react", "app.js")
	cfile := filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, react.AppJs(react.ReactData{Module: module}), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "react", "configured_api.js")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, react.ConfiguredApiJs(react.ReactData{Module: module}), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "react", "main.less")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, react.MainLess(react.ReactData{Module: module}), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "react", "package.json")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, react.PackageJson(react.ReactData{Module: module}), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "react", "webpack.config.js")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, react.WebpackConfigJs(react.ReactData{Module: module}), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "react/components", "Hello.js")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, react.HelloJs(react.ReactData{Module: module}), withoutExistOutput)
	if err != nil {
		return
	}

	appPath := strings.Replace(rootPath, build.Default.GOPATH+"/src/", "", 1)
	templatePkg := fmt.Sprintf("%s/%s/web/templates", appPath, module)
	webPkg := fmt.Sprintf("%s/%s/web", appPath, module)
	configPkg := fmt.Sprintf("%s/%s/config", appPath, module)
	httpImplPkg := fmt.Sprintf("%s/%s/%sapihttpimpl", appPath, module, module)
	data := web.WebData{
		Module:      module,
		TemplatePkg: templatePkg,
		WebPkg:      webPkg,
		HttpImplPkg: httpImplPkg,
		ConfigPkg:   configPkg,
	}

	f = filepath.Join(module, "web", "routes.go")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, web.Routes(data), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "web", "middleware.go")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, web.Middleware(data), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "web/templates", "P500.gohtml")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, web.P500(data), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "", "main.go")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, web.Main(data), withoutExistOutput)
	if err != nil {
		return
	}

	f = filepath.Join(module, "www", "index.html")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, web.Index(data), withoutExistOutput)
	if err != nil {
		return
	}

	return
}
