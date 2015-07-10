package service

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func Services(data ServiceData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n\npackage services\n\nimport (\n\t\"")
	_buffer.WriteString(gorazor.HTMLEscape(data.ApiPkg))
	_buffer.WriteString("\"\n\t\"")
	_buffer.WriteString(gorazor.HTMLEscape(data.ConfigPkg))
	_buffer.WriteString("\"\n\t\"")
	_buffer.WriteString(gorazor.HTMLEscape(data.DbPkg))
	_buffer.WriteString("\"\n)\n\n\ntype AppServiceImpl struct {\n}\n\nfunc (as *AppServiceImpl) GetUserService(session string) (us ")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("api.UserService, err error) {\n\tus = &UserServiceImpl{}\n\treturn\n}\n\ntype UserServiceImpl struct {\n}\n\nfunc (us *UserServiceImpl) PutUser(id int, name string) (user * ")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("api.User, err error) {\n\t_ = config.Env\n\tdbUser := &db.User{}\n\terr = db.DB.Where(&db.User{ID: id}).Assign(&db.User{Name: name}).FirstOrCreate(dbUser).Error\n\tuser = &")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("api.User{\n\t\tName: dbUser.Name,\n\t}\n\treturn\n}\n\nvar DefaultAppService = new(AppServiceImpl)")

	return _buffer.String()
}
