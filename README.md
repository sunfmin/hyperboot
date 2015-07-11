# Hyperboot Stack

Libraries, Tools and Frameworks, Database that the Hyperboot App Stack are using:

- [react](http://facebook.github.io/react)
- [webpack](http://webpack.github.io)
- [jQuery](http://jquery.com)
- [bootstrap](http://getbootstrap.com)
- [less](http://lesscss.org)
- [gorm](http://github.com/jinzhu/gorm)
- [gorazor](http://github.com/sipin/gorazor)
- [hypermusk](http://github.com/hypermusk/hypermusk)
- [mysql](http://dev.mysql.com)

# Installation

1. Install ORM, Templates, Router, Middleware go packages:

	```
	go get github.com/sipin/gorazor
	go get github.com/hypermusk/hypermusk
	go get github.com/jinzhu/gorm
	go get github.com/go-sql-driver/mysql
	go get github.com/codegangsta/negroni
	go get github.com/gorilla/mux
	go get github.com/phyber/negroni-gzip/gzip

	```

2. Install Hyperboot

	```
	go get github.com/sunfmin/hyperboot
	```
3. Bootstrap your App

	```
	cd $GOPATH/github.com/{yourname}/{yourapp}
	hyperboot app --module="hello"

	```

4. Install Webpack, React, jQuery, Bootstrap

	```
	brew install npm
	npm install webpack -g
	cd hello/react
	npm install .
	```
5. Create MySQL Database

	Install MySQL if you don't have it

	```
	brew install mysql
	```
	Create the development database
	```
	mysql -u root -e "CREATE DATABASE hello_development CHARSET utf8"

	```
6. Generate and Compile javascript, gorazor templates, and hypermusk generated code
	```
	cd hello
	sh refresh.sh
	```

7. Start the app
	```
	go run main.go
	```



# Directories

```
{Project Root}
	mod1
		api
		cmds
		config
		db
			models.go
			db.go
		services
		tests
		web
			templates
				Home.gohtml
			home.go
		mod1apihttpimpl
		react
			package.json
			webpack.config.js
			api.js
			app.js
			css.js
			main.less
		www
			index.html
		main.go
	mod2
		... # similiar to mod1
	{project_name}
		main.go

```
