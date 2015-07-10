package react

import (
	"bytes"
)

func AppJs(data ReactData) string {
	var _buffer bytes.Buffer
	_buffer.WriteString("\n/* @flow */\n\nrequire(\"./main.less\");\nwindow.jQuery = require('jquery');\nvar service = require(\"./configured_api.js\")\n\nvar React = require(\"react\");\nvar Hello = require(\"./components/Hello.js\")\nReact.render(<Hello service={service}/>, document.getElementById(\"react\"));")

	return _buffer.String()
}
