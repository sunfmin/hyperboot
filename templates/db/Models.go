package db

import (
	"bytes"
)

func Models(data ModelData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\npackage db\n\nimport (\n\t\"time\"\n)\n\n/*\nDescribe table columns like the following:\n\tID           int `gorm:\"primary_key\"`\n\tSubscribed   bool\n\tBirthday     time.Time\n    Age          int\n    Name         string  `sql:\"size:255\"` // Default size for string is 255, you could reset it with this tag\n    Num          int     `sql:\"AUTO_INCREMENT\"`\n    UserID       int     `sql:\"index\"` // tag `index` will create index for this field when using AutoMigrate\n    Email        string  `sql:\"type:varchar(100);unique_index\"` // Set field's sql type, tag `unique_index` will create unique index\n    Name         string `sql:\"index:idx_name_code\"` // Will create combined index if find other fields defined same name\n    Code         string `sql:\"index:idx_name_code\"` // `unique_index` also works\n*/\n\n\ntype User struct {\n\tID        int\n\tName      string\n\tCreatedAt time.Time\n\tUpdatedAt time.Time\n}")

	return _buffer.String()
}
