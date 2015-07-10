package services

import (
	"fmt"
	"github.com/sunfmin/hyperboot/templates/service"
	"go/build"
	"path/filepath"
	"strings"
)

func CreateServices(rootPath string, module string, withoutExistOutput bool) (err error) {
	err = CreateConfig(rootPath, module, "mysql", true)
	if err != nil {
		return
	}

	err = CreateDb(rootPath, module, true)
	if err != nil {
		return
	}

	err = CreateModels(rootPath, module, true)
	if err != nil {
		return
	}

	err = CreateApi(rootPath, module, true)
	if err != nil {
		return
	}

	f := filepath.Join(module, "services", "app.go")
	cfile := filepath.Join(rootPath, f)

	apiPath := strings.Replace(rootPath, build.Default.GOPATH+"/src/", "", 1)

	apiPkg := fmt.Sprintf("%s/%s/%sapi", apiPath, module, module)
	dbPkg := fmt.Sprintf("%s/%s/db", apiPath, module)
	configPkg := fmt.Sprintf("%s/%s/config", apiPath, module)

	err = createFileWithContentSkipExists(f, cfile, service.Services(service.ServiceData{Module: module, ApiPkg: apiPkg, DbPkg: dbPkg, ConfigPkg: configPkg}), withoutExistOutput)

	servicesPkg := fmt.Sprintf("%s/%s/services", apiPath, module)

	f = filepath.Join(module, "", "refresh.sh")
	cfile = filepath.Join(rootPath, f)
	err = createFileWithContentSkipExists(f, cfile, service.RefreshSh(service.ServiceData{Module: module, ApiPkg: apiPkg, ServicesPkg: servicesPkg}), withoutExistOutput)
	if err != nil {
		return
	}

	return
}

func CreateApi(rootPath string, module string, withoutExistOutput bool) (err error) {

	f := filepath.Join(module, module+"api", "api.go")
	cfile := filepath.Join(rootPath, f)

	err = createFileWithContentSkipExists(f, cfile, service.Api(service.ServiceData{Module: module}), withoutExistOutput)
	return
}
