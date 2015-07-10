package react

import (
	"bytes"
	"github.com/sipin/gorazor/gorazor"
)

func PackageJson(data ReactData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n{\n  \"name\": \"")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("\",\n  \"version\": \"0.0.1\",\n  \"description\": \"")
	_buffer.WriteString(gorazor.HTMLEscape((data.Module)))
	_buffer.WriteString("\",\n  \"main\": \"app.js\",\n  \"directories\": {\n    \"test\": \"test\"\n  },\n  \"dependencies\": {\n    \"moment\": \"^2.8.4\",\n    \"react\": \"^0.12.2\",\n    \"webpack\": \"^1.4.15\",\n    \"jsx-loader\": \"^0.12.2\",\n    \"css-loader\": \"^0.9.0\",\n    \"style-loader\": \"^0.8.2\",\n    \"less\": \"^2.2.0\",\n    \"less-loader\": \"^2.0.0\",\n    \"extract-text-webpack-plugin\": \"^0.3.6\",\n    \"file-loader\": \"^0.8.1\",\n    \"bootstrap\": \"^3.3.1\",\n    \"url-loader\": \"^0.5.5\",\n    \"jquery\": \"^2.1.3\"\n  },\n  \"devDependencies\": {},\n  \"scripts\": {},\n  \"repository\": {\n    \"type\": \"git\",\n    \"url\": \"git://yourdomain.com/yourrepository.git\"\n  },\n  \"keywords\": [\n    \"ui\",\n    \"widget\"\n  ],\n  \"author\": {\n    \"name\": \"Your Name\"\n  },\n  \"license\": \"MIT\",\n  \"bugs\": {\n    \"url\": \"https://yourdomain.com/yourrepository/issues\"\n  },\n  \"homepage\": \"https://yourdomain.com/yourrepository\"\n}")

	return _buffer.String()
}
