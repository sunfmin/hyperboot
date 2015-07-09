package services

import (
	"fmt"
	"github.com/sunfmin/hyperboot/templates/db"
	"go/build"
	"path/filepath"
	"strings"
)

func CreateDb(rootPath string, module string) (err error) {
	err = CreateConfig(rootPath, module, "mysql", true)
	if err != nil {
		return
	}

	f := filepath.Join(module, "db", "db.go")
	cfile := filepath.Join(rootPath, f)

	configPath := strings.Replace(rootPath, build.Default.GOPATH+"/src/", "", 1)
	configPkg := fmt.Sprintf("%s/%s/config", configPath, module)

	err = createFileWithContentSkipExists(f, cfile, db.Db(db.DbData{Module: module, ConfigPkg: configPkg}), false)
	return
}
