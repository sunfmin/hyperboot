package services

import (
	"github.com/sunfmin/hyperboot/templates/config"
	"path/filepath"
)

func CreateConfig(rootPath string, module string, database string, withoutExistOutput bool) (err error) {
	f := filepath.Join(module, "config", "init.go")
	cfile := filepath.Join(rootPath, f)

	err = createFileWithContentSkipExists(f, cfile, config.Init(config.ConfigData{Module: module, Database: database}), withoutExistOutput)
	return
}
