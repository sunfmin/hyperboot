package db

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Db(data DbData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\npackage db\n\nimport (\n\t\"")
	_buffer.WriteString(gorazor.HTMLEscape(data.ConfigPkg))
	_buffer.WriteString("\"\n\t\"fmt\"\n\t_ \"github.com/go-sql-driver/mysql\"\n\t\"github.com/jinzhu/gorm\"\n)\n\nvar DB gorm.DB\nvar AllModels = []interface{}{\n\t&User{},\n}\n\nfunc init() {\n\thost := config.MySQLHost\n\tdatabase := config.MySQLDatabase\n\n\tvar err error\n\tconnect := fmt.Sprintf(\"root:%s@tcp(%s:%s)/%s?parseTime=true\", config.MySQLRootPassword, host, config.MySQLPort, database)\n\t// fmt.Println(connect)\n\tDB, err = gorm.Open(\"mysql\", connect)\n\n\tif err != nil {\n\t\tpanic(fmt.Sprintf(\"Error when connect database: '%v'\", err))\n\t}\n\n\tDB.LogMode(config.Env == \"development\")\n\n\tmodels := AllModels\n\n\tfor _, m := range models {\n\t\terr = DB.AutoMigrate(m).Error\n\t\tif err != nil {\n\t\t\tpanic(fmt.Sprintf(\"Error when migrate: '%v'\", err))\n\t\t}\n\t}\n\n}")

	return _buffer.String()
}
